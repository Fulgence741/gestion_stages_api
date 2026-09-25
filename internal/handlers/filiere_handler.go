package handlers

import (
	"fmt"
	"net/http"
)

func CreateFiliere(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(response, "La route Filire fonctionne !!")
}

func ListFilieres(response http.ResponseWriter, request *http.Request) {

}

func GetFiliere(response http.ResponseWriter, request *http.Request) {

}

func UpdateFiliere(response http.ResponseWriter, request *http.Request) {

}

func DeleteFiliere(response http.ResponseWriter, request *http.Request) {

}
