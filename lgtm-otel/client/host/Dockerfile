FROM alpine:latest

#~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
# Download node_exporter
#~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
ARG NODE_EXPORTER_VERSION=1.9.0
RUN apk update && apk add --no-cache tar wget
RUN wget -qO- "https://github.com/prometheus/node_exporter/releases/download/v${NODE_EXPORTER_VERSION}/node_exporter-${NODE_EXPORTER_VERSION}.linux-amd64.tar.gz" | tar xz -C /tmp/
RUN mv /tmp/node_exporter-${NODE_EXPORTER_VERSION}.linux-amd64/node_exporter /usr/local/bin/ && \
    rm -rf /tmp/node_exporter-${NODE_EXPORTER_VERSION}.linux-amd64

#~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
# User for node_exporter
#~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
RUN adduser -D -s /bin/false node_exporter

#~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
# Service configuration
#~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
RUN mkdir -p /etc/systemd/system && \
    echo "[Unit]\n\
Description=Node Exporter\n\
After=network.target\n\
\n\
[Service]\n\
User=node_exporter\n\
Group=node_exporter\n\
Type=simple\n\
ExecStart=/usr/local/bin/node_exporter\n\
\n\
[Install]\n\
WantedBy=multi-user.target" > /etc/systemd/system/node_exporter.service

#~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
# Download shadow
#~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
RUN apk add shadow

#~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
# Expose the metric port
#~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
EXPOSE 9100

CMD ["/usr/local/bin/node_exporter"]