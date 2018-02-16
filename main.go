package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/olekukonko/tablewriter"
)

var repeat = 1

func calculate(n float64, wg *sync.WaitGroup, fun string, c chan []string) {
	defer wg.Done()

	start := time.Now()
	fmt.Printf("%v%s", fun, "\tStart\n")
	c <- []string{fun, "Start", "", ""}
	result := 0.0
	for i := 0; i < repeat; i++ {
		switch fun {
		case "Pow":
			result = math.Pow(n, 2)
		case "Sqrt":
			result = math.Sqrt(n)
		case "Sin":
			result = math.Sin(n)
		case "Cos":
			result = math.Cos(n)
		case "Tan":
			result = math.Tan(n)
		}
	}
	end := time.Now()
	fmt.Printf("%v%s%v%s%v%s", fun, "\tDone\t\t", end.Sub(start), "\t\t", result, "\n")
	c <- []string{fun, "Done", end.Sub(start).String(), strconv.FormatFloat(result, 'f', -1, 64)}
}

func calculateNormal(n float64, fun string, c chan []string) {
	start := time.Now()
	fmt.Printf("%v%s", fun, "\tStart\n")
	c <- []string{fun, "Start", "", ""}
	result := 0.0
	for i := 0; i < repeat; i++ {
		switch fun {
		case "Pow":
			result = math.Pow(n, 2)
		case "Sqrt":
			result = math.Sqrt(n)
		case "Sin":
			result = math.Sin(n)
		case "Cos":
			result = math.Cos(n)
		case "Tan":
			result = math.Tan(n)
		}
	}
	end := time.Now()
	fmt.Printf("%v%s%v%s%v%s", fun, "\tDone\t\t", end.Sub(start), "\t\t", result, "\n")
	c <- []string{fun, "Done", end.Sub(start).String(), strconv.FormatFloat(result, 'f', -1, 64)}
}

func main() {
	input := 1
	number := 0.0
	for input == 1 || input == 2 {
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"func", "action", "time", "result"})
		c := make(chan []string, 10)
		fmt.Print("Type\n")
		fmt.Print("1.Multithread\n2.Singlethread\nPlease select type (1/2, 0 exit): ")
		fmt.Scanf("%d", &input)
		if input != 1 && input != 2 {
			break
		}
		fmt.Print("Number : ")
		fmt.Scanf("%f", &number)
		fmt.Print("Number of times to repeat :")
		fmt.Scanf("%d", &repeat)
		// fmt.Print("====================================================\n")
		// fmt.Print("func\taction\t\ttime\t\tresult\n")
		// fmt.Print("====================================================\n")
		switch input {
		case 1:
			var wg sync.WaitGroup
			wg.Add(5)
			start := time.Now()
			go calculate(number, &wg, "Pow", c)
			go calculate(number, &wg, "Sqrt", c)
			go calculate(number, &wg, "Sin", c)
			go calculate(number, &wg, "Cos", c)
			go calculate(number, &wg, "Tan", c)
			wg.Wait()
			end := time.Now()
			fmt.Print("====================================================\n")
			fmt.Println("Done", end.Sub(start))
		case 2:
			start := time.Now()
			calculateNormal(number, "Pow", c)
			calculateNormal(number, "Sqrt", c)
			calculateNormal(number, "Sin", c)
			calculateNormal(number, "Cos", c)
			calculateNormal(number, "Tan", c)
			end := time.Now()
			fmt.Print("====================================================\n")
			fmt.Println("Done", end.Sub(start))
		}
		close(c)
		for i := range c {
			table.Append(i)
		}
		table.Render()
	}
}
