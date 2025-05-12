package api

import (
	"3-bin/bins"
	"3-bin/config"
	"3-bin/file"
	"3-bin/storage"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const rootURL = "https://api.jsonbin.io/v3"
var client = &http.Client{}

func makeRequest(method string, route string, name *string, body *[]byte) (*http.Request, error) {
	var payload io.Reader
	if body == nil {
		payload = nil
	} else {
		payload = bytes.NewBuffer(*body)
	}
	request, err := http.NewRequest(method, rootURL + route, payload)
	setup := *config.Setup
	request.Header.Set("X-Master-Key", setup.Key)
	if method != "GET" {
		request.Header.Set("Content-Type", "application/json")
	}
	if name != nil {
		request.Header.Set("X-Bin-Name", *name)
	}
	return request, err
}

func CreateBin(filename *string, name *string) {
	postBody, err := file.ReadFile(*filename)
	if err != nil {
		fmt.Println(err.Error())
	}
	request, _ := makeRequest("POST", "/b", name, &postBody)
	response, err := client.Do(request)

	if err != nil {
		fmt.Println(err.Error())
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err.Error()) 
		return
	}
	var rawData map[string]json.RawMessage
	json.Unmarshal(body, &rawData)
	var bin bins.Bin
	json.Unmarshal(rawData["metadata"], &bin)

	var binList bins.BinList
	dataStorage := storage.NewStorage(&binList)
	dataStorage.Load()
	dataStorage.Data.AddBin(&bin)
	err = dataStorage.Save()
	if err != nil {
		fmt.Println("Не удалось сохранить файл")
	}
}

func UpdateBin(filename *string, id *string) {
	postBody, err := file.ReadFile(*filename)
	if err != nil {
		fmt.Println(err.Error())
	}
	route := fmt.Sprintf("/b/%s", *id)
	request, _ := makeRequest("PUT", route, nil, &postBody)
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err.Error()) 
		return
	}
	var rawData map[string]json.RawMessage
	json.Unmarshal(body, &rawData)
	var bin bins.Bin
	json.Unmarshal(rawData["metadata"], &bin)

	var binList bins.BinList
	dataStorage := storage.NewStorage(&binList)
	dataStorage.Load()
	dataStorage.Data.AddBin(&bin)
	err = dataStorage.Save()
	if err != nil {
		fmt.Println("Не удалось сохранить файл")
	}
}

func DeleteBin(id *string) {
	route := fmt.Sprintf("/b/%s", *id)
	request, _ := makeRequest("DELETE", route, nil, nil)
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("ERROR:", response)
	}
}

func GetBin(id *string) {
	route := fmt.Sprintf("/b/%s", *id)
	request, _ := makeRequest("GET", route, nil, nil)
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("ERROR:", response)
		return
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err.Error()) 
		return
	}
	var rawData map[string]json.RawMessage
	json.Unmarshal(body, &rawData)
	fmt.Println(string(rawData["record"]))
	fmt.Println(string(rawData["metadata"]))
	fmt.Println("---")
}

func ListBins() {
	var bins bins.BinList
	storage := storage.NewStorage(&bins)
	storage.Load()
	for _, bin := range storage.Data.Bins {
		fmt.Println("ID:", bin.Id)
		fmt.Println("Name:", bin.Name)
		fmt.Println("Private:", bin.Private)
		fmt.Println("Created at:", bin.CreatedAt)
		fmt.Println("---")
	}
}
