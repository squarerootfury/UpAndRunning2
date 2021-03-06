package routes

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// Sends a simple welcome-message to the user.
func NoWebFrontendIndex(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	SendJsonMessage(w, http.StatusOK, true, "Welcome to UpAndRunning2! Currently the Web-Frontend is disabled, but you can still use our simple API.")
}

// Sends a simple welcome-message to the user.
func ApiIndex(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	SendJsonMessage(w, http.StatusOK, true, "Welcome to UpAndRunning2's API!")
}

// Sends a simple welcome-message to the user.
func ApiIndexV1(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	SendJsonMessage(w, http.StatusOK, true, "Welcome to UpAndRunning2's API v1!")
}
