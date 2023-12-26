package libs

import (
	"log"

	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/image"
	"github.com/johnfercher/maroto/v2/pkg/components/line"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/border"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/consts/linestyle"
	"github.com/johnfercher/maroto/v2/pkg/consts/orientation"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

func GetMaroto() core.Maroto {
	cfg := config.NewBuilder().
		WithPageNumber("Página {current} de {total}", props.RightBottom).
		WithMargins(10, 5, 10).
		Build()

	mrt := maroto.New(cfg)
	m := maroto.NewMetricsDecorator(mrt)

	err := m.RegisterHeader(getPageHeader())
	if err != nil {
		log.Fatal(err)
	}

	m.AddRows(getTitle())
	m.AddRows(getSeparatorLine(orientation.Horizontal))
	//m.AddRow(3) //separator blankline
	//m.AddRows(getCampus())
	//TableStaff(m)
	return m
}

func TableStaff(m core.Maroto) {
	headerTable := func(paragrahp string, sizeCol int, top float64) core.Col {
		return col.New(sizeCol).Add(
			text.New(paragrahp, props.Text{
				Top:    top,
				Family: "",
				Size:   11,
				Align:  align.Center,
				Color: &props.Color{
					Red:   255,
					Green: 255,
					Blue:  255,
				},
			}),
		).WithStyle(&props.Cell{
			BackgroundColor: &props.Color{
				Red:   86,
				Green: 86,
				Blue:  86,
			},
			BorderColor:     &props.Color{},
			BorderType:      border.Full,
			BorderThickness: 0,
			LineStyle:       linestyle.Solid,
		})
	}

	bodyTable := func(paragrahp string, sizeCol int, top float64, assigned bool) core.Col {
		withColorBack := &props.Color{
			Red:   255,
			Green: 255,
			Blue:  255,
		}
		c := 190
		if assigned {
			withColorBack.Red = c
			withColorBack.Green = c
			withColorBack.Blue = c
		}

		return col.New(sizeCol).Add(
			text.New(paragrahp, props.Text{
				Top:    top,
				Family: "",
				Size:   11,
				Align:  align.Center,
			}),
		).WithStyle(&props.Cell{
			BackgroundColor: withColorBack,
			BorderType:      border.Full,
			BorderThickness: 0,
			LineStyle:       linestyle.Solid,
		})
	}

	m.AddRow(
		10,
		headerTable("N°", 1, 2),
		headerTable("NOMBRES Y APELLIDOS", 5, 2),
		headerTable("CARGO", 3, 2),
		headerTable("ASIGNACIÓN DE SERVICIO", 3, 1),
	)
	//content
	m.AddRow(
		5,
		bodyTable("1", 1, 0, true),
		bodyTable("Luis E. Quispe Alegre", 5, 0, true),
		bodyTable("Analista PAD I", 3, 0, true),
		bodyTable("", 3, 0, true),
	)
	m.AddRow(
		5,
		bodyTable("2", 1, 0, false),
		bodyTable("Luis E. Quispe Alegre", 5, 0, false),
		bodyTable("Analista PAD I", 3, 0, false),
		bodyTable("", 3, 0, false),
	)
}

func getCampus() core.Row {
	return row.New(20).Add(
		col.New(4).Add(
			text.New("Ciudad Universitaria", props.Text{
				Top:   7,
				Style: fontstyle.Bold,
				Size:  11,
				Align: align.Center,
			}),
		),
		//
		col.New(3).Add(
			text.New("ID HOJA DE SERVICIO", props.Text{
				Top:    2,
				Family: "",
				Size:   11,
				Align:  align.Right,
			}),
			text.New("FECHA INICIO", props.Text{
				Top:    7,
				Family: "",
				Size:   11,
				Align:  align.Right,
			}),
			text.New("FECHA FIN", props.Text{
				Top:    12,
				Family: "",
				Size:   11,
				Align:  align.Right,
			}),
		),
		//
		col.New(5).Add(
			text.New("114065ac-0b2a-4e22-b9d5-584c6abca11b", props.Text{
				Top:       2,
				Size:      11,
				Align:     align.Center,
				Hyperlink: new(string),
			}),
			text.New("2023-12-23", props.Text{
				Top:   7,
				Size:  11,
				Align: align.Center,
			}),
			text.New("2023-12-23", props.Text{
				Top:   12,
				Size:  11,
				Align: align.Center,
			}),
		),
	)
}

func getTitle() core.Row {
	return text.NewRow(10, "OFICINA DE TECNOLOGIAS DE LA INFORMACIÓN", props.Text{
		Top:   3,
		Style: fontstyle.Bold,
		Size:  11,
		Align: align.Center,
	})
}
func getSeparatorLine(o orientation.Type) core.Row {
	return line.NewRow(2, props.Line{
		Style:         linestyle.Solid,
		Thickness:     1,
		Orientation:   o,
		OffsetPercent: 100,
		SizePercent:   100,
	})
}

func getPageHeader() core.Row {
	return row.New(20).Add(
		image.NewFromFileCol(1, "public/logofirma.png", props.Rect{
			Center:  true,
			Percent: 80,
		}),
		//m.AddRows()
		col.New(10).Add(
			text.New(`"UNIVERSIDAD NACIONAL AMAZONICA DE MADRE DE DIOS"`, props.Text{
				Top:    2,
				Family: "",
				Style:  fontstyle.Bold,
				Size:   11,
				Align:  align.Center,
			}), //"Madre de Dios Capital de la Biodiversidad del Peru"
			text.New(`"Madre de Dios Capital de la Biodiversidad del Perú"`, props.Text{
				Top:    7,
				Family: "",
				Style:  fontstyle.Bold,
				Size:   11,
				Align:  align.Center,
			}), //"Madre de Dios Capital de la Biodiversidad del Peru"
			text.New(`"Año de la unidad, la paz y el desarrollo"`, props.Text{
				Top:    11,
				Family: "",
				Style:  fontstyle.Bold,
				Size:   11,
				Align:  align.Center,
			}), //"Madre de Dios Capital de la Biodiversidad del Peru"
		),
		image.NewFromFileCol(1, "public/logofirma.png", props.Rect{
			Center:  true,
			Percent: 80,
		}),
	)
}
