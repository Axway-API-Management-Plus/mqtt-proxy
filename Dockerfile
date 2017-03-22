FROM golang:alpine
RUN apk add --no-cache make git
WORKDIR /app/src/mqtt-proxy
COPY Makefile .deps ./
RUN make deps-install

COPY . ./
RUN make

EXPOSE 1883
ENV PORT 1883
ENV MQTT_BROKER_HOST 0.0.0.0
ENV MQTT_BROKER_PORT 1883
ENV MQTT_BROKER_USERNAME guest
ENV MQTT_BROKER_PASSWORD guest
ENV AUTH_URL ""

CMD [ "/app/src/mqtt-proxy/mqtt-proxy" ]
