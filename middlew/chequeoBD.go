package middlew

import (
	"net/http"

	"github.com/cardozoqsantiago/ApiGoBackend/bd"
)

/*ChequeoBD es el middlew que me permite conocer el estado de la bd */
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConnection() == 0 {
			http.Error(rw, "Conexion perdida con la base de datos", 500)
			return
		}
		next.ServeHTTP(rw, r)
	}
}
