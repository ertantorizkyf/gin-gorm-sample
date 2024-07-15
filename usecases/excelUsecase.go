package usecases

import (
	"fmt"
	"time"

	_ "image/png"

	"github.com/xuri/excelize/v2"
)

func GenerateSampleExcel() {
	xlsx := excelize.NewFile()

	savePath := "./sample.xlsx"
	sheetName := "Sample Sheet"
	xlsx.SetSheetName(xlsx.GetSheetName(0), sheetName)

	// SET HEADER BG
	headerStyle, err := xlsx.NewStyle(&excelize.Style{
		Fill: excelize.Fill{
			Type:    "pattern",
			Pattern: 1,
			Color:   []string{"#F1F6FF"},
		},
	})
	if err != nil {
		fmt.Println("ERROR", err.Error())
	}
	xlsx.SetCellStyle(sheetName, "A1", "M7", headerStyle)

	// SET HEADER TITLE AND IMAGE
	title := "Lorem Ipsum"
	titleStyle, err := xlsx.NewStyle(&excelize.Style{
		Fill: excelize.Fill{
			Type:    "pattern",
			Pattern: 1,
			Color:   []string{"#F1F6FF"},
		},
		Font: &excelize.Font{
			Size: 27,
			Bold: true,
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
	})
	if err != nil {
		fmt.Println("ERROR", err.Error())
	}
	xlsx.MergeCell(sheetName, "A2", "M4")
	xlsx.SetCellStr(sheetName, "A2", title)
	xlsx.SetCellStyle(sheetName, "A2", "A2", titleStyle)
	if err := xlsx.AddPicture(sheetName, "B1", "./assets/logo.png", &excelize.GraphicOptions{ScaleX: 0.1, ScaleY: 0.1}); err != nil {
		fmt.Println(err)
	}

	// SET DOWNLOAD DATE
	downloadDateStyle, err := xlsx.NewStyle(&excelize.Style{
		Fill: excelize.Fill{
			Type:    "pattern",
			Pattern: 1,
			Color:   []string{"#F1F6FF"},
		},
		Font: &excelize.Font{
			Size: 8,
		},
		Alignment: &excelize.Alignment{
			Horizontal: "right",
			Vertical:   "center",
		},
	})
	now := time.Now()
	downloadDate := "downloaded by (username) at " + now.Format("02-Jan-2006 15:05")
	xlsx.SetCellStr(sheetName, "M1", downloadDate)
	xlsx.SetCellStyle(sheetName, "M1", "M1", downloadDateStyle)

	// SET TABLE TITLES
	tableTitleStyle, err := xlsx.NewStyle(&excelize.Style{
		Fill: excelize.Fill{
			Type:    "pattern",
			Pattern: 1,
			Color:   []string{"#185FFF"},
		},
		Font: &excelize.Font{
			Size:  10,
			Bold:  true,
			Color: "#FFFFFF",
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "top",
		},
		Border: []excelize.Border{
			{
				Type:  "top",
				Color: "#000000",
				Style: 4,
			},
			{
				Type:  "left",
				Color: "#000000",
				Style: 4,
			},
			{
				Type:  "right",
				Color: "#000000",
				Style: 4,
			},
			{
				Type:  "bottom",
				Color: "#000000",
				Style: 4,
			},
		},
	})
	tableTitles := []map[string]interface{}{
		{"column": "A8", "value": "No"},
		{"column": "B8", "value": "Title B"},
		{"column": "C8", "value": "Title C"},
		{"column": "D8", "value": "Title D"},
		{"column": "E8", "value": "Title E"},
		{"column": "F8", "value": "Title F"},
		{"column": "G8", "value": "Title G"},
		{"column": "H8", "value": "Title H"},
		{"column": "I8", "value": "Title I"},
		{"column": "J8", "value": "Title J"},
		{"column": "K8", "value": "Title K"},
	}
	for _, item := range tableTitles {
		title, _ := item["value"].(string)
		col, _ := item["column"].(string)
		xlsx.SetCellStr(sheetName, col, title)
	}
	lastCol, _ := tableTitles[len(tableTitles)-1]["column"].(string)
	xlsx.SetCellStyle(sheetName, "A8", lastCol, tableTitleStyle)

	// SET TABLE VALUES
	tableValueStyle, err := xlsx.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Size:  10,
			Color: "#000000",
		},
		Border: []excelize.Border{
			{
				Type:  "top",
				Color: "#000000",
				Style: 4,
			},
			{
				Type:  "left",
				Color: "#000000",
				Style: 4,
			},
			{
				Type:  "right",
				Color: "#000000",
				Style: 4,
			},
			{
				Type:  "bottom",
				Color: "#000000",
				Style: 4,
			},
		},
	})
	tableValues := []map[string]interface{}{
		{"bVal": "Lorem", "cVAl": "Ipsum", "dVal": "Dolor", "eVal": "Sit", "fVal": "Amet", "gVal": "Consectetur", "hVal": "Adipiscing", "iVal": "Elit", "jVal": "In", "kVal": "Ac"},
		{"bVal": "The", "cVAl": "Quick", "dVal": "Brown", "eVal": "Fox", "fVal": "Jumps", "gVal": "Over", "hVal": "The", "iVal": "Lazy", "jVal": "Dog", "kVal": "End"},
	}
	for idx, item := range tableValues {
		bVal, _ := item["bVal"].(string)
		cVal, _ := item["cVal"].(string)
		dVal, _ := item["dVal"].(string)
		eVal, _ := item["eVal"].(string)
		fVal, _ := item["fVal"].(string)
		gVal, _ := item["gVal"].(string)
		hVal, _ := item["hVal"].(string)
		iVal, _ := item["iVal"].(string)
		jVal, _ := item["jVal"].(string)
		kVal, _ := item["kVal"].(string)
		xlsx.SetCellValue(sheetName, fmt.Sprintf("A%d", idx+9), idx+1)
		xlsx.SetCellValue(sheetName, fmt.Sprintf("B%d", idx+9), bVal)
		xlsx.SetCellValue(sheetName, fmt.Sprintf("C%d", idx+9), cVal)
		xlsx.SetCellValue(sheetName, fmt.Sprintf("D%d", idx+9), dVal)
		xlsx.SetCellValue(sheetName, fmt.Sprintf("E%d", idx+9), eVal)
		xlsx.SetCellValue(sheetName, fmt.Sprintf("F%d", idx+9), fVal)
		xlsx.SetCellValue(sheetName, fmt.Sprintf("G%d", idx+9), gVal)
		xlsx.SetCellValue(sheetName, fmt.Sprintf("H%d", idx+9), hVal)
		xlsx.SetCellValue(sheetName, fmt.Sprintf("I%d", idx+9), iVal)
		xlsx.SetCellValue(sheetName, fmt.Sprintf("J%d", idx+9), jVal)
		xlsx.SetCellValue(sheetName, fmt.Sprintf("K%d", idx+9), kVal)
	}
	xlsx.SetCellStyle(sheetName, "A9", "K"+fmt.Sprint(8+len(tableValues)), tableValueStyle)

	err = xlsx.SaveAs(savePath)
	if err != nil {
		fmt.Println(err)
	}
}
