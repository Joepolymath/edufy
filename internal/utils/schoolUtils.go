package utils

// import (
// 	"Learnium/models"
// 	"context"
// 	"crypto/rand"
// 	"fmt"
// 	"gorm.io/gorm"
// 	"math/big"
// 	"strings"
// )

// func GenerateCode(length int, startString string) string {

// 	/* This is used to genera*/
// 	// Define the characters to choose from
// 	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// 	// Get the total number of characters
// 	charCount := big.NewInt(int64(len(chars)))

// 	// Update a buffer to store the generated string
// 	buffer := make([]byte, length)

// 	// Generate random indices and select characters from the char string
// 	for i := 0; i < length; i++ {
// 		randomIndex, _ := rand.Int(rand.Reader, charCount)
// 		buffer[i] = chars[randomIndex.Int64()]
// 	}
// 	SchoolCode := string(buffer)
// 	SchoolCode = fmt.Sprintf("%s%s", startString, strings.ToUpper(SchoolCode))

// 	// Return the generated string
// 	return SchoolCode
// }

// func GenerateSchoolCode(db *gorm.DB, ctx context.Context) string {

// 	SchoolCode := GenerateCode(6, "LM")
// 	//check if code exists
// 	SchoolCodeExist := CheckIfSchoolCodeExists(db, ctx, SchoolCode)
// 	if SchoolCodeExist {
// 		// i need to generate a new one
// 		return GenerateCode(6, "LM")
// 	}

// 	// Return the generated string
// 	return SchoolCode
// }

// func CheckIfSchoolCodeExists(db *gorm.DB, ctx context.Context, schoolCode string) bool {
// 	/* This is used to check if a schoolCode exists*/

// 	var school models.School
// 	err := db.WithContext(ctx).Model(&school).Where("school_code = ?", schoolCode).First(&school).Error

// 	// if the error is not found
// 	if err == gorm.ErrRecordNotFound {
// 		return false
// 	}
// 	return true

// }
