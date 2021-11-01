package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})
	r.HandleFunc("/api/exames_ext", use(createExames, basicAuth)).Methods("POST")
	r.HandleFunc("/api/exames_ext/{accessionNumber}", use(deleteExame, basicAuth)).Methods("DELETE")
	r.HandleFunc("/api/laudos/{accessionNumber}", use(getLaudos, basicAuth)).Methods("GET")
	r.HandleFunc("/api/medicos/{crm}", use(getAssinaturas, basicAuth)).Methods("GET")

	port := os.Getenv("PORT") //Get port from .env file, we did not specify any port so this should return an empty string when tested locally
	if port == "" {
		port = "8000" //localhost
	}

	fmt.Println(port)

	err := http.ListenAndServe(":"+port, r) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}
}
