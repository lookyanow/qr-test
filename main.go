package main

import (
	"fmt"

	qrcode "github.com/skip2/go-qrcode"
)

func main() {
	fmt.Println("This is test qr-code programm")
	err := qrcode.WriteFile("https://yandex.ru", qrcode.High, 512, "qr.png")
	if err != nil {
		fmt.Println("Error during qr code generation")
	}
}
