package main

import (
	"fmt"
	"indicator/macd"
)

func main() {

	macd := macd.NewMacd(12, 26, 9)
	var price float64 = 35.68
	f1, f2, f3 := macd.Update(price)
	fmt.Println(f1)
	fmt.Println(f2)
	fmt.Println(f3)
	price  = 37.68
	f1, f2, f3 = macd.Update(price)
	fmt.Println(f1)
	fmt.Println(f2)
	fmt.Println(f3)
	price  = 39.68
	f1, f2, f3 = macd.Update(price)
	fmt.Println(f1)
	fmt.Println(f2)
	fmt.Println(f3)

}
