FROM golang:alpine as golang-builder
WORKDIR /app
COPY . .
RUN go build -o app ./cmd/app/main.go
RUN wget -O tailwind https://github.com/tailwindlabs/tailwindcss/releases/download/v3.4.1/tailwindcss-linux-x64 && \
    chmod +x tailwind && ./tailwind -i ./templates/input.css -o ./static/styles/output.css
RUN go install github.com/a-h/templ/cmd/templ@latest && \
    templ generate -path ./templates

FROM alpine:3.19
WORKDIR /app
COPY --from=golang-builder app/static/ ./static/
COPY --from=golang-builder app/templates/*_templ.go ./templates/
COPY --from=golang-builder app/app .
CMD [ "./app" ]

