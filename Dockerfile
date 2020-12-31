FROM alpine:3.6

COPY coronator /usr/local/bin/

RUN apk --update upgrade
RUN apk --no-cache add curl tzdata

COPY ./migration /opt/core-onator/migration

WORKDIR /opt/core-onator/

EXPOSE 2019
ENTRYPOINT ["coronator", "run:api"]
