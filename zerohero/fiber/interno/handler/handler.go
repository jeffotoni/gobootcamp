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

func Put(c *fiber.Ctx) (err error) {
	c.Set("Content-Type", "application/json")
	var Response Response
	if len(c.Body()) == 0 {
		return c.Status(400).SendString(`{"msg":"error Body required"}`)
	}
	var zh mgo.ZeroHero
	err = c.BodyParser(&zh)
	if err != nil {
		Response.Msg = err.Error()
		return c.Status(400).JSON(Response)
	}

	name := c.Params("name")
	err = zh.UpdateOne(name, mgo.CollHeros)
	if err != nil {
		Response.Msg = err.Error()
		return c.Status(400).JSON(Response)
	}
	return c.Status(200).SendString(``)
}

func Delete(c *fiber.Ctx) (err error) {
	c.Set("Content-Type", "application/json")
	var Response Response
	name := c.Params("name")
	err = mgo.DeleteOne(name, mgo.CollHeros)
	if err != nil {
		Response.Msg = err.Error()
		return c.Status(400).JSON(Response)
	}
	return c.Status(204).SendString("")
}

func Get(c *fiber.Ctx) (err error) {
	c.Set("Content-Type", "application/json")
	var Response Response
	name := c.Params("name")
	fatia := c.Params("fatia")

	hero, err := mgo.FindOne(name, fatia, mgo.CollHeros)
	if err != nil {
		Response.Msg = err.Error()
		return c.Status(400).JSON(Response)
	}
	return c.Status(200).JSON(&hero)
}
