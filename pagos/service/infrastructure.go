package pagos

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// ErrInvalidArgumentPagos Es Retornado cuando uno o mas argumentos son invalidos.
var ErrInvalidArgumentPagos = errors.New("invalid argument")

type repository struct {
	Activo bool
}

func (r *repository) GenerarPago(pago Pago) (string, error) {
	resp, err := http.PostForm("http://172.16.23.61/WebServices/Tributos/DeiTransaccion.asmx/ProcesarRop", url.Values{
		"claveSerial":     {pago.claveSerial},
		"rtn":             {pago.rtn},
		"contrato":        {pago.contrato},
		"fecha":           {pago.fecha},
		"hora":            {pago.hora},
		"periodo":         {pago.periodo},
		"impuesto":        {fmt.Sprintf("%d", pago.impuesto)},
		"concepto":        {fmt.Sprintf("%d", pago.concepto)},
		"valorImpuesto":   {fmt.Sprintf("%f", pago.valorImpuesto)},
		"valorMulta":      {fmt.Sprintf("%f", pago.valorMulta)},
		"valorRecargo":    {fmt.Sprintf("%f", pago.valorRecargo)},
		"valorInteres":    {fmt.Sprintf("%f", pago.valorInteres)},
		"refBancoExterna": {pago.refBancoExterna},
		"refBancoInterna": {pago.refBancoInterna},
	})
	if err != nil {
		return "", ErrInvalidArgumentPagos
	}
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(body)
	fmt.Println(body)
	if body != nil {
		return "", ErrInvalidArgumentPagos
	}
	return "body", nil
}

// NewCargoRepository returns a new instance of a in-memory cargo repository.
func NewPagoIndividualRepository() *repository {
	return &repository{Activo: true}
}
