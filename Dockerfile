FROM golang:1.18

RUN go version

WORKDIR /auth
COPY ./ ./

RUN go mod download
RUN go build -o auth-service ./cmd/app/main.go

EXPOSE 5000

CMD [ "./auth-service" ]