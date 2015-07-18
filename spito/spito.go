package spito

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
)

const (
	SPITO_API_ADD string = "http://spi.to/api/v1/spits"
)

// delegates the shortening to the corresponding method according to the
// 'multipart' flag which specifies if we want to do a request with a Multipart
// content-type or a URLEncoded content-type
func Spitit(longUrl string, exp string, multipart bool) (string, error) {
	if multipart {
		return ShortenMultipart(longUrl, exp)
	} else {
		return ShortenURLEnc(longUrl, exp)
	}
}

// exp should be the expiry time in seconds - string format
// This function uses multipart FormData Content-Type
func ShortenMultipart(longUrl string, exp string) (string, error) {
	client := &http.Client{}

	var b bytes.Buffer
	w := multipart.NewWriter(&b)

	fw, err := w.CreateFormField("spit_type")
	if err != nil {
		return "", err
	}
	if _, err = fw.Write([]byte("url")); err != nil {
		return "", err
	}
	if fw, err = w.CreateFormField("content"); err != nil {
		return "", err
	}
	if _, err = fw.Write([]byte(longUrl)); err != nil {
		return "", err
	}
	if fw, err = w.CreateFormField("exp"); err != nil {
		return "", err
	}
	if _, err = fw.Write([]byte(exp)); err != nil {
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

// This method does the same job with the Shorten() function but uses URLEncoded-form Content-Type
func ShortenURLEnc(longUrl string, exp string) (string, error) {

	client := &http.Client{}

	parameters := url.Values{}
	parameters.Add("content", longUrl)
	parameters.Add("spit_type", "url")
	parameters.Add("exp", exp)

	req, err := http.NewRequest("POST", SPITO_API_ADD,
		bytes.NewBufferString(parameters.Encode()))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	defer resp.Body.Close()
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
	urlHash := f.(map[string]interface{})
	return urlHash["absolute_url"].(string), nil
}
