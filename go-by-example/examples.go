package go_by_example

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"math"
	"net/http"
	"os"
	"os/exec"
	"slices"
	"sync"
	"syscall"
	"time"
)

const s string = "constant"

func GbeExecingProcesses() {
	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}

	args := []string{"ls", "-l", "-a", "/Users/rjh"}

	env := os.Environ()

	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}

func GbeHTTPServer() {
	hello := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello\n")
	}

	headers := func(w http.ResponseWriter, r *http.Request) {
		for name, headers := range r.Header {
			for _, h := range headers {
				fmt.Fprintf(w, "%v: %v\n", name, h)
			}
		}
	}

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	http.ListenAndServe(":8090", nil)
}

func GbeHTTPClient() {
	resp, err := http.Get("https://gobyexample.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)

	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func GbeSHA256Hashes() {
	s := "string"
	hash := sha256.Sum256([]byte(s))
	fmt.Printf("%x\n", hash)
}

func GbeTime() {
	p := fmt.Println

	now := time.Now()
	p(now)

	nowUTC := now.UTC()
	p(nowUTC)

	then := time.Date(
		2009, 10, 31, 12, 12, 12, 0, time.UTC)
	p(then)

	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Weekday())

	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	diff := now.Sub(then)
	p(diff)

	p(diff.Hours(), "hours")
	p(diff.Hours()/24, "days")
	p(diff.Hours()/24/365, "years")
}

func GbePanic() {
	//panic("a problem")

	_, err := os.Open("/tmp/file")
	if err != nil {
		panic(err)
	}
}

func GbeWaitGroups() {
	worker := func(id int) {
		fmt.Printf("Worker %d starting \n", id)
		time.Sleep(time.Second)
		fmt.Printf("Worker %d done \n", id)
	}

	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			worker(i)
		}()
	}

	wg.Wait()
}

func GbeDefer() {
	createFile := func(p string) *os.File {
		fmt.Println("creating")
		f, err := os.Create(p)
		if err != nil {
			panic(err)
		}
		return f
	}

	writeFile := func(f *os.File) {
		fmt.Println("writing")
		fmt.Fprintln(f, "data")
	}

	closeFile := func(f *os.File) {
		fmt.Println("closing")
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}

	f := createFile("/tmp/defer.txt")
	defer closeFile(f)
	writeFile(f)
}

func GbeGoroutines() {
	f := func(from string) {
		for i := 0; i < 3; i++ {
			fmt.Println(from, ":", i)
		}
	}

	f("direct")

	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	fmt.Println("done")
}

func GbeClosures() {
	intSeq := func() func() int {
		i := 0
		return func() int {
			i++
			return i
		}
	}

	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())
}

func GbeAnonymousFunctions() {
	var counter = 1

	func(str string) {
		fmt.Println(str)
	}("romain")

	funcVar := func(str string) {
		fmt.Println(str)
	}
	funcVar("romain")

	closure := func(str string) {
		fmt.Println(str, "inside a closure")
		for i := 1; i <= 3; i++ {
			fmt.Println(counter, "counter")
			counter++
		}
	}

	fmt.Println("counter before closure", counter)
	closure("romain")
	fmt.Println("counter after closure", counter)
}

func GbeSlices() {
	var s []string
	fmt.Println("empty:", s, "len:", len(s), "cap:", cap(s))

	s = make([]string, 3)
	fmt.Println("empty:", s, "len:", len(s), "cap:", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("append:", s)
	fmt.Println("len s:", len(s))

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copied:", c)

	l := s[2:5] // include 2, exclude 5
	fmt.Println("sl1:", l)

	l = s[:5] // everything from start, exclude 5 and more
	fmt.Println("sl2:", l)

	l = s[2:] // from 3rd index to the end
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("declared:", t)

	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) { // equals
		fmt.Println("t == t2")
	}
}

func GbeArrays() {
	var a [5]int
	fmt.Println("empty", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("declared", b)

	c := [...]int{1, 2, 3, 4, 5}
	fmt.Println("declared", c)

	d := [...]int{100, 3: 400, 500} // starts again at index 3
	fmt.Println("idx:", d)
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
