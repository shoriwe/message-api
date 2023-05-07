FROM golang:1.19-alpine AS build-stage

RUN apk add upx
WORKDIR /message-api-src
COPY . .
RUN go build -o /message-api .
RUN upx /message-api

FROM alpine:latest AS release-stage

COPY --from=build-stage /message-api /message-api
# Setup folders
WORKDIR /message-api-config
WORKDIR /message-api-config/database
WORKDIR /message-api-config/configuration
WORKDIR /
# -- Environment variables
ENV DATABASE_FILE               "/message-api-config/database/database.db"
ENV SECRET                      "TESTING"
ENV FIREBASE_PROJECT_ID         "PLEASE_CONFIGURE_ME"
ENV FIREBASE_CONFIGURATION_FILE "/message-api-config/configuration/firebase-adminsdk.json"
# -- Environment variables
ENTRYPOINT [ "sh", "-c", "/message-api :5000" ]