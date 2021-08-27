package main

import (
	"database/sql"
	"encoding/gob"
	"github.com/DapperBlondie/service-monitor/internal/config"
	"github.com/DapperBlondie/service-monitor/internal/handlers"
	"github.com/DapperBlondie/service-monitor/internal/models"
	"github.com/alexedwards/scs/v2"
	"github.com/pusher/pusher-http-go"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"time"
)

var app config.AppConfig
var repo *handlers.DBRepo
var session *scs.SessionManager
var preferenceMap map[string]string
var wsClient pusher.Client

const Version = "1.0.0"
const maxWorkerPoolSize = 5
const maxJobMaxWorkers = 5

func init() {
	gob.Register(models.User{})
	_ = os.Setenv("TZ", "America/Halifax")
}

// main is the application entry point
func main() {
	// set up application
	insecurePort, err := setupApp()
	if err != nil {
		log.Fatal(err)
	}

	// close channels & db when application ends
	defer close(app.MailQueue)
	defer func(SQL *sql.DB) {
		err = SQL.Close()
		if err != nil {
			log.Println(err.Error())
		}
	}(app.DB.SQL)

	// print info
	log.Printf("******************************************")
	log.Printf("** %sService-Monitor%s v%s built in %s", "\033[31m", "\033[0m", Version, runtime.Version())
	log.Printf("**----------------------------------------")
	log.Printf("** Running with %d Processors", runtime.NumCPU())
	log.Printf("** Running on %s", runtime.GOOS)
	log.Printf("******************************************")

	// create http server
	srv := &http.Server{
		Addr:              *insecurePort,
		Handler:           routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}

	log.Printf("Starting HTTP server on port %s....", *insecurePort)

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)

	// start the server
	go func() {
		err = srv.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	<-sigChan
	log.Println("Server was shutdown ...")
}
