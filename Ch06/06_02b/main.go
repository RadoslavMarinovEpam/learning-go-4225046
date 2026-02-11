package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "http://services.explorecalifornia.org/json/tours.php"

func main() {
	fmt.Println("Network requests")
	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	checkError(err)

	req.Header.Set("User-agent", "")

	resp, err := client.Do(req)
	checkError(err)
	defer resp.Body.Close()
	fmt.Printf("Response status: %s\n", resp.Status)

	byte, err := io.ReadAll(resp.Body)
	checkError(err)
	fmt.Printf("Response body: %s\n", string(byte))
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
