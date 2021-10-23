package indodate

import (
	"strings"
	"time"
)

func LetterDate(tanggal time.Time) string {
	monthOnIndo := ""
	output := ""
	getDateFromTime := strings.Split(tanggal.String(), " ")
	parsingYMD := strings.Split(getDateFromTime[0], "-")
	year := parsingYMD[0]
	month := parsingYMD[1]
	day := parsingYMD[2]

	switch month {
	case "01":
		monthOnIndo = "Januari"
	case "02":
		monthOnIndo = "Februari"
	case "03":
		monthOnIndo = "Maret"
	case "04":
		monthOnIndo = "April"
	case "05":
		monthOnIndo = "Mei"
	case "06":
		monthOnIndo = "Juni"
	case "07":
		monthOnIndo = "Juli"
	case "08":
		monthOnIndo = "Agustus"
	case "09":
		monthOnIndo = "September"
	case "10":
		monthOnIndo = "Oktober"
	case "11":
		monthOnIndo = "November"
	case "12":
		monthOnIndo = "Desember"
	default:
		monthOnIndo = ""
	}
	output = day + " " + monthOnIndo + " " + year
	return output
}

func DateWithSlash(tanggal time.Time) string {
	output := ""
	getDateFromTime := strings.Split(tanggal.String(), " ")
	parsingYMD := strings.Split(getDateFromTime[0], "-")
	year := parsingYMD[0]
	month := parsingYMD[1]
	day := parsingYMD[2]
	output = day + "/" + month + "/" + year

	return output

}

func DateWithMin(tanggal time.Time) string {
	output := ""
	getDateFromTime := strings.Split(tanggal.String(), " ")
	parsingYMD := strings.Split(getDateFromTime[0], "-")
	year := parsingYMD[0]
	month := parsingYMD[1]
	day := parsingYMD[2]
	output = day + "-" + month + "-" + year

	return output

}
