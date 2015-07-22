package spito

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"strconv"
)

const (
	SPITO_API_ADD  string = "http://spi.to/api/v1/spits"
	SPITO_API_VIEW string = "http://spi.to/api/v1/spits/"
)

// delegates the shortening to the corresponding method according to the
// 'multipart' flag which specifies if we want to do a request with a Multipart
// content-type or a URLEncoded content-type
func Spit(content string, spitType string, exp uint64, multipart bool) (string, error) {
	if multipart {
		return ShortenMultipart(content, spitType, exp)
	} else {
		return ShortenURLEnc(content, spitType, exp)
	}
}

func _handleResponse(resp *http.Response) (string, error) {
	if resp.StatusCode != 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		return "", errors.New(resp.Status + " :: " + string(body))
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New(err.Error() + " :: " + string(body))
	}

	var f interface{}
	err = json.Unmarshal(body, &f)
	if err != nil {
		return "", err
	}
	// fmt.Println(f)
	urlHash := f.(map[string]interface{})
	return urlHash["absolute_url"].(string), nil
}

// exp should be the expiry time in seconds - string format
// This function uses multipart FormData Content-Type
func ShortenMultipart(content string, spitType string, exp uint64) (string, error) {
	client := &http.Client{}

	var b bytes.Buffer
	w := multipart.NewWriter(&b)

	fw, err := w.CreateFormField("spit_type")
	if err != nil {
		return "", err
	}
	if _, err = fw.Write([]byte(spitType)); err != nil {
		return "", err
	}
	if fw, err = w.CreateFormField("content"); err != nil {
		return "", err
	}
	if _, err = fw.Write([]byte(content)); err != nil {
		return "", err
	}
	if fw, err = w.CreateFormField("exp"); err != nil {
		return "", err
	}
	if _, err = fw.Write([]byte(strconv.FormatUint(exp, 10))); err != nil {
		return "", err
	}

	w.Close()

	req, err := http.NewRequest("POST", "http://spi.to/api/v1/spits", &b)
	if err != nil {
		return "", err
	}
	// Don't forget to set the content type, this will contain the boundary.
	req.Header.Set("Content-Type", w.FormDataContentType())
	resp, err := client.Do(req)
	defer resp.Body.Close()
	return _handleResponse(resp)
}

// This method does the same job with the Shorten() function but uses URLEncoded-form Content-Type
func ShortenURLEnc(content string, spitType string, exp uint64) (string, error) {

	client := &http.Client{}

	parameters := url.Values{}
	parameters.Add("content", content)
	parameters.Add("spit_type", spitType)
	parameters.Add("exp", strconv.FormatUint(exp, 10))

	req, err := http.NewRequest("POST", SPITO_API_ADD,
		bytes.NewBufferString(parameters.Encode()))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	defer resp.Body.Close()
	return _handleResponse(resp)
}

// This method does the same job with the Shorten() function but uses URLEncoded-form Content-Type
func View(id string) (string, error) {
	resp, err := http.Get(SPITO_API_VIEW + id)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		return "", errors.New(resp.Status + " :: " + string(body))
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New(err.Error() + " :: " + string(body))
	}
	return string(body), nil
}
