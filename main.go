package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"sync"
)

var wg sync.WaitGroup

const INSTCOUNT = 10

func fetch(ch chan string) {
	defer wg.Done()
	for i := 0; i < INSTCOUNT; i++ {
		opd1 := rand.Intn(1000)
		opd2 := rand.Intn(1000) + 1

		op := "+"

		t := rand.Intn(4)

		if t == 0 {
			op = "+"
		} else if t == 1 {
			op = "-"
		} else if t == 2 {
			op = "*"
		} else {
			op = "/"
		}

		inst := strconv.Itoa(opd1) + "," + strconv.Itoa(opd2) + "," + op

		ch <- inst
	}
	close(ch)
}

func execute(in chan string, out chan string) {
	defer wg.Done()
	for i := 0; i < INSTCOUNT; i++ {
		inst := <-in
		ops := strings.Split(inst, ",")

		op1, _ := strconv.Atoi(ops[0])
		op2, _ := strconv.Atoi(ops[1])
		op := ops[2]

		var res int

		if op == "+" {
			res = op1 + op2
		} else if op == "-" {
			res = op1 - op2
		} else if op == "*" {
			res = op1 * op2
		} else {
			res = op1 / op2
		}
		out <- inst + " = " + strconv.Itoa(res)
	}
}

func write(in chan string) {
	defer wg.Done()
	for i := 0; i < INSTCOUNT; i++ {
		res := <-in
		fmt.Println(res)
	}
}

func main() {
	chan1 := make(chan string)
	chan2 := make(chan string)

	wg.Add(3)
	go fetch(chan1)
	go execute(chan1, chan2)
	go write(chan2)
	wg.Wait()
}
