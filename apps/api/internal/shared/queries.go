package shared

import (
	"strconv"
	"strings"
)

func mergeStringsArray(builder *strings.Builder, mergeStrings *[]string) {
	for i := range len(*mergeStrings) {
		builder.WriteString((*mergeStrings)[i])
	}
}

func BuildLimitOffset(page int, pageSize int) (int, int) {
	limit := pageSize
	offset := (page - 1) * pageSize
	return limit, offset
}

func Param(tableName string, param string) string {
	var builder strings.Builder
	mergeStrings := []string{
		tableName, ".", param, ", ",
	}
	mergeStringsArray(&builder, &mergeStrings)
	return builder.String()
}

func Params(tableNames []string, params []string) string {
	if len(tableNames) == 0 || len(params) == 0 {
		panic("tableNames and params should not be empty")
	}
	if len(tableNames) != len(params) {
		panic("TableNames and params should be the same lenght")
	}
	var builder strings.Builder
	for i := range len(tableNames) {
		builder.WriteString(Param(tableNames[i], params[i]))
	}
	return builder.String()
}

func LastParam(tableName string, param string) string {
	var builder strings.Builder
	mergeStrings := []string{
		tableName, ".", param,
	}
	mergeStringsArray(&builder, &mergeStrings)
	return builder.String()
}

func AppendValue(tableArray *[]string, value string, times int) {
	for i := 0; i < times; i++ {
		*tableArray = append(*tableArray, value)
	}
}

func Select(query *string, tableNames []string, params []string, fromTable string) {
	last := len(tableNames) - 1

	var selectPart strings.Builder
	if last > 0 {
		selectPart.WriteString(Params(tableNames[:last], params[:last]))
	}
	selectPart.WriteString(LastParam(tableNames[last], params[last]))

	*query = "SELECT " + selectPart.String() + " FROM " + fromTable + " "
}

func OrderBy(query *string, sortBy string, sortOrder string) {
	var builder strings.Builder
	mergeStrings := []string{
		" ORDER BY ",
		sortBy,
		" ",
		sortOrder,
	}
	mergeStringsArray(&builder, &mergeStrings)
	*query += builder.String()
}

func JoinTable(
	query *string,
	mainTable string,
	joinTable string,
	mainTableKey string,
	joinTableKey string) {
	var builder strings.Builder
	mergeStrings := []string{
		"JOIN ",
		joinTable,
		" ON ",
		LastParam(mainTable, mainTableKey),
		" = ",
		LastParam(joinTable, joinTableKey),
		" ",
	}
	mergeStringsArray(&builder, &mergeStrings)
	*query += builder.String()
}

func Where(
	query *string,
	tableNames []string,
	tableColumns []string,
	values []string,
	conditions []string,
	operators []string) {
	var builder strings.Builder
	mergeStrings := []string{
		"WHERE ",
	}
	for i := range len(tableNames) {
		mergeStrings = append(mergeStrings, LastParam(tableNames[i], tableColumns[i]))
		mergeStrings = append(mergeStrings, " ", conditions[i], " ", "'", values[i], "'")
		if i < len(tableNames)-1 {
			mergeStrings = append(mergeStrings, " ", operators[i], " ")
		}
	}
	mergeStringsArray(&builder, &mergeStrings)
	*query += builder.String()
}

func Limit(query *string, limit int, offset int) {
	var builder strings.Builder
	mergeStrings := []string{
		" LIMIT ",
		strconv.Itoa(limit),
		" OFFSET ",
		strconv.Itoa(offset),
	}
	mergeStringsArray(&builder, &mergeStrings)
	*query += builder.String()
}
