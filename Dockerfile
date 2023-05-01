# FROM golang:1.20 AS builder
# RUN mkdir /app
# ADD . /app
# WORKDIR /app
# RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o app cmd/server/main.go

# # RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -installsuffix cgo -o /go/bin/sdm .

# FROM --platform=linux/amd64 alpine:latest AS production
# COPY --from=builder /app .
# RUN chmod +x ./app
# CMD ["./app"]

FROM golang:1.20-alpine AS builder
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app cmd/server/main.go

FROM alpine:latest AS production
COPY --from=builder /app .
RUN chmod +x ./app
CMD ["./app"]