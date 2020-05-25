package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main()  {
	router := mux.NewRouter()

	router.HandleFunc("/signup",signup).Methods("POST")
	router.HandleFunc("/login",login).Methods("POST")
	router.HandleFunc("/protected",tokenVerifyMiddleWare(protectedEndpoint)).Methods("GET")
	log.Println("Listen on port 8000")
	log.Fatal(http.ListenAndServe(":8000",router)) 
}

func signup(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("Signup invoked.")
}

func login(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("Login .")
}

func protectedEndpoint(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("ProtectedEndpoint .")
}

func tokenVerifyMiddleWare(next http.HandlerFunc) http.HandlerFunc  {
	return nil
}