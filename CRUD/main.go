package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Contacto struct {
	Nombre       string    `json:"nombre"`
	Apellido string    `json:"apellido"`
	Direccion string `json:"direccion"`
	Telefono string `json:"telefono"`
	CreatedAt   time.Time `json:"created_at"`
}

//este es donde vamos a almacenar en memoria nuestros contactos, viene siendo una simulacion de una BD
var contactStore = make(map[string]Contacto)

var id int

func GetContactHandler(w http.ResponseWriter, r *http.Request) {
	var contactos []Contacto
	for _, valor := range contactStore {
		contactos = append(contactos, valor) //aqui se tienen almacenados los contactos que habia en el map contactStore
	}
	w.Header().Set("Content-Type", "application/json") //se esta creando una cabecera http, que le indica al navegador que se le devuelve un contenido de tipo json
	j, err := json.Marshal(contactos)                      //codificando el slice de notes para pasarlo a json, j se convierte en un slice de byte
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(j)                   //se esta retornando la respuesta al usuario
}

func PostContactHandler(w http.ResponseWriter, r *http.Request) {
	var contact Contacto
	//en body viene la peticion del usuario, aqui le estamos pasando lo que va a descodificar de lo que nos manda el usuario con el Request
	// Decode(&contact) le estoy indicando que lo decodifique a la structura de contact por medio de un puntero indicandole a que estructura
	err := json.NewDecoder(r.Body).Decode(&contact)
	if err != nil {
		panic(err)
	}
	contact.CreatedAt = time.Now() //agregamos la fecha en que se creo el contacto
	id++
	k := strconv.Itoa(id) //me convierte mi id entero a string para pasarlo al map que ocupa un id
	contactStore[k] = contact

	// se esta devolviendo el mismo objeto (json) que nos envio el usuario pero con mas datos, en este caso con la fecha,
	w.Header().Set("Content-Type", "application/json")
	j, err := json.Marshal(contact)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(j)

}

func PutContactHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r) //Vars recibe como parametro un puntero a un Request	y r ya es eso
	k := vars["id"]     //se obtiene el id que mando el usuario
	var contactUpdate Contacto
	/* 	el json lo decodificamos a un formato de GO, en Decode le estamos pasando el
	puntero a la variable donde se van a pasar los datos ya decodificados */
	err := json.NewDecoder(r.Body).Decode(&contactUpdate)
	if err != nil {
		panic(err)
	}

	// se veriffica si existe el id
	if contact, ok := contactStore[k]; ok {
		contactUpdate.CreatedAt = contact.CreatedAt
		delete(contactStore, k)      //se borra el contacto viejo del map, para introducir el nuevo
		contactStore[k] = contactUpdate //contacto actualizado
	} else {
		log.Printf("No encontramos el id %s", k)
	}

	w.WriteHeader(http.StatusNoContent)

}

func DeleteContactHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	k := vars["id"]

	if _, ok := contactStore[k]; ok {
		delete(contactStore, k)
	} else {
		log.Printf("No encontramos el id %s", k)
	}
	w.WriteHeader(http.StatusNoContent)
}

func main() {

	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/api/agenda", GetContactHandler).Methods("GET")
	r.HandleFunc("/api/agenda", PostContactHandler).Methods("POST")
	r.HandleFunc("/api/agenda/{id}", PutContactHandler).Methods("PUT")
	r.HandleFunc("/api/agenda/{id}", DeleteContactHandler).Methods("DELETE")

	server := &http.Server{
		Addr:           ":8083",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Println("Listening http://localhost:8083...")
	server.ListenAndServe()

}
