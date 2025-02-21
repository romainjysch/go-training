package go_by_example

import (
	"fmt"
	"math"
	"time"
)

const s string = "constant"

func GbeBase64Encoding() {

}

func GbeArrays() {
	fmt.Println("GbeArrays")
}

func GbeSwitch() {
	i := 2
	fmt.Println("i: ", i)
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	fmt.Println(time.Now().Weekday())
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it's the weekend")
	default:
		fmt.Println("it's a weekday")
	}

	t := time.Now()
	fmt.Println(t)
	switch {
	case t.Hour() < 12:
		fmt.Println("it's before noon")
	default:
		fmt.Println("it's after noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("bool")
		case int:
			fmt.Println("int")
		case string:
			fmt.Println("string")
		default:
			fmt.Println("unknown", t)
		}
	}
	whatAmI(true)
	whatAmI("hey")

}

func GbeIfElse() {
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is even")
	}

	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("either 8 or 7 is odd")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "is less than 10")
	} else {
		fmt.Println(num, "is greater than 10")
	}
}

func GbeFor() {
	fmt.Println("first for")
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	fmt.Println("second for")
	for j := 1; j < 5; j++ {
		fmt.Println(j)
	}

	fmt.Println("third for")
	for i := range 3 {
		fmt.Println(i) // starts at 0, then 1 and 2
	}

	fmt.Println("fourth for")
	for {
		fmt.Println("loop")
		break
	}

	fmt.Println("fifth for")
	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}

func GbeConstants() {
	fmt.Println(s)

	const n = 5000000000000
	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}

func GbeVariables() {
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	d := "short"
	fmt.Println(d)
}

func GbeValues() {
	fmt.Println("go" + "lang")
	fmt.Println(7 / 3.0)
	fmt.Println(6 / 3)
	fmt.Println("1+1=", 1+1)
	fmt.Println(true && false)
	fmt.Println(true || false)
}
