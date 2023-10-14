package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func generatePhoneNumber() string {
	return fmt.Sprintf("090%08d", rand.Intn(100000000))
}

func main() {
	rand.Seed(time.Now().UnixNano()) // ランダムシードの初期化

	file, err := os.Create("sample_data.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	headers := []string{"id", "user_name", "call_1", "call_2", "call_3", "call_4", "memo"}
	writer.Write(headers)

	for i := 1; i <= 1400000; i++ {
		row := []string{
			fmt.Sprintf("%d", i),
			"user" + fmt.Sprintf("%d", i),
			generatePhoneNumber(),
			generatePhoneNumber(),
			generatePhoneNumber(),
			generatePhoneNumber(),
			"memo" + fmt.Sprintf("%d", i),
		}
		writer.Write(row)
	}
}
