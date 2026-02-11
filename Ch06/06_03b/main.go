package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

const url = "http://services.explorecalifornia.org/json/cars.php"

func main() {
	content := readHttpContent()
	//fmt.Print(content)
	cars := carsFromJSON(content)
	for _, car := range cars {
		car.Price = car.Price * float64(car.Quantity)
		price, err := strconv.ParseFloat(fmt.Sprintf("%.2f", car.Price), 64)
		checkError(err)
		fmt.Printf("cars: %s, Price: $%.2f\n", car.Name, price)
	}
}

func carsFromJSON(content string) []CartItem {
	cars := make([]CartItem, 0)
	decoder := json.NewDecoder(strings.NewReader(content))
	_, err := decoder.Token()
	checkError(err)

	var car CartItem
	for decoder.More() {
		err := decoder.Decode(&car)
		checkError(err)
		cars = append(cars, car)
	}
	return cars
}

func readHttpContent() string {
	fmt.Println("Network requests")
	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	checkError(err)

	req.Header.Set("User-Agent", "")

	resp, err := client.Do(req)
	checkError(err)
	defer resp.Body.Close()

	fmt.Printf("Response type:%T\n", resp)

	bytes, err := io.ReadAll(resp.Body)
	checkError(err)
	return string(bytes)

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

type Tour struct {
	Name, Price string
}

type CartItem struct {
	Name     string
	Price    float64
	Quantity int
}
