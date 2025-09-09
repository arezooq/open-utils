package security

import (
	"crypto/rand"
	"math/big"
)

// GenerateOTP
func GenerateOTP(length int) string {
	const digits = "0123456789"
	otp := make([]byte, length)

	for i := range otp {
		num, _ := rand.Int(rand.Reader, big.NewInt(int64(len(digits))))
		otp[i] = digits[num.Int64()]
	}

	return string(otp)
}
