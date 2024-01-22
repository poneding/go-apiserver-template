FROM alpine
ARG TARGETOS TARGETARCH

WORKDIR /app
COPY _output/bin/go-apiserver-template_${TARGETOS}_${TARGETARCH} /app/go-apiserver-template

ENTRYPOINT ["/app/go-apiserver-template"]