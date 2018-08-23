#
# MQTT/RabbitMQ
#
# based on:
# RabbitMQ Dockerfile
#
# https://github.com/dockerfile/rabbitmq
#

# Pull base image.
FROM ubuntu

# Define environment variables.
ENV RABBITMQ_LOG_BASE /data/log
ENV RABBITMQ_MNESIA_BASE /data/mnesia

# Define mount points.
VOLUME ["/data/log", "/data/mnesia"]

# Define working directory.
WORKDIR /data

# Install RabbitMQ.
RUN apt-get update
RUN DEBIAN_FRONTEND=noninteractive apt-get install -y wget gnupg

RUN wget -qO - http://www.rabbitmq.com/rabbitmq-signing-key-public.asc | apt-key add -
RUN echo "deb http://www.rabbitmq.com/debian/ testing main" > /etc/apt/sources.list.d/rabbitmq.list

RUN DEBIAN_FRONTEND=noninteractive apt-get install -y rabbitmq-server apg
RUN rm -rf /var/lib/apt/lists/*

# Add files.
ADD bin/rabbitmq-start /usr/local/bin/
RUN chmod +x /usr/local/bin/rabbitmq-start

ADD etc/rabbitmq/rabbitmq.config /etc/rabbitmq/
ADD etc/rabbitmq/rabbitmq-env.conf /etc/rabbitmq/

# initial config
RUN rabbitmq-plugins enable rabbitmq_management rabbitmq_mqtt

# Define default command.
CMD ["/usr/local/bin/rabbitmq-start"]

# Expose ports.
EXPOSE 5672
EXPOSE 15672
EXPOSE 1883
