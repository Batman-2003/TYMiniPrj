package myQRLib

import (
	"fmt"
	"log"

	"github.com/skip2/go-qrcode"
)

func CreateQRCode(data string, id uint32) {
	filename := fmt.Sprintf("./../Frontend/Resources/QRCodes/%v.png", id)
	err := qrcode.WriteFile(data, qrcode.Medium, 256, filename)
	if err != nil {
		fmt.Println("Error While Creating QR Code")
		log.Fatal(err)
	}
}
