
FROM ubuntu:latest

WORKDIR /

RUN apt-get update && apt-get install -y \
    tzdata \
    nmap \
    sqlite3

COPY ./license-backend/server/server .
COPY ./license-backend/server/.env .
COPY ./license-frontend/license/dist/spa ./static
COPY ./license-backend/server/repository/sqlite/migrations/001_create_tables.sql .
COPY ./license-backend/server/repository/sqlite/migrations/seed.sql .

RUN chmod +x server

CMD ["./server"]
