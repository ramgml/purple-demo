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

func CreateBin(filename *string, name *string) (*bins.Bin, error) {
	postBody, err := file.ReadFile(*filename)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	request, err := makeRequest("POST", "/b", name, &postBody)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	response, err := client.Do(request)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err.Error()) 
		return nil, err
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
		return nil, err
	}
	return &bin, err
}

func UpdateBin(filename *string, id *string) (*map[string]json.RawMessage, error) {
	postBody, err := file.ReadFile(*filename)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	route := fmt.Sprintf("/b/%s", *id)
	request, err := makeRequest("PUT", route, nil, &postBody)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	response, err := client.Do(request)
	if err != nil {
		return nil, err 
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err.Error()) 
		return nil, err
	}

	var rawData map[string]json.RawMessage
	err = json.Unmarshal(body, &rawData)
	return &rawData, err
}

func DeleteBin(id *string) (*map[string]json.RawMessage, error) {
	route := fmt.Sprintf("/b/%s", *id)
	request, err := makeRequest("DELETE", route, nil, nil)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var binList bins.BinList
	dataStorage := storage.NewStorage(&binList)
	dataStorage.Load()
	var newBinList bins.BinList
	for _, bin := range dataStorage.Data.Bins {
		if bin.Id != *id {
			binList.Bins = append(newBinList.Bins, bin)
		}
	}
	dataStorage.Data.Bins = newBinList.Bins
	err = dataStorage.Save()
	if err != nil {
		return nil, err

	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err.Error()) 
		return nil, err
	}

	var rawData map[string]json.RawMessage
	err = json.Unmarshal(body, &rawData)
	return &rawData, err
}

func GetBin(id *string) (*map[string]json.RawMessage, error) {
	route := fmt.Sprintf("/b/%s", *id)
	request, err := makeRequest("GET", route, nil, nil)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("ERROR:", response)
		return nil, err
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(body) 
		return nil, err
	}
	var rawData map[string]json.RawMessage
	err = json.Unmarshal(body, &rawData)
	if err != nil {
		return nil, err
	}
	return &rawData, nil
}

func ListBins() (*bins.BinList, error) {
	var bins bins.BinList
	storage := storage.NewStorage(&bins)
	err := storage.Load()
	if err != nil {
		return nil, err
	}
	return &storage.Data, err
}
