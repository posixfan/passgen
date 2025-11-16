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
	fmt.Println("Repository: https://github.com/posixfan/passgen")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  pgen                    # Generate one 14-character password")
	fmt.Println("  pgen <length>           # Generate one password with specified length")
	fmt.Println("  pgen <length> <count>   # Generate multiple passwords")
	fmt.Println("  pgen -h                 # Show this help message")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("  <length>    - Password length (minimum: 8 characters)")
	fmt.Println("  <count>     - Number of passwords to generate")
	fmt.Println("  --no-punct  - Exclude punctuation symbols")
	fmt.Println("  -h, --help  - Show help")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  pgen 16                 # Generate one 16-character password")
	fmt.Println("  pgen 12 5               # Generate five 12-character passwords")
	fmt.Println("  pgen --no-punct         # Generate password without punctuation")
	fmt.Println("  pgen 16 3 --no-punct    # Generate three 16-char passwords without punctuation")
	fmt.Println("  pgen --no-punct 16 3    # Same as above (flags can be anywhere)")
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

func parseArgs(args []string) (length int, count int, noPunct bool, help bool) {
	// Default values
	length = 14
	count = 1
	noPunct = false
	help = false

	positionalArgs := []string{}

	for _, arg := range args {
		switch arg {
		case "-h", "--help":
			help = true
			return
		case "--no-punct", "-no-punct":
			noPunct = true
		default:
			// Check if it's a positional argument (number)
			if _, err := strconv.Atoi(arg); err == nil {
				positionalArgs = append(positionalArgs, arg)
			} else {
				fmt.Printf("Error: unknown argument '%s'\n", arg)
				fmt.Println("Use 'pgen -h' for help")
				os.Exit(1)
			}
		}
	}

	// Process positional arguments
	switch len(positionalArgs) {
	case 0:
		// Use defaults
	case 1:
		if val, err := strconv.Atoi(positionalArgs[0]); err == nil {
			length = val
		}
	case 2:
		if val, err := strconv.Atoi(positionalArgs[0]); err == nil {
			length = val
		}
		if val, err := strconv.Atoi(positionalArgs[1]); err == nil {
			count = val
		}
	default:
		fmt.Println("Error: too many arguments")
		fmt.Println("Use 'pgen -h' for help")
		os.Exit(1)
	}

	return
}

func main() {
	args := os.Args[1:] // Skip program name

	// Special case: no arguments
	if len(args) == 0 {
		if err := passwordGenerator(asciiLetters+asciiDigits+asciiPunctuation, 14, 1); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		return
	}

	// Parse arguments
	length, count, noPunct, help := parseArgs(args)

	if help {
		printHelp()
		return
	}

	// Validate length
	if length < 8 {
		fmt.Println("Error: minimum password length is 8 characters")
		os.Exit(1)
	}

	// Determine character set
	var charSet string
	if noPunct {
		charSet = asciiLetters + asciiDigits
	} else {
		charSet = asciiLetters + asciiDigits + asciiPunctuation
	}

	// Generate passwords
	if err := passwordGenerator(charSet, length, count); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
