package routes

import (
	"fmt"
	"go_web/utilities"
	"go_web/validations"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func Form_get(response http.ResponseWriter, request *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/forms/form.html", utilities.Frontend))
	css_message, css_session := utilities.ReturnMsgFlash(response, request)
	data := map[string]string{
		"css":     css_session,
		"message": css_message,
	}
	tmpl.Execute(response, data)
}

func Form_post(response http.ResponseWriter, request *http.Request) {

	message := ""
	if len(request.FormValue("name")) == 0 {
		message = message + "El campo nombre es requerido"
		// fmt.Fprintln(response, message)
		// return
	}

	if validations.Regex_correo.FindStringSubmatch(request.FormValue("email")) == nil {
		message = message + "El campo email es requerido"
		// fmt.Fprintln(response, message)
		// return
	}

	if message != "" {
		utilities.CreateMsgFlash(response, request, "danger", message)
		http.Redirect(response, request, "/form", http.StatusSeeOther)
	}

	fmt.Fprintln(response, request.FormValue("name"))
}

func Form_upload_get(response http.ResponseWriter, request *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/forms/form-upload.html", utilities.Frontend))
	css_message, css_session := utilities.ReturnMsgFlash(response, request)
	data := map[string]string{
		"css":     css_session,
		"message": css_message,
	}

	tmpl.Execute(response, data)
}

func Form_upload_post(response http.ResponseWriter, request *http.Request) {
	file, handler, err := request.FormFile("foto")

	if err != nil {
		utilities.CreateMsgFlash(response, request, "danger", "Ocurrio un error")
	}

	var ext = filepath.Ext(handler.Filename)
	fmt.Print(file, ext)
	// time := strings.Split(time.Now().String(), " ")
	foto := "Nombrequequiero." + ext
	var saveFoto string = "public/uploads/fotos/" + foto

	f, errSave := os.OpenFile(saveFoto, os.O_WRONLY|os.O_CREATE, 0777)

	if errSave != nil {
		utilities.CreateMsgFlash(response, request, "danger", "Ocurrio un error inesperado")
	}

	_, errCopy := io.Copy(f, file)

	if errCopy != nil {
		utilities.CreateMsgFlash(response, request, "danger", "Ocurrio un error inesperado")
	}

	//Aca se guardaria el registro en la bd

	//Aca redireccionamos
	utilities.CreateMsgFlash(response, request, "success", "Carga de archivo exitosa")
	http.Redirect(response, request, "/form-upload", http.StatusSeeOther)

}
