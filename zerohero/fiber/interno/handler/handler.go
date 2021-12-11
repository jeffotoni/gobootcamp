package handler

import (
	mgo "interno/mongo"

	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Msg string `json:"msg"`
}

func Ping(c *fiber.Ctx) (err error) {
	c.Set("Content-Type", "application/json")
	return c.Status(200).SendString(`{"msg":"ping😍"}`)
}

func Post(c *fiber.Ctx) (err error) {
	c.Set("Content-Type", "application/json")
	var Response Response
	if len(c.Body()) == 0 {
		return c.Status(400).SendString(`{"msg":"error body empty"}`)
	}

	var zh mgo.ZeroHero
	err = c.BodyParser(&zh)
	if err != nil {
		Response.Msg = err.Error()
		return c.Status(400).JSON(Response)
	}

	if len(zh.ID) == 0 {
		return c.Status(400).SendString(`{"msg":"error ID required"}`)
	}
	if len(zh.Response) == 0 {
		return c.Status(400).SendString(`{"msg":"error Response required"}`)
	}

	err = zh.InsertOne(mgo.CollHeros)
	if err != nil {
		Response.Msg = err.Error()
		return c.Status(400).JSON(Response)
	}

	return c.Status(201).JSON(zh)
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
