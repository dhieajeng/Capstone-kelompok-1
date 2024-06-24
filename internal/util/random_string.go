package util

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func RandomStringGenerator(length int) (code string) {
	var randomizer = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	var letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")

	c := make([]string, length)
	for i := range c {
		numOrAlpha := rand.Intn(2)
		if numOrAlpha == 0 {
			c[i] = strconv.Itoa(randomizer.Intn(10))
		} else {
			c[i] = string(letters[randomizer.Intn(len(letters))])
		}

		code = strings.Join(c, "")
	}
	return
}
