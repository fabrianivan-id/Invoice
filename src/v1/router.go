package v1

import (
	"net/http"

	"esb-test/src/v1/handler"

	"github.com/go-chi/chi/v5"
)

func Router(r *chi.Mux, deps *Dependency) {
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	r.Get("/esb-test/v1/swagger/v1/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./swagger/v1/swagger.json")
	})

	r.Route("/esb-test/v1", func(v1 chi.Router) {
		// Invoice endpoints
		v1.Get("/invoice", handler.GetListInvoiceHandler(deps.Services.InvoiceSvc))
		v1.Get("/invoice/{id}", handler.GetInvoiceHandler(deps.Services.InvoiceSvc))
		v1.Post("/invoice", handler.CreateInvoiceHandler(deps.Services.InvoiceSvc))
		// v1.Delete("/invoice/{id}", handler.DeleteInvoiceHandler(deps.Services.InvoiceSvc))
		// v1.Put("/invoice/{id}", handler.UpdateInvoiceHandler(deps.Services.InvoiceSvc))
	})
}
