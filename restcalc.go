package main

import (
	"context"
	"encoding/json"
	_ "expr_rest-api/docs"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/borichevskiy/expression_calculator"
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
)

type response struct {
	Expr string `json:"expr"`
	Res  int    `json:"res"`
	Err  string `json:"err"`
}

// @Summary Evaluate
// @Description evaluate expression
// @ID eval-expr
// @Accept  json
// @Produce  json
// @Param expr query string true "Expression"
// @Success 200 {string} string
// @Failure 400,404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Failure default {object} httputil.DefaultError
// @Router /evaluate [get]
func evaluateExpression(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	expr := r.URL.Query().Get("expr")
	res, err := expression_calculator.Evaluate(expr)
	// to avoid incorrect nil error: "\u003cnil\u003e"
	var err2 string = fmt.Sprint(err)
	if err == nil {
		err2 = "nil"
	}

	rsp, _ := json.Marshal(response{Expr: expr, Res: res, Err: err2})

	w.Write(rsp)
}

// @title Expression Rest-Api
// @version 1.0
// @description This is rest-http server for expression calculator.
// @host localhost:9000
// @BasePath /evaluate
func main() {
	r := chi.NewRouter()
	r.Get("/evaluate/", evaluateExpression)
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:9000/swagger/doc.json"), //The url pointing to API definition"
	))

	var port int

	flag.IntVar(&port, "port", 9000, "Port number")
	flag.Parse()

	srv := &http.Server{
		Addr:    ":9000",
		Handler: r,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	log.Print("Server Started")

	<-done
	log.Print("Server Stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}

	defer func() {
		cancel()
	}()

	log.Print("Server Exited Properly")
}
