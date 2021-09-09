package routers

import (
	"encoding/json"
	"net/http"

	"github.com/cardozoqsantiago/ApiGoBackend/bd"
	"github.com/cardozoqsantiago/ApiGoBackend/models"
)

/*Registro es la funcion para crear en la bd el usuario*/
func Registro(rw http.ResponseWriter, r *http.Request) {
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(rw, "Error en los datos recibios"+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(rw, "El email de usuario es requerido", http.StatusBadRequest)
		return
	}
	if len(t.Password) == 0 {
		http.Error(rw, "Debe especificar una contrasenia de al menos 6 caracteres", http.StatusBadRequest)
		return
	}

	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)
	if encontrado == true {
		http.Error(rw, "ya existe un usuario con ese email", http.StatusBadRequest)
		return
	}

	_, status, err := bd.InsertarRegistro(t)
	if err != nil {
		http.Error(rw, "Ocurrio un error al intenrar ingresar el usuario"+err.Error(), http.StatusBadRequest)
	}
	if status == false {
		http.Error(rw, "No se pudo insertar el registro", http.StatusBadRequest)
	}
	rw.WriteHeader(http.StatusCreated)
}
