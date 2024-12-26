package models

import "mime/multipart"

type ComprasExtraordinarias struct {
	Titulo            string                `json:"titulo" binding:"required"`
	Aprobador         string                `json:"-"`
	Aprobador2        string                `json:"-"`
	Descripcion       string                `json:"descripcion" binding:"required"`
	Link1             string                `json:"link1" binding:"required"`
	Link2             string                `json:"link2" binding:"required"`
	Link3             string                `json:"link3" binding:"required"`
	Direccion         string                `json:"direccion" binding:"required"`
	Servicio          string                `json:"servicio" binding:"required"`
	TipoPeticion      string                `json:"tipoPeticion" binding:"required"`
	Urgencia          string                `json:"urgencia" binding:"required"`
	Proyecto          string                `json:"proyecto" binding:"required"`
	EstadoAprobacion  string                `json:"-"` // No se espera desde el cliente
	EstadoAprobacion2 string                `json:"-"` // No se espera desde el cliente
	EstadoCompra      string                `json:"-"` // No se espera desde el cliente
	Fecha             string                `json:"-"` // No se espera desde el cliente
	DatosAdjunto      *multipart.FileHeader `json:"datosAdjunto"`
	NombreAdjunto     string                `json:"nombreAdjunto"`
	Coste             float64               `json:"coste"`
	Usuario           string                `json:"usuario" binding:"required"`
}
