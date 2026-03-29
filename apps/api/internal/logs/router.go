package logs

import (
	"github.com/bborutenko/LogForge/internal/shared"
	"github.com/gin-gonic/gin"
)

func ListLogs(c *gin.Context) {
	var filterByParams shared.FilterByQueryParams
	var logsQueryParams LogsQueryParams

	if err := filterByParams.LoadFromFilterBy(c); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

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
