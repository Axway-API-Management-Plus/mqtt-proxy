FROM golang:alpine AS build
RUN apk add --no-cache make git
WORKDIR /app/src/mqtt-proxy
COPY Makefile .deps ./
RUN make deps-install

COPY . .git ./
RUN make

CMD [ "/app/src/mqtt-proxy/mqtt-proxy" ]

FROM alpine
RUN apk add --no-cache ca-certificates
COPY --from=build /app/src/mqtt-proxy/mqtt-proxy /usr/bin/mqtt-proxy
COPY ./certs ./certs

EXPOSE 1883
ENV PORT 1883
ENV MQTT_BROKER_HOST 0.0.0.0
ENV MQTT_BROKER_PORT 1883
ENV MQTT_BROKER_USERNAME guest
ENV MQTT_BROKER_PASSWORD guest
ENV AUTH_URL ""

CMD ["mqtt-proxy"]
