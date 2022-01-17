package main

import (
	"fmt"
	"github.com/tidwall/gjson"
	"log"
	"strconv"
	"os"
)

func readFileContents(json string){
	bytes, err := os.ReadFile("data.txt");
	if err != nil {
		log.Fatal(err);
	}
	json = string(bytes[:])
	//fmt.Println(json)
}

func main() {
	json:= ``
	//var json string
	//readFileContents(json)
	var countVendors int
	var countChains int
	totalCount := gjson.Get(json,"payload.totalCount")
	countChain := totalCount.String()
	countString, err := strconv.Atoi(countChain)

	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < countString; i++{
		idVendorsArray := gjson.Get(json,"payload.items."+strconv.Itoa(i)+".filteredVendorIds")
		idChain := gjson.Get(json,"payload.items."+strconv.Itoa(i)+".id")
		idVendors := idVendorsArray.Array()
		fmt.Print("https://www.delivery-club.ru/ajax/manage/?m=xml_update&id=",idChain, "\n")
		countChains++
		//fmt.Println(countChains)

		for  idx := 0; idx < len(idVendors); idx++{

			vendor := idVendors[idx].String()
			countVendors++
			//fmt.Print(countVendors)
			fmt.Print("https://www.delivery-club.ru/ajax/sd/?m=xml_update&id=",vendor, "\n")
		}
	}
}

