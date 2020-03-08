// delta calculates time in format 1-11 0-05 :6 =>1-00
//colon and minus counts as separators hh-mm == hh:mm
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func convMin(v string) int {
	entry := strings.ReplaceAll(v, "-", ":") //replace 1-10->1:10
	split := strings.Split(entry, ":")
	if len(split) == 1 {
		// 10: or :10
		if strings.HasPrefix(v, ":") {
			//:10 mins
			min, _ := strconv.Atoi(v[1:])
			return min
		} else if strings.HasSuffix(v, ":") { // parse hrs
			//10: hours
			trimmed := strings.TrimSuffix(v, ":") //10
			hrs, _ := strconv.Atoi(trimmed)
			return hrs * 60
		} else {
			hrs, _ := strconv.Atoi(v)
			return hrs * 60
		}
	} else if len(split) > 1 {
		hour, _ := strconv.Atoi(split[0])
		min, _ := strconv.Atoi(split[1])
		return hour*60 + min
	} else {
		fmt.Fprintf(os.Stderr, "wrong argument spotted: %v\n", v)
		os.Exit(1)
	}
	return -int(^uint(0)>>1) - 1
}

func conv(mins int) string {
	hh := mins / 60
	mm := mins % 60
	return fmt.Sprintf("%v-%02v", hh, mm)
}

func main() {
	//read arguments
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Printf("Usage: delta 16-11 1-01 7-01\n")
		fmt.Printf("Ans:=8-09")
		os.Exit(1)
	}

	resMins := convMin(args[0])
	for _, e := range args[1:] {
		resMins -= convMin(e)
	}

	fmt.Printf("Ans:=%vmins, (%v)\n", resMins, conv(resMins))

}
