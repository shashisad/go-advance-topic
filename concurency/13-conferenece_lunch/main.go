
package main

import (
	"fmt"
	"log"
)

// the number of attendees we need to serve lunch to
const consumerCount = 50

// foodCourses represents the types of resources to pass to the consumers
var foodCourses = []string{
	"Caprese Salad",
	"Spaghetti Carbonara",
	"Vanilla Panna Cotta",
}

// takeLunch is the consumer function for the lunch simulation
// Change the signature of this function as required
func takeLunch(name string, in []chan string, done chan<- struct{}) {
	for _, ch := range in {
		log.Printf("%s eats %s.\n", name, <-ch)
	}
	done <- struct{}{}
}

// serveLunch is the producer function for the lunch simulation.
// Change the signature of this function as required
func serveLunch(course string, out chan<- string, done <-chan struct{}) {
	for {
		select {
		case out <- course:
		case <-done:
			return
		}
	}
}

func main() {
	log.Printf("Welcome to the conference lunch! Serving %d attendees.\n",
		consumerCount)
	var courses []chan string
	doneEating := make(chan struct{})
	doneServing := make(chan struct{})
	for _, c := range foodCourses {
		ch := make(chan string)
		courses = append(courses, ch)
		go serveLunch(c, ch, doneServing)
	}
	for i := 0; i < consumerCount; i++ {
		name := fmt.Sprintf("Attendee %d", i)
		go takeLunch(name, courses, doneEating)
	}

	for i := 0; i < consumerCount; i++ {
		<-doneEating
	}
	close(doneServing)
}


