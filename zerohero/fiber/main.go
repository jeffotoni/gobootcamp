package main

import (
  "time"

  "github.com/gofiber/fiber/v2"
  "github.com/gofiber/fiber/v2/middleware/cache"

  h "interno/handler"
  mw "interno/middleware"
)

func main() {
  app := fiber.New(
    fiber.Config{
      BodyLimit: 10 * 1024 * 1024,
    })

  mw.Logger(app)
  app.Get("/ping", h.Ping)
  app.Post("/api", h.Post)
  app.Put("/api", h.Put)
  app.Delete("/api/:name", h.Delete)

  //app.Use(cache.New())
  app.Use(cache.New(cache.Config{
    Next: func(c *fiber.Ctx) bool {
      println("cache...")
      return c.Query("refresh") == "true"
    },
    Expiration:   10 * time.Second,
    CacheControl: true,
  }))
  app.Get("/api/:name/:fatia", h.Get)
  app.Get("/api/:name", h.Get)

  app.Listen("0.0.0.0:8080")
}
