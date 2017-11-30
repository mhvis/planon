package planonrpc

// Planon low-level RPC.

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mhvis/planon/jsonrpc"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type planonError struct {
	ID            int
	ExMessage     string `json:"exMessage"`
	DetailMessage string `json:"detailMessage"`
	Cause         struct {
		Class   string `json:"class"`
		Message string `json:"message"`
		ID      int
	} `json:"cause"`
	StackTrace           []interface{} `json:"stackTrace"`
	SuppressedExceptions []interface{} `json:"suppressedExceptions"`
}

// Call calls the given method with parameters on Planon.
func (p *Service) Call(endpoint, method string, params map[string]interface{}, result interface{}) error {
	// Construct request
	paramsEncoded := jsonEncodeParams(params)
	req, err := jsonrpc.EncodeRequest(p.twowayauthUrl+"/JSONrpc"+endpoint, method, paramsEncoded, p.id)
	if err != nil {
		return err
	}

	// Do request
	resp, err := p.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	defer func() {
		// Increment id at the end of the function
		p.id = p.id + 1
	}()

	// Check response Content-Type
	switch contentType(resp) {
	case "application/json":
		// JSON response, this is good
	case "text/html":
		// HTML response, assuming login screen, try to authenticate
		login := url.Values{"j_username": {p.jUsername}, "j_password": {p.jPassword}}
		resp, err = p.client.PostForm(p.twowayauthUrl+"/j_security_check", login)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		// Check if response type is now JSON, i.e. authentication succeeded (hopefully)
		if contentType(resp) != "application/json" {
			return unexpectedResponseError(resp)
		}
	default:
		// Unknown content type, throw error
		return unexpectedResponseError(resp)
	}

	// Decode JSON-RPC response
	var resultEncoded interface{}
	var error interface{}
	id, err := jsonrpc.DecodeResponse(resp, &resultEncoded, &error)
	if err != nil {
		return err
	}
	// Check RPC response message
	if error != nil {
		// Collect error data
		errorMap := error.(map[string]interface{})
		errorId := errorMap["ID"].(float64)
		errorMessage := errorMap["exMessage"].(string)
		return errors.New(fmt.Sprintf("%v %v", errorId, errorMessage))
	}
	if p.id != id {
		return errors.New("rpc response id is not equal to request id")
	}
	// Decode result
	if resultEncoded == nil {
		return nil
	}
	return json.Unmarshal([]byte(resultEncoded.(string)), result)
}

// contentType returns the content type.
func contentType(resp *http.Response) string {
	v := resp.Header.Get("Content-Type")
	vSplit := strings.SplitN(v, ";", 2)
	vTrimmed := strings.TrimSpace(vSplit[0])
	return vTrimmed
}

// jsonEncodeParams converts all elements in the map to JSON encoded strings.
func jsonEncodeParams(params map[string]interface{}) map[string]string {
	paramsEncoded := make(map[string]string)
	for k, v := range params {
		encoded, err := json.Marshal(v)
		if err != nil {
			panic(err)
		}
		paramsEncoded[k] = string(encoded)
	}
	return paramsEncoded
}

func unexpectedResponseError(resp *http.Response) error {
	errorString := fmt.Sprintf("unexpected response: %v", resp)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		errorString += ", could not read body: " + err.Error()
	} else {
		errorString += ", body: " + string(body)
	}
	return errors.New(errorString)

}
