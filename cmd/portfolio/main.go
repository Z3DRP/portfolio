package main

import "net/http"

func main() {

	mux := http.NewServeMux()
	files := http.FileServer(http.Dir(config.Static))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	mux.HandleFunc("GET /", index)
	mux.HandleFunc("/err", err)

	mux.HandleFunc("GET /resume", myResume)
	mux.HandleFunc("GET /zypher", zypherToolsGet)
	mux.HandleFunc("POST /zypher", zypherToolsPost)
	mux.HandleFunc("GET /shiftapproval", shiftApprovalManagerGet)
	mux.HandleFunc("POST /shiftapproval", shiftApprovalManagerPost)
	mux.HandleFunc("GET /schedule", myScheduleGet)
	mux.HandleFunc("POST /schedule", mySchedulePost)

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
