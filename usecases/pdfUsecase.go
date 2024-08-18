package usecases

import (
	"fmt"
	"strings"

	_ "image/png"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

type PdfUsecase struct{}

func NewPdfUsecase() PdfUsecase {
	return PdfUsecase{}
}

func (uc *PdfUsecase) GenerateSamplePDF() error {
	// Create new PDF generator
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		fmt.Println(err)
		return err
	}

	// Set global options
	pdfg.Dpi.Set(300)
	pdfg.Orientation.Set(wkhtmltopdf.OrientationPortrait)

	// Add to document
	html := "<html>Hi</html>"
	pdfg.AddPage(wkhtmltopdf.NewPageReader(strings.NewReader(html)))

	// Create PDF document in internal buffer
	err = pdfg.Create()
	if err != nil {
		fmt.Println(err)
		return err
	}

	// Write buffer contents to file on disk
	err = pdfg.WriteFile("./sample.pdf")
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("PDF Generated Successfully")
	return nil
}
