package function

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Response struct {
	Sum string `json:"sum"`
}

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", DefaultHandler).Methods("GET")
	r.HandleFunc("/{num}", AddTenHandler).Methods("GET")
	return r
}

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to add ten! Add a number to the end of the URL!"))
}

func AddTenHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	numString := vars["num"]

	num, err := strconv.Atoi(numString)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	finalNum := int64(num + 10)

	var response Response
	response.Sum = strconv.FormatInt(finalNum, 10)
	err = json.NewEncoder(w).Encode(&response)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Max-Age", "3600")
	w.Header().Set("Content-Type", "application/json")
	// create router
	router := NewRouter()
	router.ServeHTTP(w, r)
}
