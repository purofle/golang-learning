package main

import "fmt"

func main() {
	fmt.Println("扫描")
	_, err := fmt.Scanln()
	fmt.Println("扫描2")

	//error!!! no new variables on left side of :=
	// _, err := fmt.Scanln()

	if err != nil {
		return
	}
}
