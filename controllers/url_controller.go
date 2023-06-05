package controllers

import (
	"context"

	"github.com/chamanbravo/tinyurl/redis"
	"github.com/chamanbravo/tinyurl/utils"
	"github.com/gofiber/fiber/v2"
)

type URL struct {
	Url string
}

func ShortenUrl(c *fiber.Ctx) error {
	url := new(URL)
	if err := c.BodyParser(url); err != nil {
		return err
	}

	connRedis := redis.RedisConnection()
	var shortUrl string
	for {
		shortUrl = utils.UrlGenerator(4)
		val, _ := connRedis.Get(context.Background(), shortUrl).Result()
		if val == "" {
			break
		}
	}

	err := connRedis.Set(context.Background(), shortUrl, url.Url, 0).Err()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"err":   true,
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{"tinyUrl": shortUrl})
}

func Resolve(c *fiber.Ctx) error {
	shortUrl := c.Params("url")
	connRedis := redis.RedisConnection()
	url, _ := connRedis.Get(context.Background(), shortUrl).Result()
	return c.Redirect(url)
}
