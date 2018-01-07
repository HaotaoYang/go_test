package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Address struct {
	Type    string
	City    string
	Country string
}

type VCard struct {
	FirstName string
	LastName  string
	Addresses []*Address
	Remark    string
}

func main() {
	pa := &Address{"private", "LongHua", "China"}
	wa := &Address{"work", "NanShan", "China"}
	vc := VCard{"HaoTao", "Yang", []*Address{pa, wa}, "test"}
	fmt.Printf("vc: %v\n", vc)
	// JSON format:
	js, _ := json.Marshal(vc)
	fmt.Printf("JSON format: %s\n", js)
	// using an encoder:
	file, _ := os.OpenFile("vcard.json", os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()
	enc := json.NewEncoder(file)
	err := enc.Encode(vc)
	if err != nil {
		log.Println("Error in encoding json")
	}
}
