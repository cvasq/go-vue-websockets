# Stage 1 - build
FROM node:8.15-alpine AS builder
WORKDIR /app
COPY ui/package*.json ./
RUN  npm install
COPY ui/. .
RUN VUE_APP_BASE_PATH="/websocket-echo" npm run build

# Stage 2 - build go binary
FROM golang:1.16.2-alpine@sha256:12d5f94cd4d2840e538e82e26a5dfddf711b30cc98a9f6e01bcf65d7aaf7ccd8 AS webserver_build
WORKDIR /server
ADD ./server/ /server
RUN CGO_ENABLED=0 go build -ldflags '-extldflags "-static"' -o webserver

ENV USER=appuser
ENV UID=10001
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"

# Stage 2 - final build
FROM scratch as final
WORKDIR /app/html
COPY --from=builder /app/dist .
COPY --from=webserver_build /server/webserver ..
COPY --from=webserver_build /etc/passwd /etc/passwd
COPY --from=webserver_build /etc/group /etc/group
USER appuser:appuser
ENTRYPOINT ["/app/webserver"]
