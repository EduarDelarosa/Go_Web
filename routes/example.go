package routes

import (
	"go_web/utilities"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

func Home(response http.ResponseWriter, _ *http.Request) {

	tmpl := template.Must(template.ParseFiles("templates/ejemplo/home.html", utilities.Frontend))

	tmpl.Execute(response, nil)

	/*tmpl, err := template.ParseFiles("templates/ejemplo/home.html", "templates/layout/frontend.html")

	if err != nil {
		panic(err)
	} else {
		tmpl.Execute(response, nil)
	}*/
}

func Page404(response http.ResponseWriter, _ *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/ejemplo/page404.html", utilities.Frontend))

	tmpl.Execute(response, nil)
}

func About(response http.ResponseWriter, _ *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/ejemplo/about.html", utilities.Frontend))
	tmpl.Execute(response, nil)
	/*tmpl, err := template.ParseFiles("templates/ejemplo/about.html", "templates/layout/frontend.html")

	if err != nil {
		panic(err)
	} else {
		tmpl.Execute(response, nil)
	}*/
}

func Parameters(response http.ResponseWriter, request *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/ejemplo/parameters.html", utilities.Frontend))
	vars := mux.Vars(request)
	data := map[string]string{
		"id":    vars["id"],
		"slug":  vars["slug"],
		"texto": "andate pasha",
	}

	tmpl.Execute(response, data)
	/*if err != nil {
		panic(err)
	} else {
		tmpl.Execute(response, data)
	}*/
}

func ParametersQueryString(response http.ResponseWriter, request *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/ejemplo/parameters_querystring.html", utilities.Frontend))
	data := map[string]string{
		"id":   request.URL.Query().Get("id"),
		"slug": request.URL.Query().Get("slug"),
	}
	tmpl.Execute(response, data)
	/*if err != nil {
		panic(err)
	} else {
		tmpl.Execute(response, data)
	}*/
}

type Habilidad struct {
	Nombre string
}

type Datos struct {
	Nombre      string
	Edad        int
	Perfil      int
	Habilidades []Habilidad
}

func Structures(response http.ResponseWriter, _ *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/ejemplo/structures.html", utilities.Frontend))
	habilidad1 := Habilidad{"Programar"}
	habilidad2 := Habilidad{"Jugar"}
	habilidad3 := Habilidad{"Ver tv"}
	Habilidades := []Habilidad{habilidad1, habilidad2, habilidad3}
	tmpl.Execute(response, Datos{"Eduard", 27, 1, Habilidades})
	/*if err != nil {
		panic(err)
	} else {
		tmpl.Execute(response, Datos{"Eduardo", 17, 1, []})
	}*/
}

/*func Parameters(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	fmt.Fprintln(response, "Los parametros son: ID = "+vars["id"]+" Y SLUG = "+vars["slug"])
}

func ParametersQueryString(response http.ResponseWriter, request *http.Request) {
	fmt.Println(request.URL)
	fmt.Fprintln(response, request.URL)

	id := request.URL.Query().Get("id")
	slug := request.URL.Query().Get("slug")

	fmt.Fprintln(response, "Los parametros son: ID = "+id+" Y SLUG = "+slug)
}*/
