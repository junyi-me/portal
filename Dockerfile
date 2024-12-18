FROM golang:1.23 AS build
WORKDIR /app
COPY . .
RUN go build -o /bin/portal ./src

FROM ubuntu:24.10
WORKDIR /app
COPY --from=build /bin/portal /app/portal
COPY tmpl /app/tmpl
ENTRYPOINT ["/app/portal"]

