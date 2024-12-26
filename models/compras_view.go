package models

import "encoding/base64"

// ComprasExtraordinariasView estructura para visualizar registros
type ComprasExtraordinariasView struct {
	ID               int     `json:"id"`
	Titulo           string  `json:"titulo"`
	Aprobador        string  `json:"aprobador"`
	Aprobador2       string  `json:"aprobador2"`
	EstadoAprobacion string  `json:"estado_aprobacion"`
	EstadoCompra     string  `json:"estado_compra"`
	Descripcion      string  `json:"descripcion"`
	Fecha            string  `json:"fecha"`
	Link1            string  `json:"link1"`
	Link2            string  `json:"link2"`
	Link3            string  `json:"link3"`
	Direccion        string  `json:"direccion"`
	Servicio         string  `json:"servicio"`
	TipoPeticion     string  `json:"tipo_peticion"`
	Urgencia         string  `json:"urgencia"`
	Proyecto         string  `json:"proyecto"`
	DatosAdjunto     string  `json:"datos_adjunto"` // Base64 del archivo adjunto
	Coste            float64 `json:"coste"`
	Usuario          string  `json:"usuario"`
}

// ToBase64 convierte []byte a una cadena base64
func (c *ComprasExtraordinariasView) ToBase64(data []byte) {
	if len(data) > 0 {
		c.DatosAdjunto = base64.StdEncoding.EncodeToString(data)
	}
}
