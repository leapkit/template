FROM golang:1.24-alpine AS builder
RUN apk --update add build-base

WORKDIR /src/app

ADD go.mod .
RUN go mod download

ADD . .

# Downloading the tailwind binary, musl because this is an Alpine image, and
# running the build command
RUN go tool tailo download -v v4.0.6 --musl
RUN go tool tailo --i internal/system/assets/tailwind.css -o internal/system/assets/application.css

# Building the migrate command
RUN go build -tags osusergo,netgo -o bin/migrate ./cmd/migrate

# Building the app
RUN go build -tags osusergo,netgo -o bin/app ./cmd/app

FROM alpine
RUN apk add --no-cache tzdata ca-certificates

WORKDIR /bin/

# Copying binaries
COPY --from=builder /src/app/bin/app .
COPY --from=builder /src/app/bin/migrate .

# Specifying the shell to use
SHELL ["/bin/ash", "-c"]
CMD migrate && app
