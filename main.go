package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("BORING!")
	// c := make(chan string)
	// go boring("boring!", c)
	// for i := 0; i < 5; i++ {
	// 	fmt.Printf("You say: %s\n", <-c)
	// }
	c := FanIn(boringGenerator("Ann"), boringGenerator("Joe"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("You're boring! I'm leaving")
}

func boring(msg string, c chan string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintf("%s %d", msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

func boringGenerator(msg string) <-chan string { // return receive-only channel string
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}
