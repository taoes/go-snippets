package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/astaxie/beego/logs"
	"strconv"
)

func main() {
	writeExcelFile()
	readExcelFile()
}

func writeExcelFile() {
	f := excelize.NewFile()
	// 创建新数据表
	index := f.NewSheet("Sheet2")

	// 填充值
	for i := 0; i < 100; i++ {
		_ = f.SetCellValue("Sheet2", "A"+strconv.Itoa(i), i)
	}

	// 配置激活的表
	f.SetActiveSheet(index)

	// 保存excel到指定路径
	err := f.SaveAs("./Book1.xlsx")
	if err != nil {
		logs.Error("保存文件出现异常:", err)
	}
	logs.Info("Excel文件写入完成")
}

func readExcelFile() {
	fr, _ := excelize.OpenFile("/Users/zhoutao/code/python/excel/data/2019.1.xlsx")
	sheetMap := fr.GetSheetMap();
	for sheet, sheetName := range sheetMap {
		logs.Info("标序号:", sheet, "数据表名称", sheetName)
		rows, _ := fr.GetRows(sheetName)
		for _, row := range rows {
			for _, colCell := range row {
				fmt.Print(colCell, "|", "\t")
			}
			fmt.Println()
		}
	}
	logs.Info("Excel 读取完毕")
}
