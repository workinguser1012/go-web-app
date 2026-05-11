FROM golang:1.21 AS base
WORKDIR /app
COPY go.mod .
RUN go mod download
COPY . .
RUN go build -o main .

FROM gcr.io/distroless/base
WORKDIR /app
COPY --from=base /app/main .
COPY --from=base /app/static ./static
COPY --from=base /app/templates ./templates
EXPOSE 8080
CMD ["./main"]

