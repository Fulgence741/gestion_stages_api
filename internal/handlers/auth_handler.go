package handlers

import (
	"fmt"
	"net/http"
)

func AuthLogin(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(response, "Bonjour je suis grand!!!")

}

func AuthRegister(response http.ResponseWriter, request *http.Request) {

}

func AuthMe(response http.ResponseWriter, request *http.Request) {

}

func AuthLogout(response http.ResponseWriter, request *http.Request) {

}
