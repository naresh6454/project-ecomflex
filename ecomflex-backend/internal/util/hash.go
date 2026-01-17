package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// Cost for bcrypt
const bcryptCost = 12

// Initialize random seed
func init() {
	rand.Seed(time.Now().UnixNano())
}

// HashPassword hashes a password using bcrypt
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcryptCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	
	return string(bytes), nil
}

// CheckPasswordHash checks if a password matches a hash
// CheckPasswordHash checks if a password matches a hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// GenerateReferralCode generates a referral code from a name
func GenerateReferralCode(name string) string {
	// Remove spaces and convert to uppercase
	name = strings.ToUpper(strings.ReplaceAll(name, " ", ""))
	
	// Take the first 5 characters or pad with random letters if too short
	if len(name) < 5 {
		for i := len(name); i < 5; i++ {
			name += string(rune('A' + rand.Intn(26)))
		}
	}
	
	code := name[:5]
	
	// Add a random number
	code += fmt.Sprintf("%04d", rand.Intn(10000))
	
	return code
}

// GenerateRandomReferralCode generates a completely random referral code
func GenerateRandomReferralCode() string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, 9) // 9 characters
	
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	
	return string(result)
}