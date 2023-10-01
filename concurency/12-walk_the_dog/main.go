package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

const maxSeconds = 3

type Dog struct {
	name string
}

func (d Dog) fetchLeash() {
	log.Printf("%s goes to fetch leash.\n", d.name)
	randomSleep()
	log.Printf("%s has fetched leash. Woof woof!\n", d.name)
}

func (d Dog) findTreats() {
	log.Printf("%s goes to fetch treats.\n", d.name)
	randomSleep()
	log.Printf("%s has fetched the treats. Woof woof!\n", d.name)
}

func (d Dog) runOutside() {
	log.Printf("%s starts running outside.\n", d.name)
	randomSleep()
	log.Printf("%s is having fun outside. Woof woof!\n", d.name)
}

type Owner struct {
	name string
}

func (o Owner) putShoesOn() {
	log.Printf("%s starts putting shoes on.\n", o.name)
	randomSleep()
	log.Printf("%s finishes putting shoes on.\n", o.name)
}

func (o Owner) findKeys() {
	log.Printf("%s starts looking for keys.\n", o.name)
	randomSleep()
	log.Printf("%s has found keys.\n", o.name)
}

func (o Owner) lockDoor() {
	log.Printf("%s starts locking the door.\n", o.name)
	randomSleep()
	log.Printf("%s has locked the door.\n", o.name)
}

func randomSleep() {
	r := rand.Intn(maxSeconds)
	time.Sleep(time.Duration(r)*time.Second + 500*time.Millisecond)
}

func main() {
	owner := Owner{name: "Jimmy"}
	dog := Dog{name: "Lucky"}
	ownerActions := []func(){
		owner.putShoesOn,
		owner.findKeys,
		owner.lockDoor,
	}
	dogActions := []func(){
		dog.fetchLeash,
		dog.findTreats,
		dog.runOutside,
	}
	executeWalk(ownerActions, dogActions)
}

func executeWalk(ownerActions []func(), dogActions []func()) {
	var wg sync.WaitGroup
	wg.Add(2)
	defer wg.Wait()
	executeActions := func(actions []func()) {
		defer wg.Done()
		for _, a := range actions {
			a()
		}
	}
	go executeActions(ownerActions)
	go executeActions(dogActions)
}


//waitgroup

// 2023/10/01 18:12:04 Lucky goes to fetch leash.
// 2023/10/01 18:12:04 Jimmy starts putting shoes on.
// 2023/10/01 18:12:04 Jimmy finishes putting shoes on.
// 2023/10/01 18:12:04 Jimmy starts looking for keys.
// 2023/10/01 18:12:06 Lucky has fetched leash. Woof woof!
// 2023/10/01 18:12:06 Lucky goes to fetch treats.
// 2023/10/01 18:12:07 Jimmy has found keys.
// 2023/10/01 18:12:07 Jimmy starts locking the door.
// 2023/10/01 18:12:08 Jimmy has locked the door.
// 2023/10/01 18:12:09 Lucky has fetched the treats. Woof woof!
// 2023/10/01 18:12:09 Lucky starts running outside.
// 2023/10/01 18:12:09 Lucky is having fun outside. Woof woof!