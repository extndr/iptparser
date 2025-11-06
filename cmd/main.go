package main

import (
	"fmt"
	"log"

	"github.com/extndr/iptparser/internal/parser"
)

func main() {
	rules, err := parser.GetDNATRules()
	if err != nil {
		log.Fatal(err)
	}

	if len(rules) == 0 {
		fmt.Println("No DNAT rules found.")
		return
	}

	for _, r := range rules {
		fmt.Printf("%d → %s\n", r.DPort, r.Dest)
	}
}
