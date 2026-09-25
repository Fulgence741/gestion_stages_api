package handlers

import (
	"fmt"
	"net/http"
)

func CreateUsers(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(response, "La route Users Fonctionne")
}

func ListUsers(response http.ResponseWriter, request *http.Request) {

}

func GetUsers(response http.ResponseWriter, request *http.Request) {

}

func UpdateUsers(response http.ResponseWriter, request *http.Request) {

}

func DeleteUsers(response http.ResponseWriter, request *http.Request) {

}

func StatutChange(response http.ResponseWriter, request *http.Request) {

}

func PasswordChange(response http.ResponseWriter, request *http.Request) {

}
