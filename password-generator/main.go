package main

import (
	"bufio"
	"fmt"
	"math/rand"
	_ "math/rand" // The blank identifier (_) prevents the "unused import" error until you use it
	"os"
	"strconv"
	"strings"
)

const (
	lowerChars   = "abcdefghijklmnopqrstuvwxyz"
	upperChars   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digitChars   = "0123456789"
	specialChars = "!@#$%^&*()-_=+[]{}|;:,.<>?"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to the Random Password Generator")

	// Interactive Prompts (Already built for you):
	length := readInt(reader, "Enter password length (min 8) [Default 10]: ", 10)
	if length < 8 {
		fmt.Println("Length too short, defaulting to 8.")
		length = 8
	}

	useUpper := readBool(reader, "Include uppercase letters? (y/n) [Default y]: ", true)
	useDigits := readBool(reader, "Include digits? (y/n) [Default y]: ", true)
	useSymbols := readBool(reader, "Include symbols? (y/n) [Default n]: ", false)

	count := readInt(reader, "How many passwords to generate? [Default 1]: ", 1)
	if count < 1 {
		count = 1
	}

	outPath := readString(reader, "Save to file? (Enter filename) [Default: output_1.txt]: ", "output_1.txt")

	fmt.Println("\nGenerating...")

	// Execute Core Logic
	err := generateAndSavePasswords(count, length, useUpper, useDigits, useSymbols, outPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Generation Error: %v\n", err)
		os.Exit(1)
	}
}

// --- Core Cryptographic Logic (YOUR TURN) ---

func randomCharFromSet(set string) byte {
	// TODO 1: Pick a random index from the 'set' string and return the character.
	idx := rand.Intn(len(set))
	return set[idx]
}

func shuffle(b []byte) {
	// TODO 2: Randomize the order of the bytes in the slice.
	rand.Shuffle(len(b), func(i, j int) {
		b[i], b[j] = b[j], b[i]
	})
	
}

func generatePassword(length int, upper, digits, symbols bool) string {
	// TODO 3: Build your character pool based on the boolean flags.
	// TODO 4: Guarantee at least one character from each selected pool is used.
	masterPool := lowerChars
	var password []byte
	if(upper){
		masterPool = masterPool + upperChars
		password = append(password, randomCharFromSet(upperChars))
	}
	if(digits){
		masterPool = masterPool + digitChars
		password = append(password, randomCharFromSet(digitChars))
	}
	if(symbols){
		masterPool = masterPool + specialChars
		password = append(password, randomCharFromSet(specialChars))
	} 

	// TODO 5: Fill the remaining length with random characters from your combined pool.
	for len(password)< length {
		password = append(password, randomCharFromSet(masterPool))
	}

	// TODO 6: Shuffle the final password so the guaranteed characters aren't always at the front.
	shuffle(password)

	return string(password)
}

func evaluateStrength(pwd string) int {
	// TODO 7: Write logic to score the password out of 5 based on its length.
	// Then use strings.ContainsAny() against upperChars, digitChars, and specialChars to check its actual complexity.
	scr := 0
	if len(pwd) >= 8{
		scr++
	}
	if len(pwd) >= 16{
		scr++
	}
	if strings.ContainsAny(pwd, specialChars){
		scr++
	}
	if strings.ContainsAny(pwd, digitChars){
		scr++
	}
	if strings.ContainsAny(pwd, upperChars){
		scr++
	}
	return scr
}

// --- High-Level Logic (YOUR TURN) ---

func generateAndSavePasswords(count, length int, upper, digits, symbols bool, outPath string) error {
	// TODO 8: Create a loop that runs 'count' times.
	// TODO 9: Inside the loop, call generatePassword() and evaluateStrength().
	// TODO 10: Format the results nicely and store them in a single string.
	// TODO 11: Write that final string to the 'outPath' file.
	var finalString string
	for i := 0; i < count; i++ {
		currPass := generatePassword(length, upper, digits, symbols)
		currScr := evaluateStrength(currPass)
		curr := fmt.Sprintf("Password: %-25s | Strength: %d/5\n", currPass, currScr)
		finalString = finalString + curr
	}
	err := os.WriteFile(outPath, []byte(finalString), 0644)
	if err != nil {
		return err
	}
	fmt.Printf("Successfully saved %d password(s) to '%s'\n", count, outPath)
	return nil
}

// --- Prompt Helpers (DO NOT MODIFY THESE FUNCTIONS) ---

func readString(reader *bufio.Reader, prompt string, defaultVal string) string {
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if input == "" {
		return defaultVal
	}
	return input
}

func readInt(reader *bufio.Reader, prompt string, defaultVal int) int {
	input := readString(reader, prompt, "")
	if input == "" {
		return defaultVal
	}
	val, err := strconv.Atoi(input)
	if err != nil {
		fmt.Printf("Invalid number, using default (%d).\n", defaultVal)
		return defaultVal
	}
	return val
}

func readBool(reader *bufio.Reader, prompt string, defaultVal bool) bool {
	input := strings.ToLower(readString(reader, prompt, ""))
	if input == "" {
		return defaultVal
	}
	return input == "y" || input == "yes"
}


