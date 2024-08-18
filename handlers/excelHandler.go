package handlers

import (
	"github.com/ertantorizkyf/gin-gorm-sample/usecases"
	"github.com/gin-gonic/gin"
)

type ExcelHandler struct {
	excelUsecase usecases.ExcelUsecase
}

func NewExcelHandler(excelUsecase usecases.ExcelUsecase) ExcelHandler {
	return ExcelHandler{
		excelUsecase: excelUsecase,
	}
}

func (h *ExcelHandler) GenerateExcelGet(c *gin.Context) {
	err := h.excelUsecase.GenerateSampleExcel()

	if err != nil {
		c.JSON(500, gin.H{
			"message": err,
		})

		return
	}

	c.JSON(200, gin.H{
		"data": "Excel file generated successfully",
	})
}
