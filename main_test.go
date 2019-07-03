package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"

	"encoding/json"

	handlers "github.com/karlygrcm/proxy-app/api/handlers"
	middleware "github.com/karlygrcm/proxy-app/api/middleware"
	server "github.com/karlygrcm/proxy-app/api/server"
	utils "github.com/karlygrcm/proxy-app/api/utils"
)

func init() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		utils.LoadEnv()
		app := server.SetUp()
		handlers.HandlerRedirection(app)
		wg.Done()
		server.RunServer(app)
	}(wg)
	wg.Wait()
	fmt.Println("Server running...")

}

// Queue
type Queue struct {
	Domain   string
	Weight   int
	Priority int
}

type Response struct {
	Status       int            `json:"status,omitempty"`
	Response     string         `json:"result,omitempty"`
	ResponseText []ResponseText `json:"res,omitempty"`
}

type ResponseText struct {
	Domain string
}

func mainTest(t *testing.T) {
	cases := []struct {
		Domain string
		Output string
	}{
		{Domain: "alpha", Output: `["alpha"]`},
		{Domain: "", Output: "domain error"},
	}

	valuesToCompare := &Response{}
	client := http.Client{}
	for _, singleCase := range cases {
		req, err := http.NewRequest("GET", "http://localhost:8080/ping", nil)
		assert.Nil(t, err)
		req.Header.Add("domain", singleCase.Domain)

		response, err := client.Do(req)

		bytes, err := ioutil.ReadAll(response.Body)
		assert.Nil(t, err)

		err = json.Unmarshal(bytes, valuesToCompare)

		assert.Nil(t, err)
		assert.Equal(t, singleCase.Output, valuesToCompare.Response)
	}
}

func customSortingTest(t *testing.T) {
	initialQueue := []*middleware.Queue{
		{
			Domain:   "alpha",
			Weight:   4,
			Priority: 1,
		},
		{
			Domain:   "omega",
			Weight:   2,
			Priority: 2,
		},
		{
			Domain:   "beta",
			Weight:   5,
			Priority: 5,
		},
	}
	expectedQueue := []*middleware.Queue{
		{
			Domain:   "alpha",
			Weight:   5,
			Priority: 5,
		},
		{
			Domain:   "omega",
			Weight:   4,
			Priority: 1,
		},
		{
			Domain:   "beta",
			Weight:   2,
			Priority: 2,
		},
	}
	fmt.Println("test sort")
	sorted := middleware.CustomSorting(initialQueue)
	assert.Equal(t, sorted, expectedQueue)
}
