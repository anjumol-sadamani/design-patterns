package main

import (
	"fmt"
	"sync"
)

type Animal interface {
	Speak() string
	SpeakMore() string
}

type Dog struct{}

func (d Dog) Speak() string {
	return "Woof!"
}

func (d Dog) SpeakMore() string {
	return "Woof, wooof... wooof!"
}

type AnimalFactory interface {
	CreateAnimal() Animal
}

type DogFactory struct{}

func (d *DogFactory) CreateAnimal() Animal {
	return Dog{}
}

// func main() {
// 	d := DogFactory{}
// 	dog := d.CreateAnimal()
// 	fmt.Println(dog.Speak())
// 	fmt.Println(dog.SpeakMore())

// }

// The Factory Design Pattern is a creational design pattern used to abstract the instantiation process of objects.
// It provides a way to create objects without specifying the exact class of the object that will be created. This
// pattern allows for better code organization and flexibility, especially when dealing with complex object creation
// logic.

func worker(jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		results <- j * 2
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	var wg sync.WaitGroup

	for w := 1; w < 4; w++ {
		wg.Add(1)
		go worker(jobs, results, &wg)
	}

	for j := 1; j <= 1010; j++ {
		jobs <- j
		fmt.Println(<-results)
	}

	close(jobs)
	wg.Wait()
}
