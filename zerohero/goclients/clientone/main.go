package main

import (
	"bytes"
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	b, err := os.ReadFile("../../json/loki.json")
	if err != nil {
		log.Println("erro readfile:", err.Error())
		return
	}
	// payloadBuf := new(bytes.Buffer)
	// err = json.NewEncoder(payloadBuf).Encode(&zerohero)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second)*5)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "POST", "http://localhost:8080/api", bytes.NewBuffer(b))
	if err != nil {
		log.Println("error:", err.Error())
		return
	}

	req.Header.Set("Content-Type", "application/json")
	//req.Header.Set("Authorization", fmts.ConcatStr("Bearer ", ""))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("error:", err.Error())
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 201 {
		log.Println("error status diferente 200:", resp.StatusCode)
		return
	}
	println("Criado com sucesso: [status]:", resp.StatusCode)

	resp, err = http.Get("http://localhost:8080/api/loki/work")
	if err != nil {
		log.Println("error:", err.Error())
		return
	}

	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error %s", err.Error())
		return
	}

	println(string(b))
	println(".............................................")
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}
	resp, err = client.Get("http://localhost:8080/api/loki")
	if err != nil {
		log.Println("error:", err.Error())
		return
	}

	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error %s", err.Error())
		return
	}
	println(string(b))
}
