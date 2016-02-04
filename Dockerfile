FROM alpine:3.3

ADD ./atlas-meta /usr/local/bin/

ENTRYPOINT ["atlas-meta"]
