package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/person", personHandler)
	http.ListenAndServe("localhost:8080", nil)
}

var persons = make([]Person, 0)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Response struct {
	Success bool   `json:"success"`
	Data    any    `json:"data"`
	Error   string `json:"error"`
}

func addPerson(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		var response Response = Response{
			Success: false,
			Error:   err.Error(),
		}
		jsonResponse(response, http.StatusInternalServerError, w)
		return
	}

	var newPerson Person

	err = json.Unmarshal(data, &newPerson)
	if err != nil {
		var response Response = Response{
			Success: false,
			Error:   err.Error(),
		}
		jsonResponse(response, http.StatusBadRequest, w)
		return
	}

	persons = append(persons, newPerson)

	var response Response = Response{
		Success: true,
		Data:    newPerson,
	}

	jsonResponse(response, http.StatusOK, w)
}

func getAllPerson(w http.ResponseWriter, r *http.Request) {
	var response Response = Response{
		Success: true,
		Data:    persons,
	}
	jsonResponse(response, http.StatusOK, w)
}

func personHandler(l http.ResponseWriter, f *http.Request) {
	switch f.Method {
	case http.MethodGet:
		getAllPerson(l, f)
	case http.MethodPost:
		addPerson(l, f)
		// var r Response = Response{
		// 	Success: false,
		// 	Error:   "not found",
		// }

		// jsonResponse(r, http.StatusNotFound, l)
	}

}

func jsonResponse(r Response, httpCode int, h http.ResponseWriter) {
	result, err := json.Marshal(r)
	if err != nil {
		fmt.Println("error marshaling", err)
		http.Error(h, "error json endocing", http.StatusInternalServerError)
		return
	}

	h.Header().Add("Content-Type", "application/json")
	h.WriteHeader(httpCode)
	h.Write(result)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ini dari println")
	w.Write([]byte("Hello World"))
}
