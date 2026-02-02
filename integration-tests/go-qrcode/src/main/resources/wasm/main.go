package main

import (
	"unsafe"

	"github.com/skip2/go-qrcode"
)

// generateQR generates a QR code from input text and returns the PNG image
// Parameters:
//   - textPtr: pointer to input text
//   - textLen: length of input text
//   - sizePtr: pointer to write output size (will be set)
//
// Returns: pointer to PNG image data (caller must NOT call free, it's handled by GC), or 0 on error
//
//export generateQR
func generateQR(textPtr *byte, textLen int, sizePtr *int) *byte {
	// Convert pointer to Go string
	text := string(unsafe.Slice(textPtr, textLen))

	// Generate QR code as PNG (256x256 pixels, medium error correction)
	qr, err := qrcode.Encode(text, qrcode.Medium, 256)
	if err != nil {
	    // Return null pointer on error
		return nil
	}

	// Keep the QR code data alive by storing it in a global slice
	// This prevents Go's GC from collecting it before the host reads it
	buf := make([]byte, len(qr))
	copy(buf, qr)

	// Write the size to the output parameter
	*sizePtr = len(buf)

	// Return pointer to the data
	return &buf[0]
}

func main() {
	// Empty main function required for WASM
	// The actual functionality is accessed via exported functions
}
