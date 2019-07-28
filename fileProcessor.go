package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Create("sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	file.WriteString("This is a sample test")
	file.Close()
	stream, err := ioutil.ReadFile("sample.txt")
	if err != nil {
		log.Fatal(err)
	}

	readString := string(stream)
	fmt.Println(readString)

	randint := 5
	randfloat := 10.5
	randstring := "100"
	randString2 := "20.5"

	fmt.Println(float64(randint))
	fmt.Println(int(randfloat))
	newint, _ := strconv.ParseInt(randstring, 0, 64)
	fmt.Println(newint)
	newfloat, _ := strconv.ParseFloat(randString2, 64)
	fmt.Println(newfloat)
}
