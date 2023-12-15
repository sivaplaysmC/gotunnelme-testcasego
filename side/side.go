package main

import (
	"crypto"
	"fmt"

	"github.com/sec51/twofactor"
)

func main() {
	otp, _ := twofactor.NewTOTP("info@sec51.com", "Sec51", crypto.SHA1, 8)

	qrBytes, _ := otp.QR()

	fmt.Println(string(qrBytes))

	err := otp.Validate("")
	if err != nil {
	}
	// if there is an error, then the authentication failed
	// if it succeeded, then store this information and do not display the QR code ever again.
}
