package routes

type request struct {
	URL         string `json:"url"`
	CustomShort string `json:"short"`
	Expiry      string `json:"expiry"`
}
type response struct {
	URL             string `json:"url"`
	CustomShort     string `json:"short"`
	Expiry          string `json:"expiry"`
	XRateRemaining  string `json:"rate_limit"`
	XRateLimitReset string `json:"rate_limit_reset"`
}
