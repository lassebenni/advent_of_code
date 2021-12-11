package submarine

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func differences(measures []int) map[string]int {
	var res map[string]int = make(map[string]int)

	var prev int = measures[0]
	measures = measures[1:]
	for _, v := range measures {
		if v > prev {
			res["larger"] += 1
		} else if v < prev {
			res["smaller"] += 1
		} else {
			res["equal"] += 1
		}
		prev = v
	}
	return res
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

func slidingWindow(measurements []int, size int) []int {
	var sliding_measurements []int

	for i := 0; i < len(measurements)-1; i++ {
		total := 0

		// if 0+3 is larger than (3) elements left in list
		if i+size > len(measurements) {
			break
		}
		for j := 0 + i; j < i+size; j++ {
			total += measurements[j]
		}
		sliding_measurements = append(sliding_measurements, total)
	}

	return sliding_measurements
}

func ScanFloor() {
	measurements := readMeasurements("measures.txt")
	measurements = slidingWindow(measurements, 3)
	diff := differences(measurements)
	fmt.Println(diff)
}
