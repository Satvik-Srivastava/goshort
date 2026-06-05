package routes

import (
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/Satvik-Srivastava/goshort/database"
	"github.com/Satvik-Srivastava/goshort/helpers"
	"github.com/asaskevich/govalidator"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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

func ShortenURL(c *fiber.Ctx) error {
	body := new(request)
	// we need to parse the upcomming JSON request into a struct that is understood by the GO programming lang

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	// implement rate limiting(we will check if that the user can only make 10 request in 30 min)
	/*
		we will decrement the credit(10) by 1 whenever the user uses our services
	*/

	redisDBClient := database.CreateClient(1)
	defer redisDBClient.Close()
	// redis is a key-value pair databse
	val, err := redisDBClient.Get(database.Ctx, c.IP()).Result()

	if err == redis.Nil {
		// saving the data as "ip-address":"apiquota" and time left to reset
		_ = redisDBClient.Set(database.Ctx, c.IP(), os.Getenv("API_QUOTA"), 30*60*time.Second).Err()
	} else {
		val, _ = redisDBClient.Get(database.Ctx, c.IP()).Result()
		valIn, _ := strconv.Atoi(val)
		if valIn <= 0 {
			limit, _ := redisDBClient.TTL(database.Ctx, c.IP()).Result()
			return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
				"error":            "rate limit exceeded",
				"rate_limit_reset": limit / time.Nanosecond / time.Minute,
			})
		}
	}

	// check if the url given by the user is actually correct or not
	if !govalidator.IsURL(body.URL) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid url"})
	}

	// check for the domain error so that our program do not enter in infinte loop
	if !helpers.RemoveDomainError(body.URL) {
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{"error": "can't use the service"})
	}

	//enfore https, ssl
	body.URL = helpers.EnforceHTTP(body.URL)

	/*
		now we are going to implement custom url
		- we will accept the input from the user
		- check in our database that no other user is using the same url
		- if user has not sent any custom input then we need to create the shortend url from our side
	*/

	var id string
	if body.CustomShort == "" {
		id = uuid.New().String()[:6]
	} else {
		body.CustomShort = strings.TrimSpace(body.CustomShort)
		if strings.Contains(body.CustomShort, "/") {
			parts := strings.Split(strings.Trim(body.CustomShort, "/"), "/")
			body.CustomShort = parts[len(parts)-1]
		}
		id = body.CustomShort
	}

	r := database.CreateClient(0)
	defer r.Close()

	val, _ = r.Get(database.Ctx, id).Result()

	if val != "" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "URL custom short is already in use",
		})
	}

	if body.Expiry == 0 {
		body.Expiry = 24
	}

	err = r.Set(database.Ctx, id, body.URL, body.Expiry*3600*time.Second).Err()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "unable to connect to the server",
		})
	}

	domain := os.Getenv("DOMAIN")
	if !strings.HasPrefix(strings.ToLower(domain), "http://") && !strings.HasPrefix(strings.ToLower(domain), "https://") {
		domain = "http://" + domain
	}

	resp := response{
		URL:             body.URL,
		CustomShort:     "",
		Expiry:          body.Expiry,
		XRateRemaining:  10,
		XRateLimitReset: 30,
	}

	redisDBClient.Decr(database.Ctx, c.IP())
	val, _ = redisDBClient.Get(database.Ctx, c.IP()).Result()
	resp.XRateRemaining, _ = strconv.Atoi(val)

	ttl, _ := redisDBClient.TTL(database.Ctx, c.IP()).Result()
	resp.XRateLimitReset = ttl / time.Nanosecond / time.Minute
	resp.CustomShort = domain + "/" + id

	return c.Status(fiber.StatusOK).JSON(resp)
}
