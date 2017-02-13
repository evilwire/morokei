FROM alpine:3.5

# the version of the docker file
ARG VERSION
ENV VERSION $VERSION

LABEL maintainer="evilwire"
LABEL version=$VERSION

# Copy the binary
COPY build/$VERSION/morokei /opt/morokei
RUN chmod +x /opt/morokei

# Expose the appropriate ports
# 9000 is the actual port to respond to requests
# 9090 is the healthcheck port
EXPOSE 9000 9090

# Set the entry point
ENTRYPOINT ["/opt/morokei"]
