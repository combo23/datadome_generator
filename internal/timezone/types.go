package timezone

type response struct {
	TimeZone struct {
		ID        string `json:"id"`
		UTCOffset int    `json:"utcOffset"`
	} `json:"timeZone"`
}
