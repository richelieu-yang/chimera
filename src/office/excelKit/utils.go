package excelKit

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v3/src/core/errorKit"
	"regexp"
	"strconv"
)

// ConvertAxisToRowCol
/*
e.g.
("C2") =>
*/
func ConvertAxisToRowCol(axis string) (int, int, error) {
	re := regexp.MustCompile(`^([A-Z]+)(\d+)$`)
	if !re.MatchString(axis) {
		return 0, 0, errorKit.New("axis(%s) is invalid", axis)
	}
	var tmp = re.FindAllStringSubmatch(axis, -1)
	if len(tmp) != 1 {
		return 0, 0, errorKit.New("axis(%s) is invalid", axis)
	}
	var s []string = tmp[0]
	if len(s) != 3 {
		return 0, 0, errorKit.New("axis(%s) is invalid", axis)
	}

	colStr := s[1]
	rowStr := s[2]

	fmt.Println(rowStr, colStr)

	// TODO:
	return 0, 0, nil
}

// ConvertRowColToAxis
/*
e.g.
(1, 2) => "C2", nil
*/
func ConvertRowColToAxis(row, col int) (string, error) {
	str, err := ConvertColToString(col)
	if err != nil {
		return "", err
	}
	str1, err := ConvertRowToString(row)
	if err != nil {
		return "", err
	}
	return str + str1, nil
}

//func ConvertStringToRow(str string) (int, error) {
//	row, err := strconv.Atoi(str)
//	if err != nil {
//		return 0, err
//	}
//	row--
//	if row < 0 || row > MaxRow {
//		return 0, errorKit.New("row(%d) is invalid", row)
//	}
//	return row, nil
//}
//
//func ConvertStringToCol(str string) (int, error) {
//
//}

func ConvertRowToString(row int) (string, error) {
	if row < 0 || row > MaxRow {
		return "", errorKit.New("row(%d) is invalid", row)
	}
	return strconv.Itoa(row + 1), nil
}

// ConvertColToString
/*
e.g.
(0) 	=> "A", nil
(16368) => "XEO", nil
*/
func ConvertColToString(col int) (string, error) {
	if col < 0 || col > MaxCol {
		return "", errorKit.New("col(%d) is invalid", col)
	}

	str := ""
	for col >= 0 {
		// 由于确定不存在中文字符的情况，此处使用 byte 而非 rune
		var tmp = byte('A' + col%26)
		str = string(tmp) + str
		col = col/26 - 1
	}
	return str, nil
}
