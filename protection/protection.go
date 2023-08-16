package protection

import (
	"go_web/utilities"
	"net/http"
)

// Para que la funcion sea interpretada como un middleware debe recibir el parametro next http.HandlerFunc y retornar el mismo tipo
func Protection(next http.HandlerFunc) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		//Colocamos las validaciones necesarias o requeridas
		session, _ := utilities.Store.Get(request, "session-name")
		if session.Values["user_id"] != nil {
			//Permitimos continuar a la ruta solicitada
			next.ServeHTTP(response, request)
		} else {
			utilities.CreateMsgFlash(response, request, "warning", "Debes estar logeado para visualizar este contenido")
			http.Redirect(response, request, "/segurity/login", http.StatusSeeOther)
		}
		//Permitimos continuar a la ruta solicitada
		// next.ServeHTTP(response, request)
	}
}
