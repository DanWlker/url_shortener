FROM golang:bookworm AS builder

ENV APP_HOME="/go/src/url_shortener"

WORKDIR "$APP_HOME"
COPY . .

RUN go mod download
RUN go mod verify
RUN go build -o url_shortener

# Second stage
FROM golang:bookworm

ENV APP_HOME="/go/src/url_shortener"
WORKDIR "$APP_HOME"

COPY --from=builder $APP_HOME $APP_HOME

EXPOSE 8090
CMD ["./url_shortener"]
