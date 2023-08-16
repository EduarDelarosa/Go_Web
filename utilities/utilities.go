package utilities

import (
	"net/http"

	"github.com/gorilla/sessions"
	gomail "gopkg.in/gomail.v2"
)

var Frontend = "templates/layout/frontend.html"

var Store = sessions.NewCookieStore([]byte("session-name"))

func CreateMsgFlash(response http.ResponseWriter, request *http.Request, css string, message string) {
	session, err := Store.Get(request, "flash-session")
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	session.AddFlash(css, "css")
	session.AddFlash(message, "message")
	session.Save(request, response)
}

func ReturnMsgFlash(response http.ResponseWriter, request *http.Request) (string, string) {
	session, _ := Store.Get(request, "flash-session")

	fm := session.Flashes("css")
	session.Save(request, response)
	css_session := ""
	if len(fm) == 0 {
		css_session = ""
	} else {
		css_session = fm[0].(string)
	}
	fm2 := session.Flashes("message")
	session.Save(request, response)
	css_message := ""
	if len(fm2) == 0 {
		css_message = ""
	} else {
		css_message = fm2[0].(string)
	}

	return css_message, css_session
}

func SendEmail() {
	msg := gomail.NewMessage()
	msg.SetHeader("From", "noreply@agendahoras.cl")
	msg.SetHeader("To", "eduardelarosa09@gmail.com")
	msg.SetHeader("Subject", "Curso Golang")
	msg.SetBody("text/html", "<h1>Curso Golang</h1> <br> <p> Email enviado desde golang </p>")
	//msg.Attach() aca podemos enviar un adjunto
	//Ahora configuramos la conexión con SMTP
	n := gomail.NewDialer("smtp.dreamhost.com", 587, "noreply@agendahoras.cl", "khdwJAXysB")

	if err := n.DialAndSend(msg); err != nil {
		panic(err)
	}

}

func ReturnLogin(request *http.Request) (string, string) {
	session, _ := Store.Get(request, "session-name")
	user_id := ""
	user_name := ""
	if session.Values["user_id"] != nil {
		user_id_token, _ := session.Values["user_id"].(string)
		user_id = user_id_token
	}

	if session.Values["user_name"] != nil {
		user_name_token, _ := session.Values["user_name"].(string)
		user_name = user_name_token
	}

	return user_id, user_name
}
