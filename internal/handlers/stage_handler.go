package handlers

import (
	"fmt"
	"net/http"
)

func CreateStage(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(response, "La route Stage fonctionne !!")
}

func ListStages(response http.ResponseWriter, request *http.Request) {

}

func GetStage(response http.ResponseWriter, request *http.Request) {

}

func UpdateStage(response http.ResponseWriter, request *http.Request) {

}

func DeleteStage(response http.ResponseWriter, request *http.Request) {

}
