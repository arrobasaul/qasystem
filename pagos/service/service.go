package pagos

import "errors"

// ErrInvalidArgument Es Retornado cuando uno o mas argumentos son invalidos.
var ErrInvalidArgument = errors.New("invalid argument")

//Service expone el servicio de pagos
type Service interface {
	PagoIndividual(claveSerial string, rtn string, contrato string, fecha string, hora string, periodo string, impuesto int, concepto int, valorImpuesto float32, valorMulta float32, valorRecargo float32, valorInteres float32, refBancoExterna string, refBancoInterna string) (string, error)
	PagoMasivo(archivo []byte) ([]string, error)
}
type service struct {
	pagoRepository Repository
}

func (s *service) PagoIndividual(claveSerial string, rtn string, contrato string, fecha string, hora string, periodo string, impuesto int, concepto int, valorImpuesto float32, valorMulta float32, valorRecargo float32, valorInteres float32, refBancoExterna string, refBancoInterna string) (string, error) {
	pago := Pago{
		claveSerial:     claveSerial,
		rtn:             rtn,
		contrato:        contrato,
		fecha:           fecha,
		hora:            hora,
		periodo:         periodo,
		impuesto:        impuesto,
		concepto:        concepto,
		valorImpuesto:   valorImpuesto,
		valorMulta:      valorMulta,
		valorRecargo:    valorRecargo,
		valorInteres:    valorInteres,
		refBancoExterna: refBancoExterna,
		refBancoInterna: refBancoInterna,
	}
	doc, err := s.pagoRepository.GenerarPago(pago)
	if err != nil {
		return "", ErrInvalidArgument
	}
	return doc, nil
}
func (s *service) PagoMasivo(archivo []byte) ([]string, error) {
	return []string{}, nil
}

//NewService Crea un servicio de pagos con sus dependencias
func NewService(pagoRepository Repository) Service {
	return &service{
		pagoRepository: pagoRepository,
	}
}
