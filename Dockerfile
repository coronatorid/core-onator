FROM alpine:3.6

COPY coronator /usr/local/bin/

RUN apk --update upgrade
RUN apk --no-cache add curl tzdata

EXPOSE 2019
ENTRYPOINT ["coronator", "run:api"]
