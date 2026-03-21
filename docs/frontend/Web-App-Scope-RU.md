# Описание общего вида

Весь сайт на английском языке.

На главной странице проекта слева расположен сайдбар со следующими пунктами:
- Logs
- Logs Levels
- Endpoints
- User Actions
- Anomalies

В центре (если не выбраны Logs Levels, Endpoints или User Actions) отображается пустой экран.

После выбора:
- Logs Levels, Endpoints или User Actions

Из уже существующего сайдбара выдвигается новый сайдбар с выбором конкретных названий. Центр по-прежнему пустой.

После выбора:
- Logs Levels Name, Endpoints Name или User Actions Name

В центре отображаются аналитические данные по выбранному названию, подробности см. ниже.

После выбора:
- Logs или Anomalies

В центре отображаются аналитические данные, подробности см. ниже.

# Глобальные фильтры

Все страницы поддерживают query-параметр `filter_by` в формате JSON:

`filter_by={"time_from": "2026-01-01 14:30:00", "time_to": "2026-03-17 14:30:00", "user_ids": ["uuid1"], "status_codes": [500]}`

### Поля в параметре `filter_by`

- Диапазон времени (time_from: datetime, time_to: datetime)
- Список user_ids: List[string]
- Диапазон status_codes: List[int]
- Конкретные endpoints: List[string]
- Метаданные meta: HashMap[string, string]

# Страницы с аналитикой

## Просмотр логов

### URL

`localhost:3000/logs?filter_by='{JSON}&page=1&page_size=25&sort_by=timestamp&sort_order=desc`


Страница отображает сырые логи с возможностью:
- Фильтрации по всем доступным полям
- Сортировки (timestamp DESC, level, endpoint)
- Пагинации (25/50/100 записей на страницу)
- Быстрого поиска по сообщению

### Содержание

График распределения по уровням | Таблица логов | Быстрые фильтры по level/status

## Аналитика уровней

### URL

Страница выбора уровня:  
`localhost:3000/levels`

Страница по конкретному уровню:  
`localhost:3000/levels/[level_name]?filter_by='{JSON}&period=7d&compare_with_previous=true`


`localhost:3000/levels` — выбор уровня логирования (INFO, WARN, ERROR, DEBUG).

`localhost:3000/levels/[level_name]` — детальный дашборд с:
- Линейным графиком событий по времени (час/день)
- Таблицей топ эндпоинтов для данного уровня
- Выбором периода (7 дней)
- Сравнением с предыдущим периодом

## Аналитика эндпоинтов

### URL

Страница выбора эндпоинта:  
`localhost:3000/endpoints`

Страница по конкретному эндпоинту:  
`localhost:3000/endpoints/[endpoint_name]?filter_by='{JSON}'`

`localhost:3000/endpoints` — выбор эндпоинта.

`localhost:3000/endpoints/[endpoint_name]` — глубокая аналитика по эндпоинту:
- График response_time (p50/p95/p99)
- Распределение status codes (pie chart)
- Топ пользователей по частоте вызовов

## Анализ воронки пользователей

### URL

Страница выбора воронки:  
`localhost:3000/user-actions?filter_by='{JSON}'&funnel='["login","dashboard","purchase"]'`

`localhost:3000/user-actions` — выбор последовательности действий (login→dashboard→purchase).

Страница отображает:
- Конверсию между шагами воронки
- Возврат пользователей (день 1, 3, 7)
- Среднее время между действиями
- Таблицу drop-off по шагам

## Анализ аномалий и всплесков

### URL

`localhost:3000/logs/anomalies?filter_by='{JSON}'`

### Логика

- Аномалии обнаружены → список логов с отклонениями (response_time > 3σ, всплески ERROR > 200%)
- Всё стабильно → последние 100 логов с меткой "System is stable"

### Содержимое дашборда

- Тепловая карта аномалий по endpoint/hour
- Топ-5 текущих аномалий
- График отклонений Z-score
