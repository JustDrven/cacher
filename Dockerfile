FROM golang:1.24rc2-alpine3.21

WORKDIR /app
COPY . /app/

EXPOSE 8080

RUN chmod x+ /app/start.sh
RUN /app/start.sh

CMD [ "go", "run", "/app/main.go" ]