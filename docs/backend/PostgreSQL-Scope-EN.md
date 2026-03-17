# Minimal Database Schema

## Tables

| Name | Fields | Purpose |
|------|--------|---------|
| `endpoint_metrics` | `id`, `user_id` (nullable), `endpoint`, `status_code`, `response_time_ms`, `meta` (JSONB) | **Endpoint metrics**. API performance statistics (p95/p99, errors, performance). Written at request end |
| `logs` | `id`, `timestamp`, `level`, `message`, `meta` (JSONB), `endpoint_metrics_id` (FK → `endpoint_metrics.id`) | **General logs**. Raw system logs, log levels, messages. Linked to endpoint metrics |
| `user_actions` | `id`, `timestamp`, `user_id`, `action`, `session_id`, `endpoint_metrics_id` (FK → `endpoint_metrics.id`) | **Funnels**. User action sequence within session. Linked to endpoint metrics |

## Relationships

- `logs.endpoint_metrics_id` → `endpoint_metrics.id`
- `user_actions.endpoint_metrics_id` → `endpoint_metrics.id`

## Logic

One endpoint call → one `endpoint_metrics` record → multiple related records in `logs` and `user_actions`.
