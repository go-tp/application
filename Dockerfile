FROM golang:1.19.0
# ENV GOPATH /go
ENV GOPROXY=https://goproxy.cn,direct
ENV GO111MODULE="on"
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && echo 'Asia/Shanghai' >/etc/timezone
# Install fresh
RUN go install github.com/pilu/fresh@latest
EXPOSE 7777