package exampleproject

import (
	"io"
	"net/http"
)

type WeServiceExample struct {
	httpHandlerFunc *http.HandlerFunc
	httpListenAndServe *http.Server
}

func Index(w http.ResponseWriter, r *http.Request) {
	 io.WriteString(w, "Hello World")
}



