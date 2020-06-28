package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

type FizzBuzz struct {
	n           int
	wg          *sync.WaitGroup
	chanVar chan int
}

func (this *FizzBuzz) PrintLoop(passCondition func(int) bool, printString func(int)) {
	defer this.wg.Done()

	for i := 0; i <= this.n; i++ {
		if passCondition(i) {
			nextNum := <-this.chanVar
	
			if i == nextNum {
				printString(i)
				this.chanVar <- i + 1 
			} else {
				this.chanVar <- nextNum 
				i--
			}
			runtime.Gosched()
		}
	}
}

func (this *FizzBuzz) PrintFizz() {
	PassCondition := func(i int) bool { 
		return (i % 3 == 0) && (i % 5 != 0) 
	}
	PrintString := func(i int) { 
		fmt.Printf("Fizz(%d), ", i) 
	}

	this.PrintLoop(PassCondition, PrintString)
}

func (this *FizzBuzz) PrintBuzz() {
	PassCondition := func(i int) bool {
		return (i % 3 != 0) && (i % 5 == 0)
	}
	PrintString := func(i int) {
		fmt.Printf("Buzz(%d), ", i)
	}

	this.PrintLoop(PassCondition, PrintString)
}

func (this *FizzBuzz) PrintFizzBuzz() {
	PassCondition := func(i int) bool {
		return i % 15 == 0
	}
	PrintString := func(i int) {
		fmt.Printf("FizzBuzz(%d), ", i)
}

	this.PrintLoop(PassCondition, PrintString)
}

func (this *FizzBuzz) PrintNumber() {
	PassCondition := func(i int) bool {
		return (i % 3 != 0) && (i % 5 != 0)
	}
	PrintString := func(i int) {
		fmt.Printf("%d, ", i)
	}

	this.PrintLoop(PassCondition, PrintString)
}

func main() {
	start := time.Now()

	for testCase := 0; testCase <= 20; testCase++ {

		fizzbuzz := &FizzBuzz{
			n:           testCase,
			wg:          &sync.WaitGroup{},
			chanVar: make(chan int, 1),
		}

		fizzbuzz.wg.Add(4)
		go fizzbuzz.PrintFizz()
		go fizzbuzz.PrintBuzz()
		go fizzbuzz.PrintFizzBuzz()
		go fizzbuzz.PrintNumber()
		fizzbuzz.chanVar <- 0
		fizzbuzz.wg.Wait()
		close(fizzbuzz.chanVar)

		fmt.Println()
	}

	spentTime := time.Now().Sub(start)
	fmt.Println("Spent time:", spentTime)
}