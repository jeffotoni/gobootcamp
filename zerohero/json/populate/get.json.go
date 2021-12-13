package main

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	secret := os.Getenv("SUPER_HERO_API_SECRET")
	if secret == "" {
		panic("no secret found")
	}
	var result map[string]interface{}
	var err error
	resp := &http.Response{}

	for i := 585; i < 732; i++ {

		resp, err = http.Get("https://superheroapi.com/api/" + secret + "/" + strconv.Itoa(i))
		if err != nil {
			fmt.Println("not possible to get json for id ", i, " err ", err)
			return
		}
		err = json.NewDecoder(resp.Body).Decode(&result)
		if err != nil {
			fmt.Println("error  while getting decoding json ", err)
			return
		}
		fN := result["name"]
		fileName := fN.(string)
		fileName = strings.ToLower(fileName)
		fileName = strings.ReplaceAll(fileName, "-", "_")
		fileName = strings.ReplaceAll(fileName, " ", "_")
		toSave, err := json.MarshalIndent(result, "", "       ")
		if err != nil {
			fmt.Println(err)
			return
		}
		err = os.WriteFile(fileName+".json", toSave, fs.ModePerm)
		if err != nil {
			fmt.Println(err)
			return
		}
		time.Sleep(60000)
	}

}
