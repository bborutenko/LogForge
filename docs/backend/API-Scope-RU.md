# Эндпоинты

Ниже описаны эндпоинты API для получения данных на фронтенде. Формат тела запроса — JSON. 

### Поля в параметре `filter_by`

- Диапазон времени (time_from: datetime, time_to: datetime)
- Список user_ids: List[string]
- Диапазон status_codes: List[int]
- Конкретные endpoints: List[string]
- Метаданные meta: HashMap[string, string]

---

## GET /api/logs

**Метод:** `GET`  
**Назначение:** Получение списка сырых логов с фильтрацией, сортировкой и пагинацией для страницы `localhost:3000/logs`.

**Query-параметры:**
- `filter_by` — строка JSON с полями:
  - `time_from: string (datetime)`
  - `time_to: string (datetime)`
  - `user_ids: string[]`
  - `status_codes: number[]`
  - `endpoints: string[]`
  - `meta: Record<string, string>`
- `page: number` — номер страницы.
- `page_size: number` — размер страницы (25/50/100).
- `sort_by: string` — поле сортировки (`timestamp`, `level`, `endpoint`).
- `sort_order: string` — порядок сортировки (`asc` | `desc`).

---

## GET /api/logs/levels-distribution

**Метод:** `GET` 
**Назначение:** Получение агрегированных данных для графика распределения логов по уровням логирования (для страницы логов).

**Query-параметры:**
- `filter_by` — строка JSON (как выше).

---

## GET /api/levels

**Метод:** `GET`
**Назначение:** Получение списка доступных уровней логирования `localhost:3000/levels`.

---

## GET /api/levels/:level_name

**Метод:** `GET` 
**Назначение:** Получение детальной аналитики по конкретному уровню логирования для страницы `localhost:3000/levels/[level_name]`.

**Параметры пути:**
- `level_name: string` — название уровня (`INFO`, `WARN`, `ERROR`, `DEBUG` и т.д.).

**Query-параметры:**
- `filter_by` — строка JSON.
- `period: string` — выбранный период (например, `7d`).
- `compare_with_previous: boolean` — нужна ли выборка для сравнения с предыдущим периодом.

---

## GET /api/levels/:level_name/events-timeseries

**Метод:** `GET`
**Назначение:** Получение временного ряда событий для выбранного уровня (график по часам/дням).

**Параметры пути:**
- `level_name: string`

**Query-параметры:**
- `filter_by` — строка JSON.
- `bucket: string` — размер агрегирования (`hour`, `day`).

---

## GET /api/levels/:level_name/top-endpoints

**Метод:** `GET`
**Назначение:** Получение таблицы топ-эндпоинтов для выбранного уровня логирования.

**Параметры пути:**
- `level_name: string`

**Query-параметры:**
- `filter_by` — строка JSON.
- `limit: number` — количество эндпоинтов (например, 10).

---

## GET /api/endpoints

**Метод:** `GET`
**Назначение:** Получение списка всех эндпоинтов `localhost:3000/endpoints`.

---

## GET /api/endpoints/:endpoint_name

**Метод:** `GET`
**Назначение:** Получение общей агрегированной информации по конкретному эндпоинту (для хедера дашборда).

**Параметры пути:**
- `endpoint_name: string`

**Query-параметры:**
- `filter_by` — строка JSON.

---

## GET /api/endpoints/:endpoint_name/latency-timeseries

**Метод:** `GET`
**Назначение:** Получение временного ряда метрик времени ответа (p50/p95/p99) для графика.

**Параметры пути:**
- `endpoint_name: string`

**Query-параметры:**
- `filter_by` — строка JSON.
- `bucket: string` — размер агрегирования (`minute`, `hour`, `day`).

---

## GET /api/endpoints/:endpoint_name/status-distribution

**Метод:** `GET`
**Назначение:** Получение распределения HTTP status codes (для pie chart).

**Параметры пути:**
- `endpoint_name: string`

**Query-параметры:**
- `filter_by` — строка JSON.

---

## GET /api/endpoints/:endpoint_name/top-users

**Метод:** `GET`
**Назначение:** Получение топ-пользователей по частоте вызовов выбранного эндпоинта.

**Параметры пути:**
- `endpoint_name: string`

**Query-параметры:**
- `filter_by` — строка JSON.
- `limit: number` — количество пользователей (например, 10).

---

## GET /api/user-actions

**Метод:** `GET`
**Назначение:** Выбор воронки с пользовтаелями и последующее отображение с фильтрами `localhost:3000/user-actions`.

**Query-параметры:**
- `filter_by` — строка JSON (опционально).
- `funnel` — список с последовательностью воронки. Порядок важен

**Ответ содержит:**
- Конверсии между шагами
- Возврат пользователей по дням (1, 3, 7)
- Среднее время между действиями
- Данные для таблицы drop-off

---

## GET /api/logs/anomalies

**Метод:** `GET`
**Назначение:** Получение списка аномальных логов и текущих всплесков для страницы `localhost:3000/logs/anomalies`.

**Query-параметры:**
- `filter_by` — строка JSON.

---

## GET /api/logs/anomalies/heatmap

**Метод:** `GET`
**Назначение:** Получение данных для тепловой карты аномалий по `endpoint/hour`.

**Query-параметры:**
- `filter_by` — строка JSON.

