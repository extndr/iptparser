package parser

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"sort"
	"strconv"
	"strings"
)

type Rule struct {
	DPort int
	Dest  string
}

func GetDNATRules() ([]Rule, error) {
	if os.Geteuid() != 0 {
		return nil, fmt.Errorf("this tool requires root privileges")
	}

	cmd := exec.Command("iptables-save", "-t", "nat")
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	sc := bufio.NewScanner(bytes.NewReader(out))
	var rules []Rule

	for sc.Scan() {
		line := sc.Text()
		if strings.Contains(line, "DNAT") {
			r, ok := parse(line)
			if ok {
				rules = append(rules, r)
			}
		}
	}
	if err := sc.Err(); err != nil {
		return nil, err
	}

	sort.Slice(rules, func(i, j int) bool { return rules[i].DPort < rules[j].DPort })
	return rules, nil
}

func parse(line string) (Rule, bool) {
	fields := strings.Fields(line)
	var dport int
	var dest string

	for i, f := range fields {
		if f == "--dport" && i+1 < len(fields) {
			p, err := strconv.Atoi(fields[i+1])
			if err != nil {
				return Rule{}, false
			}
			dport = p
		}
		if f == "--to-destination" && i+1 < len(fields) {
			dest = fields[i+1]
		}
	}

	if dport == 0 || dest == "" {
		return Rule{}, false
	}
	return Rule{DPort: dport, Dest: dest}, true
}
