package webserver

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/GolangUnited/helloweb/internal/handlers"
	"github.com/gorilla/mux"
)

var (
	httpStatus int
)

type handler struct {
}

func NewHandler() handlers.Handler {
	return &handler{}
}

func (h *handler) Register(router *mux.Router) {
	router.HandleFunc("/bad", BadHandler).Methods("GET")
	router.HandleFunc("/name/{PARAM}", NameHandler).Methods("GET")
	router.HandleFunc("/data", MessageHandler).Methods("POST")
	router.HandleFunc("/headers", SumHandler).Methods("POST")
}

func BadHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
}

func NameHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello, %v!", vars["PARAM"])
}

func MessageHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Fprintf(w, "I got message:\n%v", string(body))
}

func SumHandler(w http.ResponseWriter, r *http.Request) {

	httpStatus = http.StatusOK

	a := r.Header.Get("a")
	b := r.Header.Get("b")

	sum, err := Sum(a, b)

	if err != nil {
		httpStatus = http.StatusBadRequest
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		httpStatus = http.StatusBadRequest
		log.Fatalln(err)
	}

	w.Header().Set("a+b", sum)
	w.WriteHeader(httpStatus)

	if len(body) == 0 {
		fmt.Fprintf(w, "Empty body")
	}

}

func Sum(a, b string) (string, error) {

	x1, err := strconv.Atoi(a)

	if err != nil {
		log.Fatalln(err)
		return "", err
	}

	x2, err := strconv.Atoi(b)

	if err != nil {
		log.Fatalln(err)
		return "", err
	}

	sum := x1 + x2

	return strconv.Itoa(sum), nil
}
