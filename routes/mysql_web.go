package routes

import (
	"fmt"
	"go_web/connection"
	"go_web/models"
	"go_web/utilities"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

func MySql_list(response http.ResponseWriter, request *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/my_sql/home.html", utilities.Frontend))
	css_message, css_session := utilities.ReturnMsgFlash(response, request)

	// Connection to db
	connection.Connetion()

	sql := "SELECT id, nombre, correo, telefono, fecha FROM clientes ORDER BY id DESC;"
	clients := models.Clients{}
	datos, err := connection.Db.Query(sql)
	if err != nil {
		fmt.Println(err)
	}

	defer connection.CloseConnection()

	for datos.Next() {
		dato := models.Client{}
		datos.Scan(&dato.Id, &dato.Nombre, &dato.Telefono, &dato.Correo, &dato.Fecha)
		clients = append(clients, dato)
	}

	data := models.ClientsHttp{
		Css:     css_session,
		Message: css_message,
		Datos:   clients,
	}

	tmpl.Execute(response, data)
}

func Mysql_Create(response http.ResponseWriter, request *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/my_sql/create_mysql.html", utilities.Frontend))
	css_message, css_session := utilities.ReturnMsgFlash(response, request)

	data := map[string]string{
		"css":     css_session,
		"message": css_message,
	}

	tmpl.Execute(response, data)
}

func Mysql_Create_Post(response http.ResponseWriter, request *http.Request) {
	message := ""
	if len(request.FormValue("nombre")) == 0 {
		message = message + "El campo nombre es requerido"
	} else if len(request.FormValue("correo")) == 0 {
		message = message + "El campo correo es requerido"
	} else if len(request.FormValue("telefono")) == 0 {
		message = message + "El campo telefono es requerido"
	}

	if message != "" {
		utilities.CreateMsgFlash(response, request, "danger", message)
		http.Redirect(response, request, "/mysql", http.StatusFound)
	}

	connection.Connetion()
	defer connection.CloseConnection()

	// now := time.Now()
	// fecha := now.Format("2000-01-02 15:00:00")

	sql := "INSERT INTO clientes(id, nombre, telefono, correo, fecha) VALUES(null,?,?,?,?);"
	result, err := connection.Db.Exec(sql, request.FormValue("nombre"), request.FormValue("telefono"), request.FormValue("correo"), "2023-07-30 19:27:00")

	if err != nil {
		panic(err)
	}

	fmt.Println(result)
	fmt.Println("Registro exitoso")
	http.Redirect(response, request, "/mysql", http.StatusFound)

}

func Mysql_Edit(response http.ResponseWriter, request *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/my_sql/edit.html", utilities.Frontend))
	css_message, css_session := utilities.ReturnMsgFlash(response, request)

	vars := mux.Vars(request)
	// fmt.Println(vars["id"])

	//Coneccion bd
	connection.Connetion()
	defer connection.CloseConnection()
	//Validar si existe Cliente
	sql := "SELECT id, nombre, correo, telefono, fecha FROM clientes where id=?;"
	datos, err := connection.Db.Query(sql, vars["id"])
	if err != nil {
		panic(err)
	}

	var dato models.Client
	for datos.Next() {
		errResult := datos.Scan(&dato.Id, &dato.Nombre, &dato.Telefono, &dato.Correo, &dato.Fecha)
		if errResult != nil {
			panic(errResult)
		}
	}

	data := models.ClientHttp{
		Css:     css_session,
		Message: css_message,
		Datos:   dato,
	}

	tmpl.Execute(response, data)
}

func Mysql_Edit_Post(response http.ResponseWriter, request *http.Request) {
	fmt.Println("editar")
	vars := mux.Vars(request)
	message := ""
	if len(request.FormValue("nombre")) == 0 {
		message = message + "El campo nombre es requerido"
	} else if len(request.FormValue("correo")) == 0 {
		message = message + "El campo correo es requerido"
	} else if len(request.FormValue("telefono")) == 0 {
		message = message + "El campo telefono es requerido"
	}

	if message != "" {
		utilities.CreateMsgFlash(response, request, "danger", message)
		http.Redirect(response, request, "/mysql", http.StatusFound)
	}

	connection.Connetion()
	defer connection.CloseConnection()

	// now := time.Now()
	// fecha := now.Format("2000-01-02 15:00:00")

	sql := "UPDATE clientes SET nombre=?, correo=?, telefono=?, fecha=? where id=?;"
	result, err := connection.Db.Exec(sql, request.FormValue("nombre"), request.FormValue("telefono"), request.FormValue("correo"), "2023-07-30 19:27:00", vars["id"])

	if err != nil {
		panic(err)
	}

	fmt.Println(result)
	fmt.Println("Actualización exitosa")
	http.Redirect(response, request, "/mysql", http.StatusFound)
}

func Mysql_Delete(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	// fmt.Println(vars["id"])

	//Coneccion bd
	connection.Connetion()
	defer connection.CloseConnection()
	//Validar si existe Cliente
	sql := "SELECT id, nombre, correo, telefono, fecha FROM clientes where id=?;"
	result, err := connection.Db.Query(sql, vars["id"])
	if err != nil {
		panic(err)
	}

	for result.Next() {
		var dato models.Client
		errResult := result.Scan(&dato.Id, &dato.Nombre, &dato.Telefono, &dato.Correo, &dato.Fecha)

		if errResult != nil {
			panic(errResult)
		}

		if dato.Id != 0 {
			sql := "DELETE FROM clientes WHERE id=?;"
			resultDelete, errDelete := connection.Db.Exec(sql, vars["id"])

			if errDelete != nil {
				panic(errDelete)
			}

			fmt.Println(resultDelete)

			http.Redirect(response, request, "/mysql", http.StatusFound)
		}
	}
}
