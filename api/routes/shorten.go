package routes

import (
	"time"

	"github.com/Satvik-Srivastava/goshort/helpers"
	"github.com/gofiber/fiber/v2"
)

/*
create a struct so that the frontend can expect what he will get, this make our code very stable
*/

// defining my request structure
type request struct {
	URL         string        `json:"url"`
	CustomShort string        `json:"custom_short"`
	Expiry      time.Duration `json:"expiry"`
}

// defining my response structure
type response struct {
	URL             string        `json:"url"`
	CustomShort     string        `json:"custom_short"`
	Expiry          time.Duration `json:"expiry"`
	XRateRemaining  int           `json:"rate_limit"` // we dont want our frontend to make unlimited number of request
	XRateLimitReset time.Duration `json:"rate_limit_reset"`
}

func ShortenURL(c *fiber.Ctx) error{
	body := new(request)
	// we need to parse the upcomming JSON request into a struct that is understood by the GO programming lang

	if err := c.BodyParser(&body); err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"cannot parse JSON"})
	}

	// implement rate limiting(we will check if that the user can only make 10 request in 30 min)
	/*
	we will decrement the credit(10) by 1 whenever the user uses our services
	*/	



	// check if the url given by the user is actually correct or not
	if !govalidator.IsURL(body.URL){
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"invalid url"})
	}

	// check for the domain error so that our program do not enter in infinte loop
	if !helpers.RemoveDomainError(body.URL){
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{"error":"can't use the service"})
	}

	//enfore https, ssl
	body.URL = helpers.EnforceHTTP(body.URL)
}