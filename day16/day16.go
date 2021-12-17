package main

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

const (
	SUM     = "000"
	PRODUCT = "001"
	MIN     = "010"
	MAX     = "011"
	LITERAL = "100"
	GREATER = "101"
	LESSER  = "110"
	EQUALS  = "111"
)

func main() {
	data, _ := ioutil.ReadFile("input.txt")
	bitsTransmission := string(data)
	decodeString, _ := hex.DecodeString(bitsTransmission)
	byteString := strings.ReplaceAll(fmt.Sprintf("%08b", decodeString), " ", "")
	byteString = byteString[1 : len(byteString)-1]

	_, result := decodePackage(byteString)
	fmt.Println("Part 1:", totalVersion)
	fmt.Println("Part 2:", result)
}

var totalVersion int64 = 0

func decodePackage(bitsPackage string) (string, int64) {
	version := bitsPackage[0:3]
	typePackage := bitsPackage[3:6]
	bitsPackage = bitsPackage[6:]

	versionNumber, _ := strconv.ParseInt(version, 2, 64)
	totalVersion += versionNumber

	if typePackage == LITERAL {
		return literalPacket(bitsPackage)
	}
	lengthType := bitsPackage[:1]
	if lengthType == "0" {
		return operatorMode0(bitsPackage[1:], typePackage)
	}
	return operatorMode1(bitsPackage[1:], typePackage)
}

func literalPacket(bitsPackage string) (string, int64) {
	buildNumberString := ""
	numberBits := bitsPackage
	total := 0
	for true {
		total += 5
		number := numberBits[0:5]
		buildNumberString += number[1:5]
		if number[0:1] == "0" {
			break
		}
		numberBits = numberBits[5:]
	}
	numberFound, _ := strconv.ParseInt(buildNumberString, 2, 64)
	return bitsPackage[total:], numberFound
}

func operatorMode0(bitsPackage string, typePackage string) (string, int64) {
	totalLengthInBits, _ := strconv.ParseInt(bitsPackage[:15], 2, 64)
	subPackets := bitsPackage[15:]
	tmp := subPackets[:totalLengthInBits]

	var results []int64

	for len(tmp) > 0 {
		subpart, result := decodePackage(tmp)
		tmp = subpart
		results = append(results, result)
	}
	return subPackets[totalLengthInBits:], computeResult(results, typePackage)
}

func operatorMode1(bitsPackage string, typePackage string) (string, int64) {
	numberOfSubPackages, _ := strconv.ParseInt(bitsPackage[:11], 2, 64)
	subPackets := bitsPackage[11:]

	var results []int64
	for i := 0; i < int(numberOfSubPackages); i++ {
		subpart, result := decodePackage(subPackets)
		subPackets = subpart
		results = append(results, result)
	}
	return subPackets, computeResult(results, typePackage)
}

func computeResult(results []int64, mode string) int64 {
	if mode == SUM {
		var sum int64 = 0
		for _, result := range results {
			sum += result
		}
		return sum
	}
	if mode == PRODUCT {
		var product int64 = 1
		for _, result := range results {
			product *= result
		}
		return product
	}
	if mode == MIN {
		var min int64 = math.MaxInt64
		for _, result := range results {
			if result < min {
				min = result
			}
		}
		return min
	}
	if mode == MAX {
		var max int64 = 0
		for _, result := range results {
			if result > max {
				max = result
			}
		}
		return max
	}
	if mode == GREATER {
		if results[0] > results[1] {
			return 1
		}
		return 0
	}
	if mode == LESSER {
		if results[0] < results[1] {
			return 1
		}
		return 0
	}
	if mode == EQUALS {
		if results[0] == results[1] {
			return 1
		}
		return 0
	}
	return 0
}
