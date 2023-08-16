package main

import (
	"fmt"
	"go_web/protection"
	"go_web/routes"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	r := mux.NewRouter()
	// Rutas
	r.HandleFunc("/", routes.Home)
	r.HandleFunc("/about", routes.About)
	r.HandleFunc("/parameters/{id}/{slug}", routes.Parameters)
	r.HandleFunc("/parameters-query", routes.ParametersQueryString)
	r.HandleFunc("/structures", routes.Structures)

	r.HandleFunc("/form", routes.Form_get)
	r.HandleFunc("/form-post", routes.Form_post).Methods("POST")

	r.HandleFunc("/form-upload", routes.Form_upload_get)
	r.HandleFunc("/form-upload-post", routes.Form_upload_post).Methods("POST")

	r.HandleFunc("/resources", routes.Resources)
	r.HandleFunc("/pdf", routes.Resources_pdf)
	r.HandleFunc("/generate", routes.Resources_pdf_generate)
	r.HandleFunc("/excel", routes.Resources_generate_excel)
	r.HandleFunc("/qr", routes.Resources_generate_QR)
	r.HandleFunc("/email", routes.Resources_SendEmail)

	r.HandleFunc("/cliente-http", routes.ClienteHttp)
	r.HandleFunc("/cliente-http/create", routes.ClienteHttpCreate)
	r.HandleFunc("/cliente-http/create-post", routes.ClienteHttpCreate_post).Methods("POST")
	r.HandleFunc("/cliente-http/edit/{id:.*}", routes.EditClientHttp)
	r.HandleFunc("/cliente-http/edit-post/{id:.*}", routes.EditClientHttp_post).Methods("POST")
	r.HandleFunc("/cliente-http/delete/{id:.*}", routes.DeleteClientHttp)

	r.HandleFunc("/mysql", routes.MySql_list)
	r.HandleFunc("/mysql/create", routes.Mysql_Create)
	r.HandleFunc("/mysql/create-post", routes.Mysql_Create_Post).Methods("POST")
	r.HandleFunc("/mysql/edit/{id:.*}", routes.Mysql_Edit)
	r.HandleFunc("/mysql/edit-post/{id:.*}", routes.Mysql_Edit_Post).Methods("POST")
	r.HandleFunc("/mysql/delete/{id:.*}", routes.Mysql_Delete)

	r.HandleFunc("/segurity/register", routes.Segurity_register)
	r.HandleFunc("/segurity/register-post", routes.Segurity_register_post).Methods("POST")

	r.HandleFunc("/segurity/login", routes.Register_login)
	r.HandleFunc("/segurity/login-post", routes.Register_login_post).Methods("POST")
	r.HandleFunc("/segurity/protegida", protection.Protection(routes.Protegida))
	r.HandleFunc("/segurity/logout", protection.Protection(routes.Segurity_logout))

	//Archivos estáticos hacia mux
	s := http.StripPrefix("/public/", http.FileServer(http.Dir("./public/")))
	r.PathPrefix("/public/").Handler(s)

	//Pagina 404
	r.NotFoundHandler = http.HandlerFunc(routes.Page404)

	//Ejecucion del servidor
	errorVariables := godotenv.Load()

	if errorVariables != nil {
		panic(errorVariables)
	}

	server := &http.Server{
		Addr:         "192.168.1.10:" + os.Getenv("PORT"),
		Handler:      r,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("Servidor corriendo en el puerto: http://192.168.1.10:" + os.Getenv("PORT"))
	log.Fatal(server.ListenAndServe())

	/*http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(response, "Hello World!")
	})

	fmt.Println("Servidor corriendo en el puerto: http://192.168.1.10:8080")
	log.Fatal(http.ListenAndServe("192.168.1.10:8080", nil))*/
}
