package pagos

import (
	"time"

	"github.com/go-kit/kit/log"
)

type loggingService struct {
	logger log.Logger
	Service
}

// NewLoggingService returns a new instance of a logging Service.
func NewLoggingService(logger log.Logger, s Service) Service {
	return &loggingService{logger, s}
}

func (s *loggingService) PagoIndividual(claveSerial string, rtn string, contrato string, fecha string, hora string, periodo string, impuesto int, concepto int, valorImpuesto float32, valorMulta float32, valorRecargo float32, valorInteres float32, refBancoExterna string, refBancoInterna string) (id string, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "PagoIndividual",
			"claveSerial", claveSerial,
			"rtn", rtn,
			"contrato", contrato,
			"fecha", fecha,
			"hora", hora,
			"periodo", periodo,
			"impuesto", impuesto,
			"concepto", concepto,
			"valorImpuesto", valorImpuesto,
			"valorMulta", valorMulta,
			"valorRecargo", valorRecargo,
			"valorInteres", valorInteres,
			"refBancoExterna", refBancoExterna,
			"refBancoInterna", refBancoInterna,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.PagoIndividual(claveSerial, rtn, contrato, fecha, hora, periodo, impuesto, concepto, valorImpuesto, valorMulta, valorRecargo, valorInteres, refBancoExterna, refBancoInterna)
}
