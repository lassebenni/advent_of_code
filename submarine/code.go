package submarine

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func printDifferences(measures []int) {
	var prev int = measures[0]
	var larger int = 0

	measures = measures[1:]
	for _, v := range measures {
		var diff = v - prev
		if diff > 0 {
			fmt.Println("increased by ", diff)
			larger += 1
		} else if diff < 0 {
			fmt.Println("decreased by ", diff)
		} else {
			fmt.Println("no change")
		}
		prev = v
	}
	fmt.Println("There were ", larger, "larger changes")
}

func readMeasurements(filename string) []int {
	content, err := ioutil.ReadFile("measures.txt")

	if err != nil {
		log.Fatal(err)
	}

	measures_str := strings.Split(string(content), "\n")

	var measures []int
	for _, measure := range measures_str {
		var measure_int, _ = strconv.Atoi(measure)
		measures = append(measures, measure_int)
	}
	return measures
}

func ScanFloor() {
	measurements := readMeasurements("measures.txt")
	printDifferences(measurements)
}
