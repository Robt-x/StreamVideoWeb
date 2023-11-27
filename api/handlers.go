package main

import (
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	io.WriteString(w, "CreateUser")
}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//uname := p.ByName("user_name")
	io.WriteString(w, "Login")

}
