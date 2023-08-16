package models

type Category struct {
	Id     int
	Nombre string
	Slug   string
}

type Categories []Category

type Client struct {
	Id       int
	Nombre   string
	Correo   string
	Telefono string
	Fecha    string
}

type Clients []Client

type ClientsHttp struct {
	Css     string
	Message string
	Datos   Clients
}

type ClientHttp struct {
	Css     string
	Message string
	Datos   Client
}

type Usuario struct {
	Id       int
	Nombre   string
	Correo   string
	Telefono string
	Password string
}

type Usuarios []Usuario
