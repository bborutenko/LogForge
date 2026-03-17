# Минимальная схема БД

## Таблицы

| Название | Поля | Назначение |
|----------|------|------------|
| `endpoint_metrics` | `id`, `user_id` (nullable), `endpoint`, `status_code`, `response_time_ms`, `meta` (JSONB) | **Метрики эндпоинтов**. Статистика работы API (p95/p99, ошибки, производительность). Записывается в конце запроса |
| `logs` | `id`, `timestamp`, `level`, `message`, `meta` (JSONB), `endpoint_metrics_id` (FK → `endpoint_metrics.id`) | **Общие логи**. Сырые логи системы, уровни логирования, сообщения. Связь с метриками эндпоинта |
| `user_actions` | `id`, `timestamp`, `user_id`, `action`, `session_id`, `endpoint_metrics_id` (FK → `endpoint_metrics.id`) | **Воронки**. Последовательность действий пользователя в сессии. Связь с метриками эндпоинта |

## Связи

- `logs.endpoint_metrics_id` → `endpoint_metrics.id`
- `user_actions.endpoint_metrics_id` → `endpoint_metrics.id`

## Логика

Один вызов эндпоинта → одна запись в `endpoint_metrics` → несколько связанных записей в `logs` и `user_actions`.
