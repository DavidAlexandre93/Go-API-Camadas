package routes

import (
    "log"
    "net/http"
    "controllers/controllers"
)

func HandleResquest() {
    http.HandleFunc("/", controllers.Home)
    log.Fatal(http.ListenAndServe(":8000", nil))
}