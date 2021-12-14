package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/google/uuid"
	"github.com/rs/cors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	pathNames = map[string]string{
		"image":       "image",
		"powerstats":  "powerstats",
		"biography":   "biography",
		"appearance":  "appearance",
		"work":        "work",
		"connections": "connections",
	}
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
	sc   count32
	once sync.Once

	ambiente   string
	pool       = 80
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

type Config struct {
	Srv     string
	DB      string
	Host    string
	User    string
	Pass    string
	Options string
}

type count32 int32

func (c *count32) Increment() int32 {
	return atomic.AddInt32((*int32)(c), 1)
}

func (c *count32) Get() int32 {
	return atomic.LoadInt32((*int32)(c))
}

func (c *count32) Zerar() int32 {
	return atomic.AddInt32((*int32)(c), -80)
}

func timePoolLimit() {
	sc.Increment()
	count := sc.Get()
	if int(count) == pool {
		sc.Zerar()
		time.Sleep(time.Second * 15)
	}
}

func New(srv, host, db, user, pass, options string) Config {
	return Config{
		Srv:     srv,
		Host:    host,
		DB:      db,
		User:    user,
		Pass:    pass,
		Options: options,
	}
}

func (c Config) Connect() (*mongo.Collection, context.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Minute*8))
	defer cancel()
	once.Do(func() {
		if session == nil {
			//if session == nil {
			connstr := c.Srv + "://" + c.User + ":" + c.Pass + "@" + c.Host + "/" + c.DB + "?" + c.Options
			optcli := options.Client().ApplyURI(connstr)
			//optcli.SetMaxPoolSize(80)
			session, err = mongo.NewClient(optcli)
			if err != nil {
				log.Println("error connect:", err)
				return
			}

			err = session.Connect(ctx)
			if err != nil {
				log.Println("Error client.Connect:", err)
				return
			}
			collection = session.Database(c.DB).Collection(CollHeros)
		}
	})
	return collection, ctx
}

// InsertOne criar doc em mongodb
// criar o index para deixar unique o campo no banco
// db.heros.createIndex( { "name": 1 }, { unique: true } )
func (zh ZeroHero) InsertOne(c Config) (err error) {
	collectionx, _ := c.Connect()
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Minute*10))
	defer cancel()
	timePoolLimit()

	zh.UUID = uuid.New().String()
	zh.Name = strings.ToLower(zh.Name)
	zh.Name = strings.ReplaceAll(zh.Name, " ", "")
	result, err := collectionx.InsertOne(ctx, zh, options.InsertOne())
	if err != nil {
		return
	}

	if len(result.InsertedID.(string)) > 0 {
		return
	}
	return errors.New("Error ao tentar inserir no mongoDB")
}

// FindOne responsavel por buscar nosso do heros
func FindOne(name, fatia string, c Config) (mzh interface{}, err error) {
	collectionx, _ := c.Connect()
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Minute*10))
	defer cancel()
	timePoolLimit()

	mzh = nil
	name = strings.ToLower(name)
	var zh ZeroHero
	err = collectionx.FindOne(ctx, bson.M{"name": name}).Decode(&zh)
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
func DeleteOne(name string, c Config) (err error) {
	collectionx, _ := c.Connect()
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Minute*10))
	defer cancel()
	timePoolLimit()

	name = strings.ToLower(name)
	res, err := collectionx.DeleteOne(ctx, bson.M{"name": name})
	if err != nil {
		return
	}

	if res.DeletedCount > 0 {
		return
	}
	return
}

// UpdateOne responsavel por atualizar nosso do heros
func (zh ZeroHero) UpdateOne(name string, c Config) (err error) {
	collectionx, _ := c.Connect()
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Minute*10))
	defer cancel()
	timePoolLimit()

	name = strings.ToLower(name)

	var zht ZeroHero
	err = collectionx.FindOne(
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
	res, err := collectionx.UpdateOne(ctx, filter, update, options.Update().SetUpsert(true))
	if err != nil {
		return
	}

	if res.MatchedCount == 0 {
		err = errors.New("não existe o registro para atualização")
		return
	}
	return
}

