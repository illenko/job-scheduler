CREATE SOURCE CONNECTOR `postgres-source` WITH(
    "connector.class"='io.confluent.connect.jdbc.JdbcSourceConnector',
    "connection.url"='jdbc:postgresql://psql:5432/postgres?user=postgres&password=postgres',
    "mode"='timestamp',
    "timestamp.column.name"='create_time',
    "topic.prefix"='',
    "table.whitelist"='public.job_schedule',
    "key"='job_id');