package types

import "time"

type Hash map[string]string

type Request struct {
	URL         string        `json:"url"`
	CustomShort string        `json:"custom-short"`
	Expiry      time.Duration `json:"expiry"`
}

type Response struct {
	URL             string        `json:"url"`
	CustomShort     string        `json:"custom-short"`
	Expiry          time.Duration `json:"expiry"`
	XRateRemaining  int           `json:"rate-limit"`
	XRateLimitReset time.Duration `json:"rate-limit-rest"`
}
