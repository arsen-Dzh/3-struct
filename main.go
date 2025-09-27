package main

import (
	"demo/3-struct/bins"
	"fmt"
)

func main() {
	for {
		var isCreated string
		fmt.Println("Создать bin yes/no?")
		fmt.Scan(&isCreated)
		if isCreated != "yes" {
			break
		}
		name := promptData("Введите имя bina")
		bin, err := bins.NewBin(name)
		if err != nil {
			return
		}
		fmt.Print(bin)
	}
}

func promptData(prompt string) string {
	fmt.Print(prompt + ": ")
	var res string
	fmt.Scanln(&res)
	return res
}
