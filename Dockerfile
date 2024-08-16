FROM golang:bookworm AS builder

WORKDIR /app
COPY . /app/

RUN go mod download
RUN go mod verify
RUN CGO_ENABLED=0 go build -o /go/bin/app

FROM gcr.io/distroless/static-debian12

COPY --from=builder go/bin/app /
COPY --from=builder app/templates /templates

CMD ["/app"]
