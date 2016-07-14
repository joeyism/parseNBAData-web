package main

import(
	"fmt"
)

func showError(err error){
	if err != nil{
		fmt.Println("error: ", err)
	}
}
