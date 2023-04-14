FROM registry.docker.ir/library/golang:1.19 AS build-stage

WORKDIR /app

COPY . .
RUN go get -v -d ./...
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /file-service -v github.com/xoltawn/simple-file-storage-file-service

FROM build-stage AS run-test-stage
RUN go test -v ./...

FROM registry.docker.ir/library/alpine:3 AS build-release-stage

WORKDIR /

COPY --from=build-stage /file-service /file-service

EXPOSE 50051

USER root:root
VOLUME [ "./images" ]

ENTRYPOINT ["/file-service"]