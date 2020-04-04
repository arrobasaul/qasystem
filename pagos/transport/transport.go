package pagos

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	kitlog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"

//MakeHandler manejador para el servicio pagos
func MakeHandler(ps Service, logger kitlog.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(encodeError),
	}
	PagoIndividualHadler := kithttp.NewServer(
		makePagoIndividualEndPoint(ps),
		decodePagoIndividualRequest,
		encodeResponse,
		opts...,
	)
	r := mux.NewRouter()
	r.Handle("/pagoindividual/v1/pagoindividual", PagoIndividualHadler).Methods("POST")
	return r
}
func decodePagoIndividualRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var body struct {
		ClaveSerial     string  `json:"claveSerial"`
		Rtn             string  `json:"rtn"`
		Contrato        string  `json:"contrato"`
		Fecha           string  `json:"fecha"`
		Hora            string  `json:"hora"`
		Periodo         string  `json:"periodo"`
		Impuesto        int     `json:"impuesto"`
		Concepto        int     `json:"concepto"`
		ValorImpuesto   float32 `json:"valorImpuesto"`
		ValorMulta      float32 `json:"valorMulta"`
		ValorRecargo    float32 `json:"valorRecargo"`
		ValorInteres    float32 `json:"valorInteres"`
		RefBancoExterna string  `json:"refBancoExterna"`
		RefBancoInterna string  `json:"refBancoInterna"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}
	return pagoIndividualRequest{
		claveSerial:     body.ClaveSerial,
		rtn:             body.Rtn,
		contrato:        body.Contrato,
		fecha:           body.Fecha,
		hora:            body.Hora,
		periodo:         body.Periodo,
		impuesto:        body.Impuesto,
		concepto:        body.Concepto,
		valorImpuesto:   body.ValorImpuesto,
		valorMulta:      body.ValorMulta,
		valorRecargo:    body.ValorRecargo,
		valorInteres:    body.ValorInteres,
		refBancoExterna: body.RefBancoExterna,
		refBancoInterna: body.RefBancoInterna,
	}, nil
}
func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		encodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

type errorer interface {
	error() error
}

// encode errors from business-logic
func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	switch err {
	case cargo.ErrUnknown:
		w.WriteHeader(http.StatusNotFound)
	case ErrInvalidArgument:
		w.WriteHeader(http.StatusBadRequest)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
