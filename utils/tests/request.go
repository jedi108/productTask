package tests

import (
	"bytes"
	"net/http"
)

type sendReq struct {
	Response *http.Response
	Body     string
}

func SendRequest(typeRequest string, uri string, jsonByte []byte) sendReq {
	ioEncryptJson := bytes.NewReader(jsonByte)

	req, err := http.NewRequest(typeRequest, uri, ioEncryptJson)
	req.Header.Set("Content-Type", "application/octet-stream")
	req.Header.Set("accept", "application/json, text/plain, */*")

	client := &http.Client{}
	resp, err := client.Do(req)

	var sendResponse = sendReq{}

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		panic(err)
	}
	sendResponse.Response = resp

	sendResponse.Body = string(buf.Bytes())
	//T.Log(sendResponse.Body)
	return sendResponse
}
