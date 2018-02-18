packae main 

import (
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/sqlite"

)

var db *gorm.DB
var err error

type User struct {
	gorm.Model
	Name string
	Email string
}

func InitialMigration() {
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("Falló la conexión a base de datos")
	}
	defer db.Close()

	db.AutoMigrate(&User{})
}

func AllUsers(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("No pudo conectarse a la base de datos")
	}

	defer db.Close()

	var users []User
	db.Find(&users)
	json.NewEncoder(w).Encode(user)
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

