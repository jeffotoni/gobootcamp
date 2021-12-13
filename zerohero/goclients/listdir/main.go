package main

import (
    "bytes"
    "context"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"
    "time"
)

func main() {
    var dir string = "../../json"
    files, err := ioutil.ReadDir(dir)
    if err != nil {
        log.Fatal(err)
    }

    for _, f := range files {
        fmt.Println(f.Name())
        namefile := dir + "/" + f.Name()
        dat, err := os.ReadFile(namefile)
        if err != nil {
            log.Println("error:", err)
            continue
        }
        insertOne(f.Name(), dat)
    }
}

func insertOne(name string, dat []byte) {
    ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second)*5)
    defer cancel()

    req, err := http.NewRequestWithContext(ctx, "POST", "http://localhost:8080/api", bytes.NewBuffer(dat))
    if err != nil {
        log.Println("error:", err.Error())
        return
    }

    req.Header.Set("Content-Type", "application/json")
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
    println("Criado com sucesso: [status]:", resp.StatusCode, " heroi:", name)
}
