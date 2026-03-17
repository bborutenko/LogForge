# Endpoints

The endpoints below are for getting data to the frontend. Request body format — JSON.

### Fields in the `filter_by` parameter

- Time range (time_from: datetime, time_to: datetime)
- List of user_ids: List[string]
- Range of status_codes: List[int]
- Specific endpoints: List[string]
- Metadata meta: HashMap[string, string]

---

## GET /api/logs

**Method:** `GET`  
**Path:** `/api/logs`  
**Purpose:** Get raw logs list with filtering, sorting and pagination for page `localhost:3000/logs`.

**Query parameters:**
- `filter_by` — JSON string with fields:
  - `time_from: string (datetime)`
  - `time_to: string (datetime)`
  - `user_ids: string[]`
  - `status_codes: number[]`
  - `endpoints: string[]`
  - `meta: Record<string, string>`
- `page: number` — page number.
- `page_size: number` — page size (25/50/100).
- `sort_by: string` — sort field (`timestamp`, `level`, `endpoint`).
- `sort_order: string` — sort order (`asc` | `desc`).

---

## GET /api/logs/levels-distribution

**Method:** `GET`  
**Path:** `/api/logs/levels-distribution`  
**Purpose:** Get aggregated data for log levels distribution chart (for logs page).

**Query parameters:**
- `filter_by` — JSON string (as above).

---

## GET /api/levels

**Method:** `GET`  
**Path:** `/api/levels`  
**Purpose:** Get list of available log levels for `localhost:3000/levels`.

---

## GET /api/levels/:level_name

**Method:** `GET`  
**Path:** `/api/levels/:level_name`  
**Purpose:** Get detailed analytics for specific log level for page `localhost:3000/levels/[level_name]`.

**Path parameters:**
- `level_name: string` — level name (`INFO`, `WARN`, `ERROR`, `DEBUG` etc.).

**Query parameters:**
- `filter_by` — JSON string.
- `period: string` — selected period (e.g. `7d`).
- `compare_with_previous: boolean` — need previous period data for comparison.

---

## GET /api/levels/:level_name/events-timeseries

**Method:** `GET`  
**Path:** `/api/levels/:level_name/events-timeseries`  
**Purpose:** Get events time series for selected level (hour/day chart).

**Path parameters:**
- `level_name: string`

**Query parameters:**
- `filter_by` — JSON string.
- `bucket: string` — aggregation size (`hour`, `day`).

---

## GET /api/levels/:level_name/top-endpoints

**Method:** `GET`  
**Path:** `/api/levels/:level_name/top-endpoints`  
**Purpose:** Get top endpoints table for selected log level.

**Path parameters:**
- `level_name: string`

**Query parameters:**
- `filter_by` — JSON string.
- `limit: number` — number of endpoints (e.g. 10).

---

## GET /api/endpoints

**Method:** `GET`  
**Path:** `/api/endpoints`  
**Purpose:** Get all endpoints list for `localhost:3000/endpoints`.

---

## GET /api/endpoints/:endpoint_name

**Method:** `GET`  
**Path:** `/api/endpoints/:endpoint_name`  
**Purpose:** Get general aggregated info for specific endpoint (for dashboard header).

**Path parameters:**
- `endpoint_name: string`

**Query parameters:**
