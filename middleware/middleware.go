package middleware

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// RumpRunCLIInterface is the interface for making requests to the deferpanic rumprun api
type RumpRunCLIInterface interface {
	Postit(b []byte, url string) (result string, err error)
}

// RumpRunCLIImplementation is the base struct for making requests to the deferpanic rumprun api
type RumpRunCLIImplementation struct {
	Token string
}

const (
	// httpTooManyRequests is htto status code for too many requests error
	httpTooManyRequests = 429
)

var _ RumpRunCLIInterface = &RumpRunCLIImplementation{}

// NewRumpRunCLIImplementation instantiates and returns a new deferpanic rumprun cli
func NewRumpRunCLIImplementation(token string) *RumpRunCLIImplementation {
	cli := &RumpRunCLIImplementation{
		Token: token,
	}

	return cli
}

// Postit posts an api request w/b body to url and sets appropriate headers
func (c *RumpRunCLIImplementation) Postit(b []byte, url string) (result string, err error) {
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