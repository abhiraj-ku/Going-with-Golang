package routes

import (
	"github.com/abhiraj-ku/url_shot_go/database"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

func ResolveURL(c *fiber.Ctx) error {

	// get the url from the
	url := c.Params("url")

	// check the db , if the original url is matched with this
	// increment the redirect counter and redirect to original url
	// else return the error message saying url does not match

	r := database.CreateClient(0)
	defer r.Close()

	value, err := r.Get(database.Ctx, url).Result()
	if err == redis.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "short url not found in the Db",
		})
	} else if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot connect to db",
		})
	}

	// increment the counter
	rIncr := database.CreateClient(1)
	defer rIncr.Close()

	_ = rIncr.Incr(database.Ctx, "counter")

	// redirect to original url
	return c.Redirect(value, 301)

}
