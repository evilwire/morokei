FROM alpine:3.5

LABEL maintainer="evilwire"

# Copy the binary
COPY build/morokei /opt/morokei

# Expose the appropriate ports
EXPOSE 9000 9090

# Set the entry point
ENTRYPOINT "morokei"
