/*
 * Mandatory exercise 1
 *
 * Mandatory exercse 1
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package main

import (
	"log"
	"net/http"
)

func main() {
	log.Printf("Server started")

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
