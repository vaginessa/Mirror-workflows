FROM alpine/git:latest AS gitimport
RUN git clone -b FindMyDeviceServer --single-branch https://github.com/schklom/Mirror-workflows.git /fmd


FROM golang:alpine AS builder
# We know from test that $GOPATH=/go
WORKDIR /go/src/fmd
RUN curl -s https://raw.githubusercontent.com/objectbox/objectbox-go/main/install.sh | bash

COPY --from=gitimport /fmd/ /go/src/fmd/
WORKDIR /go/src/fmd/cmd
#RUN go build fmdserver.go
RUN go build -ldflags '-w -s' -o fmdserver

RUN mkdir -p /fmd/web
VOLUME /fmd
RUN mv /go/src/fmd/web/ /fmd/web/
RUN mv /go/src/fmd/cmd/fmdserver /fmd/fmdserver

#FROM golang:alpine
#VOLUME /fmd
#COPY --from=builder /go/src/fmd/web /fmd/web/
#COPY --from=builder /go/src/fmd/cmd/fmdserver /fmd/

WORKDIR /fmd

# https://gitlab.com/Nulide/findmydeviceserver/-/issues/3
# HTTP
EXPOSE 1020/tcp
# HTTPS
EXPOSE 1008/tcp

ENTRYPOINT [ "/fmd/fmdserver" ]
