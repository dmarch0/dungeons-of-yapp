FROM golang:1.20.5-alpine as builder

WORKDIR /build

COPY go.mod .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /main ./cmd/main.go

FROM scratch

COPY --from=builder main /bin/main

ENTRYPOINT ["/bin/main"]