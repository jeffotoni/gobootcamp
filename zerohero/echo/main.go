package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type Response struct {
	Msg string `json:"msg"`
}

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

func main() {
	e := echo.New()
	e.GET("/ping", Ping)
	e.POST("/", Post)
	e.PUT("/", Put)
	e.DELETE("/", Delete)
	e.GET("/:name", Get)
	e.GET("/:name/:fatia", GetHero)

	e.Logger.SetLevel(log.INFO)
	e.Start("0.0.0.0:8081")
}

func Ping(c echo.Context) (err error) {
	return c.JSON(http.StatusCreated, Response{
		Msg: `ping😍`,
	})

}

func Post(c echo.Context) (err error) {
	var zh ZeroHero
	if err := c.Bind(&zh); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, &zh)
}

func Put(c echo.Context) (err error) {
	return c.JSON(http.StatusCreated, Response{
		Msg: `put`,
	})
}

func Delete(c echo.Context) (err error) {
	return c.JSON(http.StatusCreated, Response{
		Msg: `delete`,
	})
}

func Get(c echo.Context) (err error) {
	return c.JSON(http.StatusCreated, Response{
		Msg: `get`,
	})
}

func GetHero(c echo.Context) (err error) {
	name := c.Param("name")
	fatia := c.Param("fatia")
	println("name:", name, " fatia:", fatia)
	return c.JSON(http.StatusCreated, Response{
		Msg: `getHero`,
	})
}
