package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type bingoSheet struct {
	rows     [][]string
	columns  [][]string
	allItems []string
}

func (receiver *bingoSheet) hasBingo(data []string) bool {
	return hasBingo(data, receiver.columns) || hasBingo(data, receiver.rows)
}

func hasBingo(data []string, fields [][]string) bool {
	for _, row := range fields {
		matches := 0
		for _, item := range row {
			for _, bingoItem := range data {
				if bingoItem == item {
					matches++
				}
			}
		}
		if matches == 5 {
			return true
		}
	}
	return false
}

func (receiver *bingoSheet) computeScore(data []string) int {
	var filteredItems []string
	for _, item := range receiver.allItems {
		if !contains(data, item) {
			filteredItems = append(filteredItems, item)
		}
	}

	sum := 0
	for _, item := range filteredItems {
		atoi, _ := strconv.Atoi(item)
		sum += atoi
	}
	lastNumber, _ := strconv.Atoi(data[len(data)-1])
	return lastNumber * sum
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func main() {
	data, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(data), "\n\n")

	bingoNumbers := strings.Split(lines[0], ",")

	var sheets []bingoSheet
	for _, unparsedSheet := range lines[1:] {
		unparsedRows := strings.Split(unparsedSheet, "\n")

		var rows [][]string
		for _, row := range unparsedRows {
			parsedRow := strings.Fields(row)
			rows = append(rows, parsedRow)
		}
		var columns [][]string
		for i := 0; i < len(rows[0]); i++ {
			var column []string
			for j := 0; j < len(rows); j++ {
				column = append(column, rows[j][i])
			}
			columns = append(columns, column)
		}

		var allItems []string
		for _, row := range rows {
			allItems = append(allItems, row...)
		}

		sheet := bingoSheet{
			allItems: allItems,
			rows:     rows,
			columns:  columns,
		}
		sheets = append(sheets, sheet)
	}

	fmt.Println("Part1: ", day1(bingoNumbers, sheets))
	fmt.Println("Part2: ", day2(bingoNumbers, sheets))
}

func day1(bingoNumbers []string, sheets []bingoSheet) int {
	for i := range bingoNumbers {
		toValidate := bingoNumbers[0:i]

		for _, sheet := range sheets {
			if sheet.hasBingo(toValidate) {
				return sheet.computeScore(toValidate)
			}
		}
	}
	return 0
}

func day2(bingoNumbers []string, sheets []bingoSheet) int {
	var lastScore int
	maxLength := 0
	for _, sheet := range sheets {
		for i := range bingoNumbers {
			toValidate := bingoNumbers[0:i]
			if sheet.hasBingo(toValidate) {
				if maxLength < len(toValidate) {
					maxLength = len(toValidate)
					lastScore = sheet.computeScore(toValidate)
				}
				break
			}
		}
	}
	return lastScore
}
