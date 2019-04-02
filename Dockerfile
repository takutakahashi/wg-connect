FROM golang
ADD . /src
RUN cd /src/cmd \
 && go build

CMD ["/src/cmd/cmd", "client"]
