package CalculatorMethods

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func CalcBMI(height, weight float64) (float64, error) {
	if height <= 0 || weight <= 0 {
		return 0, errors.New("height or weight is less than or equal to 0")
	}

	return weight / math.Pow(height, 2) * 703, nil
}

func GoAgain(kb *bufio.Reader) (bool, error) {
	var res bool

	if kb == nil {
		return false, errors.New("kb is nil")
	}

	for {
		fmt.Print("Would you like to calculate another BMI (Yes/No) ")
		input, err := kb.ReadString('\n')

		if err != nil {
			return false, err
		}

		input = strings.ToLower(strings.TrimSuffix(input, "\n"))

		if input == "yes" {
			res = true
			break
		}

		if input == "no" {
			res = false
			break
		}
	}

	return res, nil
}

func ReadName(kb *bufio.Reader) (string, error) {
	if kb == nil {
		return "", errors.New("kb is nil")
	}

	fmt.Print("Please enter the person's name ")
	name, err := kb.ReadString('\n')

	return strings.TrimSuffix(name, "\n"), err
}

func ReadInfo(kb *bufio.Reader, typ string) (float64, error) {
	if kb == nil {
		return 0, errors.New("kb is nil")
	}

	if typ == "" {
		return 0, errors.New("typ is empty")
	}

	fmt.Printf("Please enter your %s ", typ)
	str, err := kb.ReadString('\n')

	if err != nil {
		return 0, err
	}

	data, err := strconv.ParseFloat(strings.TrimSuffix(str, "\n"), 64)
	return data, err
}

func DisplayResults(name string, height, weight, bmi float64) {
	var res string

	if bmi < 18.5 {
		res = "underweight"
	} else if bmi < 25 {
		res = "normal"
	} else if bmi < 30 {
		res = "overweight"
	} else {
		res = "obese"
	}

	fmt.Printf("%s with a weight of %.6f and a height of %.6f, your BMI is %.6f you are %s.\n\n", name, weight, height, bmi, res)
}
