# Минимальная схема БД

## Таблицы

| Название | Поля | Назначение |
|----------|------|------------|
| `logs` | `id`, `timestamp`, `level`, `message`, `user_id`, `endpoint`, `status_code`, `response_time_ms`, `meta` (JSONB) | **Основная**. Все эндпоинты. Сырые логи + аналитика (уровни, эндпоинты, аномалии, фильтры) |
| `user_actions` | `id`, `timestamp`, `user_id`, `action`, `endpoint`, `session_id` | **Воронки**. `/api/user-actions`, последовательность действий, retention, конверсии |
