package tests

import (
	"testing"

	"github.com/combo23/datadome_generator/internal/datadome"
)

func TestPayload(t *testing.T) {
	t.Log("Testing payload")
	task := datadome.DataDomePayload{UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36"}
	task.JsType = "ch"
	payload, err := task.GetPayload()
	if err != nil {
		t.Error(err)
	}
	t.Log(payload)
	task.JsType = "le"
	payload, err = task.GetPayload()
	if err != nil {
		t.Error(err)
	}
	t.Log(payload)
}
