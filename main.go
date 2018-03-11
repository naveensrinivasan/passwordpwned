package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/mattevans/pwned-passwords"
)

func main() {
	client := hibp.NewClient()
	passwd := flag.String("password", "pwd.csv", "chrome password csv export")
	flag.Parse()
	file, err := os.Open(*passwd)
	if err != nil {
		log.Fatal(err)

	}
	defer file.Close()

	reader := csv.NewReader(file)
	m := make(map[string]int)
	rows, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	for _, p := range rows {
		// Check to see if your given string is compromised.
		pwned, e := client.Pwned.Compromised(p[3])
		if err != nil {
			log.Fatal(e)
		}
		if pwned {
			if _, ok := m[p[3]]; ok {
				m[p[3]]++
			} else {
				m[p[3]] = 1
			}
		}
	}
	sortMap(m)
	for no := range m {
		fmt.Println(no)
	}

}
func sortMap(m map[string]int) {
	n := map[int][]string{}
	var a []int
	for k, v := range m {
		n[v] = append(n[v], k)
	}
	for k := range n {
		a = append(a, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	for _, k := range a {
		for _, s := range n[k] {
			fmt.Printf("%s, %d\n", s, k)
		}
	}
}
