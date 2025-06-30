package main

import ("fmt"; "log"; "net/http")

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Go HTTP Server!")
}

func main() {
	PORT := ":7999"
	// Register th handler for the "/" path
	http.HandleFunc("/", helloHandler)
	fmt.Println("Server listening on " + PORT)
	log.Fatal(http.ListenAndServe(PORT, nil))
}
