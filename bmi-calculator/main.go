package main

import (
	"bmi-calculator/CalculatorMethods"
	"bufio"
	"os"
)

func main() {
	kb := bufio.NewReader(os.Stdin)
	var name string
	var height, weight float64

	for {

		name, _ = CalculatorMethods.ReadName(kb)
		height, _ = CalculatorMethods.ReadInfo(kb, "height")
		weight, _ = CalculatorMethods.ReadInfo(kb, "weight")
		bmi, _ := CalculatorMethods.CalcBMI(height, weight)
		CalculatorMethods.DisplayResults(name, height, weight, bmi)

		if redo, _ := CalculatorMethods.GoAgain(kb); !redo {
			break
		}
	}
}
