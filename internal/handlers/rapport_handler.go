package handlers

import (
	"fmt"
	"net/http"
)

func CreateRapport(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(response, "La route Rapport fonctionne")
}

func ListRapports(response http.ResponseWriter, request *http.Request) {

}

func GetRapport(response http.ResponseWriter, request *http.Request) {

}

func UpdateRapport(response http.ResponseWriter, request *http.Request) {

}

func DeleteRapport(response http.ResponseWriter, request *http.Request) {

}

func GetStageRapport(response http.ResponseWriter, request *http.Request) {

}
