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
**Путь:** `/api/logs`  
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
**Путь:** `/api/logs/levels-distribution`  
**Назначение:** Получение агрегированных данных для графика распределения логов по уровням логирования (для страницы логов).

**Query-параметры:**
- `filter_by` — строка JSON (как выше).

---

## GET /api/levels

**Метод:** `GET`  
**Путь:** `/api/levels`  
**Назначение:** Получение списка доступных уровней логирования `localhost:3000/levels`.

---

## GET /api/levels/:level_name

**Метод:** `GET`  
**Путь:** `/api/levels/:level_name`  
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
**Путь:** `/api/levels/:level_name/events-timeseries`  
**Назначение:** Получение временного ряда событий для выбранного уровня (график по часам/дням).

**Параметры пути:**
- `level_name: string`

**Query-параметры:**
- `filter_by` — строка JSON.
- `bucket: string` — размер агрегирования (`hour`, `day`).

---

## GET /api/levels/:level_name/top-endpoints

**Метод:** `GET`  
**Путь:** `/api/levels/:level_name/top-endpoints`  
**Назначение:** Получение таблицы топ-эндпоинтов для выбранного уровня логирования.

**Параметры пути:**
- `level_name: string`

**Query-параметры:**
- `filter_by` — строка JSON.
- `limit: number` — количество эндпоинтов (например, 10).

---

## GET /api/endpoints

**Метод:** `GET`  
**Путь:** `/api/endpoints`  
**Назначение:** Получение списка всех эндпоинтов `localhost:3000/endpoints`.

---

## GET /api/endpoints/:endpoint_name

**Метод:** `GET`  
**Путь:** `/api/endpoints/:endpoint_name`  
**Назначение:** Получение общей агрегированной информации по конкретному эндпоинту (для хедера дашборда).

**Параметры пути:**
- `endpoint_name: string`

**Query-параметры:**
- `filter_by` — строка JSON.

---

## GET /api/endpoints/:endpoint_name/latency-timeseries

**Метод:** `GET`  
**Путь:** `/api/endpoints/:endpoint_name/latency-timeseries`  
**Назначение:** Получение временного ряда метрик времени ответа (p50/p95/p99) для графика.

**Параметры пути:**
- `endpoint_name: string`

**Query-параметры:**
- `filter_by` — строка JSON.
- `bucket: string` — размер агрегирования (`minute`, `hour`, `day`).

---

## GET /api/endpoints/:endpoint_name/status-distribution

**Метод:** `GET`  
**Путь:** `/api/endpoints/:endpoint_name/status-distribution`  
**Назначение:** Получение распределения HTTP status codes (для pie chart).

**Параметры пути:**
- `endpoint_name: string`

**Query-параметры:**
- `filter_by` — строка JSON.

---

## GET /api/endpoints/:endpoint_name/top-users

**Метод:** `GET`  
**Путь:** `/api/endpoints/:endpoint_name/top-users`  
**Назначение:** Получение топ-пользователей по частоте вызовов выбранного эндпоинта.

**Параметры пути:**
- `endpoint_name: string`

**Query-параметры:**
- `filter_by` — строка JSON.
- `limit: number` — количество пользователей (например, 10).

---

## GET /api/user-actions

**Метод:** `GET`  
**Путь:** `/api/user-actions`  
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
**Путь:** `/api/logs/anomalies`  
**Назначение:** Получение списка аномальных логов и текущих всплесков для страницы `localhost:3000/logs/anomalies`.

**Query-параметры:**
- `filter_by` — строка JSON.

---

## GET /api/logs/anomalies/heatmap

**Метод:** `GET`  
**Путь:** `/api/logs/anomalies/heatmap`  
**Назначение:** Получение данных для тепловой карты аномалий по `endpoint/hour`.

**Query-параметры:**
- `filter_by` — строка JSON.

---

## GET /api/logs/anomalies/top

**Метод:** `GET`  
**Путь:** `/api/logs/anomalies/top`  
**Назначение:** Получение топ‑5 текущих аномалий для отображения на дашборде.

**Query-параметры:**
- `filter_by` — строка JSON.
- `limit: number` — количество аномалий (по умолчанию 5).

---

## GET /api/logs/anomalies/zscore-timeseries

**Метод:** `GET`  
**Путь:** `/api/logs/anomalies/zscore-timeseries`  
**Назначение:** Получение временного ряда Z-score отклонений для графика.

**Query-параметры:**
- `filter_by` — строка JSON.
- `metric: string` — метрика, по которой считается Z-score (например, `errors_count`, `response_time`).
