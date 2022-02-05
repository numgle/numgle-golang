FROM golang:1.17.6-alpine3.15 AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN CGO_ENABLED=0 go build -tags=nomsgpack -o /bin/server ./cmd/server

FROM gcr.io/distroless/static
ENV GIN_MODE=release
COPY --from=build /bin/server /
CMD ["/server"]