FROM golang:1.19 as base

ARG SERVICE_PORT

FROM base AS dev

WORKDIR /app

COPY social-service/go.mod .

RUN go mod download

COPY . .

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

CMD ["air"]

EXPOSE $SERVICE_PORT

FROM base as prod

WORKDIR /app

COPY . .

RUN go build -o /bin/server ./main.go

CMD ["/bin/server"]

EXPOSE $SERVICE_PORT