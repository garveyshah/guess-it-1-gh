package utils

import (
	"bufio"
	"fmt"
	"guess/stats"
	"math"
	"os"
	"strconv"
)



func Reader() {
	fileScan := bufio.NewScanner(os.Stdin)
	fileScan.Split(bufio.ScanLines)
	var input []float64

	fmt.Println("\033[032mEnter value (or 'exit' to finish): \033[0m")
	for fileScan.Scan() {
		line := fileScan.Text()

		// if line == "exit" {
		// 	break
		// }

		value, err := strconv.ParseFloat(line, 64)
		if err != nil {
			fmt.Printf("error parsing float from line '%v': %v", line, err)
			return
		}

		input = append(input, value)

		if len(input) > 1 {
			lowerB, upperB := stats.Guesser(input)
			fmt.Printf("%d %d\n", lowerB, upperB)
		} else if len(input) == 0 {
			fmt.Printf("%v %v\n", math.NaN(), math.NaN())
		}

		// Check if any error occurred during scanning
		if err := fileScan.Err(); err != nil {
			fmt.Printf("error readering input: %v", err)
			return
		}
	}

}
