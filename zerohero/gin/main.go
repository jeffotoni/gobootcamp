package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
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
	router := gin.Default()
	router.GET("/ping", Ping)
	router.POST("/api", Post)
	router.PUT("/api/:name", Put)
	router.DELETE("/api/:name", Delete)
	router.GET("/api/:name", Get)
	router.GET("/api/:name/:fatia", Get)

	router.Run("0.0.0.0:8080")
}

func Ping(c *gin.Context) {
	c.JSON(http.StatusCreated, Response{
		Msg: `ping😍`,
	})

}

func Post(c *gin.Context) {
	var zh ZeroHero
	if err := c.BindJSON(&zh); err != nil {
		log.Println(err)
		return
	}
	c.JSON(http.StatusCreated, &zh)
}

func Put(c *gin.Context) {
	c.JSON(http.StatusCreated, Response{
		Msg: `put`,
	})
}

func Delete(c *gin.Context) {
	c.JSON(http.StatusCreated, Response{
		Msg: `delete`,
	})
}

func Get(c *gin.Context) {
	name := c.Param("name")
	fatia := c.Param("fatia")
	println("name:", name, " fatia:", fatia)
	c.JSON(http.StatusCreated, Response{
		Msg: `get`,
	})
}
