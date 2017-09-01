package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/aashishthite/message_parser/pkg/parse"
	"github.com/gorilla/mux"
)

const (
	DEFAULT_PORT = ":8080"
	timeout      = 15 * time.Second
)

type Handlers struct {
	parser parse.Parser
}

func New() *Handlers {
	return &Handlers{
		parser: parse.NewParser(),
	}
}

func (h *Handlers) ListenAndServe() error {
	port := os.Getenv("PORT")

	if port == "" {
		port = DEFAULT_PORT
	}
	log.Println("Listening on port", port)
	return http.ListenAndServe(":"+port, h.endpoints())
}

func (h *Handlers) endpoints() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)

	addHandler(router, "GET", "/", h.Health)

	addHandler(router, "POST", "/parse", h.Parse)

	return router
}

// Health check endpoint
func (h *Handlers) Health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "All OK")
}

func (h *Handlers) Parse(w http.ResponseWriter, r *http.Request) {
	var postReq postParseReq

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&postReq)
	if err != nil || postReq.Msg == "" {
		ErrorRespondAndLog(w, r, err, "Failed to decode the request body", http.StatusBadRequest)
		return
	}
	result, err := h.parser.Parse(postReq.Msg)
	if err != nil {
		ErrorRespondAndLog(w, r, err, "Unable to parse message", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(result)
}

//Helper Funcs
func addHandler(r *mux.Router, method, pattern string, hand func(w http.ResponseWriter, r *http.Request)) {
	r.HandleFunc(pattern, requestLogger(hand)).Methods(method)
}

func requestLogger(hand func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("source:%s\tmethod:%s\tpath:%s\n", strings.Split(r.RemoteAddr, ":")[0], r.Method, r.URL.Path)

		done := make(chan bool, 1)
		go func() {
			hand(w, r)
			done <- true
		}()

		select {
		case <-done:
			return
		case <-time.After(timeout - time.Second):
			log.Printf("TIMEOUT ERROR: source:%s\tmethod:%s\tpath:%s", strings.Split(r.RemoteAddr, ":")[0], r.Method, r.URL.Path)
		}
	}
}

func ErrorRespondAndLog(w http.ResponseWriter, r *http.Request, err error, errStr string, code int) {
	log.Printf("ERROR at path=\"%s\" err=\"%s\" code:\"%d\"", r.RequestURI, err, code)
	http.Error(w, errStr, code)
}
