package routes

import (
	"time"

	"github.com/abhiraj-ku/url_shot_go/database"
	"github.com/abhiraj-ku/url_shot_go/helpers"
	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
)

type request struct {
	URL         string `json:"url"`
	CustomShort string `json:"short"`
	Expiry      string `json:"expiry"`
}
type response struct {
	URL             string        `json:"url"`
	CustomShort     string        `json:"short"`
	Expiry          time.Duration `json:"expiry"`
	XRateRemaining  string        `json:"rate_limit"`
	XRateLimitReset time.Duration `json:"rate_limit_reset"`
}

func ShortenURL(c *fiber.Ctx) error {
	// check for incoming body request
	body := new(request)

	// parse the body to json
	err := c.BodyParser(&body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	// implement the rate limiting
	r2 := database.CreateClient(1)
	defer r2.Close()

	// check for the input is actual url or not

	if !govalidator.IsURL(body.URL) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid URL",
		})
	}

	// check for domain error so that they don't enter the infinite loop
	if !helpers.RemoveDomainError(body.URL) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Bad Url Request",
		})
	}

	// enforce strict Https rule
	body.URL = helpers.EnforceHTTP(body.URL)

}
