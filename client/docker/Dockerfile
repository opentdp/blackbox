FROM alpine AS builder

ADD . /ifs
RUN sh /ifs/deploy

########

FROM alpine

LABEL version="v0.2.0" \
    maintainer="wang@rehiy.com"

COPY --from=builder /usr/local/bin/tdp-* /usr/local/bin/

ENV NODE_NAME= \
    NODE_OWNER= \
    NODE_REGION= \
    NODE_ISP= \
    NODE_BANNER=

HEALTHCHECK --interval=30s --timeout=5s \
    CMD wget -spider http://localhost:9115/-/healthy || exit 1

ENTRYPOINT ["tdp-server"]

EXPOSE 9000
