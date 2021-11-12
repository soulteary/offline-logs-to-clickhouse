package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	monthLabels := []string{"\"Jan\"", "\"Feb\"", "\"Mar\"", "\"Apr\"", "\"May\"", "\"Jun\"", "\"Jul\"", "\"Aug\"", "\"Sep\"", "\"Oct\"", "\"Nov\"", "\"Dec\""}
	monthValues := []string{"\"01\"", "\"02\"", "\"03\"", "\"04\"", "\"05\"", "\"06\"", "\"07\"", "\"08\"", "\"09\"", "\"10\"", "\"11\"", "\"12\""}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		src := scanner.Text()
		dest := ""
		for i, monthLabel := range monthLabels {
			if strings.Contains(src, monthLabel) {
				dest = strings.Replace(src, monthLabel, monthValues[i], -1)
				break
			}
		}
		fmt.Println(dest)
	}
}
