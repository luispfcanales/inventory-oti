package services

import (
	"log"

	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/luispfcanales/inventory-oti/libs"
	"github.com/luispfcanales/inventory-oti/ports"
)

type PDFSrv struct {
	//repo   ports.StorageFileService
}

func (pdfsrv *PDFSrv) PreviewRenderPDF() []byte {
	m := libs.GetMaroto()

	pdfsrv.rowSpace(m)
	pdfsrv.generateTableStaff(m)

	d, err := m.Generate()
	if err != nil {
		log.Fatal(err)
	}

	return d.GetBytes()
}

func (pdfsrv *PDFSrv) List(_ string) {
	panic("not implemented") // TODO: Implement
}

func NewPDFSrv() ports.FileService {
	return &PDFSrv{}
}

func (pdfsrv *PDFSrv) generateTableStaff(cm core.Maroto) {
	libs.TableStaff(cm)
}
func (pdfsrv *PDFSrv) rowSpace(cm core.Maroto) {
	cm.AddRow(3)
}
