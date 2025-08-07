package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/skip2/go-qrcode"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("QR Code Generator")
	fmt.Print("Enter the URL you want to encode: ")
	
	url, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	url = strings.TrimSpace(url)

	if url == "" {
		fmt.Println("Error: URL cannot be empty")
		return
	}


	filename := generateFilenameFromURL(url) + ".png"
	fmt.Printf("\nGenerating QR code for: %s\nSaving as: %s\n", url, filename)
	
	err = generateQRCode(url, filename)
	if err != nil {
		fmt.Println("Error generating QR code:", err)
		return
	}

	fmt.Printf("QR code generated successfully as '%s'\n", filename)
}

func generateQRCode(url string, filename string) error {
	return qrcode.WriteFile(url, qrcode.Medium, 256, filename)
}

func generateFilenameFromURL(url string) string {
	// Cleaning
	cleanURL := strings.TrimPrefix(strings.TrimPrefix(url, "http://"), "https://")

	cleanURL = strings.TrimPrefix(cleanURL, "www.")
	

	if idx := strings.Index(cleanURL, "?"); idx != -1 {
		cleanURL = cleanURL[:idx]
	}
	
	cleanURL = strings.TrimRight(cleanURL, "/")
	
	reg := regexp.MustCompile(`[^\w\-\.]`)
	cleanURL = reg.ReplaceAllString(cleanURL, "_")
	
	reg = regexp.MustCompile(`_+`)
	cleanURL = reg.ReplaceAllString(cleanURL, "_")
	

	if len(cleanURL) > 50 {
		cleanURL = cleanURL[:50]
	}
	
	return cleanURL
}
