# Specifies a parent image
FROM golang:1.19.2-bullseye AS build_stage

# Creates an app directory to hold your app’s source code
WORKDIR /app

# Copies everything from your root directory into /app
COPY ./sampleApp/ .

# Installs Go dependencies
RUN go mod download
RUN go version
# Builds your app with optional configuration
RUN go build -o godocker
RUN ls -la /app
FROM alpine:latest

WORKDIR /app
RUN pwd
COPY --from=build_stage /app/godocker godocker
RUN chmod +x godocker
RUN ls -la /app

# Tells Docker which network port your container listens on
EXPOSE 80

# Specifies the executable command that runs when the container starts
ENTRYPOINT ["./godocker"]
