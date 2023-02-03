package main

import (
	"auth/src/configs"
	"auth/src/handlers"
	"auth/src/repository/database/mongo"
	"auth/src/router"
	"auth/src/services"
	"context"
	"flag"
	"github.com/danborodin/go-logd"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func init() {

}

func main() {

	logFilePath := flag.String("l", "", "log file")
	flag.Parse()

	var l *logd.Logger

	if *logFilePath != "" {
		logFile, err := os.OpenFile(*logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			log.Fatal(err)
		}
		l = logd.NewLogger(logFile, log.Ldate|log.Ltime|log.Llongfile)
	} else {
		l = logd.NewLogger(os.Stdout, log.Ldate|log.Ltime|log.Llongfile)
	}

	defer func(l *logd.Logger) {
		err := l.Close()
		if err != nil {
			l.ErrPrintln(err)
		}
	}(l)

	//init configs
	configs.Conf = configs.New(l)
	configs.InitConfig(configs.Conf)

	//create db
	db := mongo.NewDB(l)
	defer db.Disconnect()

	//create services
	srvc := services.NewServices(l, db)

	//createhandlers and router
	rhandler := handlers.NewRouterHandlers(l, srvc)
	router := router.CreateRouter(l, rhandler)

	//create server
	server := &http.Server{
		Addr:    ":6969",
		Handler: router,
	}

	//wait for os signals
	connsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, syscall.SIGINT, syscall.SIGTERM)
		<-sigint

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		err := server.Shutdown(ctx)
		if err != nil {
			l.ErrPrintln(err)
		}

		l.InfoPrintln("shutting down")

		close(connsClosed)
	}()

	l.InfoPrintln("Starting server")
	if err := server.ListenAndServe(); err != nil {
		l.ErrPrintln(err)
	}

	<-connsClosed

	os.Exit(0)
}
