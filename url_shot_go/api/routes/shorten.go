package routes

import "time"

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
