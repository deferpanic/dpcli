package main

import (
	"encoding/json"
	"github.com/deferpanic/dpcli/middleware"
	"io/ioutil"
	"net"
	"net/http"
	"testing"
)

func TestProjectsNew(t *testing.T) {

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}

		var i Image
		err = json.Unmarshal(body, &i)
		if err != nil {
			t.Error(err)
		}

		if i.Name != "bob" {
			t.Error("wrong project name")
		}

		if i.Source != "https://github.com/vsukhin/phprump" {
			t.Error("wrong project name")
		}

	})

	// set listener
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Error("http not listening")
	}

	APIBase = "http://" + l.Addr().String()

	go http.Serve(l, mux)

	cli = middleware.NewCLIImplementation("token")

	projects := &Projects{}

	projects.New("bob", "php", "https://github.com/vsukhin/phprump", "")

}
