package handlers

import (
	"github.com/ertantorizkyf/gin-gorm-sample/usecases"
	"github.com/gin-gonic/gin"
)

type PdfHandler struct {
	pdfUsecase usecases.PdfUsecase
}

func NewPdfHandler(pdfUsecase usecases.PdfUsecase) PdfHandler {
	return PdfHandler{
		pdfUsecase: pdfUsecase,
	}
}

func (h *PdfHandler) GeneratePDFGet(c *gin.Context) {
	err := h.pdfUsecase.GenerateSamplePDF()
	if err != nil {
		c.JSON(500, gin.H{
			"message": err,
		})

		return
	}

	c.JSON(200, gin.H{
		"data": "PDF generated successfully",
	})
}
