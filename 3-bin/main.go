package main

import (
	"3-bin/api"
	"flag"
	"fmt"
)

func main() {
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
		api.CreateBin(filename, name)
		fmt.Println("Бин создан")
	case *isUpdate:
		api.UpdateBin(filename, id)
		fmt.Println("Бин обновлен")
	case *isDelete:
		api.DeleteBin(id)
		fmt.Println("Бин удален")
	case *isGet:
		api.GetBin(id)
	case *isList:
		api.ListBins()
	default:
		flag.Usage()
	}
}
