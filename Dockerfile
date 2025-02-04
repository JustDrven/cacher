FROM golang:1.24rc2-alpine3.21

WORKDIR /app
COPY . /app/

EXPOSE 8080

CMD [ "bash", "app/start.sh" ]