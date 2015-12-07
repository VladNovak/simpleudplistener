FROM ubuntu

ENV SERVER /server
RUN mkdir $SERVER
WORKDIR $SERVER

ADD . $SERVER

CMD ./simpletcplistener
