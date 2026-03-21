import csv
import random
import json
import uuid
import time
from datetime import datetime, timedelta

# Настройки
TOTAL_RECORDS = 10_000_000
BATCH_SIZE = 100_000 

# Справочники
ENDPOINTS = ['/api/v1/auth', '/api/v1/users', '/api/v1/billing', '/api/v1/products', '/api/v1/checkout']
METHODS = ['GET', 'POST', 'PUT', 'DELETE', 'PATCH']
STATUS_CODES = [200, 201, 400, 401, 403, 404, 500]
LOG_LEVELS = ['INFO', 'DEBUG', 'WARN', 'ERROR']
ACTIONS = ['login', 'logout', 'view_item', 'add_to_cart', 'purchase', 'click_banner']
OS_LIST = ['Linux', 'Windows', 'macOS', 'Android', 'iOS']

def generate_mock_data():
    print(f"Запуск генерации: {TOTAL_RECORDS:,} строк для каждой таблицы.")
    print("Порядок полей синхронизирован с Go slice 'endpointMetricsRowNames'.")
    start_time_script = time.time()

    with open('endpoint_metrics.csv', 'w', newline='', encoding='utf-8') as f_em, \
         open('logs.csv', 'w', newline='', encoding='utf-8') as f_logs, \
         open('user_actions.csv', 'w', newline='', encoding='utf-8') as f_ua:

        writer_em = csv.writer(f_em)
        writer_logs = csv.writer(f_logs)
        writer_ua = csv.writer(f_ua)

        # 1. Заголовки (строго по твоему списку)
        writer_em.writerow(["id", "endpoint", "user_id", "method", "status_code", "response_time_ms", "meta", "created_at"])
        writer_logs.writerow(['timestamp', 'level', 'message', 'meta', 'endpoint_metrics_id'])
        writer_ua.writerow(['timestamp', 'user_id', 'action', 'session_id', 'endpoint_metrics_id'])

        batch_em, batch_logs, batch_ua = [], [], []
        base_time = datetime.now() - timedelta(days=30)

        for i in range(1, TOTAL_RECORDS + 1):
            # Уникальный таймстемп (шаг 100мс) для PK
            current_time = base_time + timedelta(milliseconds=100 * i)
            ts_str = current_time.strftime('%Y-%m-%d %H:%M:%S.%f%z') + '+00'
            
            # Общий user_id для этого "события"
            current_user = f"usr_{random.randint(1, 100000)}"

            # --- Таблица: endpoint_metrics ---
            em_id = i
            endpoint = random.choice(ENDPOINTS)
            method = random.choice(METHODS)
            status = random.choice(STATUS_CODES)
            resp_time = random.randint(10, 2000)
            meta_em = json.dumps({"os": random.choice(OS_LIST), "ip": f"192.168.1.{random.randint(1, 254)}"})
            
            # Порядок: id, endpoint, user_id, method, status_code, response_time_ms, meta, created_at
            batch_em.append([em_id, endpoint, current_user, method, status, resp_time, meta_em, ts_str])

            # --- Таблицы Logs и User Actions (Связь 1:N через случайный ID) ---
            # Выбираем случайную метрику из уже существующих/будущих ID
            random_metric_link = random.randint(1, TOTAL_RECORDS)

            # Logs
            level = random.choice(LOG_LEVELS)
            msg = f"Request {method} {endpoint} handled"
            meta_l = json.dumps({"trace_id": str(uuid.uuid4())[:8]})
            batch_logs.append([ts_str, level, msg, meta_l, random_metric_link])

            # User Actions
            act = random.choice(ACTIONS)
            sess = f"sess_{random.getrandbits(32)}"
            batch_ua.append([ts_str, current_user, act, sess, random_metric_link])

            # Запись батча
            if i % BATCH_SIZE == 0:
                writer_em.writerows(batch_em)
                writer_logs.writerows(batch_logs)
                writer_ua.writerows(batch_ua)
                batch_em.clear(); batch_logs.clear(); batch_ua.clear()
                print(f"Обработано {i:,} / {TOTAL_RECORDS:,}...")

        # Записываем остатки
        if batch_em:
            writer_em.writerows(batch_em)
            writer_logs.writerows(batch_logs)
            writer_ua.writerows(batch_ua)

    print(f"\nГотово! Файлы созданы за {time.time() - start_time_script:.2f} сек.")

if __name__ == "__main__":
    generate_mock_data()