package planon

// Planon low-level RPC.

import (
	"encoding/json"
	"errors"
	"github.com/mhvis/planon/jsonrpc"
	"net/http"
	"net/url"
	"strings"
)

// Call calls the given method with parameters on Planon.
func (p *Planon) Call(endpoint, method string, params map[string]interface{}, result interface{}) error {
	// Construct request
	paramsEncoded := jsonEncodeParams(params)
	req, err := jsonrpc.JSONRPCEncodeRequest(p.jsonRpcUrl+endpoint, method, paramsEncoded, p.id)
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
		resp, err = p.client.PostForm(p.jSecurityCheckUrl, login)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		// Check if response type is now JSON, i.e. authentication succeeded (hopefully)
		if contentType(resp) != "application/json" {
			return errors.New("Planon authentication failed")
		}
	default:
		// Unknown content type, throw error
		return errors.New("got unknown response")
	}

	// Decode JSON-RPC response
	var resultEncoded string
	var error interface{}
	id, err := jsonrpc.JSONRPCDecodeResponse(resp, &resultEncoded, &error)
	if err != nil {
		return err
	}
	// Check RPC response message
	if error != nil {
		// Todo return more useful error message
		return errors.New("rpc error")
	}
	if p.id != id {
		return errors.New("rpc response id is not equal to request id")
	}
	// Decode result
	return json.Unmarshal([]byte(resultEncoded), result)
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
