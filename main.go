package main

import (
	"fmt"
	"sync"
)

type Animal interface {
	Speak() string
}

type Dog struct{}

func (d *Dog) Speak() string {
	return "Woof!"
}

type AnimalFactory interface {
	CreateAnimal() Animal
}

type DogFactory struct{}

func (d *DogFactory) CreateAnimal() Animal {
	return &Dog{}
}

 func main() {
	d := DogFactory{}
	dog := d.CreateAnimal()
	fmt.Println(dog.Speak())
 }

//worker pool design pattern
// func worker(jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for j := range jobs {
// 		results <- j * 2
// 	}
// }

// func main() {
// 	jobs := make(chan int, 100)
// 	results := make(chan int, 100)
// 	var wg sync.WaitGroup

// 	for w := 1; w < 4; w++ {
// 		wg.Add(1)
// 		go worker(jobs, results, &wg)
// 	}

// 	for j := 1; j <= 1010; j++ {
// 		jobs <- j
// 		fmt.Println(<-results)
// 	}

// 	close(jobs)
// 	wg.Wait()
// }
