package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var _ CliInterface = &CliImplementation{}

var Cli *CliImplementation

// CliInterface is the interface for making requests to the deferpanic api
type CliInterface interface {
	Postit(b []byte, url string) (result string, err error)
}

// CliImplementation is the base struct for making requests to the deferpanic api
type CliImplementation struct {
	Token string
}

const (
	// httpTooManyRequests is http status code for too many requests error
	httpTooManyRequests = 429
)

// NewCliImplementation instantiates and returns a new deferpanic API cli
func NewCliImplementation(token string) *CliImplementation {
	cli := &CliImplementation{
		Token: token,
	}

	return cli
}

// Postit posts an api request w/b body to url and sets appropriate headers
func (c *CliImplementation) Postit(b []byte, url string) (result string, err error) {
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

	err = statusToHuman(resp.StatusCode)
	if err != nil {
		return string(body), err
	}

	return string(body), nil
}

// GetJSON does a get request and returns json
func (c *CliImplementation) GetJSON(url string, iface interface{}) (err error) {
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

	err = statusToHuman(resp.StatusCode)
	if err != nil {
		return err
	}

	return json.NewDecoder(resp.Body).Decode(iface)
}

// PostJSON does a POST request and returns JSON
func (c *CliImplementation) PostJSON(b []byte, url string, iface interface{}) (err error) {
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

	err = statusToHuman(resp.StatusCode)
	if err != nil {
		return err
	}

	return json.NewDecoder(resp.Body).Decode(iface)
}

// GrabFile downloads to location
// TODO - refactor me
func (c *CliImplementation) GrabFile(b []byte, url string, fname string) (err error) {
	defer func() {
		if rec := recover(); rec != nil {
			err := fmt.Sprintf("%q", rec)
			log.Println(err)
		}
	}()

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	req.Header.Set("DP-APIToken", c.Token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 302 {
		newLoc := resp.Header.Get("Location")
		return c.GrabRedirectFile(newLoc, fname)
	}

	sz := resp.Header.Get("Content-Length")
	fmt.Println("\033[1;32msaving download " + sz + " bytes to " + fname + "\033[0m")

	out, err := os.Create(fname)
	if err != nil {
		log.Println(err)
	}

	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.Println(err)
	}

	return statusToHuman(resp.StatusCode)
}

func (c *CliImplementation) GrabRedirectFile(url string, fname string) (err error) {
	req, err := http.NewRequest("GET", url, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	sz := resp.Header.Get("Content-Length")
	fmt.Println("\033[1;32msaving download " + sz + " bytes to " + fname + "\033[0m")

	out, err := os.Create(fname)
	if err != nil {
		log.Println(err)
	}

	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.Println(err)
	}

	return statusToHuman(resp.StatusCode)
}

func statusToHuman(status int) error {
	switch status {
	case http.StatusBadRequest:
		log.Println(http.StatusText(status))
		return errors.New("wrong using of API method")
	case http.StatusUnauthorized:
		log.Println(http.StatusText(status))
		return errors.New("wrong or invalid API token")
	case http.StatusNotFound:
		log.Println(http.StatusText(status))
		return errors.New("API method was not found")
	case httpTooManyRequests:
		log.Println("too many requests - you are being rate limited")
		return errors.New("too many requests - you are being rate limited")
	case http.StatusInternalServerError:
		log.Println(http.StatusText(status))
		return errors.New("internal service error")
	case http.StatusServiceUnavailable:
		log.Println(http.StatusText(status))
		return errors.New("service not available")
	default:
		return nil
	}
}
