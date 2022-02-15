FROM alpine:latest

WORKDIR  /app
COPY gink .
ENV TZ=Asia/Shanghai
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories \
    && apk add --no-cache tzdata \
    && ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone \
    && chmod +x gink

EXPOSE 80
CMD ["-env","local"]
ENTRYPOINT ["./gink"]
