package webserver

import "net/http"

func MiWebServer() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":3000", nil)
}
func home(writer http.ResponseWriter, request *http.Request) {
	http.ServeFile(writer, request, "./webserver/index.html")
}
