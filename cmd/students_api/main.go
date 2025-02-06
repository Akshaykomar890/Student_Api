package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Akshaykomar890/Students_Api/internal/config"
)

func main(){
	//load config
	cfg := config.MustLoad()
	//databasesetup
	//router
	router :=http.NewServeMux()
	router.HandleFunc("GET /",func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to students api"))
	})
	//server
	server:=http.Server{
		Addr: cfg.HttpServer.Address,
		Handler: router , //router
	}
	fmt.Printf("Server Statred %s",cfg.HttpServer.Address)
	error:=server.ListenAndServe()
	if error != nil {
		log.Fatal("Failed to start server")
	}


}