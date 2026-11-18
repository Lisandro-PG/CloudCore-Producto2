package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "desconocido"
	}

	fmt.Fprintf(w, "Hola desde el pod %s\n", hostname)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Servidor iniciado en el puerto 8080")
	http.ListenAndServe(":8080", nil)
}
