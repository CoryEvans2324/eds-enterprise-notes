FROM golang:1.16.6-alpine AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download


COPY config/ ./config/
COPY database/ ./database/
COPY middleware/ ./middleware/
COPY models/ ./models/
COPY routes/ ./routes/

COPY main.go ./

RUN go build -o /server main.go

# Build the style sheet
FROM node:16-alpine3.14 AS nodebuild

WORKDIR /data
COPY package*.json ./

RUN npm ci

COPY tailwind.css ./
COPY *.config.js ./
COPY web/ ./web/

RUN npm run build

# The image that runs the application
FROM alpine:latest

WORKDIR /app

COPY --from=build /server ./
COPY web/ ./web/
COPY --from=nodebuild /data/web/static/style.css ./web/static/style.css

EXPOSE 80

ENTRYPOINT [ "./server" ]