package main

import (
    "bytes"
    "context"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"
    "sync"
    "time"
    // mgo "github.com/jeffotoni/gobootcamp/zerohero/core/mongo"
)

var dir string = "../../json"
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
        case res := <-chann:
            fmt.Println("hero:", res)
        case _ = <-readDone:
            close(chann)
            end = time.Since(start).String()
            break readloop
        }
    }

    println("done:", end)
}

func isDir(path string) (b bool) {
    fileInfo, err := os.Stat(path)
    if err != nil {
        // error handling
        return
    }
    if fileInfo.IsDir() {
        // is a directory
        return true
    } else {
        return
        // is not a directory
    }
}

func Process(file string, chann chan string, wg *sync.WaitGroup) {
    defer wg.Done()
    if file != "populate" {
        time.Sleep(time.Millisecond * 300)
        namefile := dir + "/" + file
        dat, err := os.ReadFile(namefile)
        if err != nil {
            log.Println("error:", err)
            return
        }
        insertOne(file, dat)
        // InsertOnePkg(file, dat)
        // ZeroHeroGet(file)
        // ZeroHeroGetPkg(file)
        chann <- file
    }
}

func insertOne(name string, dat []byte) {
    ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second)*30)
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
        //log.Println("error status diferente 200:", resp.StatusCode)
        log.Println("error status diferente 200:", resp.StatusCode)
        return
    }
    // println("==> sucesso: [status]:", resp.StatusCode, " heroi:", name)
}

func InsertOnePkg(file string, dat []byte) {
    var zh ZeroHero
    err := json.Unmarshal(dat, &zh)
    if err != nil {
        log.Println(err)
        return
    }
    err = zh.InsertOne(c)
    if err != nil {
        log.Println(err)
        return
    }
}

func ZeroHeroGetPkg(name string) {
    i, err := FindOne("hulk", "", c)
    fmt.Println(err)
    fmt.Println(i)
}

func ZeroHeroGet(name string) {
    tr := &http.Transport{
        MaxIdleConns:       10,
        IdleConnTimeout:    30 * time.Second,
        DisableCompression: true,
    }
    client := &http.Client{Transport: tr}
    resp, err := client.Get("https://zerohero.s3apis.com/api/loki/work")
    if err != nil {
        log.Println("error:", err.Error())
        return
    }

    b, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Printf("Error %s", err.Error())
        return
    }
    println(string(b))
}
