FROM golang:1.21.3 AS build

WORKDIR /app

COPY go.mod go.sum ./
COPY . .

RUN go mod download

RUN CGO_ENABLED=0 go build -o apique

FROM alpine:latest

WORKDIR /app

COPY --from=build /app/apique .

EXPOSE 8061

CMD [ "./apique" ]