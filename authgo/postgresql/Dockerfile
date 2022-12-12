FROM postgres:15.1-alpine

COPY ./init.sql /docker-entrypoint-initdb.d/

ENV POSTGRES_PASSWORD "moffy0223"
ENV POSTGRES_USER "moffy"
ENV POSTGRES_DB "chatuser"

EXPOSE 5432