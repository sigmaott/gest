FROM confluentinc/cp-ksql-cli:latest

COPY ssai.sql /app/ssai.sql

CMD ["ksql", "$KSQL_SERVER_URL", "/app/migrate.sql"]