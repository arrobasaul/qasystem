package pagos

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type pagoIndividualRequest struct {
	ClaveSerial     string
	Rtn             string
	Contrato        string
	Fecha           string
	Hora            string
	Periodo         string
	Impuesto        int
	Concepto        int
	ValorImpuesto   float32
	ValorMulta      float32
	ValorRecargo    float32
	ValorInteres    float32
	RefBancoExterna string
	RefBancoInterna string
}
type pagoIndividualResponse struct {
	Documento string
	Err       error
}

func (res *pagoIndividualResponse) err() error { return res.Err }

func makePagoIndividualEndPoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(pagoIndividualRequest)
		id, err := s.PagoIndividual(req.ClaveSerial, req.Rtn, req.Contrato, req.Fecha, req.Hora, req.Periodo, req.Impuesto, req.Concepto, req.ValorImpuesto, req.ValorMulta, req.ValorRecargo, req.ValorInteres, req.RefBancoExterna, req.RefBancoInterna)
		return pagoIndividualResponse{Documento: id, Err: err}, nil
	}
}

type pagoMasivoRequest struct {
	Archivo []byte
}
type pagoMasivoResponse struct {
	Documentos []string
	Err        error
}

func (res *pagoMasivoResponse) err() error { return res.Err }

func makePagoMasivoEndPoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(pagoMasivoRequest)
		id, err := s.PagoMasivo(req.Archivo)
		return pagoMasivoResponse{Documentos: id, Err: err}, nil
	}
}
