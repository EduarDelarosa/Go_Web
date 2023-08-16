package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go_web/models"
	"go_web/utilities"
	"html/template"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

var Token string = "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpZCI6MzYsImlhdCI6MTY4OTI3OTIwMSwiZXhwIjoxNjkxODcxMjAxfQ._OnjCx_COQZl8V29sl_-9a2hMolHq4H39PU2crFzeWs"

func ClienteHttp(response http.ResponseWriter, _ *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/cliente_http/home.html", utilities.Frontend))
	//Conexion a la API
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://www.api.tamila.cl/api/categorias", nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Authorization", Token)
	reg, err2 := client.Do(req)
	if err2 != nil {
		fmt.Println(err2)
	}
	defer reg.Body.Close()
	fmt.Println(reg.Status)
	//convertimos la info en un slice
	body, errBody := io.ReadAll(reg.Body)
	if errBody != nil {
		fmt.Println(errBody)
	}
	dataInfo := models.Categories{}
	errJson := json.Unmarshal(body, &dataInfo)
	if errJson != nil {
		fmt.Println(errJson)
	}
	//Retorno
	data := map[string]models.Categories{
		"data": dataInfo,
	}
	// fmt.Println(data)
	//css_message, css_session := utilities.ReturnMsgFlash(response, request)

	// data := map[string]string{
	// 	"css":     css_session,
	// 	"message": css_message,
	// }

	tmpl.Execute(response, data)
}

func ClienteHttpCreate(response http.ResponseWriter, request *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/cliente_http/create_client.html", utilities.Frontend))
	css_message, css_session := utilities.ReturnMsgFlash(response, request)
	data := map[string]string{
		"css":     css_session,
		"message": css_message,
	}
	tmpl.Execute(response, data)
}

func ClienteHttpCreate_post(response http.ResponseWriter, request *http.Request) {
	message := ""
	if len(request.FormValue("name")) == 0 {
		message = message + "El campo Nombre categoria es requerido"
	}

	if message != "" {
		utilities.CreateMsgFlash(response, request, "danger", message)
		http.Redirect(response, request, "/cliente-http/create", http.StatusFound)
	}
	data := map[string]string{
		"nombre": request.FormValue("name"),
	}
	jsonData, _ := json.Marshal(data)
	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://www.api.tamila.cl/api/categorias", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Authorization", Token)
	reg, err2 := client.Do(req)
	if err2 != nil {
		fmt.Println(err2)
	}
	defer reg.Body.Close()
	fmt.Println(reg.Status)
	utilities.CreateMsgFlash(response, request, "success", "Categoria creada")
	http.Redirect(response, request, "/cliente-http", http.StatusFound)
}

func EditClientHttp(response http.ResponseWriter, request *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/cliente_http/edit_client.html", utilities.Frontend))
	// css_message, css_session := utilities.ReturnMsgFlash(response, request)
	client := &http.Client{}
	vars := mux.Vars(request)
	req, err := http.NewRequest("GET", "https://www.api.tamila.cl/api/categorias/"+vars["id"], nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Authorization", Token)
	reg, err2 := client.Do(req)
	if err2 != nil {
		fmt.Println(err2)
	}
	defer reg.Body.Close()
	body, errBody := io.ReadAll(reg.Body)
	if errBody != nil {
		fmt.Println(errBody)
	}
	dataInfo := models.Category{}
	errJson := json.Unmarshal(body, &dataInfo)
	if errJson != nil {
		fmt.Println(errJson)
	}
	data := map[string]string{
		"id":     vars["id"],
		"nombre": dataInfo.Nombre,
		"slug":   dataInfo.Slug,
	}
	fmt.Println(reg.Status)
	tmpl.Execute(response, data)
}

func EditClientHttp_post(response http.ResponseWriter, request *http.Request) {
	message := ""
	if len(request.FormValue("name")) == 0 {
		message = message + "El campo Nombre esta vacio"
	}

	if message != "" {
		fmt.Println("Nombre esta vacio")
	}

	vars := mux.Vars(request)
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://www.api.tamila.cl/api/categorias/"+vars["id"], nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Authorization", Token)
	reg, errClientDo := client.Do(req)
	if errClientDo != nil {
		fmt.Println(errClientDo)
	}
	defer reg.Body.Close()

	body, errBody := io.ReadAll(reg.Body)
	if errBody != nil {
		fmt.Println(errBody)
	}

	dataInfo := models.Category{}
	jsonErr := json.Unmarshal(body, &dataInfo)
	if jsonErr != nil {
		fmt.Println(jsonErr)
	}
	//Editar registro
	data := map[string]string{
		"nombre": request.FormValue("name"),
	}
	jsonValue, _ := json.Marshal(data)
	fmt.Println(bytes.NewBuffer(jsonValue))
	reqPut, errPut := http.NewRequest("PUT", "https://www.api.tamila.cl/api/categorias/"+vars["id"], bytes.NewBuffer(jsonValue))
	reqPut.Header.Set("Authorization", Token)
	if errPut != nil {
		fmt.Println(errPut)
	}
	reg2, err3 := client.Do(reqPut)
	if err3 != nil {
		fmt.Println(err3)
	}
	defer reg2.Body.Close()
	http.Redirect(response, request, "/cliente-http", http.StatusSeeOther)
}

func DeleteClientHttp(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://www.api.tamila.cl/api/categorias/"+vars["id"], nil)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Set("Authorization", Token)
	reg, errReg := client.Do(req)
	if errReg != nil {
		fmt.Println(errReg)
	}
	defer reg.Body.Close()
	body, errBody := io.ReadAll(reg.Body)
	if errBody != nil {
		fmt.Println(errBody)
	}
	dataInfo := models.Category{}
	errJson := json.Unmarshal(body, &dataInfo)
	if errJson != nil {
		fmt.Println(errJson)
	}

	//Elimino registro
	req2, err2 := http.NewRequest("DELETE", "https://www.api.tamila.cl/api/categorias/"+vars["id"], nil)
	req2.Header.Set("Authorization", Token)
	if err2 != nil {
		fmt.Println(err2)
	}
	regDelete, errDelete := client.Do(req2)
	if errDelete != nil {
		fmt.Println(errDelete)
	}
	defer regDelete.Body.Close()

	http.Redirect(response, request, "/cliente-http", http.StatusSeeOther)
}
