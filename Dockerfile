FROM ubuntu

ENV SERVER /server
RUN mkdir $SERVER
WORKDIR $SERVER

RUN apt-get update -qq && apt-get install -y golang

ADD . $SERVER

RUN go build

CMD ./server udp
