package main

import (
	"fmt"
	"net/http"
)

// func main() {
// 	fmt.Println("Belajar coding golang melanjutkan mempelajari golang")
// }

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	var message = "Ini adalah web App pertama saya"
	w.Write([]byte(message))
}

func handlerHello(w http.ResponseWriter, r *http.Request) {
	var message = "Hello saya sedang belajar"
	w.Write([]byte(message))
}

func handleBelajar(w http.ResponseWriter, r *http.Request) {
	var pesan = "saya menggunakan variabel pakai bahasa indonesia"
	w.Write([]byte(pesan))
}

func main() {
	http.HandleFunc("/", handlerIndex)
	http.HandleFunc("/index", handlerIndex)
	http.HandleFunc("/hello", handlerHello)
	http.HandleFunc("/belajar", handleBelajar)

	var address = "localhost:8989"
	fmt.Printf("server started at %s\n", address)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
