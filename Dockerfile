FROM alpine/git as builder
RUN git clone https://gitlab.com/Nulide/findmydeviceserver.git /

FROM golang
COPY --from=builder /findmydeviceserver .
WORKDIR cmd
RUN go build -race -ldflags "-extldflags '-static'" -o fmdserver

# https://gitlab.com/Nulide/findmydeviceserver/-/issues/3
# HTTP
EXPOSE 1020/tcp
# HTTPS
EXPOSE 1008/tcp

RUN ls -alh
#VOLUME ??
CMD fmdserver