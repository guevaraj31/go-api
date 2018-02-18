packae main 

import (
	"fmt"
	"net/http"

)

func AllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Todos los usuarios de Endpoint Hit")
}

func NewUser(w http.ResponseWriter, r *http.Request) {
	fmtFprintf(w, "Nuevo usuario de Endpoint Hit")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Elimina usuario de Endpoint Hit")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Actualiza usuario Endpoint Hit")
}

