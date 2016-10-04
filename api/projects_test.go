package api

import (
	"encoding/json"
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

		if i.Compiler != "rumprun" {
			t.Error("wrong compiler name")
		}

		if i.Source != "https://github.com/deferpanic/php_example" {
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

	Cli = NewCliImplementation("token")

	projects := &Projects{}

	projects.New("bob", "php", "rumprun", "https://github.com/deferpanic/php_example", "")

}
