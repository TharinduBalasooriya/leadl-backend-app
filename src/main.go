package main

import (
	"log"
	"net/http"
	"os"

	"github.com/TharinduBalasooriya/LogAnalyzerBackend/src/routes"
	"github.com/gorilla/handlers"
	"github.com/joho/godotenv"
)

// LoadEnv /*
func LoadEnv() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

}

/*
 Entry point
*/
func main() {

	//Starting the API server
	router := routes.LogRoutes()

	os.MkdirAll("temp",0755)
	os.MkdirAll("localstorage",0755)
	os.MkdirAll("debug_env",0755)

	//Load the env file
	LoadEnv()
	http.Handle("/api/", router)
	//log.Println("Server Started localhost :3000")
	log.Fatal(http.ListenAndServe(":3000", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"http://localhost:4200"}))(router)))



}
