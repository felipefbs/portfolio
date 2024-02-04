FROM golang:alpine as golang-builder
WORKDIR /app
COPY . .
RUN wget -O tailwind https://github.com/tailwindlabs/tailwindcss/releases/download/v3.4.1/tailwindcss-linux-x64 && \
    chmod +x tailwind && ./tailwind -i ./templates/input.css -o ./static/styles/output.css --minify
RUN go install github.com/a-h/templ/cmd/templ@latest && \
    templ generate -path ./templates
RUN go build -o app ./cmd/app/main.go

FROM scratch
WORKDIR /app
COPY --from=golang-builder app/static/ ./static/
COPY --from=golang-builder app/app .
EXPOSE 8080
CMD [ "./app" ]