---

## GET /api/logs/anomalies/top

**Метод:** `GET`
**Назначение:** Получение топ‑5 текущих аномалий для отображения на дашборде.

**Query-параметры:**
- `filter_by` — строка JSON.
- `limit: number` — количество аномалий (по умолчанию 5).

---

## GET /api/logs/anomalies/zscore-timeseries

**Метод:** `GET`
**Назначение:** Получение временного ряда Z-score отклонений для графика.

**Query-параметры:**
- `filter_by` — строка JSON.
- `metric: string` — метрика, по которой считается Z-score (например, `errors_count`, `response_time`).

---

## POST /api/logs

**Метод:** `POST`
**Назначение:** Создание одной записи лога.

**Тело запроса (JSON):**
- `level` (string, обязательный) — уровень логирования (например, ERROR, INFO).
- `message` (string, обязательный) — текст сообщения.
- `meta` (object, необязательный) — дополнительные метаданные (например, service, trace_id).
- `endpoint_metrics_id` (string, необязательный) — UUID связанной метрики эндпоинта.

**Обязательные:** level, message
**Необязательные:** meta, endpoint_metrics_id

Пример запроса:
```json
{
  "level": "ERROR",
  "message": "Payment service timeout",
  "meta": {
    "service": "billing",
    "trace_id": "abc123"
  },
  "endpoint_metrics_id": "550e8400-e29b-41d4-a716-446655440000"
}
```

---

## POST /api/logs/batch

**Метод:** `POST`
**Назначение:** Сохранение списка записей логов за один раз.

**Тело запроса (JSON):**
- `logs` (object, обязательный) — список всех параметров логов состоящий из
  - `level` (string, обязательный) — уровень логирования (например, ERROR, INFO).
  - `message` (string, обязательный) — текст сообщения.
  - `meta` (object, необязательный) — дополнительные метаданные (например, service, trace_id).
  - `endpoint_metrics_id` (string, необязательный) — UUID связанной метрики эндпоинта.

**Обязательные:** logs[].level, logs[].message
**Необязательные:** logs[].meta, logs[].endpoint_metrics_id

Пример запроса:
```json
{
  "logs": [
    {
      "level": "ERROR",
      "message": "Payment timeout",
      "meta": {"service": "billing"}
    },
    {
      "level": "INFO",
      "message": "Payment processed",
      "endpoint_metrics_id": "uuid-123"
    }
  ]
}
```

---

## POST /api/endpoint-metrics

**Метод:** `POST`
**Назначение:** Создание одной метрики эндпоинта.

**Обязательные:** endpoint, status_code, response_time_ms
**Необязательные:** user_id, meta

Пример запроса:
```json
{
  "endpoint": "/api/payments/create",
  "status_code": 504,
  "response_time_ms": 12500,
  "user_id": "550e8400-e29b-41d4-a716-446655440000",
  "meta": {
    "trace_id": "abc123",
    "env": "prod"
  }
}
```

---

## POST /api/endpoint-metrics/batch

**Метод:** `POST`
**Назначение:** Создание нескольких записей метрик эндпоинтов.

**Обязательные:** metrics[].endpoint, metrics[].status_code, metrics[].response_time_ms
**Необязательные:** metrics[].user_id, metrics[].meta

Пример запроса:
```json
{
  "metrics": [
    {
      "endpoint": "/api/payments/create",
      "status_code": 504,
      "response_time_ms": 12500
    },
    {
      "endpoint": "/api/users/list",
      "status_code": 200,
      "response_time_ms": 150,
      "user_id": "uuid-456"
    }
  ]
}
```

---

## POST /api/user-actions

**Метод:** `POST`
**Назначение:** Создание одной записи действия пользователя.

**Обязательные:** user_id, action
**Необязательные:** endpoint_metrics_id

Пример запроса:
```json
{
  "user_id": "550e8400-e29b-41d4-a716-446655440000",
  "action": "login",
  "endpoint_metrics_id": "550e8400-e29b-41d4-a716-446655440001"
}

```

---

## POST /api/user-actions/batch

**Метод:** `POST`
**Назначение:** Создание нескольких записей действий пользователей.

**Обязательные:** actions[].user_id, actions[].action
**Необязательные:** actions[].endpoint_metrics_id

Пример запроса:
```json
{
  "actions": [
    {
      "user_id": "uuid-123",
      "action": "login"
    },
    {
      "user_id": "uuid-123",
      "action": "dashboard_view",
      "endpoint_metrics_id": "uuid-metrics-1"
    }
  ]
}
```

---

## POST /api/endpoint

**Метод:** `POST`
**Назначение:** Создание полной записи эндпоинта со связанными логами и действиями пользователей.

**Обязательные:** endpoint_metrics
**Необязательные:** logs (массив), user_actions (массив) — хотя бы одно из полей не должно быть пустым

Пример запроса:
```json
{
  "metrics": {
    "endpoint": "/api/payments/create",
    "status_code": 200,
    "response_time_ms": 250,
    "user_id": "uuid-123"
  },
  "logs": [
    {
      "level": "INFO",
      "message": "Payment processed successfully"
    }
  ],
  "user_actions": [
    {
      "user_id": "uuid-123",
      "action": "purchase_completed"
    }
  ]
}

```