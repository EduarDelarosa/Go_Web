package routes

import (
	"go_web/connection"
	"go_web/models"
	"go_web/utilities"
	"go_web/validations"
	"net/http"
	"strconv"
	"text/template"

	"golang.org/x/crypto/bcrypt"
)

func Segurity_register(response http.ResponseWriter, request *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/segurity/register.html", utilities.Frontend))
	css_message, css_session := utilities.ReturnMsgFlash(response, request)

	data := map[string]string{
		"css":     css_session,
		"message": css_message,
	}

	tmpl.Execute(response, data)
}

func Segurity_register_post(response http.ResponseWriter, request *http.Request) {
	message := ""
	if len(request.FormValue("nombre")) == 0 {
		message = message + "El campo nombre es requerido <br>"
	}

	if validations.Regex_correo.FindStringSubmatch(request.FormValue("correo")) == nil {
		message = message + "El campo correo es requerido <br>"
	}

	if !validations.ValidarPassword(request.FormValue("password")) {
		message = message + "La contraseña debe tener al menos 1 número, una mayúscula y un largo entre 6 y 20 caracteres <br>"
	}

	if message != "" {
		utilities.CreateMsgFlash(response, request, "danger", message)
		http.Redirect(response, request, "/segurity/register", http.StatusSeeOther)
		return
	}

	connection.Connetion()
	defer connection.CloseConnection()
	sql := "INSERT INTO usuarios VALUES(null, ?,?,?,?);"
	//Generamos el hash de contraseña
	//p2gHNiENUw
	costo := 8
	bytes, _ := bcrypt.GenerateFromPassword([]byte(request.FormValue("password")), costo)

	_, err := connection.Db.Exec(sql, request.FormValue("nombre"), request.FormValue("correo"), request.FormValue("telefono"), string(bytes))

	if err != nil {
		panic(err)
	}

	utilities.CreateMsgFlash(response, request, "success", "Se creó el registro exitosamente")
	http.Redirect(response, request, "/segurity/register", http.StatusSeeOther)
}

func Register_login(response http.ResponseWriter, request *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/segurity/login.html", utilities.Frontend))

	css_message, css_session := utilities.ReturnMsgFlash(response, request)

	data := map[string]string{
		"css":     css_session,
		"message": css_message,
	}

	tmpl.Execute(response, data)
}

func Register_login_post(response http.ResponseWriter, request *http.Request) {
	message := ""
	if validations.Regex_correo.FindStringSubmatch(request.FormValue("correo")) == nil {
		message = message + "El campo correo es requerido <br>"
	}

	if !validations.ValidarPassword(request.FormValue("password")) {
		message = message + "La contraseña debe tener al menos 1 número, una mayúscula y un largo entre 6 y 20 caracteres <br>"
	}

	if message != "" {
		utilities.CreateMsgFlash(response, request, "danger", message)
		http.Redirect(response, request, "/segurity/login", http.StatusSeeOther)
		return
	}

	//Nos conectamos a la bd
	connection.Connetion()
	defer connection.CloseConnection()

	//Validar que el correo exista
	//p2gHNiENUw
	sql := "SELECT id, nombre, correo, telefono, password FROM usuarios WHERE correo=?;"

	userData, err := connection.Db.Query(sql, request.FormValue("correo"))
	if err != nil {
		panic(err)
	}

	var dato models.Usuario

	for userData.Next() {
		errNext := userData.Scan(&dato.Id, &dato.Nombre, &dato.Correo, &dato.Telefono, &dato.Password)
		if errNext != nil {
			utilities.CreateMsgFlash(response, request, "danger", "Las credenciales no son validas")
			http.Redirect(response, request, "/segurity/login", http.StatusSeeOther)
		}
	}
	//Comparar hash
	passwordBytes := []byte(request.FormValue("password"))
	passwordDB := []byte(dato.Password)

	errPassword := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)
	if errPassword != nil {
		utilities.CreateMsgFlash(response, request, "danger", "Las credenciales no son validas")
		http.Redirect(response, request, "/segurity/login", http.StatusSeeOther)
	} else {
		session, _ := utilities.Store.Get(request, "session-name")
		str := strconv.Itoa(dato.Id)
		session.Values["user_id"] = str
		session.Values["user_name"] = dato.Nombre
		errSession := session.Save(request, response)
		if errSession != nil {
			http.Error(response, errSession.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(response, request, "/segurity/protegida", http.StatusSeeOther)
	}
}

func Protegida(response http.ResponseWriter, request *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/segurity/protegida.html", utilities.Frontend))
	css_message, css_session := utilities.ReturnMsgFlash(response, request)
	user_id, user_name := utilities.ReturnLogin(request)
	data := map[string]string{
		"css":       css_session,
		"message":   css_message,
		"user_id":   user_id,
		"user_name": user_name,
	}

	tmpl.Execute(response, data)
}

func Segurity_logout(response http.ResponseWriter, request *http.Request) {
	session, _ := utilities.Store.Get(request, "session-name")
	session.Values["user_id"] = nil
	session.Values["user_name"] = nil
	err2 := session.Save(request, response)
	if err2 != nil {
		http.Error(response, err2.Error(), http.StatusInternalServerError)
		return
	}
	utilities.CreateMsgFlash(response, request, "primary", "Se a cerrado tu sesión!")
	http.Redirect(response, request, "/segurity/login", http.StatusSeeOther)
}
