package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func run() (err error) {
	stdin, _ := ioutil.ReadAll(os.Stdin)
	lines := strings.Split(string(stdin), "\n")

	t := ""
	result := map[string]map[string]string{}

	for _, l := range lines {
		if l == "" {
			continue
		}

		fields := strings.Fields(l)
		parts := strings.Split(fields[0], ".")
		t = fields[2]

		if _, ok := result[parts[0]]; !ok {
			result[parts[0]] = make(map[string]string)
		}
		result[parts[0]][strings.Join(parts[1:], ".")] = fields[1]

	}

	for key, values := range result {
		fmt.Printf("[%s]\n", key)
		for k, v := range values {
			fmt.Println(k, v)
		}
	}

	fmt.Printf("at %s\n", t)
	return
}

func main() {
	err := run()
	if err != nil {
		panic(err)
	}
}
