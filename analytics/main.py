import pandas as pd
from sqlalchemy import create_engine
import matplotlib.pyplot as plt
import os
from dotenv import load_dotenv

load_dotenv(dotenv_path="../.env")

user = os.getenv("DB_USER")
password = os.getenv("DB_PASSWORD")
host = os.getenv("DB_HOST", "localhost")
port = os.getenv("DB_PORT", "5432")
dbname = os.getenv("DB_NAME")

DB_URL = f"postgresql://{user}:{password}@{host}:{port}/{dbname}"

def check_db():

    try:
        engine = create_engine(DB_URL)

        df = pd.read_sql("SELECT * FROM checks", engine)

        if df.empty:
            print("База поключена, но данных нет")
            return

        print("Данные получены")
        print(f"Всего проверок базе: {len(df)}")

        print("\n Статистика по сайтам:")
        stats = df.groupby('url').agg({
            'status_code': lambda x: (x==200).mean() * 100,
            'response_time_ms': 'mean'
        }).rename(columns={'status_code': 'Uptime %', 'response_time_ms': 'Avg Latency (ms)'})

        print(stats)

        df['created_at'] = pd.to_datetime(df['created_at'])

        plt.figure(figsize = (12,6))
        for url in df[df['status_code'] == 200]['url'].unique():
            subset = df[df['url'] == url]
            plt.plot(subset['created_at'], subset['response_time_ms'], marker = 'o', label = url)
        plt.title('Response Time trend (Success Checks)')
        plt.xlabel('Time')
        plt.ylabel('Response Time (ms)')
        plt.legend()
        plt.grid(True, alpha = 0.3)
        plt.xticks(rotation = 0)
        plt.tight_layout()

        print("\nРисую график... Закрой окно графика, чтобы завершить программу.")
        plt.show()

    except Exception as e:
        print(f"Ошибка подключения: {e}")


if __name__ == "__main__":
    check_db()


