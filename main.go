package main

import (
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

type shorten struct {
	Url string
}

type db struct {
	Url   string
	Short string
}

var myDb []db
var letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func generateRand(n int) string {
	b := make([]string, n)
	for i := range b {
		b[i] = string(letters[rand.Intn(len(letters))])
	}
	randString := strings.Join(b, "")
	return randString
}

func shortenUrl(c *fiber.Ctx) error {
	req := new(shorten)
	err := c.BodyParser(req)
	if err != nil {
		return err
	}
	var randString string
	for {
		randString = generateRand(4)
		var findDup string
		for i := range myDb {
			if myDb[i].Short == randString {
				findDup = myDb[i].Short
			}
		}
		if findDup != randString {
			break
		}
	}
	myDb = append(myDb, db{Url: req.Url, Short: randString})
	return c.JSON(fiber.Map{"tinyUrl": randString})
}

func resolve(c *fiber.Ctx) error {
	url := c.Params("url")
	for i := range myDb {
		if myDb[i].Short == url {
			return c.Redirect(myDb[i].Url)
		}
	}
	return c.SendString("not found")
}

func main() {
	rand.Seed(time.Now().UnixNano())
	app := fiber.New()
	app.Static("/", "./public")
	app.Post("/shorten", shortenUrl)
	app.Get("/:url", resolve)
	log.Fatal(app.Listen(":3000"))
}
