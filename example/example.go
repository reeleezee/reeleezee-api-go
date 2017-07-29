/*
Reeleezee API sample program

Licensed under MIT license
(c) 2017 Reeleezee BV
*/
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"reeleezee-api-go/misc"
	"reeleezee-api-go/net"
)

// --------------------------------------------------------------------
// Config, external configuration in settings.json with default values
// --------------------------------------------------------------------
type Config struct {
	Uri      string
	Username string
	Password string
}

var config = Config{
	Uri:      "https://portal.reeleezee.nl/api/v1",
	Username: "username",
	Password: "password",
}

// --------------------------------------------------------------------
// init, read json configuration
// --------------------------------------------------------------------
func init() {
	// configFile := "./settings.json"
	// fixed path as vscode may run program from temp directory
	configFile := filepath.Join(os.Getenv("GOPATH"), "src/reeleezee-api-go/example/settings.json")
	if _, err := os.Stat(configFile); !os.IsNotExist(err) {
		file, _ := os.Open(configFile)
		defer file.Close()
		decoder := json.NewDecoder(file)
		err := decoder.Decode(&config)
		if err != nil {
			fmt.Println("error:", err)
			os.Exit(1)
		}
	}
}

func main() {
	fmt.Println("------------------------------ GetUserInfo ------------------------------")
	GetUserInfo()
	fmt.Println("\n------------------------------ GetProducts ------------------------------")
	GetProducts()
	fmt.Println("\n------------------------------ PutProduct ------------------------------")
	PutProduct()
}

// --------------------------------------------------------------------
// Products
// --------------------------------------------------------------------
func GetProducts() {
	route := "/Products"
	client := new(net.ApiClient)
	client.Init(config.Uri, config.Username, config.Password)

	for route != "" {
		status, data := client.Get(route)
		if status == 200 && client.IsJSON {
			var j misc.DynamicJsonValues
			json.Unmarshal(data, &j)
			for _, product := range j.Value {
				fmt.Printf("%-38s %-40s %-20s %.2f\n", product["id"],
					misc.MaxString(product["Description"], 40),
					misc.MaxString(product["SearchName"], 20), product["Price"])
			}
			route = j.GetNextLink()
		} else {
			fmt.Println("status:", status)
			route = ""
		}
	}
}

func PutProduct() {
	route := "/Products/"
	client := new(net.ApiClient)
	client.Init(config.Uri, config.Username, config.Password)

	product := make(map[string]interface{})
	guid := misc.PseudoUuidV4()
	product["id"] = guid
	product["Description"] = "New product from API"
	product["SearchName"] = "New product from API"
	product["Comment"] = "This product is created by the Go API client with id: " + guid
	product["Price"] = 12.52

	data, _ := json.Marshal(product)
	status, data := client.Put(route+guid, data)
	if status == 200 && client.IsJSON {
		json.Unmarshal(data, &product)
		fmt.Printf("%-38s %-40s %-20s %.2f\n", product["id"],
			misc.MaxString(product["Description"], 40),
			misc.MaxString(product["SearchName"], 20), product["Price"])
	} else {
		fmt.Println("status:", status)
	}
}

// --------------------------------------------------------------------
// UserInfo
// --------------------------------------------------------------------
func GetUserInfo() {
	route := "/UserInfo?$expand=*"
	client := new(net.ApiClient)
	client.Init(config.Uri, config.Username, config.Password)
	status, data := client.Get(route)
	if status == 200 && client.IsJSON {
		var out bytes.Buffer
		json.Indent(&out, data, "", "    ")
		fmt.Printf(out.String())
	} else {
		fmt.Println("status:", status)
	}
}
