package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	mgo "interno/mongo"
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

func Post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	defer r.Body.Close()
	if http.MethodPost != strings.ToUpper(r.Method) {
		w.WriteHeader(http.StatusMethodNotAllowed)
		jsonstr := `{"msg":"O método permitido é Post!"}`
		w.Write([]byte(jsonstr))
		return
	}

	var zh mgo.ZeroHero
	err := json.NewDecoder(r.Body).Decode(&zh)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		jsonstr := `{"msg":"` + err.Error() + `"}`
		w.Write([]byte(jsonstr))
		return
	}

	err = zh.InsertOne(mgo.CollHeros)
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

	hero, err := mgo.FindOne(name, fatia, mgo.CollHeros)
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
	err := mgo.DeleteOne(name, mgo.CollHeros)
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

	var zh mgo.ZeroHero
	err := json.NewDecoder(r.Body).Decode(&zh)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		jsonstr := `{"msg":"` + err.Error() + `"}`
		w.Write([]byte(jsonstr))
		return
	}

	rup := r.URL.Path
	lastInd := strings.LastIndex(rup, "/")
	name := rup[lastInd+1:]

	err = zh.UpdateOne(name, mgo.CollHeros)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		jsonstr := `{"msg":"` + err.Error() + `"}`
		w.Write([]byte(jsonstr))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(""))
}
