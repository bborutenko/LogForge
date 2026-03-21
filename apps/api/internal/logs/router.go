package logs

import (
	"github.com/bborutenko/LogForge/internal/shared"
	"github.com/gin-gonic/gin"
)

func ListLogs(c *gin.Context) {
	var filterByParams shared.FilterByQueryParams
	var logsQueryParams LogsQueryParams

	filterByParams.TimeFrom = c.Query("time_from")
	filterByParams.TimeTo = c.Query("time_to")
	filterByParams.UserIDs = c.QueryArray("user_ids")
	filterByParams.StatusCodes = shared.ParseStringArrayIntoIntArray(c.QueryArray("status_codes"))
	filterByParams.Endpoints = c.QueryArray("endpoints")
	filterByParams.Meta = c.QueryMap("meta")

	logsQueryParams.Page = shared.ParseStringIntoInt(c.Query("page"))
	logsQueryParams.PageSize = shared.ParseStringIntoInt(c.Query("page_size"))
	logsQueryParams.SortBy = c.Query("sort_by")
	logsQueryParams.SortOrder = c.Query("sort_order")

	if err := logsQueryParams.CheckForEmptyParams(); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	result, err := ListLogsQuery(logsQueryParams, filterByParams)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": result})
}
