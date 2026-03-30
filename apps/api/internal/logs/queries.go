package logs

import (
	"context"
	"strconv"

	"github.com/rs/zerolog/log"

	"github.com/bborutenko/LogForge/internal/config"
	"github.com/bborutenko/LogForge/internal/endpoint_metrics"
	"github.com/bborutenko/LogForge/internal/shared"
)

func buildListLogsQuery(params LogsQueryParams, filterBy shared.FilterByQueryParams) (string, bool) {
	joinTable := len(filterBy.Endpoints) != 0 || len(filterBy.UserIDs) != 0 || len(filterBy.StatusCodes) != 0

	l := "log_forge.logs"
	em := "log_forge.endpoint_metrics"
	query := ""

	whereTableNames := []string{}
	whereColumnNames := []string{}
	whereConditions := []string{}
	whereOperators := []string{}
	whereAllValues := []string{}

	limit, offset := shared.BuildLimitOffset(params.Page, params.PageSize)

	queryTables := []string{}
	queryParams := []string{"timestamp", "level", "message", "meta"}

	shared.AppendValue(&queryTables, l, len(queryParams))

	if joinTable {
		appendQueryParams := []string{"endpoint", "user_id", "method", "status_code", "response_time_ms", "meta"}
		shared.AppendValue(&queryTables, em, len(appendQueryParams))
		queryParams = append(queryParams, appendQueryParams...)
	}

	shared.Select(&query, queryTables, queryParams, l)

	if joinTable {
		shared.JoinTable(&query, l, em, "endpoint_metrics_id", "id")
	}

	for _, userID := range filterBy.UserIDs {
		whereTableNames = append(whereTableNames, em)
		whereColumnNames = append(whereColumnNames, "user_id")
		whereConditions = append(whereConditions, "=")
		whereAllValues = append(whereAllValues, userID)
	}

	for _, endpoint := range filterBy.Endpoints {
		whereTableNames = append(whereTableNames, em)
		whereColumnNames = append(whereColumnNames, "endpoint")
		whereConditions = append(whereConditions, "=")
		whereAllValues = append(whereAllValues, endpoint)
	}

	for _, statusCode := range filterBy.StatusCodes {
		whereTableNames = append(whereTableNames, em)
		whereColumnNames = append(whereColumnNames, "status_code")
		whereConditions = append(whereConditions, "=")
		whereAllValues = append(whereAllValues, strconv.Itoa(statusCode))
	}

	for i := 0; i < len(whereAllValues)-1; i++ {
		whereOperators = append(whereOperators, "OR")
	}

	if len(whereAllValues) > 0 {
		shared.Where(&query, whereTableNames, whereColumnNames, whereAllValues, whereConditions, whereOperators)
	}

	shared.OrderBy(&query, params.SortBy, params.SortOrder)
	shared.Limit(&query, limit, offset)

	log.Debug().Str("query", query).Msg("Built SQL query")

	return query + ";", joinTable
}

func ListLogsQuery(
	params LogsQueryParams,
	filterBy shared.FilterByQueryParams,
) ([]DisplayLogSchema, error) {
	var logs []DisplayLogSchema
	query, joinTable := buildListLogsQuery(params, filterBy)

	rows, err := config.DBPool.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var item DisplayLogSchema

		if joinTable {
			var em endpoint_metrics.DisplayEndpointMetricsSchema
			if err := rows.Scan(
				&item.Timestamp,
				&item.Level,
				&item.Message,
				&item.Meta,
				&em.Endpoint,
				&em.UserID,
				&em.Method,
				&em.StatusCode,
				&em.ResponseTimeMs,
				&em.Meta,
			); err != nil {
				return nil, err
			}
			item.EndpointMetrics = &em
		} else {
			if err := rows.Scan(
				&item.Timestamp,
				&item.Level,
				&item.Message,
				&item.Meta,
			); err != nil {
				return nil, err
			}
		}

		logs = append(logs, item)
	}

	return logs, nil
}
