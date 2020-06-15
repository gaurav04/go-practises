package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	mydata := []byte("I am gaurav harsola")

	err := ioutil.WriteFile("loc.data", mydata, 0777)

	if err != nil {
		fmt.Println(err.Error)
	}

	data, err := ioutil.ReadFile("loc.data")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Print(string(data))

	f, err := os.OpenFile("loc.data", os.O_APPEND|os.O_WRONLY, 0777)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString("\n new data that wasn't there originally\n"); err != nil {
		panic(err)
	}

	data, err = ioutil.ReadFile("loc.data")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Print(string(data))
}
