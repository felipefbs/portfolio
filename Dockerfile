FROM golang:alpine
WORKDIR /app
COPY . .
RUN go build -o app ./cmd/app/main.go
CMD [ "./app" ]