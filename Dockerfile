FROM registry.docker.ir/library/golang:1.19 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /file-service

FROM build-stage AS run-test-stage
RUN go test -v ./...

FROM registry.docker.ir/library/alpine:3 AS build-release-stage

WORKDIR /

COPY --from=build-stage /file-service /file-service

EXPOSE 8080

USER 1001:1001

ENTRYPOINT ["/file-service"]