package logs

import (
	"errors"
	"time"

	"github.com/rs/zerolog/log"
)

type LogsQueryParams struct {
	Page      int    `form:"page" binding:"required,min=1"`
	PageSize  int    `form:"page_size" binding:"required,oneof=25 50 100"`
	SortBy    string `form:"sort_by" binding:"omitempty,oneof=timestamp level endpoint"`
	SortOrder string `form:"sort_order" binding:"omitempty,oneof=asc desc"`
}

func (l *LogsQueryParams) CheckForEmptyParams() error {
	if l.Page == 0 {
		log.Warn().Msg("Page parameter is missing")
		return errors.New("page parameter is missing")
	} else if l.PageSize == 0 {
		log.Warn().Msg("Page size parameter is missing")
		return errors.New("page_size parameter is missing")
	} else if l.SortBy == "" {
		log.Warn().Msg("Sort by parameter is missing")
		return errors.New("sort_by parameter is missing")
	} else if l.SortOrder == "" {
		log.Warn().Msg("Sort order parameter is missing")
		return errors.New("sort_order parameter is missing")
	}
	return nil
}

type DisplayLogSchema struct {
	Timestamp         time.Time              `json:"timestamp"`
	Level             string                 `json:"level"`
	Message           string                 `json:"message"`
	Meta              map[string]interface{} `json:"meta,omitempty"`
	EndpointMetricsID *int64                 `json:"endpoint_metrics_id,omitempty"`
}
