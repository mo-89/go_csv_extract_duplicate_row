package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: <program_name> <path_to_csv>")
		return
	}
	csvFileName := os.Args[1]

	startTime := time.Now()

	file, err := os.Open(csvFileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	// 電話番号の出現回数をカウントするマップ
	phoneCounts := make(map[string]int)

	for _, line := range lines[1:] { // ヘッダーを除外
		for _, phone := range line[2:6] {
			phoneCounts[phone]++
		}
	}

	headers := lines[0]
	duplicatedRows := [][]string{headers}

	for _, line := range lines[1:] {
		for _, phone := range line[2:6] {
			if phoneCounts[phone] > 1 {
				duplicatedRows = append(duplicatedRows, line)
				break
			}
		}
	}

	outFile, err := os.Create(fmt.Sprintf("answer_%s.csv", time.Now().Format("20060102150405")))
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	writer := csv.NewWriter(outFile)
	defer writer.Flush()

	writer.WriteAll(duplicatedRows)

	elapsedTime := time.Since(startTime)
	fmt.Printf("Elapsed Time: %v\n", elapsedTime)
}
