# General Layout Description

The entire site is in English.

On the main project page, there is a sidebar on the left with the following options:
- Logs
- Logs Levels
- Endpoints
- User Actions
- Anomalies

In the center (if Logs Levels, Endpoints, or User Actions are not selected) there is an empty screen.

After selecting:
- Logs Levels, Endpoints, or User Actions

A new sidebar slides out from the existing one and provides a selection of specific names. The center is still empty.

After selecting:
- Logs Levels Name, Endpoints Name, or User Actions Name

The center displays analytics data for the selected name, see details below.

After selecting:
- Logs or Anomalies

The center displays analytics data, see details below.

# Global Filters

All pages support the `filter_by` query parameter in JSON format:

`filter_by={"time_from": "2026-01-01 14:30:00", "time_to": "2026-03-17 14:30:00", "user_ids": ["uuid1"], "status_codes": [500]}`

Fields in the `filter_by` parameter:
- Time range (time_from: datetime, time_to: datetime)
- List of user_ids: List[string]
- Range of status_codes: List[int]
- Specific endpoints: List[string]
- Metadata meta: HashMap[string, string]

# Analytics Pages

## Log View

### **URL**

`localhost:3000/logs?filter_by='{JSON}'`

The page displays raw logs with the ability to:
- Filter by all available fields
- Sort (timestamp DESC, level, endpoint)
- Paginate (25/50/100 records per page)
- Perform quick search by message
- Export to CSV

### **Content**

Level distribution chart | Logs table | Quick filters by level/status

## Levels Analytics

### **URL**

Level selection page:  
`localhost:3000/levels`

Specific level page:  
`localhost:3000/levels/[level_name]?filter_by='{JSON}'`

`localhost:3000/levels` — log level selection (INFO, WARN, ERROR, DEBUG).

`localhost:3000/levels/[level_name]` — detailed dashboard with:
- Line chart of events over time (hour/day)
- Table of top endpoints for this level
- Period selection (7 days)
- Comparison with the previous period

## Endpoints Analytics

### **URL**

Endpoint selection page:  
`localhost:3000/endpoints`

Specific endpoint page:  
`localhost:3000/endpoints/[endpoint_name]?filter_by='{JSON}'`

`localhost:3000/endpoints` — endpoint selection.

`localhost:3000/endpoints/[endpoint_name]` — deep analytics for the endpoint:
- Response_time chart (p50/p95/p99)
- Status codes distribution (pie chart)
- Top users by call frequency

## User Funnel Analysis

### **URL**

Funnel selection page:  
`localhost:3000/user-actions?filter_by='{JSON}'&pipeline='["login","dashboard","purchase"]'`

`localhost:3000/user-actions` — selection of the sequence of actions (login→dashboard→purchase).

The page displays:
- Conversion between funnel steps
- User return (day 1, 3, 7)
- Average time between actions
- Drop-off table by steps

## Anomalies and Spikes Analysis

### **URL**

`localhost:3000/logs/anomalies?filter_by='{JSON}'`

### **Logic**

- Anomalies detected → List of logs with deviations (response_time > 3σ, ERROR spikes > 200%)
- Everything is stable → Last 100 logs with the label "System is stable"

### **Dashboard Content**

- Anomalies heatmap by endpoint/hour
- Top 5 current anomalies
- Z-score deviation chart
- Auto-refresh every 30 seconds
