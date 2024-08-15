FROM golang:bookworm

RUN go install github.com/air-verse/air@latest
RUN go install github.com/DanWlker/url_shortener@docker_test

ENV DATABASE_URL=""

ENV APP_HOME="/go/src/url_shortener"
RUN mkdir -p "$APP_HOME"
WORKDIR "$APP_HOME"

EXPOSE 8090

CMD ["air"]
