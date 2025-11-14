// Console-based cryptographic password generator
package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"strconv"
)

const (
	asciiLetters     = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	asciiDigits      = "0123456789"
	asciiPunctuation = "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"
)

func printHelp() {
	fmt.Println("Password Generator - creates secure random passwords")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  passgen                    # Generate one 14-character password")
	fmt.Println("  passgen <length>           # Generate one password with specified length")
	fmt.Println("  passgen <length> <count>   # Generate multiple passwords")
	fmt.Println("  passgen -h                 # Show this help message")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("  <length>  - Password length (minimum: 8 characters)")
	fmt.Println("  <count>   - Number of passwords to generate")
	fmt.Println("  -h        - Show help")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  passgen 16                 # Generate one 16-character password")
	fmt.Println("  passgen 12 5               # Generate five 12-character passwords")
}

// secureRandomInt generates a cryptographically secure random number in [0, max)
func secureRandomInt(max int) (int, error) {
	if max <= 0 {
		return 0, fmt.Errorf("max must be positive")
	}
	n, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		return 0, err
	}
	return int(n.Int64()), nil
}

// secureRandomChoice securely selects a random character from a string
func secureRandomChoice(chars string) (rune, error) {
	if len(chars) == 0 {
		return 0, fmt.Errorf("empty character set")
	}
	idx, err := secureRandomInt(len(chars))
	if err != nil {
		return 0, err
	}
	return []rune(chars)[idx], nil
}

func passwordGenerator(symbols string, length int, quantity int) error {
	lettersAndDigits := asciiLetters + asciiDigits

	for ; quantity > 0; quantity-- {
		// Build password securely
		password := make([]rune, length)

		// First character - must be a letter or digit
		firstChar, err := secureRandomChoice(lettersAndDigits)
		if err != nil {
			return fmt.Errorf("failed to generate first character: %v", err)
		}
		password[0] = firstChar

		// Remaining characters - any from available symbols
		for i := 1; i < length; i++ {
			char, err := secureRandomChoice(symbols)
			if err != nil {
				return fmt.Errorf("failed to generate character at position %d: %v", i, err)
			}
			password[i] = char
		}

		fmt.Println(string(password))
	}
	return nil
}

func main() {
	asciiPrintable := asciiLetters + asciiDigits + asciiPunctuation
	args := os.Args

	// Check for help flag
	if len(args) == 2 && (args[1] == "-h" || args[1] == "--help") {
		printHelp()
		return
	}

	switch len(args) {
	case 1:
		// No arguments - generate one 14-character password
		if err := passwordGenerator(asciiPrintable, 14, 1); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	case 2:
		// One argument - password length
		if length, err := strconv.Atoi(args[1]); err == nil {
			if length < 8 {
				fmt.Println("Error: minimum password length is 8 characters")
				os.Exit(1)
			}
			if err := passwordGenerator(asciiPrintable, length, 1); err != nil {
				fmt.Printf("Error: %v\n", err)
				os.Exit(1)
			}
		} else {
			fmt.Println("Error: password length must be a number")
			os.Exit(1)
		}
	case 3:
		// Two arguments - length and count
		if length, err := strconv.Atoi(args[1]); err == nil {
			if length < 8 {
				fmt.Println("Error: minimum password length is 8 characters")
				os.Exit(1)
			}
			if count, err := strconv.Atoi(args[2]); err == nil {
				if err := passwordGenerator(asciiPrintable, length, count); err != nil {
					fmt.Printf("Error: %v\n", err)
					os.Exit(1)
				}
			} else {
				fmt.Println("Error: password count must be a number")
				os.Exit(1)
			}
		} else {
			fmt.Println("Error: password length must be a number")
			os.Exit(1)
		}
	default:
		// Too many arguments
		fmt.Println("Error: too many arguments")
		fmt.Println("Use 'passgen -h' for help")
		os.Exit(1)
	}
}
