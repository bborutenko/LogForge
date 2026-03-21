package logs

import (
	"context"
	"fmt"

	"github.com/bborutenko/LogForge/internal/config"
	"github.com/bborutenko/LogForge/internal/shared"
)

func buildListLogsQuery(params LogsQueryParams, filterBy shared.FilterByQueryParams) string {
	limit, offset := BuildLimitOffset(params.Page, params.PageSize)

	// todo: add filters to the query
	query := `SELECT timestamp, level, message, meta, endpoint_metrics_id 
		FROM log_forge.logs
		ORDER BY %s %s
		LIMIT %d OFFSET %d;`

	query = fmt.Sprintf(query, params.SortBy, params.SortOrder, limit, offset)
	return query
}

func ListLogsQuery(params LogsQueryParams, filterBy shared.FilterByQueryParams) ([]DisplayLogSchema, error) {
	var logs []DisplayLogSchema
	query := buildListLogsQuery(params, filterBy)
	rows, err := config.DBPool.Query(context.Background(), query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var log DisplayLogSchema
		if err := rows.Scan(&log.Timestamp, &log.Level, &log.Message, &log.Meta, &log.EndpointMetricsID); err != nil {
			return nil, err
		}
		logs = append(logs, log)
	}

	return logs, nil
}
