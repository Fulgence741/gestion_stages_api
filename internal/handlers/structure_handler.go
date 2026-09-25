package handlers

import (
	"fmt"
	"net/http"
)

func CreateStructure(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(response, "La route Strcuture marche !!")
}

func ListStructures(response http.ResponseWriter, request *http.Request) {

}

func GetStructure(response http.ResponseWriter, request *http.Request) {

}

func UpdateStructure(response http.ResponseWriter, request *http.Request) {

}

func DeleteStructure(response http.ResponseWriter, request *http.Request) {

}

func ListMaitreStageStructure(response http.ResponseWriter, request *http.Request) {

}
