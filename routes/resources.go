package routes

import (
	"encoding/base64"
	"fmt"
	"go_web/utilities"
	"html/template"
	"log"
	"net/http"

	"github.com/jung-kurt/gofpdf"
	qrcode "github.com/skip2/go-qrcode"
	"github.com/xuri/excelize/v2"
)

func Resources(response http.ResponseWriter, request *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/resources/resources.html", utilities.Frontend))

	css_message, css_session := utilities.ReturnMsgFlash(response, request)

	data := map[string]string{
		"css":     css_session,
		"message": css_message,
	}

	tmpl.Execute(response, data)
}

func Resources_pdf(response http.ResponseWriter, request *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/resources/pdf.html", utilities.Frontend))

	css_message, css_session := utilities.ReturnMsgFlash(response, request)

	data := map[string]string{
		"css":     css_session,
		"message": css_message,
	}

	tmpl.Execute(response, data)
}

func Resources_pdf_generate(response http.ResponseWriter, request *http.Request) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "PDF generado utilizando GOFPDF - Golang")
	response.Header().Set("Content-Type", "application/pdf")
	response.Header().Set("Content-Disposition", "attachment; filename=example.pdf")
	err := pdf.Output(response)
	// err := pdf.OutputFileAndClose("hello.pdf") //Se genera el pdf en la raiz del proyecto
	if err != nil {
		utilities.CreateMsgFlash(response, request, "danger", "Error al generar pdf")
		http.Redirect(response, request, "/form", http.StatusSeeOther)
	} else {
		utilities.CreateMsgFlash(response, request, "success", "pdf creado con exito")
		http.Redirect(response, request, "/resources", http.StatusSeeOther)
	}
}

func Resources_generate_excel(response http.ResponseWriter, request *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/resources/excel.html", utilities.Frontend))
	//Excel
	f := excelize.NewFile()

	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	index, err := f.NewSheet("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}

	f.SetCellValue("Sheet1", "A1", "ID")
	f.SetCellValue("Sheet1", "B1", "Nombre")

	f.SetActiveSheet(index)

	if err := f.SaveAs("ExcelPrueba.xlsx"); err != nil {
		fmt.Println(err)
	}

	css_message, css_session := utilities.ReturnMsgFlash(response, request)

	data := map[string]string{
		"css":     css_session,
		"message": css_message,
	}

	tmpl.Execute(response, data)
}

func Resources_generate_QR(response http.ResponseWriter, request *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/resources/qr.html", utilities.Frontend))
	//Generamos qr
	imgCoding, err := qrcode.Encode("https://eduardelarosa.github.io/PortafolioWeb/Portafolio2.0/", qrcode.High, 256)
	if err != nil {
		log.Fatal("Error al generar codigo QR", err)
	}
	img := base64.StdEncoding.EncodeToString(imgCoding)
	//Retornar

	css_message, css_session := utilities.ReturnMsgFlash(response, request)
	data := map[string]string{
		"css":     css_session,
		"message": css_message,
		"image":   img,
	}

	tmpl.Execute(response, data)
}

func Resources_SendEmail(response http.ResponseWriter, request *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/resources/send_email.html", utilities.Frontend))
	//Envio de correo
	utilities.SendEmail()
	css_message, css_session := utilities.ReturnMsgFlash(response, request)
	data := map[string]string{
		"css":     css_session,
		"message": css_message,
	}

	tmpl.Execute(response, data)
}
