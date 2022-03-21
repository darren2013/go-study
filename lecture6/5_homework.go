package main

import "fmt"

func main() {

	strArr := [...]string{"I", "am", "stupid", "and", "weak"}

	fmt.Println("update before ", strArr)

	for i, v := range strArr {
		fmt.Println(i, v)

		switch v {
		case "stupid":
			strArr[i] = "smart"
		case "weak":
			strArr[i] = "strong"
		}
	}

	fmt.Println("update after ", strArr)
}
