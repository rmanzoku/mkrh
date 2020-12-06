package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	n         = 1
	preformat = true
)

func run() (err error) {
	stdin, _ := ioutil.ReadAll(os.Stdin)
	lines := strings.Split(string(stdin), "\n")

	var t time.Time
	result := make(map[string]map[string]string)

	for _, l := range lines {
		if l == "" {
			continue
		}

		fields := strings.Fields(l)
		parts := strings.Split(fields[0], ".")
		ti, _ := strconv.Atoi(fields[2])
		t = time.Unix(int64(ti), 0)

		subject := strings.Join(parts[0:n], ".")
		key := strings.Join(parts[n:], ".")

		if _, ok := result[subject]; !ok {
			result[subject] = make(map[string]string)
		}
		result[subject][key] = fields[1]

	}

	if preformat {
		fmt.Println("```")
	}

	fmt.Printf("at %s\n", t)
	for subject, values := range result {
		fmt.Printf("[%s]\n", subject)

		total := 0.0
		cnt := 0
		for k, v := range values {
			f, _ := strconv.ParseFloat(v, 84)
			total += f
			cnt++

			fmt.Println(k, v)
		}
		if cnt > 1 {
			fmt.Println("Total: ", total)
		}
	}

	if preformat {
		fmt.Println("```")
	}
	return
}

func main() {
	flag.BoolVar(&preformat, "p", true, "preformatted ```")
	flag.IntVar(&n, "n", 1, "field number")
	flag.Parse()
	err := run()
	if err != nil {
		panic(err)
	}
}
