FROM golang:1.27.3-alpine3.23 AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/scope ./cmd/scope/scope.go

FROM alpine:3.23

WORKDIR /build

COPY --from=build /app/score ./score

CMD ["./scope"]