// JWT
// CORS
// LOGGER
// RATE LIMIT
// CACHE
// INSTRUMENTACAO
// BANCO DE DADOS
// GERAÇÃO DE LOGS SAÍDA
// Basic Auth
// Key Auth
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

var c Config

func init() {
	c = Config{
		Srv:     os.Getenv("MGO_SRV"),
		DB:      os.Getenv("MGO_DB"),
		Host:    os.Getenv("MGO_HOST"),
		User:    os.Getenv("MGO_USER"),
		Pass:    os.Getenv("MGO_PASSWORD"),
		Options: os.Getenv("MGO_OPTS"),
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(w http.ResponseWriter,
		r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong😍"))
	})

	//mux.HandleFunc("/", Use(Service, Logger()))
	mux.HandleFunc("/", Service)

	handler := cors.Default().Handler(mux)
	s := &http.Server{
		Addr:           "0.0.0.0:8080",
		Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20, // 1MB
	}
	log.Println("\033[1;44mRunning on http://0.0.0.0:8080 (Press CTRL+C to quit)\033[0m")
	log.Fatal(s.ListenAndServe())
}

func Service(w http.ResponseWriter, r *http.Request) {
	split := strings.Split(r.URL.Path, "/")
	switch r.Method {
	case http.MethodPost:
		if len(split) > 2 {
			http.NotFound(w, r)
			return
		}
		if split[1] != "api" {
			http.NotFound(w, r)
			return
		}
		Post(w, r)
	case http.MethodGet:
		if len(split) < 3 {
			http.NotFound(w, r)
			return
		}
		if split[1] != "api" {
			http.NotFound(w, r)
			return
		}
		Get(w, r)
	case http.MethodDelete:
		if len(split) != 3 {
			http.NotFound(w, r)
			return
		}
		if split[1] != "api" {
			http.NotFound(w, r)
			return
		}
		Delete(w, r)
	case http.MethodPut:
		if len(split) != 3 {
			http.NotFound(w, r)
			return
		}
		if split[1] != "api" {
			http.NotFound(w, r)
			return
		}
		Put(w, r)
	default:
		http.NotFound(w, r)
	}
}

const (
	_  = 1 << (iota * 10)
	Kb // decimal 1024
	Mb // decimal 104857

	Red int = iota
	Orange
	Yellow
	Green
	Blue
	Indigo
	Violet
)

func Post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	defer r.Body.Close()
	if http.MethodPost != strings.ToUpper(r.Method) {
		w.WriteHeader(http.StatusMethodNotAllowed)
		jsonstr := `{"msg":"O método permitido é Post!"}`
		w.Write([]byte(jsonstr))
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, Kb)
	var zh ZeroHero
	err = json.NewDecoder(r.Body).Decode(&zh)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		jsonstr := `{"msg":"` + err.Error() + `"}`
		w.Write([]byte(jsonstr))
		return
	}

	err = zh.InsertOne(c)
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

	if http.MethodGet != strings.ToUpper(r.Method) {
		w.WriteHeader(http.StatusMethodNotAllowed)
		jsonstr := `{"msg":"O método permitido é Get!"}`
		w.Write([]byte(jsonstr))
		return
	}

	rup := r.URL.Path
	lastInd := strings.LastIndex(rup, "/")
	name := rup[lastInd+1:]

	_, ok := pathNames[name]
	fatia := ""
	if ok {
		fatia = name
		split := strings.Split(rup, "/")
		if len(split) > 2 {
			name = split[2]
		} else {
			w.WriteHeader(http.StatusNotFound)
			return
		}
	}
	hero, err := FindOne(name, fatia, c)
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

	err = DeleteOne(name, c)
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

	err = zh.UpdateOne(name, c)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		jsonstr := `{"msg":"` + err.Error() + `"}`
		w.Write([]byte(jsonstr))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(""))
}
