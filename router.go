package main

import(
    	"net/http"
    	"github.com/gorilla/mux"
        "log"
        //"fmt"
)

 

//RouterStart роутер
func RouterStart() {
    var router = mux.NewRouter()  
    static := http.StripPrefix("/static/", http.FileServer(http.Dir("./Files/Static/")))
    router.PathPrefix("/static/").Handler(static).Methods("GET")
   
    //router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {StartHandler(w, r,"Index")}).Methods("GET")
    router.HandleFunc("/manifest.mf", ManifestHandler)
    router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {StartHandler(w, r,"Main")}).Methods("GET")
    router.HandleFunc("/sale", func(w http.ResponseWriter, r *http.Request) {StartHandler(w, r,"Main")}).Methods("GET")
    //router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {StartHandler(w, r,"Main")}).Methods("GET")
    router.HandleFunc("/forum", func(w http.ResponseWriter, r *http.Request) {StartHandler(w, r,"Main")}).Methods("GET")
    //router.HandleFunc("/settings", func(w http.ResponseWriter, r *http.Request) {StartHandler(w, r,"Main")}).Methods("GET")
    //router.HandleFunc("/training", func(w http.ResponseWriter, r *http.Request) {StartHandler(w, r,"Main")}).Methods("GET")
    router.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {MainHandler(w, r,"Main")}).Methods("GET") //Login  
    router.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {StartHandler(w, r,"Logout")}).Methods("GET")  
    router.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {StartHandler(w, r,"Main")}).Methods("GET")       
    router.HandleFunc("/api/{method}", func(w http.ResponseWriter, r *http.Request) {ApiHandler(w, r)})
    router.HandleFunc("/websocket", ClientSockServer) 

    log.Println("Server started") 
    http.ListenAndServe(":80", router)
}


