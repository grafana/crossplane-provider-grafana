FROM gcr.io/distroless/static:nonroot@sha256:aafe993665bf6200db38a02c11a9351e782a84b6a45e7f33b6695638c88e3931

ARG TARGETOS
ARG TARGETARCH

ADD "bin/${TARGETOS}_${TARGETARCH}/provider" /usr/local/bin/provider

USER 65532
EXPOSE 8080

ENTRYPOINT ["provider"]
