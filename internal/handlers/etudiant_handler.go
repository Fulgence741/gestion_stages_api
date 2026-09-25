package handlers

import (
	"fmt"
	"net/http"
)

func CreateEtudiant(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(response, "La route Etudiant marche !!")
}

func ListEtudiant(response http.ResponseWriter, request *http.Request) {

}

func Getetudiant(response http.ResponseWriter, request *http.Request) {

}

func UpdateEtudiant(response http.ResponseWriter, request *http.Request) {

}

func DeleteEudiant(response http.ResponseWriter, request *http.Request) {

}

func GetEtudiantStatge(response http.ResponseWriter, request *http.Request) {

}
