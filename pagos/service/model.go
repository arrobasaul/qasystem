package pagos
//Pago modelo de Pago
type Pago struct {
	claveSerial string
	rtn string
	contrato string
	fecha string
	hora string
	periodo string
	impuesto int
	concepto int
	valorImpuesto float32
	valorMulta float32
	valorRecargo float32
	valorInteres float32
	refBancoExterna string
	refBancoInterna string
}
//Repository interface que sera implementada en la capa de transporte
type Repository interface {
	GenerarPago(pago Pago)(string, error)
}