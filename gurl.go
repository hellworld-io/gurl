package main

import (
	"fmt"
	"net/http"
)

func main() {
	res, err := http.Get("https://api.github.com/users/octocat/orgs")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res)
}
