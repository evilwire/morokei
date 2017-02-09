FROM alpine:3.5

# the version of the docker file
ARG VERSION

LABEL maintainer="evilwire"
LABEL version=$VERSION

# Copy the binary
COPY build/$VERSION/morokei /opt/morokei

# Expose the appropriate ports
EXPOSE 9000 9090

# Set the entry point
ENTRYPOINT "morokei"
