package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strings"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/GuilhermeCaruso/bellt"
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

	// user       = os.Getenv("MGO_USER")
	// senha      = os.Getenv("MGO_PASSWORD")
	// mgoUri     = os.Getenv("MGO_HOST")
	// mgoSrv     = os.Getenv("MGO_SRV")
	// mgoOptions = os.Getenv("MGO_OPTS")

	//mgoOptions = "retryWrites=true&w=majority"
	//mgoUriDocker = "mongodb.local.com:27017"

	connectStr = mgoSrv + "://" + user + ":" + senha + "@" + mgoUri + "/" + MgoDb + "?" + mgoOptions
)

func init() {
	// capturando ambiente atraves da compilacao
	// ela ira fazer com que nosso servico comunique com
	// mongo dentro do container
	session, err = mongo.NewClient(options.Client().ApplyURI(connectStr))
	if err != nil {
		log.Println("error connect:", err)
		println("Configura as variaveis.....")
		println("MGO_USER=root")
		println("MGO_PASSWORD=senha123")
		println("MGO_HOST=localhost:27017")
		println("...........................")
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

// JWT
// CORS
// LOGGER
// RATE LIMIT
// CACHE
// INSTRUMENTACAO
// BANCO DE DADOS
// GERAÇÃO DE LOGS SAÍDA

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	b := bellt.NewRouter()
	b.HandleFunc("/api/{name}", Use(Get, Logger()), "GET")
	b.HandleFunc("/api/{name}/{fatia}", Get, "GET")
	b.HandleFunc("/api", Post, "POST")
	//b.HandleFunc("/api", Put, "PUT")
	//b.HandleFunc("/api", Delete, "DELETE")
	log.Println("\033[1;44mRunning on http://0.0.0.0:8080 (Press CTRL+C to quit)\033[0m")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}

type Middleware func(http.HandlerFunc) http.HandlerFunc

func Logger() Middleware {

	// Create a new Middleware
	return func(f http.HandlerFunc) http.HandlerFunc {

		// Define the http.HandlerFunc
		return func(w http.ResponseWriter, r *http.Request) {

			// Do middleware things
			start := time.Now()
			defer func() {
				log.Printf(
					"\033[5m%s \033[0;103m%s\033[0m \033[0;93m%s\033[0m \033[1;44m%s \033[0m%s \033[1;44m%s\033[0m",
					r.Method,
					r.RequestURI,
					r.RemoteAddr,
					r.Header.Get("User-Agent"),
					r.URL.Path,
					time.Since(start))
			}()
			// Call the next middleware/handler in chain
			f(w, r)
		}
	}
}

// Use applies middlewares to a http.HandlerFunc
func Use(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}

func Post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	defer r.Body.Close()
	if http.MethodPost != strings.ToUpper(r.Method) {
		w.WriteHeader(http.StatusMethodNotAllowed)
		jsonstr := `{"msg":"O método permitido é Post!"}`
		w.Write([]byte(jsonstr))
		return
	}

	var zh ZeroHero
	err = json.NewDecoder(r.Body).Decode(&zh)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		jsonstr := `{"msg":"` + err.Error() + `"}`
		w.Write([]byte(jsonstr))
		return
	}

	err = zh.InsertOne(CollHeros)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		jsonstr := `{"msg":"` + err.Error() + `"}`
		w.Write([]byte(jsonstr))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(""))
}

func Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	rv := bellt.RouteVariables(r)
	name := fmt.Sprintf("%s", rv.GetVar("name"))
	fatia := fmt.Sprintf("%s", rv.GetVar("fatia"))

	if http.MethodGet != strings.ToUpper(r.Method) {
		w.WriteHeader(http.StatusMethodNotAllowed)
		jsonstr := `{"msg":"O método permitido é Get!"}`
		w.Write([]byte(jsonstr))
		return
	}

	hero, err := FindOne(name, fatia, CollHeros)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		jsonstr := `{"msg":"` + err.Error() + `"}`
		w.Write([]byte(jsonstr))
		return
	}

	b, err := json.Marshal(&hero)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		jsonstr := `{"msg":"` + err.Error() + `"}`
		w.Write([]byte(jsonstr))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if http.MethodDelete != strings.ToUpper(r.Method) {
		w.WriteHeader(http.StatusMethodNotAllowed)
		jsonstr := `{"msg":"O método permitido é Delete!"}`
		w.Write([]byte(jsonstr))
		return
	}

	rup := r.URL.Path
	lastInd := strings.LastIndex(rup, "/")
	name := rup[lastInd+1:]

	err = DeleteOne(name, CollHeros)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		jsonstr := `{"msg":"` + err.Error() + `"}`
		w.Write([]byte(jsonstr))
		return
	}

	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte(""))
}

func Put(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	defer r.Body.Close()
	if http.MethodPut != strings.ToUpper(r.Method) {
		w.WriteHeader(http.StatusMethodNotAllowed)
		jsonstr := `{"msg":"O método permitido é Put!"}`
		w.Write([]byte(jsonstr))
		return
	}

	var zh ZeroHero
	err = json.NewDecoder(r.Body).Decode(&zh)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		jsonstr := `{"msg":"` + err.Error() + `"}`
		w.Write([]byte(jsonstr))
		return
	}

	rup := r.URL.Path
	lastInd := strings.LastIndex(rup, "/")
	name := rup[lastInd+1:]

	err = zh.UpdateOne(name, CollHeros)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		jsonstr := `{"msg":"` + err.Error() + `"}`
		w.Write([]byte(jsonstr))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(""))
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
