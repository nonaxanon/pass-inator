package main

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
)

const (
	minPasswordLength = 6
	lowercaseChars    = "abcdefghijklmnopqrstuvwxyz"
	uppercaseChars    = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numberChars       = "0123456789"
	specialChars      = "!@#$%^&*()_+-=[]{}|;:,.<>?"
)

// PasswordConfig holds the configuration for password generation
type PasswordConfig struct {
	Length          int
	UseLowercase    bool
	UseUppercase    bool
	UseNumbers      bool
	UseSpecialChars bool
}

// secureRandomInt generates a cryptographically secure random integer in [0, max)
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

// secureRandomBytes generates n cryptographically secure random bytes
func secureRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	return b, err
}

// validateConfig ensures the password configuration is valid
func validateConfig(config PasswordConfig) error {
	if config.Length < minPasswordLength {
		return fmt.Errorf("password length must be at least %d characters", minPasswordLength)
	}
	if !config.UseLowercase && !config.UseUppercase && !config.UseNumbers && !config.UseSpecialChars {
		return fmt.Errorf("at least one character type must be selected")
	}
	return nil
}

// generatePassword creates a password based on the provided configuration
func generatePassword(config PasswordConfig) (string, error) {
	if err := validateConfig(config); err != nil {
		return "", err
	}

	// Build character set based on configuration
	var charSet string
	if config.UseLowercase {
		charSet += lowercaseChars
	}
	if config.UseUppercase {
		charSet += uppercaseChars
	}
	if config.UseNumbers {
		charSet += numberChars
	}
	if config.UseSpecialChars {
		charSet += specialChars
	}

	// Ensure at least one character from each selected type
	var password strings.Builder
	if config.UseLowercase {
		idx, err := secureRandomInt(len(lowercaseChars))
		if err != nil {
			return "", fmt.Errorf("failed to generate random index: %w", err)
		}
		password.WriteByte(lowercaseChars[idx])
	}
	if config.UseUppercase {
		idx, err := secureRandomInt(len(uppercaseChars))
		if err != nil {
			return "", fmt.Errorf("failed to generate random index: %w", err)
		}
		password.WriteByte(uppercaseChars[idx])
	}
	if config.UseNumbers {
		idx, err := secureRandomInt(len(numberChars))
		if err != nil {
			return "", fmt.Errorf("failed to generate random index: %w", err)
		}
		password.WriteByte(numberChars[idx])
	}
	if config.UseSpecialChars {
		idx, err := secureRandomInt(len(specialChars))
		if err != nil {
			return "", fmt.Errorf("failed to generate random index: %w", err)
		}
		password.WriteByte(specialChars[idx])
	}

	// Fill the rest of the password with random characters
	remainingLength := config.Length - password.Len()
	for i := 0; i < remainingLength; i++ {
		idx, err := secureRandomInt(len(charSet))
		if err != nil {
			return "", fmt.Errorf("failed to generate random index: %w", err)
		}
		password.WriteByte(charSet[idx])
	}

	// Shuffle the password using Fisher-Yates algorithm with crypto/rand
	passwordStr := password.String()
	passwordRunes := []rune(passwordStr)
	for i := len(passwordRunes) - 1; i > 0; i-- {
		j, err := secureRandomInt(i + 1)
		if err != nil {
			return "", fmt.Errorf("failed to shuffle password: %w", err)
		}
		passwordRunes[i], passwordRunes[j] = passwordRunes[j], passwordRunes[i]
	}

	return string(passwordRunes), nil
}

func readUserInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func readYesNo(prompt string) bool {
	for {
		input := strings.ToLower(readUserInput(prompt))
		if input == "y" || input == "yes" {
			return true
		}
		if input == "n" || input == "no" {
			return false
		}
		fmt.Println("Please enter 'y' or 'n'")
	}
}

func main() {
	fmt.Println("Welcome to Pass-inator - Your Secure Password Generator")
	fmt.Println("-----------------------------------------------------")

	// Get password length
	lengthStr := readUserInput(fmt.Sprintf("Enter password length (minimum %d): ", minPasswordLength))
	length, err := strconv.Atoi(lengthStr)
	if err != nil {
		fmt.Printf("Error: Invalid length. Using minimum length of %d\n", minPasswordLength)
		length = minPasswordLength
	}

	config := PasswordConfig{
		Length:          length,
		UseLowercase:    readYesNo("Include lowercase letters? (y/n): "),
		UseUppercase:    readYesNo("Include uppercase letters? (y/n): "),
		UseNumbers:      readYesNo("Include numbers? (y/n): "),
		UseSpecialChars: readYesNo("Include special characters? (y/n): "),
	}

	// Generate and display password
	password, err := generatePassword(config)
	if err != nil {
		fmt.Printf("Error generating password: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("\nYour generated password is:")
	fmt.Println("------------------------")
	fmt.Println(password)
	fmt.Println("------------------------")
}
