FROM ubuntu:22.10

COPY ./main ./
COPY ./Netflix-Dataset-2021.xlsx ./dataset/

ENTRYPOINT ["./main","postgres://netflix_user:TGxCOWiDbpcwgSM@netflix-postgres.local:5432/netflix?sslmode=disable"]

