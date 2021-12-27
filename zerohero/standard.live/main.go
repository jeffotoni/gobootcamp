// primeiro
package main

// segundo momento
import (
  "encoding/json"
  "log"
  "net/http"
  "strings"
  "time"
)

const (
  _ = 1 << (iota * 10)
  KB
  MB
)

type ZeroHero struct {
  Name        string      `json:"name"`
  ID          string      `json:"id"`
  Appearance  Appearance  `json:"appearance"`
  Biography   Biography   `json:"biography"`
  Connections Connections `json:"connections"`
  Image       Image       `json:"image"`
  Powerstats  Powerstats  `json:"powerstats"`
  Response    string      `json:"response"`
  Work        Work        `json:"work"`
}
type Appearance struct {
  EyeColor  string   `json:"eye-color"`
  Gender    string   `json:"gender"`
  HairColor string   `json:"hair-color"`
  Height    []string `json:"height"`
  Race      string   `json:"race"`
  Weight    []string `json:"weight"`
}
type Biography struct {
  Aliases         []string `json:"aliases"`
  Alignment       string   `json:"alignment"`
  AlterEgos       string   `json:"alter-egos"`
  FirstAppearance string   `json:"first-appearance"`
  FullName        string   `json:"full-name"`
  PlaceOfBirth    string   `json:"place-of-birth"`
  Publisher       string   `json:"publisher"`
}
type Connections struct {
  GroupAffiliation string `json:"group-affiliation"`
  Relatives        string `json:"relatives"`
}
type Image struct {
  URL string `json:"url"`
}
type Powerstats struct {
  Combat       string `json:"combat"`
  Durability   string `json:"durability"`
  Intelligence string `json:"intelligence"`
  Power        string `json:"power"`
  Speed        string `json:"speed"`
  Strength     string `json:"strength"`
}
type Work struct {
  Base       string `json:"base"`
  Occupation string `json:"occupation"`
}

// main - api rEST
func main() {
  // server mux
  mux := http.NewServeMux()

  // route ping
  mux.HandleFunc("/ping",
    func(w http.ResponseWriter, r *http.Request) {
      // MIDDLEWARE
      log.Println("ping [success]")

      w.WriteHeader(http.StatusOK)
      w.Write([]byte("ping😍"))
    })

  // /api routa -> PUT,POST,GET,DELETE
  mux.HandleFunc("/", Service)
  // mux.HandleFunc("/api", Service).Get()
  // mux.HandleFunc("/api", Service).Delete()
  // mux.HandleFunc("/api", Service).Put()

  s := &http.Server{
    Addr:         "0.0.0.0:8080",
    Handler:      mux,
    ReadTimeout:  10 * time.Second,
    WriteTimeout: 1 * time.Second,
    // MaxHeaderBytes: 0,
  }

  log.Println("Server Run.. 8080")
  s.ListenAndServe()

  // PERSISTIR MONGODB(BASE DA DADOS NOSQL)
}

// Service Responsavel por tratar nosos metodos
// Service(w http.ResponseWriter, r *http.Request) {
func Service(w http.ResponseWriter, r *http.Request) {
  // implementar a regra rotas
  // post
  // put
  // get
  // delete

  // api
  // api/hulk => retorna JSON
  // api/hulk/work => retorna fatia JSON

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
  case http.MethodPut:
    // api/hulk
    // api/superman
    if len(split) != 3 {
      http.NotFound(w, r)
      return
    }
    if split[1] != "api" {
      http.NotFound(w, r)
      return
    }
    Put(w, r)

  case http.MethodDelete:
    // api/hulk
    // api/superman
    if len(split) != 3 {
      http.NotFound(w, r)
      return
    }
    if split[1] != "api" {
      http.NotFound(w, r)
      return
    }
    Delete(w, r)
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
  default:
    http.NotFound(w, r)
  }
}

// Post - receber json com HERO
// Post(w http.ResponseWriter, r *http.Request) {
func Post(w http.ResponseWriter, r *http.Request) {
  defer r.Body.Close()
  // Logger() transforma em middleware
  w.Header().Set("Content-type", "application/json")

  //println("kb:", KB)
  // metodo
  if http.MethodPost != strings.ToUpper(r.Method) {
    w.WriteHeader(http.StatusBadRequest)
    w.Write([]byte(`{"msg":"error method ivalid only POST"}`))
    return
  }

  ioutil.Readfile
  // FIBER => FASTHTTP => OUTRA FORMA IMPLENTAR != net/http
  // gin => net/http

  // JSON
  // conteudo do JSON (payload) => r.Body
  // struct, map, interface
  // var m map[string]interface{}
  //println(Red)
  // limit no JSON Body...
  r.Body = http.MaxBytesReader(w, r.Body, MB)
  var zh ZeroHero

  err := json.NewDecoder(r.Body).Decode(&zh)
  if err != nil {
    log.Println("log:", err) // retorna no meu server
    w.WriteHeader(http.StatusBadRequest)
    w.Write([]byte(`{"msg":"Error=>` + err.Error() + `"}`)) // para o client
    return
  }

  if len(zh.ID) == 0 {
    w.WriteHeader(http.StatusBadRequest)
    w.Write([]byte(`{"msg":"field ID required!"}`))
    return
  }

  if len(zh.Name) == 0 {
    w.WriteHeader(http.StatusBadRequest)
    w.Write([]byte(`{"msg":"field name required!"}`))
    return
  }

  b, err := json.Marshal(&zh)
  if err != nil {
    w.WriteHeader(http.StatusBadRequest)
    w.Write([]byte(`{"msg":"field name required!"}`))
    return
  }

  w.WriteHeader(http.StatusOK)
  w.Write(b)
}

func Put(w http.ResponseWriter, r *http.Request) {
  // Logger() transforma em middleware
  w.WriteHeader(http.StatusOK)
  w.Write([]byte("put"))
  return
}

func Delete(w http.ResponseWriter, r *http.Request) {
  // Logger() transforma em middleware
  w.WriteHeader(http.StatusOK)
  w.Write([]byte("Delete"))
  return
}

func Get(w http.ResponseWriter, r *http.Request) {
  // Logger() transforma em middleware
  // JSON content-type ==
  // BANCO PARA LER E RETORNAR O JSON
  w.WriteHeader(http.StatusOK)
  w.Write([]byte("Get"))
  return
}

func Logger() {
  log.Println("ping [success]")
}
