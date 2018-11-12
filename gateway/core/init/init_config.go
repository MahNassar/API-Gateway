package init

import (
	"api_gateway/gateway/core"
	"io/ioutil"
	"fmt"
	"os"
	"encoding/json"
)

func ReadConfig() core.Router {
	file, e := ioutil.ReadFile("config/config.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	var root core.JsonRoot
	json.Unmarshal(file, &root)
	// init app config from .env
	return root.Router
}
