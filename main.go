package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/JuanCManchegoH/goprisma/handlers"
	"github.com/JuanCManchegoH/goprisma/middleware"
	"github.com/JuanCManchegoH/goprisma/server"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")
	JWT_SECRET := os.Getenv("JWT_SECRET")
	DB_URI := os.Getenv("DB_URI")

	s, err := server.NewServer(context.Background(), &server.Config{
		Port:      ":" + PORT,
		JWTSecret: JWT_SECRET,
		DbURI:     DB_URI,
	})
	if err != nil {
		log.Fatal(err)
	}

	s.Start(BindRoutes)
}

func BindRoutes(s server.Server, r *mux.Router) {
	r.Use(middleware.CheckAuthMiddleware(s))
	r.HandleFunc("/welcome", handlers.HomeHandler(s)).Methods(http.MethodGet)

	//Auth
	r.HandleFunc("/signup", handlers.SignUpHandler(s)).Methods(http.MethodPost)
	r.HandleFunc("/login", handlers.LoginHandler(s)).Methods(http.MethodPost)

	// User routes
	r.HandleFunc("/profile", handlers.ProfileHandler(s)).Methods(http.MethodGet)
	r.HandleFunc("/user/update/{id}", handlers.UpdateUserHandler(s)).Methods(http.MethodPatch)
	r.HandleFunc("/user/delete/{id}", handlers.DeleteUserHandler(s)).Methods(http.MethodDelete)
}
