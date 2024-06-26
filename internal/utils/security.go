package utils

import (
	"log"
	"math/rand"
	"regexp"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func HashSecret(secret string) string {
	/* This is used to hash a secret key and return the hashed secret in string format*/
	hashedSecret, err := bcrypt.GenerateFromPassword([]byte(secret), 12)
	if err != nil {
		log.Panicln(err)
	}
	return string(hashedSecret)

}

func VerifySecret(hashedSecret string, plainSecret string) (bool, error) {
	/*
		This is used to compare a plain secret key and it's hashed version to determine if they are same
	*/
	err := bcrypt.CompareHashAndPassword(
		[]byte(hashedSecret),
		[]byte(plainSecret),
	)
	check := true
	if err != nil {
		check = false
		return check, err
	}
	return check, nil
}

func GenerateOTP(length int) string {
	rand.Seed(time.Now().UnixNano())
	// Generate a random 4-digit number
	digits := "0123456789"

	// Build the random string by selecting random digits
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = digits[rand.Intn(len(digits))]
	}

	return string(result)
}

func IsValidEmail(email string) bool {
	// Define a regular expression for a basic email validation
	// This is a simplified version and may not cover all edge cases
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	// Compile the regular expression
	re := regexp.MustCompile(emailRegex)

	// Use the regular expression to match the email format
	return re.MatchString(email)
}
