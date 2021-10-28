package routers

import (
	"encoding/json"
	"net/http"

	"github.com/Juanantogg/server-go-twittor/db"
	"github.com/Juanantogg/server-go-twittor/models"
)

func Registry(w http.ResponseWriter, r *http.Request) {
	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El email del usuario es requerido", http.StatusBadRequest)
		return
	}

	if len(t.Password) < 6 {
		http.Error(w, "Debe especificar una contraseña de al menos 6 caracteres", http.StatusBadRequest)
		return
	}

	_, found, _ := db.UserExists(t.Email)

	if found {
		http.Error(w, "Ya existe un usuario registrado con este email", http.StatusBadRequest)
		return
	}

	_, status, err := db.InsertRecord(t)

	if err != nil {
		http.Error(w, "Ocurrió un error al intentar realizar el registro de usuario"+err.Error(), http.StatusBadRequest)
		return
	}

	if !status {
		http.Error(w, "No se ha logrado insertar el registro del usuario", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
