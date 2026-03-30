import asyncio
import asyncpg
import json
import os
import random
import time
from datetime import datetime, timedelta

# Настройки из переменных окружения
DATABASE_URL = os.environ.get("DATABASE_URL")
if not DATABASE_URL:
    # Если переменная не задана, можно вставить строку подключения для теста:
    # DATABASE_URL = "postgresql://user:password@localhost:5432/logforge"
    raise ValueError("Необходимо установить переменную окружения DATABASE_URL")

TOTAL_RECORDS = 500_000_000
BATCH_SIZE = 100_000 

# Справочники
ENDPOINTS = ['/api/v1/auth', '/api/v1/users', '/api/v1/billing', '/api/v1/products', '/api/v1/checkout']
METHODS = ['GET', 'POST', 'PUT', 'DELETE', 'PATCH']
STATUS_CODES = [200, 201, 400, 401, 403, 404, 500]
LOG_LEVELS = ['INFO', 'DEBUG', 'WARN', 'ERROR']
ACTIONS = ['login', 'logout', 'view_item', 'add_to_cart', 'purchase', 'click_banner']
OS_LIST = ['Linux', 'Windows', 'macOS', 'Android', 'iOS']

async def generate_and_insert():
    print(f"🚀 Запуск генерации {TOTAL_RECORDS:,} записей в схему 'log_forge'...")
    start_time_script = time.time()
    
    # Подключение к базе
    conn = await asyncpg.connect(DATABASE_URL)
    
    try:
        # Устанавливаем search_path, чтобы asyncpg видел таблицы без префикса схемы
        await conn.execute('SET search_path TO log_forge, public')
        
        batch_em = []
        batch_logs = []
        batch_ua = []
        base_time = datetime.now() - timedelta(days=30)

        for i in range(1, TOTAL_RECORDS + 1):
            # Таймстемп (используем объекты datetime для asyncpg)
            current_time = base_time + timedelta(milliseconds=100 * i)
            current_user = f"usr_{random.randint(1, 100000)}"

            # --- 1. Данные для endpoint_metrics ---
            em_id = i
            endpoint = random.choice(ENDPOINTS)
            method = random.choice(METHODS)
            status = random.choice(STATUS_CODES)
            # В вашей схеме double precision, random.uniform или int подойдут
            resp_time = float(random.randint(10, 2000)) 
            meta_em = json.dumps({
                "os": random.choice(OS_LIST), 
                "ip": f"192.168.1.{random.randint(1, 254)}"
            })
            
            batch_em.append((
                em_id, endpoint, current_user, method, 
                status, resp_time, meta_em, current_time
            ))

            # --- 2. Данные для logs и user_actions ---
            # Связываем через ID метрики
            random_metric_link = i 

            # Logs
            level = random.choice(LOG_LEVELS)
            msg = f"Request {method} {endpoint} handled"
            meta_l = json.dumps({"trace_id": f"{random.getrandbits(32):x}"})
            batch_logs.append((current_time, level, msg, meta_l, random_metric_link))

            # User Actions
            act = random.choice(ACTIONS)
            sess = f"sess_{random.getrandbits(32)}"
            batch_ua.append((current_time, current_user, act, sess, random_metric_link))

            # Вставка батча при достижении BATCH_SIZE
            if i % BATCH_SIZE == 0:
                async with conn.transaction():
                    # Используем эффективный COPY для каждой таблицы
                    await conn.copy_records_to_table(
                        'endpoint_metrics', 
                        records=batch_em, 
                        columns=["id", "endpoint", "user_id", "method", "status_code", "response_time_ms", "meta", "created_at"]
                    )
                    await conn.copy_records_to_table(
                        'logs', 
                        records=batch_logs,
                        columns=['timestamp', 'level', 'message', 'meta', 'endpoint_metrics_id']
                    )
                    await conn.copy_records_to_table(
                        'user_actions', 
                        records=batch_ua,
                        columns=['timestamp', 'user_id', 'action', 'session_id', 'endpoint_metrics_id']
                    )
                
                elapsed = time.time() - start_time_script
                print(f"✅ Записано {i:,} / {TOTAL_RECORDS:,} строк. Прошло: {elapsed:.2f}s")
                
                # Очистка списков для следующего батча
                batch_em.clear()
                batch_logs.clear()
                batch_ua.clear()

        # Вставка остатков (если TOTAL_RECORDS не кратен BATCH_SIZE)
        if batch_em:
            async with conn.transaction():
                await conn.copy_records_to_table('endpoint_metrics', records=batch_em, columns=["id", "endpoint", "user_id", "method", "status_code", "response_time_ms", "meta", "created_at"])
                await conn.copy_records_to_table('logs', records=batch_logs, columns=['timestamp', 'level', 'message', 'meta', 'endpoint_metrics_id'])
                await conn.copy_records_to_table('user_actions', records=batch_ua, columns=['timestamp', 'user_id', 'action', 'session_id', 'endpoint_metrics_id'])

        # После ручной вставки ID нужно синхронизировать Postgres Sequence
        print("🔄 Синхронизация sequence...")
        await conn.execute("SELECT setval('log_forge.endpoint_metrics_id_seq', (SELECT max(id) FROM log_forge.endpoint_metrics))")

    except Exception as e:
        print(f"❌ Ошибка при вставке: {e}")
    finally:
        await conn.close()

    total_time = time.time() - start_time_script
    print(f"\n✨ Завершено успешно!")
    print(f"Всего вставлено: {TOTAL_RECORDS:,} наборов данных.")
    print(f"Общее время: {total_time:.2f} секунд.")

if __name__ == "__main__":
    asyncio.run(generate_and_insert())