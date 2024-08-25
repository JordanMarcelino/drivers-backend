FROM golang:1.19 AS build-stage

WORKDIR /app

COPY . .

RUN go mod download && CGO_ENABLED=0 GOOS=linux go build -v -o /main ./cmd/api/main.go

# Lean image
FROM gcr.io/distroless/base-debian11:latest-amd64 AS build-release-stage

WORKDIR /

COPY --from=build-stage /main /main
COPY --from=build-stage /app/.env /.env

EXPOSE 8000

USER nonroot:nonroot

ENTRYPOINT [ "/main" ]