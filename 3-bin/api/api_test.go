package api_test

import (
	"3-bin/api"
	"3-bin/config"
	"strings"
	"testing"
)

func setUp() {
	err := config.LoadEnv(".env.test")
	if err != nil {
		panic(err.Error())
	}
	config.Setup = config.NewConfig()
}

func TestCreateBin(t *testing.T) {
	setUp()
	filename := "./../my.json"
	name := "test-bin"
	bin, _ := api.CreateBin(&filename, &name)
	defer api.DeleteBin(&bin.Id)
	if bin.Name != name {
		t.Errorf("Expected %v, got %v", name, bin.Name)
	}
	
}

func TestUpdateBin(t *testing.T) {
	setUp()
	filename := "./../my.json"
	name := "test-bin"
	bin, _ := api.CreateBin(&filename, &name)
	defer api.DeleteBin(&bin.Id)

	responseData, _ := api.UpdateBin(&filename, &bin.Id)
	data := *responseData

	if !strings.Contains(string(data["metadata"]), bin.Id) {
		t.Errorf("Expected %v, got %v", bin.Id, string(data["metadata"]))
	}
}

func TestGetBin(t *testing.T) {
	setUp()
	filename := "./../my.json"
	name := "test-bin"
	bin, _ := api.CreateBin(&filename, &name)
	defer api.DeleteBin(&bin.Id)


	responseData, _ := api.GetBin(&bin.Id)
	data := *responseData

	if !strings.Contains(string(data["metadata"]), bin.Id) {
		t.Errorf("Expected %v, got %v", bin.Id, string(data["metadata"]))
	}
	
}

func TestDeleteBin(t *testing.T) {
	setUp()
	filename := "./../my.json"
	name := "test-bin"
	bin, _ := api.CreateBin(&filename, &name)
	defer api.DeleteBin(&bin.Id)

	responseData, _ := api.DeleteBin(&bin.Id)
	data := *responseData

	if !strings.Contains(string(data["metadata"]), bin.Id) {
		t.Errorf("Expected %v, got %v", bin.Id, string(data["metadata"]))
	}
	
}

func TestListBins(t *testing.T) {
	setUp()
	filename := "./../my.json"
	name := "test-bin"
	bin, _ := api.CreateBin(&filename, &name)
	defer api.DeleteBin(&bin.Id)

	bins, _ := api.ListBins()
	if len(bins.Bins) == 0 {
		t.Errorf("Expected %v, got %v", 1, len(bins.Bins))
	}
}
