FROM alpine
MAINTAINER Vitalii Hurin <vitalii.hurin@masterofcode.com>
ENV APP=test-service PORT=80
EXPOSE 80
ADD $APP /
ADD runit-start /sbin/
RUN mkdir /etc/container_environment && \
    chmod a+x /sbin/runit-start && \
    mkdir /etc/runit_init.d && \
    mkdir /log && \
    mkdir -p /etc/service/$APP/log && \
    echo "#!/bin/sh -e" > /etc/service/$APP/log/run && \
    echo "exec 2>&1" >> /etc/service/$APP/log/run && \
    echo "exec chpst svlogd -tt /log" >> /etc/service/$APP/log/run && \
    echo "#!/bin/sh -e" > /etc/service/$APP/run && \
    echo "exec 2>&1" >> /etc/service/$APP/run && \
    echo "exec chpst /$APP" >> /etc/service/$APP/run && \
    echo "http://dl-cdn.alpinelinux.org/alpine/edge/testing" > /etc/apk/repositories && \
    chmod a+x /$APP && \
    chmod -R a+x /log && \
    chmod -R a+x /etc/service/$APP && \
    apk add --no-cache runit
CMD ["/sbin/runit-start"]
