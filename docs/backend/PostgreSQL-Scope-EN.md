# Minimal Database Schema

## Tables

| Name | Fields | Purpose |
|------|--------|---------|
| `logs` | `id`, `timestamp`, `level`, `message`, `user_id`, `endpoint`, `status_code`, `response_time_ms`, `meta` (JSONB) | **Core**. All endpoints. Raw logs + analytics (levels, endpoints, anomalies, filters) |
| `user_actions` | `id`, `timestamp`, `user_id`, `action`, `endpoint`, `session_id` | **Funnels**. `/api/user-actions`, action sequences, retention, conversions |
