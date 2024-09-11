package routes

import (
	"net/http"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir(config.Static))
	mux.Handle("/static/", http.StripPrefix("/static/", files))
	mux.HandleFunc("GET /", indexPage)
	mux.HandleFunc("/err", errPage)
	mux.HandleFunc("GET /resume", myResume)
	mux.HandleFunc("GET /zypher/tools", fetchZypherTools)
	mux.HandleFunc("GET /zypher", calculcateZypherCypher)
	mux.HandleFunc("GET /zyphash", calculateZyphash)
	mux.HandleFunc("GET /shiftapproval", fetchShiftApprovalManager)
	mux.HandleFunc("POST /shiftapproval", updateShiftApproval)

	// TODO add authentication so only I update my schedule
	mux.HandleFunc("GET /schedule", fetchMySchedule)
	mux.HandleFunc("POST /schedule", updateMySchedule)

	return mux
}

func indexPage(w http.ResponseWriter, r *http.Request) {

}

func errPage(w http.ResponseWriter, r *http.Request) {

}

func myResume(w http.ResponseWriter, r *http.Request) {

}

func fetchZypherTools(w http.ResponseWriter, r *http.Request) {

}

func calculcateZypherCypher(w http.ResponseWriter, r *http.Request) {

}

func calculateZyphash(w http.ResponseWriter, r *http.Request) {

}

func fetchShiftApprovalManager(w http.ResponseWriter, r *http.Request) {

}

func updateShiftApproval(w http.ResponseWriter, r *http.Request) {

}

func fetchMySchedule(w http.ResponseWriter, r *http.Request) {

}

func updateMySchedule(w http.ResponseWriter, r *http.Request) {

}

func internalServerErrorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("500 internal server error"))
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("404 Not Found"))
}
