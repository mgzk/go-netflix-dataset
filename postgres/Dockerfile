FROM postgres:latest

EXPOSE 5432
COPY ./sql/init.sql /docker-entrypoint-initdb.d/
