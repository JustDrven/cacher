FROM golang:1.24rc2-alpine3.21

WORKDIR /app

COPY src/ /app/

EXPOSE 8080

ENV API_KEY=321

CMD [ "go", "run", "/app/main.go" ]