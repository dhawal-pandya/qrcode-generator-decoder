package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	skip2qrcode "github.com/skip2/go-qrcode"
	tuotooqrcode "github.com/tuotoo/qrcode"
)

// generateQRCode generates a QR code using the skip2qrcode package.
func generateQRCode(text, filename string) error {
	// Generate QR code and write to file
	err := skip2qrcode.WriteFile(text, skip2qrcode.Highest, 2048, filename)
	if err != nil {
		return err
	}
	return nil
}

// decodeQRCode decodes a QR code using the tuotooqrcode package.
func decodeQRCode(filename string) (string, error) {
	// Open the QR code image file
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Decode the QR code using tuotooqrcode
	qrCode, err := tuotooqrcode.Decode(file)
	if err != nil {
		return "", err
	}

	// Return the decoded content
	return qrCode.Content, nil
}

func main() {
	// Command line flags
	mode := flag.String("mode", "generate", "Mode: 'generate' to create a QR code, 'decode' to decode a QR code.")
	input := flag.String("input", "", "Input text for generating QR code or input filename for decoding.")
	output := flag.String("output", "qrcode.png", "Output filename for QR code generation or decoding result.")

	flag.Parse()

	switch *mode {
	case "generate":
		if *input == "" {
			fmt.Println("Error: Please provide input text using the '-input' flag for QR code generation.")
			return
		}

		// Generate QR code
		err := generateQRCode(*input, *output)
		if err != nil {
			log.Fatal("Error generating QR code:", err)
		}
		fmt.Printf("QR code for \"%s\" generated and saved to %s\n", *input, *output)

	case "decode":
		if *input == "" {
			fmt.Println("Error: Please provide input filename using the '-input' flag for QR code decoding.")
			return
		}

		// Decode QR code
		decodedText, err := decodeQRCode(*input)
		if err != nil {
			log.Fatal("Error decoding QR code:", err)
		}
		fmt.Printf("Decoded text from QR code: %s\n", decodedText)

	default:
		fmt.Println("Error: Invalid mode. Use 'generate' to create a QR code or 'decode' to decode a QR code.")
	}
}

// go run main.go -mode=generate -input="https://dhawal-pandya.github.io/" -output=site.png
// go run main.go -mode=decode -input=site.png
