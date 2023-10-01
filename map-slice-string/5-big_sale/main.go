package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"sort"
)

const path = "items.json"

// SaleItem represents the item part of the big sale.
type SaleItem struct {
	Name           string  `json:"name"`
	OriginalPrice  float64 `json:"originalPrice"`
	ReducedPrice   float64 `json:"reducedPrice"`
	SalePercentage float64
}

// matchSales adds the sales procentage of the item
// and sorts the array accordingly.
func matchSales(budget float64, items []SaleItem) []SaleItem {
	var mi []SaleItem
	for _, si := range items {
		if si.ReducedPrice <= budget {
			si.SalePercentage = -(si.ReducedPrice - si.OriginalPrice) /
				si.OriginalPrice * 100
			mi = append(mi, si)
		}
	}
	sort.Slice(mi, func(i, j int) bool {
		return mi[i].SalePercentage > mi[j].SalePercentage
	})

	return mi
}

func main() {
	budget := flag.Float64("budget", 0.0,
		"The max budget you want to shop with.")
	flag.Parse()
	items := importData()
	matchedItems := matchSales(*budget, items)
	printItems(matchedItems)
}

// printItems prints the items and their sales.
func printItems(items []SaleItem) {
	log.Println("The BIG sale has started with our amazing offers!")
	if len(items) == 0 {
		log.Println("No items found.:( Try increasing your budget.")
	}
	for i, r := range items {
		log.Printf("[%d]:%s is %.2f OFF! Get it now for JUST %.2f!\n",
			i, r.Name, r.SalePercentage, r.ReducedPrice)
	}
}

// importData reads the raffle entries from file and
// creates the entries slice.
func importData() []SaleItem {
	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	var data []SaleItem
	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Fatal(err)
	}

	return data
}


// $ go run main.go -budget 100
// 2023/10/01 17:07:01 The BIG sale has started with our amazing offers!
// 2023/10/01 17:07:01 [0]:Microwave 5000 is 66.67 OFF! Get it now for JUST 99.99!
// 2023/10/01 17:07:01 [1]:Kettle 2000 is 23.95 OFF! Get it now for JUST 55.63!
// 2023/10/01 17:07:01 [2]:Blender 1000 is 20.10 OFF! Get it now for JUST 79.50!
// $ go run main.go -budget 1000
// 2023/10/01 17:07:09 The BIG sale has started with our amazing offers!
// 2023/10/01 17:07:09 [0]:Microwave 5000 is 66.67 OFF! Get it now for JUST 99.99!
// 2023/10/01 17:07:09 [1]:WasherDryer 3000 is 40.00 OFF! Get it now for JUST 599.95!
// 2023/10/01 17:07:09 [2]:Fridge 4000 is 28.57 OFF! Get it now for JUST 499.95!
// 2023/10/01 17:07:09 [3]:Kettle 2000 is 23.95 OFF! Get it now for JUST 55.63!
// 2023/10/01 17:07:09 [4]:Blender 1000 is 20.10 OFF! Get it now for JUST 79.50!