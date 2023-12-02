package datadome

import (
	"encoding/json"
	"math/rand"
	"time"

	"github.com/justhyped/OrderedForm"
)

func (t *DataDomePayload) GetPayload() (string, error) {
	t.rand = *rand.New(rand.NewSource(time.Now().UnixNano()))
	form := new(OrderedForm.OrderedForm)

	if t.JsType == "le" {
		t.GetEvents()
	}
	request, err := extractPathFromURL(t.Referer)
	if err != nil {
		return "", err
	}
	jsData, err := t.GetJSData()
	if err != nil {
		return "", err
	}
	form.Set("jsData", jsData)
	if t.JsType == "ch" {
		form.Set("eventCounters", "[]")
	} else {
		events, err := json.Marshal(t.events)
		if err != nil {
			return "", err
		}
		form.Set("eventCounters", string(events))
	}
	ddk, ddversion := getDDInfo(t.Host)
	form.Set("jsType", t.JsType)
	form.Set("cid", t.Cid)
	form.Set("ddk", ddk)
	form.Set("Referer", t.Referer)
	form.Set("request", request)
	form.Set("responsePage", t.ResponsePage)
	form.Set("ddv", ddversion)
	return form.URLEncode(), nil
}
