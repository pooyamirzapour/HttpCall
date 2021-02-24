package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	urlStr := "http://dummy.restapiexample.com/api/v1/employee/1"
	client := &http.Client{}
	r, _ := http.NewRequest(http.MethodGet, urlStr, nil) // URL-encoded payload
	resp, _ := client.Do(r)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		fmt.Println(string(body))
	}

}
