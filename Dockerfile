FROM golang:1.26-bookworm
WORKDIR /app
COPY go.mod main.go main_test.go ./
RUN CGO_ENABLED=0 go build -o server . \
    && CGO_ENABLED=0 go test -c -o app.test . \
    && useradd --create-home --uid 10001 app
COPY scripts/check.sh scripts/check.sh
USER app
ENV PORT=3000
EXPOSE 3000
CMD ["./server"]