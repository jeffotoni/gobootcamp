package main

import (
  "os"

  "github.com/gofiber/fiber/v2"
  "github.com/gofiber/fiber/v2/middleware/logger"
)

var (
  CollHeros = "heros"
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

//Logger log
func Logger(app *fiber.App) {
  //app.Use(mw.Logger("${time} ${method} ${path} - ${ip} - ${status} - ${latency}\n"))
  app.Use(logger.New(logger.Config{
    Format:     "${pid} ${time} ${method} ${path} - ${ip} - ${status} - ${latency}\n",
    TimeFormat: "02-Jan-2006 15:04:05",
    Output:     os.Stdout,
  }))
  return
}

func main() {
  app := fiber.New(fiber.Config{BodyLimit: 10 * 1024 * 1024})
  Logger(app)
  // app.Use(mw.Logger("${time} ${method} ${path} - ${ip} - ${status} - ${latency}\n"))

  app.Get("/ping", Ping)
  app.Post("/", Post)
  app.Put("/", Put)
  app.Delete("/", Delete)
  app.Get("/:name", Get)
  app.Get("/:name/:fatia", GetHero)

  app.Listen(":8081")
}

func Ping(c *fiber.Ctx) error {
  c.Set("Content-Type", "application/json")
  return c.Status(200).SendString(`{"msg":"pong"}`)
}

func Post(c *fiber.Ctx) (err error) {
  c.Set("Content-Type", "application/json")
  if len(c.Body()) > 0 {
    return c.Status(400).SendString(`{"msg":"body empty"}`)
  }

  var zh ZeroHero
  err = c.BodyParser(&zh)
  if err != nil {
    return c.Status(400).SendString(`{"msg":"error parse, json:` + err.Error() + `"}`)
  }

  // err = zh.InsertOne(CollHeros)
  // if err != nil {
  //  return c.Status(400).SendString(`{"msg":"` + err.Error() + `"}`)
  // }

  return c.Status(201).SendString("")
}

func Put(c *fiber.Ctx) error {
  c.Set("Content-Type", "application/json")
  return c.Status(200).SendString(`{"msg":"put"}`)
}

func Delete(c *fiber.Ctx) error {
  c.Set("Content-Type", "application/json")
  return c.Status(200).SendString(`{"msg":"delete"}`)
}

func Get(c *fiber.Ctx) error {
  c.Set("Content-Type", "application/json")
  println("param:", c.Params("name"))
  return c.Status(200).SendString(`{"msg":"Get"}`)
}

func GetHero(c *fiber.Ctx) error {
  c.Set("Content-Type", "application/json")
  println("param1:", c.Params("name"))
  println("param2:", c.Params("fatia"))
  return c.Status(200).SendString(`{"msg":"GetHero"}`)
}
