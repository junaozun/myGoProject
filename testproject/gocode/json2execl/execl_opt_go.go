package main

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/tealeg/xlsx"
)

func Export() {
	file := xlsx.NewFile()                      // NewWriter 创建一个Excel写操作实例
	sheet, err := file.AddSheet("student_list") // 表实例
	if err != nil {
		fmt.Printf(err.Error())
	}
	stus := getStudents() // add data

	headers := []*HeaderColumn{
		{Field: "Name", Title: "姓名"},
		{Field: "Age", Title: "年龄"},
		{Field: "Phone", Title: "电话"},
		{Field: "Gender", Title: "性别"},
		{Field: "Mail", Title: "邮箱"},
	}
	style := map[string]float64{
		"Name":   2.0,
		"Age":    2.0,
		"Phone":  2.0,
		"Gender": 2.0,
		"Mail":   2.0,
	}
	sheet, _ = SetHeader(sheet, headers, style)

	for _, stu := range stus {
		data := make(map[string]string)
		data["Name"] = stu.Name
		data["Age"] = strconv.Itoa(stu.Age)
		data["Phone"] = stu.Phone
		data["Gender"] = stu.Gender
		data["Mail"] = stu.Mail

		row := sheet.AddRow()
		row.SetHeightCM(0.8)
		for _, field := range headers {
			row.AddCell().Value = data[field.Field]
		}
	}
	outFile := "./out_student.xlsx"
	err = file.Save(outFile)
	if err != nil {
		fmt.Printf(err.Error())
	}
	fmt.Println("export success")
}

func getStudents() []Student {
	students := make([]Student, 0)
	for i := 0; i < 10; i++ {
		stu := Student{}
		stu.Name = "mmmd" + strconv.Itoa(i+1)
		stu.Mail = stu.Name + "@studygolang.com"
		stu.Phone = "10086" + strconv.Itoa(i)
		stu.Age = 18 + i
		if i%2 == 0 {
			stu.Gender = "女"
		} else {
			stu.Gender = "男"
		}

		students = append(students, stu)
	}
	return students
}

// SetHeader 写模式下，设置字段表头和字段顺序
// 参数 header 为表头和字段映射关系，如：HeaderColumn{Field:"title", Title: "标题", Requre: true}
// 参数 width  为表头每列的宽度，单位 CM：map[string]float64{"title": 0.8}
func SetHeader(sheet *xlsx.Sheet, header []*HeaderColumn, width map[string]float64) (*xlsx.Sheet, error) {
	if len(header) == 0 {
		return nil, errors.New("Excel.SetHeader 错误: 表头不能为空")
	}

	// 表头样式
	style := xlsx.NewStyle()

	font := xlsx.DefaultFont()
	font.Bold = true

	alignment := xlsx.DefaultAlignment()
	alignment.Vertical = "center"

	style.Font = *font
	style.Alignment = *alignment

	style.ApplyFont = true
	style.ApplyAlignment = true

	// 设置表头字段
	row := sheet.AddRow()
	row.SetHeightCM(1.0)
	row_w := make([]string, 0)
	for _, column := range header {
		row_w = append(row_w, column.Field)
		cell := row.AddCell()
		cell.Value = column.Title
		cell.SetStyle(style) // 设置单元样式
	}

	// 表格列，宽度
	if len(row_w) > 0 {
		for k, v := range row_w {
			if width[v] > 0.0 {
				sheet.SetColWidth(k, k, width[v]*10)
			}
		}
	}

	return sheet, nil
}

type Student struct {
	Name   string
	Age    int
	Phone  string
	Gender string
	Mail   string
}

// HeaderColumn 表头字段定义
type HeaderColumn struct {
	Field string // 字段，数据映射到的数据字段名
	Title string // 标题，表格中的列名称
}
