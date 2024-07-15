package utils

import "fmt"

func PrintQR(qr []byte) {
    width := 256 // Width of the QR code

    // Print QR code to terminal
    for y := 0; y < width; y++ {
        for x := 0; x < width; x++ {
            // Calculate index of the QR code byte slice
            index := y*width + x

            // Check if the byte at index is 0 (black)
            if index < len(qr) && qr[index] == 0 {
                fmt.Print("██") // Print block character for "black" pixels
            } else {
                fmt.Print("  ") // Print white spaces for "white" pixels
            }
        }
        fmt.Println() // Move to the next line after each row

    }
}
