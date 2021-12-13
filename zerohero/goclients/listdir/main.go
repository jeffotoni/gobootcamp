package main

import (
    "bytes"
    "context"
    "encoding/json"
    "io/ioutil"
    "log"
    "net/http"
    "os"
    "sync"
    "time"

    mgo "github.com/jeffotoni/gobootcamp/zerohero/core/mongo"
)

var dir string = "../../json"

func main() {
    start := time.Now()
    end := ""
    println("start:", start.Format("2006-01-02 15:04:05"))
    files, err := ioutil.ReadDir(dir)
    if err != nil {
        log.Fatal(err)
    }

    var wg sync.WaitGroup
    chann := make(chan string, len(files))
    readDone := make(chan bool)
    for _, f := range files {
        wg.Add(1)
        go Process(f.Name(), chann, &wg)
    }

    // for v := range chann {
    //     fmt.Println("channel:", v)
    // }

    // for i := 0; i < len(files); i++ {
    //     fmt.Println(<-chann)
    //     //time.Sleep(time.Millisecond * 50)
    // }

    go func() {
        wg.Wait()
        readDone <- true
    }()

readloop:
    for {
        select {
        case _ = <-chann:
            // fmt.Printf("result=%#v", res)
        case _ = <-readDone:
            close(chann)
            end = time.Since(start).String()
            break readloop
        }
    }

    println("done:", end)
}

func Process(file string, chann chan string, wg *sync.WaitGroup) {
    defer wg.Done()
    if file != "populate" {
        namefile := dir + "/" + file
        dat, err := os.ReadFile(namefile)
        if err != nil {
            log.Println("error:", err)
            return
        }
        insertOne(file, dat)
        //InsertOnePkg(file, dat)
        chann <- file
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
    // println("==> sucesso: [status]:", resp.StatusCode, " heroi:", name)
}

func InsertOnePkg(file string, dat []byte) {
    var zh mgo.ZeroHero
    err := json.Unmarshal(dat, &zh)
    if err != nil {
        log.Println(err)
        return
    }
    err = zh.InsertOne("heros")
    if err != nil {
        log.Println(err)
        return
    }
}
