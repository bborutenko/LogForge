package config

var endpointMetricsRowNames = []string{
	"id", "endpoint", "user_id", "method", "status_code", "response_time_ms", "meta", "created_at",
}
var endpointMetricsRowValues = []string{
	"SERIAL",
	"TEXT NOT NULL",
	"TEXT",
	"TEXT",
	"INT NOT NULL",
	"FLOAT NOT NULL",
	"JSONB",
	"TIMESTAMPTZ NOT NULL DEFAULT NOW()",
}

var logsRowNames = []string{
	"timestamp", "level", "message", "meta", "endpoint_metrics_id",
}

var logsRowValues = []string{
	"TIMESTAMPTZ PRIMARY KEY DEFAULT NOW()",
	"TEXT NOT NULL",
	"TEXT NOT NULL",
	"JSONB",
	"BIGINT",
}

var userActionsRowNames = []string{
	"timestamp", "user_id", "action", "session_id", "endpoint_metrics_id",
}

var userActionsRowValues = []string{
	"TIMESTAMPTZ PRIMARY KEY DEFAULT NOW()",
	"TEXT",
	"TEXT NOT NULL",
	"TEXT",
	"BIGINT",
}
