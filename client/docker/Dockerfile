FROM alpine AS builder

ADD . /ifs
RUN sh /ifs/deploy

########

FROM alpine

LABEL version="v0.1.0" \
    maintainer="wang@rehiy.com"

COPY --from=builder /usr/local/bin/tdp-* /usr/local/bin/

ENV NODE_NAME= \
    NODE_OWNER= \
    NODE_REGION= \
    NODE_ISP= \
    NODE_BANNER=

ENTRYPOINT ["tdp-server"]

EXPOSE 9115
