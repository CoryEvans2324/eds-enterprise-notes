FROM golang:1.16.6-alpine AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download


COPY config/ ./config/
COPY database/ ./database/
COPY middleware/ ./middleware/
COPY models/ ./models/
COPY templates/ ./templates/

RUN go build -o /server .


# The image that runs the application
FROM alpine:latest

COPY --from=build /server /server
COPY web/ /web/

EXPOSE 8000

ENTRYPOINT [ "/server" ]