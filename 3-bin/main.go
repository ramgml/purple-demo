package main

import (
	"3-bin/api"
	"3-bin/config"
	"flag"
	"fmt"
)

func main() {
	config.LoadEnv(".env")
	isCreate := flag.Bool("create", false, "Создать бин")
	isUpdate := flag.Bool("update", false, "Изменить бин")
	isDelete := flag.Bool("delete", false, "Удалить бин по ID")
	isGet := flag.Bool("get", false, "Получить бин по ID")
	isList := flag.Bool("list", false, "Получить список бинов")

	id := flag.String("id", "", "ID бина")
	filename := flag.String("file", "", "Файл для загрузки")
	name := flag.String("name", "", "Название")

	flag.Parse()
	switch true {
	case *isCreate:
		bin, err := api.CreateBin(filename, name)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Printf("Бин %s создан", bin.Name)
		}
	case *isUpdate:
		responseData, err := api.UpdateBin(filename, id)
		data := *responseData
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Printf("Бин %s обновлен", data["metadata"])
		}
	case *isDelete:
		_, err := api.DeleteBin(id)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("Бин удален")
		}
	case *isGet:
		data, err := api.GetBin(id)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			raw := *data
			fmt.Println(string(raw["record"]))
			fmt.Println(string(raw["metadata"]))
			fmt.Println("---")
		}

	case *isList:
		bins, err := api.ListBins()
		if err != nil {
			fmt.Println(err.Error())
		}
		for _, bin := range bins.Bins {
			fmt.Println("ID:", bin.Id)
			fmt.Println("Name:", bin.Name)
			fmt.Println("Private:", bin.Private)
			fmt.Println("Created at:", bin.CreatedAt)
			fmt.Println("---")
		}
	default:
		flag.Usage()
	}
}
