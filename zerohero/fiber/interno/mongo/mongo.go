package mongo

import (
	"context"
	"errors"
	"log"
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
	ambiente     string
	session      *mongo.Client
	collection   *mongo.Collection
	err          error
	MgoDb        = "zerohero"
	CollHeros    = "heros"
	user         = "root"
	senha        = "senha123"
	mgoUri       = "127.0.0.1:27017"
	mgoUriDocker = "mongodb.local.com:27017"
	port         = "27017"
	mgoOptions   = "authSource=admin&readPreference=primary&appname=MongoDB%20Compass&ssl=false"
	connectStr   = "mongodb://" + user + ":" + senha + "@" + mgoUri + "/" + MgoDb + "?" + mgoOptions
)

func init() {
	// capturando ambiente atraves da compilacao
	// ela ira fazer com que nosso servico comunique com
	// mongo dentro do container
	if ambiente == "docker" {
		println("ambiente docker....")
		connectStr = "mongodb://" + user + ":" + senha + "@" + mgoUriDocker + "/" + MgoDb + "?" + mgoOptions
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
func FindOne(name, fatia string, collname string) (mzh map[string]interface{}, err error) {
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
		mzh1 := make(map[string]interface{}, 1)
		mzh1[fatia] = zh.Image
		mzh = mzh1
	case "powerstats":
		mzh1 := make(map[string]interface{}, 1)
		mzh1[fatia] = zh.Powerstats
		mzh = mzh1
	case "biography":
		mzh1 := make(map[string]interface{}, 1)
		mzh1[fatia] = zh.Biography
		mzh = mzh1
	case "appearance":
		mzh1 := make(map[string]interface{}, 1)
		mzh1[fatia] = zh.Appearance
		mzh = mzh1
	case "work":
		mzh1 := make(map[string]interface{}, 1)
		mzh1[fatia] = zh.Work
		mzh = mzh1
	case "connections":
		mzh1 := make(map[string]interface{}, 1)
		mzh1[fatia] = zh.Connections
		mzh = mzh1
	default:
		mzh1 := make(map[string]interface{}, 1)
		mzh1["zerohero"] = zh
		mzh = mzh1
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
