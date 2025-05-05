FROM golang:1.24-alpine AS build

WORKDIR /src

COPY go.mod .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o books ./main

FROM alpine:edge

WORKDIR /src

COPY --from=build /src/books /src/main/books

RUN chmod +x /src/main/books

RUN apk --no-cache add ca-certificates

EXPOSE 8080

ENTRYPOINT [ "/src/main/books" ]
