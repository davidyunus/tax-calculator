package httpserver

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/davidyunus/tax-calculator/internal/httpserver/controller"
	"github.com/davidyunus/tax-calculator/internal/tax"
)

// HTTPServer ...
type HTTPServer struct {
	taxController *controller.TaxController
}

// NewServer ...
func NewServer(taxService *tax.Service) *HTTPServer {
	taxCtl := controller.NewTaxController(
		taxService,
	)

	return &HTTPServer{
		taxController: taxCtl,
	}
}

// Serve ...
func (hs *HTTPServer) Serve() {
	r := hs.compileRouter()

	log.Printf("Listen to port 9090. Go to : http://127.0.0.1:9090")
	srv := http.Server{Addr: ":9090", Handler: r}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Listen: %v\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit

	log.Println("Shutdown server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server shutdown: ", err)
	}
	log.Println("Server exiting")
}
