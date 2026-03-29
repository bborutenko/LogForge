package endpoint_metrics

type DisplayEndpointMetricsSchema struct {
	UserID         string                 `json:"user_id"`
	Endpoint       string                 `json:"endpoint"`
	Method         string                 `json:"method"`
	StatusCode     int                    `json:"status_code"`
	ResponseTimeMs float64                `json:"response_time_ms"`
	Meta           map[string]interface{} `json:"meta,omitempty"`
}
