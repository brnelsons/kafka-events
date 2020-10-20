# Multistage build
FROM golang:1.15-alpine as build
# Update and install deps
RUN apk update && apk add ca-certificates && apk add git make
# Create source directory for project
ADD . /src/
WORKDIR /src
# Disable cross-compile
ENV CGO_ENABLED=0
# Run tests
RUN make test

# Build
RUN make build
RUN chmod -R 777 /src/kafka-events

# Build final image artifact
FROM alpine:3.10
RUN apk update && apk add ca-certificates
COPY --from=build /src/kafka-events /app/
COPY --from=build /src/migrations/* /migrations/
CMD ["/app/kafka-events"]