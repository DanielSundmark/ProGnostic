package main

import (
	"fmt"
	"time"
//	"bufio"
//	"os"
	"io/ioutil"
	)

type Settings struct {

}

type Person struct {
	name string
	salary float64
	socialCostsMultiplier float64
}

type Project struct {
	name string
	startDate time.Time
	endDate time.Time
	budget float64
}

type Allocation struct {
	resource Person
	project Project
	startDate time.Time
	endDate time.Time
	percentage float64
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {

//	var	persons [0]Person
//	var projects [0]Project
//	var allocations [0]Allocation

//	reader := bufio.NewReader(os.Stdin)
	

	dat, err := ioutil.ReadFile("stl.pgs")
	check(err)
    fmt.Print(string(dat))

	// 	Read Perons, Projects and Alloctions from File 
//	for i < 10 {
 //   	fmt.Print("Enter text: ")
  //  	text, _ := reader.ReadString('\n')
 //   	fmt.Println(text)
 //   	i++
//
//    	if (text == "quit\n"){
//	    	fmt.Println("murvel")
//    		i = 10
//    	}

//	}

}

