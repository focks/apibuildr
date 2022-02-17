package tpl

func ServerTemplate() []byte {
	return []byte(`
package cmd 
import (
	"context"
	"flag"
	"fmt"
	"github.com/focks/apibuildr"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var Router = mux.NewRouter()

type Server struct {
	s *http.Server
}

func (server Server) run() {
	if err := server.s.ListenAndServe(); err != nil {
		panic(err)
	}
}

func (server Server) shutdown(ctx context.Context) {
	_ = server.s.Shutdown(ctx)
}

func Exec() {
	var wait time.Duration
	var port string

	// capture the required flags
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.StringVar(&port, "port", "8080", "port number")
	flag.Parse()

	server := newServer(port, time.Second*15, Router)

	go server.run()

	fmt.Println(fmt.Sprintf("Started server on port %s", port))

	c := make(chan os.Signal, 1)

	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)
	// Block until we receive our signal.
	<-c
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	fmt.Println("Shutting down server")
	server.shutdown(ctx)
	os.Exit(0)
}

func init() {
	Router.NotFoundHandler = apibuildr.FourZeroFour()
}

func newServer(port string, timeout time.Duration, handler *mux.Router) *Server {
	return &Server{s: &http.Server{
		Addr:         fmt.Sprintf("0.0.0.0:%s", port),
		WriteTimeout: timeout,
		ReadTimeout:  timeout,
		IdleTimeout:  timeout,
		Handler:      handler,
	}}
}
`)
}
