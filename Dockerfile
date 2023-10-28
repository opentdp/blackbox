FROM rehiy/webox:slim

LABEL version="1.0.0" \
      maintainer="wang@rehiy.com"

ADD initfs /ifs
RUN sh /ifs/deploy

ENV NODE_NAME= \
    NODE_OWNER= \
    NODE_REGION= \
    NODE_ISP= \
    NODE_BANNER=

ENTRYPOINT ["/sbin/init"]

EXPOSE 9115
