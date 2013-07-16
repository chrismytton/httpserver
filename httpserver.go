package main
import (
	"net/http"
	"os"
)
func main() {
	panic(http.ListenAndServe(":"+os.Getenv("PORT"), http.FileServer(http.Dir("."))))
}
