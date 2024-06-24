package main

import (
	"bufio"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

const envFilePath = ".env"

func generateRandomKey(length int) (string, error) {
	bytes := make([]byte, length/2)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func main() {
	randomKey, err := generateRandomKey(32)
	if err != nil {
		fmt.Println("Error generating random key:", err)
		return
	}

	file, err := os.Open(envFilePath)
	if err != nil {
		fmt.Println("Error opening .env file:", err)
		return
	}
	defer file.Close()

	var lines []string
	var secretKeyFound bool
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "JWT_SECRET_KEY=") {
			line = fmt.Sprintf("JWT_SECRET_KEY=%s", randomKey)
			secretKeyFound = true
		}
		lines = append(lines, line)
	}
	if !secretKeyFound {
		lines = append(lines, fmt.Sprintf("JWT_SECRET_KEY=%s", randomKey))
	}

	file, err = os.Create(envFilePath)
	if err != nil {
		fmt.Println("Error creating .env file:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			fmt.Println("Error writing to .env file:", err)
			return
		}
	}
	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flushing to .env file:", err)
		return
	}

	fmt.Println(".env file updated successfully with random JWT_SECRET_KEY:", randomKey)
}
