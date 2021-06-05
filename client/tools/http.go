package tools

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

/*------------------------------*/

func SendJSON(resp http.ResponseWriter, obj interface{}) {

	data, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		http.Error(resp, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	resp.Header().Set("Content-Type", "application/json")
	resp.Write(data)
}

/*------------------------------*/

func GetRequest(URL string, xAPIKey string) ([]byte, error) {
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return nil, CodeError{http.StatusInternalServerError, "Something went wrong!"}
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-KEY", xAPIKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, CodeError{http.StatusInternalServerError, "Did not receive a response from Scootin Server!"}
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, CodeError{resp.StatusCode, resp.Status}
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, CodeError{http.StatusInternalServerError, "Something went wrong!"}
	}

	return body, nil
}

/*------------------------------*/

func PostRequest(URL string, postBody []byte, xAPIKey string) ([]byte, error) {

	req, err := http.NewRequest("POST", URL, bytes.NewBuffer(postBody))
	if err != nil {
		return nil, CodeError{http.StatusInternalServerError, "Something went wrong!"}
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-KEY", xAPIKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, CodeError{http.StatusInternalServerError, "Did not receive a response from Scootin Server!"}
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, CodeError{resp.StatusCode, resp.Status}
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, CodeError{http.StatusInternalServerError, "Something went wrong!"}
	}

	return body, nil
}

/*------------------------------*/

type ClosingBuffer struct {
	*bytes.Buffer
}

func (cb *ClosingBuffer) Close() error {
	return nil
}

func ReadAll(rc io.ReadCloser) ([]byte, error) {
	defer rc.Close()

	if cb, ok := rc.(*ClosingBuffer); ok {
		return cb.Bytes(), nil
	}

	return ioutil.ReadAll(rc)
}

/*------------------------------*/
