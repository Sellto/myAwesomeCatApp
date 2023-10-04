ARG BASE_GO_IMAGE=golang:alpine3.15
ARG BASE_IMAGE=alpine:3.15

FROM ${BASE_GO_IMAGE} as build-generator
WORKDIR /app
COPY go.sum .
COPY go.mod .
RUN go mod download
COPY . /app
RUN go build -o build/main cmd/main.go

FROM ${BASE_IMAGE}
COPY --from=build-generator /app/build /app

CMD ["/app/main"]