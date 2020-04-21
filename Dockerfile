FROM golang:1.14-alpine AS server-build
WORKDIR /github.com/wasabi315/todo-server
COPY go.mod go.sum ./
RUN go mod download
COPY ./ ./
RUN go build -o todo-server

FROM alpine:3.9
WORKDIR /app
COPY --from=server-build /github.com/wasabi315/todo-server/todo-server ./

EXPOSE 3000

ENTRYPOINT ./todo-server
