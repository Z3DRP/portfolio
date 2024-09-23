package routes

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Z3DRP/portfolio/config"
	"github.com/spf13/viper"
)

func readConfig() (*config.ZServerConfig, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetConfigType("yml")
	var configs config.Configurations

	if err := viper.ReadInConfig(); err != nil {
		// TODO set this to log instead of print
		emsg := fmt.Sprintf("Error reading config file, %s", err)
		fmt.Print(emsg)
		return nil, errors.New(emsg)
	}

	err := viper.Unmarshal(&configs)
	if err != nil {
		// TODO set this to log instead
		emsg := fmt.Sprintf("Unable to decode into struct, %v", err)
		fmt.Print(emsg)
		return nil, errors.New(emsg)
	}

	return &configs.ZServer, nil
}

func NewRouter() http.Handler {
	cfig, err := readConfig()
	if err != nil {
		return nil
	}
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir(cfig.Static))
	mux.Handle("/static/", http.StripPrefix("/static/", files))
	mux.HandleFunc("GET /", indexPage)
	// mux.HandleFunc("/err", errPage)
	mux.HandleFunc("GET /resume", myResume)
	// mux.HandleFunc("GET /zypher/tools", fetchZypherTools)
	// mux.HandleFunc("GET /zypher", calculcateZypherCypher)
	// mux.HandleFunc("GET /zyphash", calculateZyphash)
	// mux.HandleFunc("GET /shiftapproval", fetchShiftApprovalManager)
	// mux.HandleFunc("POST /shiftapproval", updateShiftApproval)

	// TODO add authentication so only I update my schedule
	// mux.HandleFunc("GET /schedule", fetchMySchedule)
	// mux.HandleFunc("POST /schedule", updateMySchedule)

	return mux
}

func indexPage(w http.ResponseWriter, r *http.Request) {

}

// func errPage(w http.ResponseWriter, r *http.Request) {

// }

func myResume(w http.ResponseWriter, r *http.Request) {

}

// func fetchZypherTools(w http.ResponseWriter, r *http.Request) {

// }

// func calculcateZypherCypher(w http.ResponseWriter, r *http.Request) {

// }

// func calculateZyphash(w http.ResponseWriter, r *http.Request) {

// }

// func fetchShiftApprovalManager(w http.ResponseWriter, r *http.Request) {

// }

// func updateShiftApproval(w http.ResponseWriter, r *http.Request) {

// }

// func fetchMySchedule(w http.ResponseWriter, r *http.Request) {

// }

// func updateMySchedule(w http.ResponseWriter, r *http.Request) {

// }

func internalServerErrorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("500 internal server error"))
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("404 Not Found"))
}
