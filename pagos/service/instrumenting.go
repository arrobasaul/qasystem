package pagos

import (
	"time"

	"github.com/go-kit/kit/metrics"
)

type instrumentingService struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	Service
}

// NewInstrumentingService returns an instance of an instrumenting Service.
func NewInstrumentingService(counter metrics.Counter, latency metrics.Histogram, s Service) Service {
	return &instrumentingService{
		requestCount:   counter,
		requestLatency: latency,
		Service:        s,
	}
}

func (s *instrumentingService) PagoIndividual(claveSerial string, rtn string, contrato string, fecha string, hora string, periodo string, impuesto int, concepto int, valorImpuesto float32, valorMulta float32, valorRecargo float32, valorInteres float32, refBancoExterna string, refBancoInterna string) (id string, err error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "pago").Add(1)
		s.requestLatency.With("method", "pago").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.PagoIndividual(claveSerial, rtn, contrato, fecha, hora, periodo, impuesto, concepto, valorImpuesto, valorMulta, valorRecargo, valorInteres, refBancoExterna, refBancoInterna)
}
