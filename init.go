// Flag parsing
// Implementation of Set method for Targets and Numbers
package main

import (
	"errors"
	"flag"
	"strconv"
	"strings"
)

func (t *Targets) Set(value string) error {
	if len(*t) > 0 {
		return errors.New("Targets already set")
	}
	for _, tt := range strings.Split(value, ",") {
		target, err := strconv.Atoi(tt)
		if err != nil {
			return err
		}
		*t = append(*t, cInt(target))
	}
	return nil
}

func (n *Numbers) Set(value string) error {
	if len(*n) > 0 {
		return errors.New("Targets already set")
	}
	for _, tn := range strings.Split(value, ",") {
		target, err := strconv.Atoi(tn)
		if err != nil {
			return err
		}
		*n = append(*n, cUint(target))
	}
	return nil
}

func init() {
	flag.Var(&targets, "targets", "comma separated list of targets")
	flag.Var(&numbers, "numbers", "comma separated list of numbers")
}
