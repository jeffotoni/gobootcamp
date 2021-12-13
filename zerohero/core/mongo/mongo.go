package mongo

import (
	"context"
	"errors"
	"log"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ZeroHero struct {
	Response    string      `json:"response"`
	ID          string      `json:"id"`
	UUID        string      `json:"uuid,omitempty" bson:"_id"`
	Name        string      `json:"name"`
	Powerstats  Powerstats  `json:"powerstats"`
	Biography   Biography   `json:"biography"`
	Appearance  Appearance  `json:"appearance"`
	Work        Work        `json:"work"`
	Connections Connections `json:"connections"`
	Image       Image       `json:"image"`
}
type Powerstats struct {
	Intelligence string `json:"intelligence"`
	Strength     string `json:"strength"`
	Speed        string `json:"speed"`
	Durability   string `json:"durability"`
	Power        string `json:"power"`
	Combat       string `json:"combat"`
}
type Biography struct {
	FullName        string   `json:"full-name"`
	AlterEgos       string   `json:"alter-egos"`
	Aliases         []string `json:"aliases"`
	PlaceOfBirth    string   `json:"place-of-birth"`
	FirstAppearance string   `json:"first-appearance"`
	Publisher       string   `json:"publisher"`
	Alignment       string   `json:"alignment"`
}
type Appearance struct {
	Gender    string   `json:"gender"`
	Race      string   `json:"race"`
	Height    []string `json:"height"`
	Weight    []string `json:"weight"`
	EyeColor  string   `json:"eye-color"`
	HairColor string   `json:"hair-color"`
}
type Work struct {
	Occupation string `json:"occupation"`
	Base       string `json:"base"`
}
type Connections struct {
	GroupAffiliation string `json:"group-affiliation"`
	Relatives        string `json:"relatives"`
}
type Image struct {
	URL string `json:"url"`
}

var (
	ambiente   string
	session    *mongo.Client
	collection *mongo.Collection
	err        error
	MgoDb      = "zerohero"
	CollHeros  = "heros"

	user       = "root"
	senha      = "senha123"
	mgoUri     = "localhost:27017"
	mgoSrv     = "mongodb"
	mgoOptions = "authSource=admin&readPreference=primary&appname=MongoDB%20Compass&ssl=false"

	//mgoOptions = "retryWrites=true&w=majority"
	//mgoUriDocker = "mongodb.local.com:27017"

	connectStr = mgoSrv + "://" + user + ":" + senha + "@" + mgoUri + "/" + MgoDb + "?" + mgoOptions
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	if ambiente == "cloudrun" {
		user = os.Getenv("MGO_USER")
		senha = os.Getenv("MGO_PASSWORD")
		mgoUri = os.Getenv("MGO_HOST")
		mgoSrv = os.Getenv("MGO_SRV")
		mgoOptions = os.Getenv("MGO_OPTS")
		connectStr = mgoSrv + "://" + user + ":" + senha + "@" + mgoUri + "/" + MgoDb + "?" + mgoOptions
	}
	session, err = mongo.NewClient(options.Client().ApplyURI(connectStr))
	if err != nil {
		log.Println("error connect:", err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*10))
	defer cancel()

	err := session.Connect(ctx)
	if err != nil {
		log.Println("Error client.Connect:", err)
		return
	}
}

// InsertOne criar doc em mongodb
// criar o index para deixar unique o campo no banco
// db.heros.createIndex( { "name": 1 }, { unique: true } )
func (zh ZeroHero) InsertOne(collname string) (err error) {
	collection = session.Database(MgoDb).Collection(collname)

	ctx, cancel2 := context.WithTimeout(context.Background(), time.Duration(time.Second*6))
	defer cancel2()

	zh.UUID = uuid.New().String()
	zh.Name = strings.ToLower(zh.Name)
	zh.Name = strings.ReplaceAll(zh.Name, " ", "")

	result, err := collection.InsertOne(ctx, zh, options.InsertOne())
	if err != nil {
		return
	}

	if len(result.InsertedID.(string)) > 0 {
		return
	}
	return errors.New("Error ao tentar inserir no mongoDB")
}

// FindOne responsavel por buscar nosso do heros
func FindOne(name, fatia string, collname string) (mzh interface{}, err error) {
	mzh = nil
	collection = session.Database(MgoDb).Collection(collname)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*6))
	defer cancel()
	name = strings.ToLower(name)
	var zh ZeroHero
	err = collection.FindOne(ctx, bson.M{"name": name}).Decode(&zh)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return
		}
		return
	}

	switch fatia {
	case "image":
		mzh = zh.Image
	case "powerstats":
		mzh = zh.Powerstats
	case "biography":
		mzh = zh.Biography
	case "appearance":
		mzh = zh.Appearance
	case "work":
		mzh = zh.Work
	case "connections":
		mzh = zh.Connections
	default:
		mzh = zh
	}
	return
}

// DeleteOne responsavel por deletar nosso do heros
func DeleteOne(name string, collname string) (err error) {
	collection = session.Database(MgoDb).Collection(collname)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*6))
	defer cancel()
	name = strings.ToLower(name)
	res, err := collection.DeleteOne(ctx, bson.M{"name": name})
	if err != nil {
		return
	}

	if res.DeletedCount > 0 {
		return
	}
	return
}

// UpdateOne responsavel por atualizar nosso do heros
func (zh ZeroHero) UpdateOne(name string, collname string) (err error) {
	collection = session.Database(MgoDb).Collection(collname)
	ctx, cancel2 := context.WithTimeout(context.Background(), time.Duration(time.Second*6))
	defer cancel2()

	name = strings.ToLower(name)

	var zht ZeroHero
	err = collection.FindOne(
		ctx,
		bson.M{"name": bson.M{"$eq": name}},
	).Decode(&zht)

	if err != nil {
		return err
	}

	zh.UUID = zht.UUID
	zh.Name = zht.Name
	filter := bson.M{"_id": zh.UUID}
	update := bson.M{"$set": zh}
	res, err := collection.UpdateOne(ctx, filter, update, options.Update().SetUpsert(true))
	if err != nil {
		return
	}

	if res.MatchedCount == 0 {
		err = errors.New("não existe o registro para atualização")
		return
	}
	return
}
