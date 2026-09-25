package handlers

import (
	"fmt"
	"net/http"
)

func CreateEvaluation(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(response, "La route Evaluationmarche !!")
}

func ListEvaluations(response http.ResponseWriter, request *http.Request) {

}

func GetEvaluation(response http.ResponseWriter, request *http.Request) {

}

func UpdateEvaluation(response http.ResponseWriter, request *http.Request) {

}

func DeleteEvaluation(response http.ResponseWriter, request *http.Request) {

}

func ListRapportEvaluations(response http.ResponseWriter, request *http.Request) {

}
