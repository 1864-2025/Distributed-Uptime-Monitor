CREATE TABLE IF NOT EXISTS checks (
      id SERIAL PRIMARY KEY,
      url TEXT NOT NULL,
      status_code INTEGER,
      response_time_ms INTEGER,
      created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);