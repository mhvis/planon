// Package jsonrpc provides helpers for encoding and decoding JSON-RPC 1.0 messages.
package jsonrpc

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// rpcRequest is the format for JSON RPC requests
type rpcRequest struct {
	Id     int         `json:"id"`
	Method string      `json:"method"`
	Params interface{} `json:"params"`
}

// rpcResponse is the format for JSON RPC responses
type rpcResponse struct {
	Id     int         `json:"id"`
	Error  interface{} `json:"error"`
	Result interface{} `json:"result"`
}

// Decode decodes the given JSON object and returns the RPC properties.
func Decode(data []byte, result, error interface{}) (id int, err error) {
	resp := rpcResponse{
		Error:  error,
		Result: result,
	}
	err = json.Unmarshal(data, &resp)
	id = resp.Id
	return
}

// DecodeResponse decodes the JSON RPC object in the response.
func DecodeResponse(resp *http.Response, result, error interface{}) (id int, err error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	// Todo better logging
	log.Println("received:", string(body))
	return Decode(body, result, error)
}

// Encode encodes given arguments to a JSON object.
func Encode(method string, params interface{}, id int) ([]byte, error) {
	req := rpcRequest{id, method, params}
	return json.Marshal(req)
}

// EncodeRequest creates an HTTP request using given parameters.
func EncodeRequest(url, method string, params interface{}, id int) (*http.Request, error) {
	data, err := Encode(method, params, id)
	if err != nil {
		return nil, err
	}
	// Todo: better logging
	log.Println("encoded:", string(data))

	req, err := http.NewRequest("POST", url, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	return req, nil
}
