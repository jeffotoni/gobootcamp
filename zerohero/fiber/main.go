package main

import (
  "os"

  "github.com/gofiber/fiber/v2"
  "github.com/gofiber/fiber/v2/middleware/logger"

  h "interno/handler"
)

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
  app := fiber.New(
    fiber.Config{
      BodyLimit: 10 * 1024 * 1024,
    })
  Logger(app)
  app.Get("/ping", h.Ping)
  app.Post("/", h.Post)
  app.Put("/", h.Put)
  app.Delete("/", h.Delete)
  app.Get("/:name", h.Get)
  app.Get("/:name/:fatia", h.GetHero)

  app.Listen(":8081")
}
