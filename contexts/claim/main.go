package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {

	res, err := http.Get("https://google.com")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(len(body))
}
