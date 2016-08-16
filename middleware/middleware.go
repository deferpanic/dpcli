package middleware

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// CLIInterface is the interface for making requests to the deferpanic rumprun api
type CLIInterface interface {
	Postit(b []byte, url string) (result string, err error)
}

// CLIImplementation is the base struct for making requests to the deferpanic rumprun api
type CLIImplementation struct {
	Token string
}

const (
	// httpTooManyRequests is http status code for too many requests error
	httpTooManyRequests = 429
)

var _ CLIInterface = &CLIImplementation{}

// NewCLIImplementation instantiates and returns a new deferpanic rumprun cli
func NewCLIImplementation(token string) *CLIImplementation {
	cli := &CLIImplementation{
		Token: token,
	}

	return cli
}

// Postit posts an api request w/b body to url and sets appropriate headers
func (c *CLIImplementation) Postit(b []byte, url string) (result string, err error) {
	defer func() {
		if rec := recover(); rec != nil {
			err := fmt.Sprintf("%q", rec)
			log.Println(err)
		}
	}()

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	req.Header.Set("DP-APIToken", c.Token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return "", err
	}

	switch resp.StatusCode {
	case http.StatusBadRequest:
		log.Println(http.StatusText(resp.StatusCode))
		return string(body), errors.New("wrong using of API method")
	case http.StatusUnauthorized:
		log.Println(http.StatusText(resp.StatusCode))
		return string(body), errors.New("wrong or invalid API token")
	case http.StatusNotFound:
		log.Println(http.StatusText(resp.StatusCode))
		return string(body), errors.New("API method was not found")
	case httpTooManyRequests:
		log.Println("too many requests - you are being rate limited")
		return string(body), errors.New("too many requests - you are being rate limited")
	case http.StatusInternalServerError:
		log.Println(http.StatusText(resp.StatusCode))
		return string(body), errors.New("internal service error")
	case http.StatusServiceUnavailable:
		log.Println(http.StatusText(resp.StatusCode))
		return string(body), errors.New("service not available")
	default:
	}

	return string(body), nil
}

// GetJSON does a get request and returns json
func (c *CLIImplementation) GetJSON(url string, iface interface{}) (err error) {
	defer func() {
		if rec := recover(); rec != nil {
			err := fmt.Sprintf("%q", rec)
			log.Println(err)
		}
	}()

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("DP-APIToken", c.Token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusBadRequest:
		log.Println(http.StatusText(resp.StatusCode))
		err = errors.New("wrong using of API method")
	case http.StatusUnauthorized:
		log.Println(http.StatusText(resp.StatusCode))
		err = errors.New("wrong or invalid API token")
	case http.StatusNotFound:
		log.Println(http.StatusText(resp.StatusCode))
		err = errors.New("API method was not found")
	case httpTooManyRequests:
		log.Println("too many requests - you are being rate limited")
		err = errors.New("too many requests - you are being rate limited")
	case http.StatusInternalServerError:
		log.Println(http.StatusText(resp.StatusCode))
		err = errors.New("internal service error")
	case http.StatusServiceUnavailable:
		log.Println(http.StatusText(resp.StatusCode))
		err = errors.New("service not available")
	default:
	}

	return json.NewDecoder(resp.Body).Decode(iface)
}

// PostJSON does a POST request and returns JSON
func (c *CLIImplementation) PostJSON(b []byte, url string, iface interface{}) (err error) {
	defer func() {
		if rec := recover(); rec != nil {
			err := fmt.Sprintf("%q", rec)
			log.Println(err)
		}
	}()

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	req.Header.Set("DP-APIToken", c.Token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusBadRequest:
		log.Println(http.StatusText(resp.StatusCode))
		err = errors.New("wrong using of API method")
	case http.StatusUnauthorized:
		log.Println(http.StatusText(resp.StatusCode))
		err = errors.New("wrong or invalid API token")
	case http.StatusNotFound:
		log.Println(http.StatusText(resp.StatusCode))
		err = errors.New("API method was not found")
	case httpTooManyRequests:
		log.Println("too many requests - you are being rate limited")
		err = errors.New("too many requests - you are being rate limited")
	case http.StatusInternalServerError:
		log.Println(http.StatusText(resp.StatusCode))
		err = errors.New("internal service error")
	case http.StatusServiceUnavailable:
		log.Println(http.StatusText(resp.StatusCode))
		err = errors.New("service not available")
	default:
	}

	return json.NewDecoder(resp.Body).Decode(iface)
}
