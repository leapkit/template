FROM golang:1.24-alpine AS builder
RUN apk --update add build-base

WORKDIR /src/app
# Adding go.mod and go.sum and downloading dependencies first
# This is done to leverage Docker layer caching
ADD go.* .
RUN go mod download

# Downloading the tailwind binary, musl because this is an Alpine image.
# This is done first to leverage Docker layer caching
RUN go tool tailo download -v v4.0.6 --musl

ADD . .

# Generating the Tailwind CSS styles with the tailwind binary previously downloaded.
RUN go tool tailo --i internal/system/assets/tailwind.css -o internal/system/assets/application.css

# Building the migrate command with necessary tags
RUN go build -tags osusergo,netgo -o bin/migrate ./cmd/migrate

# Building the app with necessary tags
RUN go build -tags osusergo,netgo -o bin/app ./cmd/app

FROM alpine
RUN apk add --no-cache tzdata ca-certificates

WORKDIR /bin/

# Copying binaries to /bin from the builder stage
COPY --from=builder /src/app/bin/app .
COPY --from=builder /src/app/bin/migrate .

# Specifying the shell to use
SHELL ["/bin/ash", "-c"]
CMD migrate && app
