package handlers

import (
	"fmt"
	"net/http"
)

func CreateEtablissement(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(response, "La route Etablissement marche!!")
}

func ListEtablissement(response http.ResponseWriter, request *http.Request) {

}

func GetEtablissement(response http.ResponseWriter, request *http.Request) {

}

func UpdateEtablissement(response http.ResponseWriter, request *http.Request) {

}

func DeleteEtablissement(response http.ResponseWriter, request *http.Request) {

}

func ListerFilieresEtab(response http.ResponseWriter, request *http.Request) {

}

func AddFiliereEtab(response http.ResponseWriter, request *http.Request) {

}

func DeleteFiliereEtab(response http.ResponseWriter, request *http.Request) {

}
