package planon

import (
	"reflect"
	"testing"
)

func TestJSONEncodeParams(t *testing.T) {
	params := map[string]interface{}{
		"bool":   true,
		"int":    5,
		"string": "string",
		"object": struct {
			Int    int    `json:"int"`
			String string `json:"string"`
		}{3, "hai"},
	}

	expected := map[string]string{
		"bool":   "true",
		"int":    "5",
		"string": "\"string\"",
		"object": "{\"int\":3,\"string\":\"hai\"}",
	}
	actual := jsonEncodeParams(params)

	if !reflect.DeepEqual(expected, actual) {
		t.Error("incorrectly encoded:", actual)
	}
}
