package ascii_web

import "strings"

func AsciiTableMaker(rawData string) [][]string {
	dataTmp := strings.Split(rawData, "\n")
	asciiTable := [][]string{}
	tmp := []string{}
	for _, l := range dataTmp {
		if len(l) > 0 {
			tmp = append(tmp, l)
		} else {
			if len(tmp) > 0 {
				asciiTable = append(asciiTable, tmp)
				tmp = []string{}
			}
		}
	}
	return asciiTable
}
