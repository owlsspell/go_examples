package main

import (
	"fmt"
	"strings"
)

var cipherText string = "CSOITEUIWUIZNSROCNKFD"
var keyword string = "GOLANG"
var message string = ""
var keyIndex int = 0

var text = "your message goes here"

// func main() {
// 	repeatCount := len(cipherText) / len(keyword)
// 	remainder := len(cipherText) % len(keyword)

// 	result := strings.Repeat(keyword, repeatCount) + keyword[:remainder]
// 	for i, w := range cipherText {
// 		message += string(decrypt(w, rune(result[i])))
// 	}
// 	fmt.Println(message)
// }

// func decrypt(cipherChar, keyChar rune) rune {
// 	alphabetLength := 26
// 	encryptedIndex := int(cipherChar - 'A')
// 	keyIndex := int(keyChar - 'A')
// 	decryptedIndex := (encryptedIndex - keyIndex + alphabetLength) % alphabetLength
// 	return 'A' + rune(decryptedIndex)
// }

func decrypt() {
	for i := 0; i < len(cipherText); i++ {
		c := cipherText[i] + 'A'
		k := keyword[keyIndex] + 'A'
		r := (c-k+26)%26 + 'A'
		message += string(r)

		keyIndex++
		keyIndex %= len(keyword)
	}
	fmt.Println(message)
}

func encrypt() {
	newText := strings.Replace(strings.ToUpper(text), " ", "", -1)
	fmt.Println(newText)
	for i := 0; i < len(newText); i++ {
		c := newText[i] + 'A'
		k := keyword[keyIndex] - 'A'
		r := (c+k)%26 + 'A'
		message += string(r)

		keyIndex++
		keyIndex %= len(keyword)
	}
	fmt.Println(message)

}

func main() {
	encrypt()
}
