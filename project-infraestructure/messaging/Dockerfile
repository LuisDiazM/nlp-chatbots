FROM nats:2.6.2-alpine
COPY nats-server.conf /etc/nats/nats-server.conf
EXPOSE 4222
CMD ["nats-server", "--config", "/etc/nats/nats-server.conf"]
