package main
import (
	"io"
	"log"
	"net/http"
)
func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "C语言中文网\n")
}
func main() {
	http.HandleFunc("/hello", HelloServer)
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}