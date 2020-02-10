package main
import (
	"log"
	"net/http"
)
func handler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("C语言中文网\n"))
}
func main() {
	http.HandleFunc("/", handler)
	log.Printf("监听 1234 端口成功，可以通过 https://127.0.0.1:1234/ 访问")
	err := http.ListenAndServeTLS(":1234", "cert.pem", "key.pem", nil)
	if err != nil {
		log.Fatal(err)
	}
}