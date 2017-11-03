package jsonrpc

import (
	"io/ioutil"
	"testing"
)

func TestDecode(t *testing.T) {
	jsonResponse := `{"result": "hallo", "error": {"message": "error message"}, "id": 2}`
	var result string
	var error interface{}
	id, err := Decode([]byte(jsonResponse), &result, &error)
	if err != nil {
		t.Error(err)
	}

	// Check result
	if result != "hallo" {
		t.Error("invalid result value:", result)
	}

	// Check error object
	errorMap := error.(map[string]interface{})
	errorMessage := errorMap["message"].(string)
	if errorMessage != "error message" {
		t.Error("invalid error object:", error, errorMap, errorMessage)
	}

	// Check id
	if id != 2 {
		t.Error("invalid id:", id)
	}
}

func TestDecodeObject(t *testing.T) {
	jsonResponse := `{"result":"{\"serverversion\":\"201311.00.30.01\",\"version\":\"1.1.9.0\"}","error":null,"id":0}`
	var result string
	var error interface{}
	_, err := Decode([]byte(jsonResponse), &result, &error)
	if err != nil {
		t.Error(err)
	}

	// Check result
	expected := "{\"serverversion\":\"201311.00.30.01\",\"version\":\"1.1.9.0\"}"
	if result != expected {
		t.Error("invalid result value:", result)
	}
}

func TestEncode(t *testing.T) {
	params := struct {
		AParam       string `json:"a_param"`
		AnotherParam bool   `json:"another_param"`
	}{"aValue", false}
	data, err := Encode("test", params, 2)
	if err != nil {
		t.Error(err)
	}

	// Check resulting JSON string
	dataString := string(data)
	if dataString != `{"id":2,"method":"test","params":{"a_param":"aValue","another_param":false}}` {
		t.Error("incorrect JSON result:", dataString)
	}
}

func TestEncodeRequest(t *testing.T) {
	req, err := EncodeRequest("http://ankie.catspalace", "miauw", nil, 0)
	if err != nil {
		t.Error(err)
	}
	// Check content type
	contentType := req.Header.Get("Content-Type")
	if contentType != "application/json" {
		t.Error("incorrect Content-Type:", contentType)
	}

	// Check body (JSON string)
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		t.Error(err)
	}
	jsonBody := string(body)
	if jsonBody != `{"id":0,"method":"miauw","params":null}` {
		t.Error("incorrect JSON body:", jsonBody)
	}

	// Check url
	url := req.URL.String()
	if url != "http://ankie.catspalace" {
		t.Error("incorrect URL:", url)
	}
}
