package common

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"mime/multipart"
	"strconv"
	"strings"
	"time"
)

func NewExcel(sheetName, filePath string) (*Excel, error) {
	e := new(Excel)
	err := e.Init(sheetName, filePath)
	if err != nil {
		return nil, err
	}
	return e, nil
}

type Excel struct {
	file     *xlsx.File
	sheet    *xlsx.Sheet
	rowNames []string
	filePath string
}

func (this *Excel) Init(sheetName, filePath string) (err error) {
	this.filePath = filePath
	this.file = xlsx.NewFile()
	if len(sheetName) <= 0 {
		sheetName = "Sheet1"
	}
	this.sheet, err = this.file.AddSheet(sheetName)

	if err != nil {
		return
	}
	return
}

func (this *Excel) AddHeader(data []string) error {
	row := this.sheet.AddRow()
	for i := 0; i < len(data); i++ {
		row.AddCell().SetValue(data[i])
	}
	if err := this.file.Save(this.filePath); err != nil {
		return err
	}
	this.rowNames = data
	return nil
}

func (this *Excel) Append(data []string) error {
	row := this.sheet.AddRow()
	for i := 0; i < len(data); i++ {
		row.AddCell().SetValue(data[i])
	}
	if err := this.file.Save(this.filePath); err != nil {
		return err
	}
	return nil
}

func (this *Excel) AppendCache(data []string) error {
	row := this.sheet.AddRow()
	for i := 0; i < len(data); i++ {
		row.AddCell().SetValue(data[i])
	}
	//if err := this.file.Save(this.filePath); err != nil {
	//	return err
	//}
	return nil
}

func (this *Excel) Sync() error {
	if err := this.file.Save(this.filePath); err != nil {
		return err
	}
	return nil
}

func (this *Excel) AppendCells(Cells []*xlsx.Cell) error {
	row := this.sheet.AddRow()
	for i := 0; i < len(Cells); i++ {
		row.AddCell().SetValue(Cells[i].Value)
	}
	if err := this.file.Save(this.filePath); err != nil {
		return err
	}
	return nil
}

func CellsToArr(Cells []*xlsx.Cell) []string {
	arr := make([]string, 0, len(Cells))
	for i := 0; i < len(Cells); i++ {
		if i==0 {
			arr = append(arr, convertToFormatDay(Cells[i].Value))
			continue
		}
		arr = append(arr, Cells[i].Value)
	}
	return arr
}

// excel日期字段格式化 yyyy-mm-dd
func convertToFormatDay(excelDaysString string)string{
	// 2006-01-02 距离 1900-01-01的天数
	baseDiffDay := 38719 //在网上工具计算的天数需要加2天，什么原因没弄清楚
	curDiffDay := excelDaysString
	b,_ := strconv.Atoi(curDiffDay)
	// 获取excel的日期距离2006-01-02的天数
	realDiffDay := b - baseDiffDay
	//fmt.Println("realDiffDay:",realDiffDay)
	// 距离2006-01-02 秒数
	realDiffSecond := realDiffDay * 24 * 3600
	//fmt.Println("realDiffSecond:",realDiffSecond)
	// 2006-01-02 15:04:05距离1970-01-01 08:00:00的秒数 网上工具可查出
	baseOriginSecond := 1136185445
	resultTime := time.Unix(int64(baseOriginSecond + realDiffSecond), 0).Format("2006-01-02")
	return resultTime
}

func ReadFromFile(excelFileName , sheetName string) ([]string, *xlsx.Sheet, error) {
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		return nil, nil, err
	}
	if len(sheetName) <= 0 {
		sheetName = "Sheet1"
	}
	sheet, ok := xlFile.Sheet[sheetName]
	if !ok {
		return nil, nil, fmt.Errorf("nofind Sheet1")
	}

	rowNames := make([]string, 0)
	for _, cell := range sheet.Rows[0].Cells {
		text := cell.String()
		text = strings.Replace(text, " ", "", -1)
		rowNames = append(rowNames, text)
	}
	if len(rowNames) <= 1 {
		return nil, nil, fmt.Errorf("in.xlsx format err")
	}
	return rowNames, sheet, nil
}

func ReadFromReader(r multipart.File ,size int64,  sheetName string) ([]string, *xlsx.Sheet, error) {
	xlFile, err := xlsx.OpenReaderAt(r, size)
	if err != nil {
		return nil, nil, err
	}
	if len(sheetName) <= 0 {
		sheetName = "Sheet1"
	}
	sheet, ok := xlFile.Sheet[sheetName]
	if !ok {
		return nil, nil, fmt.Errorf("nofind Sheet1")
	}

	rowNames := make([]string, 0)
	for _, cell := range sheet.Rows[0].Cells {
		text := cell.String()
		text = strings.Replace(text, " ", "", -1)
		rowNames = append(rowNames, text)
	}
	if len(rowNames) <= 1 {
		return nil, nil, fmt.Errorf("in.xlsx format err")
	}
	return rowNames, sheet, nil
}