package routes

import "time"

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
