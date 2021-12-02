FROM alpine
RUN apk update
#RUN apk add go
#ENV GO111MODULE=on
#ENV GOPATH=/root/go
#ENV GOBIN=/root/go/bin
#ENV PATH=$PATH:$GOBIN
EXPOSE 8000
COPY aoc_linux /aoc
CMD [ "/aoc", "server", "-p", "8000" ]