// $ go run main.go
// 2023/10/01 18:17:32 Welcome to the conference lunch! Serving 50 attendees.
// 2023/10/01 18:17:32 Attendee 49 eats Caprese Salad.
// 2023/10/01 18:17:32 Attendee 49 eats Spaghetti Carbonara.
// 2023/10/01 18:17:32 Attendee 49 eats Vanilla Panna Cotta.
// 2023/10/01 18:17:32 Attendee 0 eats Caprese Salad.
// 2023/10/01 18:17:32 Attendee 0 eats Spaghetti Carbonara.
// 2023/10/01 18:17:32 Attendee 0 eats Vanilla Panna Cotta.
// 2023/10/01 18:17:32 Attendee 48 eats Caprese Salad.
// 2023/10/01 18:17:32 Attendee 48 eats Spaghetti Carbonara.
// 2023/10/01 18:17:32 Attendee 48 eats Vanilla Panna Cotta.
// 2023/10/01 18:17:32 Attendee 1 eats Caprese Salad.
// 2023/10/01 18:17:32 Attendee 1 eats Spaghetti Carbonara.
// 2023/10/01 18:17:32 Attendee 1 eats Vanilla Panna Cotta.
// 2023/10/01 18:17:32 Attendee 2 eats Caprese Salad.
// 2023/10/01 18:17:32 Attendee 3 eats Caprese Salad.
// 2023/10/01 18:17:32 Attendee 4 eats Caprese Salad.
// 2023/10/01 18:17:32 Attendee 5 eats Caprese Salad.
// 2023/10/01 18:17:32 Attendee 6 eats Caprese Salad.
// 2023/10/01 18:17:32 Attendee 7 eats Caprese Salad.
// 2023/10/01 18:17:32 Attendee 8 eats Caprese Salad.
// 2023/10/01 18:17:32 Attendee 9 eats Caprese Salad.
// 2023/10/01 18:17:32 Attendee 10 eats Caprese Salad.
// 2023/10/01 18:17:32 Attendee 11 eats Caprese Salad.
// 2023/10/01 18:17:32 Attendee 12 eats Caprese Salad.
// 2023/10/01 18:17:32 Attendee 13 eats Caprese Salad.
// 2023/10/01 18:17:32 Attendee 14 eats Caprese Salad.
// 2023/10/01 18:17:32 Attendee 15 eats Caprese Salad.
// 2023/10/01 18:17:32 Attendee 16 eats Caprese Salad.
// 2023/10/01 18:17:32 Attendee 17 eats Caprese Salad.
// 2023/10/01 18:17:32 Attendee 18 eats Caprese Salad.
// 2023/10/01 18:17:32 Attendee 19 eats Caprese Salad.
// 2023/10/01 18:17:32 Attendee 20 eats Caprese Salad.
// 2023/10/01 18:17:32 Attendee 21 eats Caprese Salad.
// 2023/10/01 18:17:32 Attendee 22 eats Caprese Salad.
// 2023/10/01 18:17:32 Attendee 23 eats Caprese Salad.
// 2023/10/01 18:17:32 Attendee 24 eats Caprese Salad.
// 2023/10/01 18:17:32 Attendee 25 eats Caprese Salad.
// 2023/10/01 18:17:32 Attendee 26 eats Caprese Salad.
// 2023/10/01 18:17:32 Attendee 27 eats Caprese Salad.
// 2023/10/01 18:17:32 Attendee 28 eats Caprese Salad.
// 2023/10/01 18:17:32 Attendee 29 eats Caprese Salad.
// 2023/10/01 18:17:32 Attendee 30 eats Caprese Salad.
// 2023/10/01 18:17:32 Attendee 31 eats Caprese Salad.
// 2023/10/01 18:17:32 Attendee 32 eats Caprese Salad.
// 2023/10/01 18:17:32 Attendee 33 eats Caprese Salad.
// 2023/10/01 18:17:32 Attendee 34 eats Caprese Salad.
// 2023/10/01 18:17:32 Attendee 35 eats Caprese Salad.
// 2023/10/01 18:17:32 Attendee 36 eats Caprese Salad.
// 2023/10/01 18:17:32 Attendee 37 eats Caprese Salad.
// 2023/10/01 18:17:32 Attendee 38 eats Caprese Salad.
// 2023/10/01 18:17:32 Attendee 39 eats Caprese Salad.
// 2023/10/01 18:17:32 Attendee 40 eats Caprese Salad.
// 2023/10/01 18:17:32 Attendee 41 eats Caprese Salad.
// 2023/10/01 18:17:32 Attendee 42 eats Caprese Salad.
// 2023/10/01 18:17:32 Attendee 43 eats Caprese Salad.
// 2023/10/01 18:17:32 Attendee 44 eats Caprese Salad.
// 2023/10/01 18:17:32 Attendee 45 eats Caprese Salad.
// 2023/10/01 18:17:32 Attendee 46 eats Caprese Salad.
// 2023/10/01 18:17:32 Attendee 47 eats Caprese Salad.
// 2023/10/01 18:17:32 Attendee 47 eats Spaghetti Carbonara.
// 2023/10/01 18:17:32 Attendee 47 eats Vanilla Panna Cotta.
// 2023/10/01 18:17:32 Attendee 2 eats Spaghetti Carbonara.
// 2023/10/01 18:17:32 Attendee 2 eats Vanilla Panna Cotta.
// 2023/10/01 18:17:32 Attendee 3 eats Spaghetti Carbonara.
// 2023/10/01 18:17:32 Attendee 4 eats Spaghetti Carbonara.
// 2023/10/01 18:17:32 Attendee 5 eats Spaghetti Carbonara.
// 2023/10/01 18:17:32 Attendee 6 eats Spaghetti Carbonara.
// 2023/10/01 18:17:32 Attendee 7 eats Spaghetti Carbonara.
// 2023/10/01 18:17:32 Attendee 8 eats Spaghetti Carbonara.
// 2023/10/01 18:17:32 Attendee 9 eats Spaghetti Carbonara.
// 2023/10/01 18:17:32 Attendee 10 eats Spaghetti Carbonara.
// 2023/10/01 18:17:32 Attendee 11 eats Spaghetti Carbonara.
// 2023/10/01 18:17:32 Attendee 12 eats Spaghetti Carbonara.
// 2023/10/01 18:17:32 Attendee 13 eats Spaghetti Carbonara.
// 2023/10/01 18:17:32 Attendee 14 eats Spaghetti Carbonara.
// 2023/10/01 18:17:32 Attendee 15 eats Spaghetti Carbonara.
// 2023/10/01 18:17:32 Attendee 16 eats Spaghetti Carbonara.
// 2023/10/01 18:17:32 Attendee 17 eats Spaghetti Carbonara.
// 2023/10/01 18:17:32 Attendee 18 eats Spaghetti Carbonara.
// 2023/10/01 18:17:32 Attendee 19 eats Spaghetti Carbonara.
// 2023/10/01 18:17:32 Attendee 20 eats Spaghetti Carbonara.
// 2023/10/01 18:17:32 Attendee 21 eats Spaghetti Carbonara.
// 2023/10/01 18:17:32 Attendee 22 eats Spaghetti Carbonara.
// 2023/10/01 18:17:32 Attendee 23 eats Spaghetti Carbonara.
// 2023/10/01 18:17:32 Attendee 24 eats Spaghetti Carbonara.
// 2023/10/01 18:17:32 Attendee 25 eats Spaghetti Carbonara.
// 2023/10/01 18:17:32 Attendee 26 eats Spaghetti Carbonara.
// 2023/10/01 18:17:32 Attendee 27 eats Spaghetti Carbonara.
// 2023/10/01 18:17:32 Attendee 28 eats Spaghetti Carbonara.
// 2023/10/01 18:17:32 Attendee 29 eats Spaghetti Carbonara.
// 2023/10/01 18:17:32 Attendee 30 eats Spaghetti Carbonara.
// 2023/10/01 18:17:32 Attendee 31 eats Spaghetti Carbonara.
// 2023/10/01 18:17:32 Attendee 32 eats Spaghetti Carbonara.
// 2023/10/01 18:17:32 Attendee 33 eats Spaghetti Carbonara.
// 2023/10/01 18:17:32 Attendee 34 eats Spaghetti Carbonara.
// 2023/10/01 18:17:32 Attendee 35 eats Spaghetti Carbonara.
// 2023/10/01 18:17:32 Attendee 36 eats Spaghetti Carbonara.
// 2023/10/01 18:17:32 Attendee 37 eats Spaghetti Carbonara.
// 2023/10/01 18:17:32 Attendee 38 eats Spaghetti Carbonara.
// 2023/10/01 18:17:32 Attendee 39 eats Spaghetti Carbonara.
// 2023/10/01 18:17:32 Attendee 40 eats Spaghetti Carbonara.
// 2023/10/01 18:17:32 Attendee 41 eats Spaghetti Carbonara.
// 2023/10/01 18:17:32 Attendee 42 eats Spaghetti Carbonara.
// 2023/10/01 18:17:32 Attendee 43 eats Spaghetti Carbonara.
// 2023/10/01 18:17:32 Attendee 44 eats Spaghetti Carbonara.
// 2023/10/01 18:17:32 Attendee 45 eats Spaghetti Carbonara.
// 2023/10/01 18:17:32 Attendee 46 eats Spaghetti Carbonara.
// 2023/10/01 18:17:32 Attendee 46 eats Vanilla Panna Cotta.
// 2023/10/01 18:17:32 Attendee 3 eats Vanilla Panna Cotta.
// 2023/10/01 18:17:32 Attendee 4 eats Vanilla Panna Cotta.
// 2023/10/01 18:17:32 Attendee 5 eats Vanilla Panna Cotta.
// 2023/10/01 18:17:32 Attendee 6 eats Vanilla Panna Cotta.
// 2023/10/01 18:17:32 Attendee 7 eats Vanilla Panna Cotta.
// 2023/10/01 18:17:32 Attendee 8 eats Vanilla Panna Cotta.
// 2023/10/01 18:17:32 Attendee 9 eats Vanilla Panna Cotta.
// 2023/10/01 18:17:32 Attendee 10 eats Vanilla Panna Cotta.
// 2023/10/01 18:17:32 Attendee 11 eats Vanilla Panna Cotta.
// 2023/10/01 18:17:32 Attendee 12 eats Vanilla Panna Cotta.
// 2023/10/01 18:17:32 Attendee 13 eats Vanilla Panna Cotta.
// 2023/10/01 18:17:32 Attendee 14 eats Vanilla Panna Cotta.
// 2023/10/01 18:17:32 Attendee 15 eats Vanilla Panna Cotta.
// 2023/10/01 18:17:32 Attendee 16 eats Vanilla Panna Cotta.
// 2023/10/01 18:17:32 Attendee 17 eats Vanilla Panna Cotta.
// 2023/10/01 18:17:32 Attendee 18 eats Vanilla Panna Cotta.
// 2023/10/01 18:17:32 Attendee 19 eats Vanilla Panna Cotta.
// 2023/10/01 18:17:32 Attendee 20 eats Vanilla Panna Cotta.
// 2023/10/01 18:17:32 Attendee 21 eats Vanilla Panna Cotta.
// 2023/10/01 18:17:32 Attendee 22 eats Vanilla Panna Cotta.
// 2023/10/01 18:17:32 Attendee 23 eats Vanilla Panna Cotta.
// 2023/10/01 18:17:32 Attendee 24 eats Vanilla Panna Cotta.
// 2023/10/01 18:17:32 Attendee 25 eats Vanilla Panna Cotta.
// 2023/10/01 18:17:32 Attendee 26 eats Vanilla Panna Cotta.
// 2023/10/01 18:17:32 Attendee 27 eats Vanilla Panna Cotta.
// 2023/10/01 18:17:32 Attendee 28 eats Vanilla Panna Cotta.
// 2023/10/01 18:17:32 Attendee 29 eats Vanilla Panna Cotta.
// 2023/10/01 18:17:32 Attendee 30 eats Vanilla Panna Cotta.
// 2023/10/01 18:17:32 Attendee 31 eats Vanilla Panna Cotta.
// 2023/10/01 18:17:32 Attendee 32 eats Vanilla Panna Cotta.
// 2023/10/01 18:17:32 Attendee 33 eats Vanilla Panna Cotta.
// 2023/10/01 18:17:32 Attendee 34 eats Vanilla Panna Cotta.
// 2023/10/01 18:17:32 Attendee 35 eats Vanilla Panna Cotta.
// 2023/10/01 18:17:32 Attendee 36 eats Vanilla Panna Cotta.
// 2023/10/01 18:17:32 Attendee 37 eats Vanilla Panna Cotta.
// 2023/10/01 18:17:32 Attendee 38 eats Vanilla Panna Cotta.
// 2023/10/01 18:17:32 Attendee 39 eats Vanilla Panna Cotta.
// 2023/10/01 18:17:32 Attendee 40 eats Vanilla Panna Cotta.
// 2023/10/01 18:17:32 Attendee 41 eats Vanilla Panna Cotta.
// 2023/10/01 18:17:32 Attendee 42 eats Vanilla Panna Cotta.
// 2023/10/01 18:17:32 Attendee 43 eats Vanilla Panna Cotta.
// 2023/10/01 18:17:32 Attendee 44 eats Vanilla Panna Cotta.
// 2023/10/01 18:17:32 Attendee 45 eats Vanilla Panna Cotta.
// $ 