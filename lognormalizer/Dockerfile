FROM alpine:3.14

ENV LANG en_US.UTF-8
ENV LANGUAGE en_US.UTF-8
ENV LC_ALL=en_US.UTF-8

RUN echo '' > /etc/apk/repositories && \
    echo "https://mirror.tuna.tsinghua.edu.cn/alpine/v3.14/main"         >> /etc/apk/repositories && \
    echo "https://mirror.tuna.tsinghua.edu.cn/alpine/v3.14/community"    >> /etc/apk/repositories && \
    echo "166.111.206.63    mirror.tuna.tsinghua.edu.cn"                >> /etc/hosts && \
    echo "185.31.17.249     github.com"                                 >> /etc/hosts && \
    echo "Asia/Shanghai" > /etc/timezone

RUN apk --no-cache --no-progress update && \
    apk --no-cache --no-progress upgrade

RUN apk add liblognorm
