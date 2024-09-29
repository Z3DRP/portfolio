package routes

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/Z3DRP/portfolio/config"
	"github.com/Z3DRP/portfolio/internal/models"
	zlg "github.com/Z3DRP/portfolio/internal/zlogger"
	"github.com/spf13/viper"
)

var logfile = zlg.NewLogFile(
	zlg.WithFilename(fmt.Sprintf("%v/routes/%v", config.LogPrefix, config.LogName)),
)
var logger = zlg.NewLogger(
	logfile,
	zlg.WithJsonFormatter(true),
	zlg.WithLevel("trace"),
	zlg.WithReportCaller(true),
)

func readConfig() (*config.ZServerConfig, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetConfigType("yml")
	var configs config.Configurations

	if err := viper.ReadInConfig(); err != nil {
		emsg := fmt.Sprintf("error reading config file, %s", err)
		logger.MustDebug(emsg)
		return nil, errors.New(emsg)
	}

	err := viper.Unmarshal(&configs)
	if err != nil {
		emsg := fmt.Sprintf("unable to decode into struct, %v", err)
		logger.MustDebug(emsg)
		return nil, errors.New(emsg)
	}

	return &configs.ZServer, nil
}

func NewRouter() http.Handler {
	cfig, err := readConfig()
	if err != nil {
		logger.MustDebug(fmt.Sprintf("error parsing config file, %v", err))
		return nil
	}
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir(cfig.Static))
	mux.Handle("public/static/", http.StripPrefix("public/static/", files))
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
	skills := models.TempSkillFactory()
	myDetails := models.TempDetailFactory()
	jsonSkills, err := json.Marshal(skills)
	if err != nil {
		logger.MustDebug(fmt.Sprintf("error encoding skills to json, %v", err))
	}

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
