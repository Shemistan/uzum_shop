package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
)

// Просто оставлю здесь
func main() {
	fmt.Println(generate())
}

func generate() string {
	// Генерируем случайные байты/home/yourname/Desktop/Documents/shop/shop1
	bytes := make([]byte, 18)

	_, err := rand.Read(bytes)
	if err != nil {
		log.Println(err)
	}
	// Кодируем байты в base64
	key := base64.StdEncoding.EncodeToString(bytes)

	// Выводим ключи
	return key
}
