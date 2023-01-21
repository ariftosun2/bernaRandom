package main

import (
	"fmt"
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func randomGulme(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func gulmeYazisi(length int) string {
	return randomGulme(length, charset)
}

func main() {
	var bernaSayi int
	bernaSayi = 14

	bernaSabitHarf := "Ç"
	var bernaRandomGulmesi string

	bernaRandomGulmesi = gulmeYazisi(bernaSayi)

	fmt.Println(bernaSabitHarf + bernaRandomGulmesi)

}
