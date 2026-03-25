import pandas as pd
from sqlalchemy import create_engine

DB_URL = "postgresql://user:password@localhost:5432/uptime_monitor"

def check_db():

    try:
        engine = create_engine(DB_URL)

        df = pd.read_sql("SELECT * FROM checks", engine)

        if df.empty:
            print("База поключена, но данных нет")
            return

        print("Данные получены")
        print(f"Всего проверок базе{len(df)}")

        print("\n Статистика по сайтам:")
        stats = df.groupby('url').agg({
            'status_code': lambda x: (x==200).mean * 100,
            'response_time_ms': 'mean'
        }).rename(columns={'status_code': 'Uptime %', 'response_time_ms': 'Avg Latency (ms)'})

        print(stats)

    except Exception as e:
        print(f"Ошибка подключения: {e}")


if __name__ == "__main__":
    check_db()


