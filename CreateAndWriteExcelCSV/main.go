package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

func main() {
	// f := excelize.NewFile()
	// // sheet2, _ := f.NewSheet("Sheet2")

	// f.SetCellValue("Sheet1", "A1", "Rool.No")

	// f.SetCellValue("Sheet1", "B1", "Student Name")
	// f.SetCellValue("Sheet1", "C1", "Course")

	// f.SetCellValue("Sheet1", "A2", "1")
	// f.SetCellValue("Sheet1", "B2", "Ravic Kumac")
	// f.SetCellValue("Sheet1", "C2", "B.Tech")

	// f.SetCellValue("Sheet1", "A3", "2")
	// f.SetCellValue("Sheet1", "B3", "TK NHU")
	// f.SetCellValue("Sheet1", "C3", "GOLANG DO")
	// // f.SetActiveSheet(sheet2)
	// err := f.SaveAs("MyStudent.xlsx")

	// if err != nil {
	// 	fmt.Println("err = ", err)
	// } else {
	// 	fmt.Println("File Created Welcome")

	// }
	

	// f, err := excelize.OpenFile("MyStudent.xlsx")

	// if err != nil {
	// 	fmt.Println("err open file = ", err)
	// }

	// rows, err := f.GetRows("Sheet1")

	// if err != nil {
	// 	fmt.Println("err get rows = ", err)
	// }

	// for _, row := range rows {
	// 	for _, col := range row {
	// 		fmt.Print(col, "\t")
	// 	}
	// 	fmt.Println()
	// }

	file, err := os.OpenFile("diamonds.csv", os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Panic(err)
	}

	df := dataframe.ReadCSV(file)
	fmt.Println(df)

	fprice := df.Filter(dataframe.F{Colname: "age", Comparator: series.Greater, Comparando: 3})
	fmt.Println(fprice)

	khfiler := df.Filter(dataframe.F{Colname: "address", Comparator: series.Eq, Comparando: "nt"})
	fmt.Println(khfiler)
}
