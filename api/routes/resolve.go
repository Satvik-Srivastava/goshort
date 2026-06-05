/*
when a user enter a url the shorten.go file create a short url for the user provided
url so we need when the user click on the shortend url it should redirect the user to the original
website for this we will store the original url in our database
*/

package routes

import (
	"github.com/Satvik-Srivastava/goshort/database"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

func ResolveURL(c *fiber.Ctx) error {
	url := c.Params("url")

	// CreateClient is the function created in the database.go file
	r := database.CreateClient(0)
	defer r.Close()

	// check in the database whether the url is present or not
	value, err := r.Get(database.Ctx, url).Result()
	if err == redis.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "short not found in the database",
		})
	}
	if err != nil { // in case there is some error in the database connection
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "cannot connect to Database",
		})
	}

	// if everything goes smooth we will redirect the user to the website
	rInr := database.CreateClient(1)
	defer rInr.Close()

	_ = rInr.Incr(database.Ctx, "counter")
	return c.Redirect(value, 301)
}
