package shared

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

type FilterByQueryParams struct {
	TimeFrom    string            `json:"time_from"`
	TimeTo      string            `json:"time_to"`
	UserIDs     []string          `json:"user_ids"`
	StatusCodes []int             `json:"status_codes"`
	Endpoints   []string          `json:"endpoints"`
	Meta        map[string]string `json:"meta"`
}

func (f *FilterByQueryParams) LoadFromFilterBy(c *gin.Context) error {
	f.TimeFrom = ""
	f.TimeTo = ""
	f.UserIDs = []string{}
	f.StatusCodes = []int{}
	f.Endpoints = []string{}
	f.Meta = map[string]string{}

	raw := c.Query("filter_by")
	if raw == "" {
		return nil
	}

	var parsed FilterByQueryParams
	if err := json.Unmarshal([]byte(raw), &parsed); err != nil {
		return fmt.Errorf("invalid filter_by json: %w", err)
	}

	f.TimeFrom = parsed.TimeFrom
	f.TimeTo = parsed.TimeTo

	if parsed.UserIDs != nil {
		f.UserIDs = parsed.UserIDs
	}
	if parsed.StatusCodes != nil {
		f.StatusCodes = parsed.StatusCodes
	}
	if parsed.Endpoints != nil {
		f.Endpoints = parsed.Endpoints
	}
	if parsed.Meta != nil {
		f.Meta = parsed.Meta
	}

	return nil
}
