package datadome

import "math/rand"

type DataDomePayload struct {
	Host         string
	UserAgent    string
	JsType       string
	Referer      string
	ResponsePage string
	Cid          string
	IP           string //for tz data
	rand         rand.Rand
	events       eventCounters
	jsdata       jsDataCH
}

type Resolution struct {
	Width  int
	Height int
}
