import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"testing"

	handlers "github.com/karlygrcm/proxy-app/api/handlers"
	server "github.com/karlygrcm/proxy-app/api/server"
	utils "github.com/karlygrcm/proxy-app/api/utils"
	"github.com/stretchr/testify/assert"
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
