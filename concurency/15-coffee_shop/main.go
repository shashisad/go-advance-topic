package main

import (
	"fmt"
	"log"
	"sync"
)

// setup constants
const baristaCount = 3
const customerCount = 20
const maxOrderCount = 40

// the total amount of drinks that the bartenders have made
type coffeeShop struct {
	orderCount int
	orderLock  sync.Mutex

	orderCoffee  chan struct{}
	finishCoffee chan struct{}
	closeShop    chan struct{}
}

// registerOrder ensures that the order made by the baristas is counted
func (p *coffeeShop) registerOrder() {
	p.orderLock.Lock()
	defer p.orderLock.Unlock()
	p.orderCount++
	if p.orderCount == maxOrderCount {
		close(p.closeShop)
	}
}

// barista is the resource producer of the coffee shop
func (p *coffeeShop) barista(name string) {
	for {
		select {
		case <-p.orderCoffee:
			p.registerOrder()
			log.Printf("%s makes a coffee.\n", name)
			p.finishCoffee <- struct{}{}
		case <-p.closeShop:
			log.Printf("%s stops working. Bye!\n", name)
			return
		}
	}
}

// customer is the resource consumer of the coffee shop
func (p *coffeeShop) customer(name string) {
	for {
		select {
		case p.orderCoffee <- struct{}{}:
			log.Printf("%s orders a coffee!", name)
			<-p.finishCoffee
			log.Printf("%s enjoys a coffee!\n", name)
		case <-p.closeShop:
			log.Printf("%s leaves the shop! Bye!\n", name)
			return
		}
	}
}

func main() {
	log.Println("Welcome to the Level Up Go coffee shop!")
	orderCoffee := make(chan struct{}, baristaCount)
	finishCoffee := make(chan struct{}, baristaCount)
	closeShop := make(chan struct{})
	p := coffeeShop{
		orderCoffee:  orderCoffee,
		finishCoffee: finishCoffee,
		closeShop:    closeShop,
	}
	for i := 0; i < baristaCount; i++ {
		go p.barista(fmt.Sprint("Barista-", i))
	}
	for i := 0; i < customerCount; i++ {
		go p.customer(fmt.Sprint("Customer-", i))
	}
	<-closeShop
	log.Println("The Level Up Go coffee shop has closed! Bye!")
}


// $ go run main.go
// 2023/10/01 18:23:52 Welcome to the Level Up Go coffee shop!
// 2023/10/01 18:23:52 Customer-19 orders a coffee!
// 2023/10/01 18:23:52 Barista-0 makes a coffee.
// 2023/10/01 18:23:52 Customer-19 enjoys a coffee!
// 2023/10/01 18:23:52 Customer-19 orders a coffee!
// 2023/10/01 18:23:52 Barista-0 makes a coffee.
// 2023/10/01 18:23:52 Customer-19 enjoys a coffee!
// 2023/10/01 18:23:52 Customer-19 orders a coffee!
// 2023/10/01 18:23:52 Barista-0 makes a coffee.
// 2023/10/01 18:23:52 Customer-19 enjoys a coffee!
// 2023/10/01 18:23:52 Customer-19 orders a coffee!
// 2023/10/01 18:23:52 Barista-0 makes a coffee.
// 2023/10/01 18:23:52 Customer-19 enjoys a coffee!
// 2023/10/01 18:23:52 Customer-19 orders a coffee!
// 2023/10/01 18:23:52 Barista-0 makes a coffee.
// 2023/10/01 18:23:52 Customer-19 enjoys a coffee!
// 2023/10/01 18:23:52 Customer-19 orders a coffee!
// 2023/10/01 18:23:52 Customer-8 orders a coffee!
// 2023/10/01 18:23:52 Barista-1 makes a coffee.
// 2023/10/01 18:23:52 Barista-1 makes a coffee.
// 2023/10/01 18:23:52 Barista-1 makes a coffee.
// 2023/10/01 18:23:52 Barista-1 makes a coffee.
// 2023/10/01 18:23:52 Barista-1 makes a coffee.
// 2023/10/01 18:23:52 Barista-1 makes a coffee.
// 2023/10/01 18:23:52 Customer-0 orders a coffee!
// 2023/10/01 18:23:52 Customer-0 enjoys a coffee!
// 2023/10/01 18:23:52 Barista-2 makes a coffee.
// 2023/10/01 18:23:52 Customer-1 orders a coffee!
// 2023/10/01 18:23:52 Customer-1 enjoys a coffee!
// 2023/10/01 18:23:52 Customer-2 orders a coffee!
// 2023/10/01 18:23:52 Customer-2 enjoys a coffee!
// 2023/10/01 18:23:52 Customer-3 orders a coffee!
// 2023/10/01 18:23:52 Customer-3 enjoys a coffee!
// 2023/10/01 18:23:52 Barista-0 makes a coffee.
// 2023/10/01 18:23:52 Barista-0 makes a coffee.
// 2023/10/01 18:23:52 Barista-0 makes a coffee.
// 2023/10/01 18:23:52 Customer-19 enjoys a coffee!
// 2023/10/01 18:23:52 Customer-4 orders a coffee!
// 2023/10/01 18:23:52 Customer-4 enjoys a coffee!
// 2023/10/01 18:23:52 Customer-8 enjoys a coffee!
// 2023/10/01 18:23:52 Customer-5 orders a coffee!
// 2023/10/01 18:23:52 Customer-5 enjoys a coffee!
// 2023/10/01 18:23:52 Customer-6 orders a coffee!
// 2023/10/01 18:23:52 Customer-6 enjoys a coffee!
// 2023/10/01 18:23:52 Customer-7 orders a coffee!
// 2023/10/01 18:23:52 Customer-7 enjoys a coffee!
// 2023/10/01 18:23:52 Customer-13 orders a coffee!
// 2023/10/01 18:23:52 Barista-1 makes a coffee.
// 2023/10/01 18:23:52 Barista-1 makes a coffee.
// 2023/10/01 18:23:52 Barista-1 makes a coffee.
// 2023/10/01 18:23:52 Barista-1 makes a coffee.
// 2023/10/01 18:23:52 Barista-1 makes a coffee.
// 2023/10/01 18:23:52 Customer-9 orders a coffee!
// 2023/10/01 18:23:52 Customer-9 enjoys a coffee!
// 2023/10/01 18:23:52 Barista-2 makes a coffee.
// 2023/10/01 18:23:52 Customer-10 orders a coffee!
// 2023/10/01 18:23:52 Customer-10 enjoys a coffee!
// 2023/10/01 18:23:52 Customer-11 orders a coffee!
// 2023/10/01 18:23:52 Customer-11 enjoys a coffee!
// 2023/10/01 18:23:52 Customer-12 orders a coffee!
// 2023/10/01 18:23:52 Customer-12 enjoys a coffee!
// 2023/10/01 18:23:52 Barista-0 makes a coffee.
// 2023/10/01 18:23:52 Barista-0 makes a coffee.
// 2023/10/01 18:23:52 Barista-0 makes a coffee.
// 2023/10/01 18:23:52 Customer-16 orders a coffee!
// 2023/10/01 18:23:52 Customer-16 enjoys a coffee!
// 2023/10/01 18:23:52 Customer-13 enjoys a coffee!
// 2023/10/01 18:23:52 Customer-14 orders a coffee!
// 2023/10/01 18:23:52 Customer-14 enjoys a coffee!
// 2023/10/01 18:23:52 Customer-15 orders a coffee!
// 2023/10/01 18:23:52 Customer-15 enjoys a coffee!
// 2023/10/01 18:23:52 Customer-17 orders a coffee!
// 2023/10/01 18:23:52 Customer-17 enjoys a coffee!
// 2023/10/01 18:23:52 Customer-18 orders a coffee!
// 2023/10/01 18:23:52 Barista-1 makes a coffee.
// 2023/10/01 18:23:52 Barista-1 makes a coffee.
// 2023/10/01 18:23:52 Barista-1 makes a coffee.
// 2023/10/01 18:23:52 Barista-1 makes a coffee.
// 2023/10/01 18:23:52 Barista-1 makes a coffee.
// 2023/10/01 18:23:52 Customer-0 orders a coffee!
// 2023/10/01 18:23:52 Customer-0 enjoys a coffee!
// 2023/10/01 18:23:52 Barista-2 makes a coffee.
// 2023/10/01 18:23:52 Customer-1 orders a coffee!
// 2023/10/01 18:23:52 Customer-1 enjoys a coffee!
// 2023/10/01 18:23:52 Customer-2 orders a coffee!
// 2023/10/01 18:23:52 Customer-2 enjoys a coffee!
// 2023/10/01 18:23:52 Customer-3 orders a coffee!
// 2023/10/01 18:23:52 Customer-3 enjoys a coffee!
// 2023/10/01 18:23:52 Barista-0 makes a coffee.
// 2023/10/01 18:23:52 Barista-0 makes a coffee.
// 2023/10/01 18:23:52 Barista-0 makes a coffee.
// 2023/10/01 18:23:52 Customer-19 orders a coffee!
// 2023/10/01 18:23:52 Customer-19 enjoys a coffee!
// 2023/10/01 18:23:52 Customer-18 enjoys a coffee!
// 2023/10/01 18:23:52 Customer-4 orders a coffee!
// 2023/10/01 18:23:52 Customer-4 enjoys a coffee!
// 2023/10/01 18:23:52 Customer-8 orders a coffee!
// 2023/10/01 18:23:52 Customer-8 enjoys a coffee!
// 2023/10/01 18:23:52 Customer-5 orders a coffee!
// 2023/10/01 18:23:52 Customer-5 enjoys a coffee!
// 2023/10/01 18:23:52 Customer-6 orders a coffee!
// 2023/10/01 18:23:52 Barista-1 makes a coffee.
// 2023/10/01 18:23:52 Barista-1 makes a coffee.
// 2023/10/01 18:23:52 Barista-1 makes a coffee.
// 2023/10/01 18:23:52 Barista-1 makes a coffee.
// 2023/10/01 18:23:52 Barista-1 makes a coffee.
// 2023/10/01 18:23:52 Customer-7 orders a coffee!
// 2023/10/01 18:23:52 Customer-7 enjoys a coffee!
// 2023/10/01 18:23:52 Customer-7 leaves the shop! Bye!
// 2023/10/01 18:23:52 Customer-18 leaves the shop! Bye!
// 2023/10/01 18:23:52 Barista-2 makes a coffee.
// 2023/10/01 18:23:52 Customer-9 orders a coffee!
// 2023/10/01 18:23:52 Customer-9 enjoys a coffee!
// 2023/10/01 18:23:52 Customer-9 orders a coffee!
// 2023/10/01 18:23:52 Customer-9 enjoys a coffee!
// 2023/10/01 18:23:52 Customer-9 leaves the shop! Bye!
// 2023/10/01 18:23:52 Customer-10 orders a coffee!
// 2023/10/01 18:23:52 Customer-10 enjoys a coffee!
// 2023/10/01 18:23:52 Customer-10 leaves the shop! Bye!
// 2023/10/01 18:23:52 Customer-11 orders a coffee!
// 2023/10/01 18:23:52 Customer-11 enjoys a coffee!
// 2023/10/01 18:23:52 Customer-11 leaves the shop! Bye!
// 2023/10/01 18:23:52 Barista-0 makes a coffee.
// 2023/10/01 18:23:52 Barista-0 stops working. Bye!
// 2023/10/01 18:23:52 Customer-12 orders a coffee!
// 2023/10/01 18:23:52 Customer-12 enjoys a coffee!
// 2023/10/01 18:23:52 Customer-12 leaves the shop! Bye!
// 2023/10/01 18:23:52 Customer-6 enjoys a coffee!
// 2023/10/01 18:23:52 Customer-6 leaves the shop! Bye!
// 2023/10/01 18:23:52 Customer-16 orders a coffee!
// 2023/10/01 18:23:52 Customer-13 orders a coffee!
// 2023/10/01 18:23:52 Customer-14 orders a coffee!
// 2023/10/01 18:23:52 Customer-15 orders a coffee!
// 2023/10/01 18:23:52 Customer-5 leaves the shop! Bye!
// 2023/10/01 18:23:52 Customer-8 leaves the shop! Bye!
// 2023/10/01 18:23:52 Customer-4 leaves the shop! Bye!
// 2023/10/01 18:23:52 Customer-2 leaves the shop! Bye!
// 2023/10/01 18:23:52 Customer-19 leaves the shop! Bye!
// 2023/10/01 18:23:52 Customer-3 leaves the shop! Bye!
// 2023/10/01 18:23:52 Customer-0 leaves the shop! Bye!
// 2023/10/01 18:23:52 Customer-1 leaves the shop! Bye!
// 2023/10/01 18:23:52 Customer-17 leaves the shop! Bye!
// 2023/10/01 18:23:52 The Level Up Go coffee shop has closed! Bye!
// $ 