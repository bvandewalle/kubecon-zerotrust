FROM ubuntu:xenial
MAINTAINER Aporeto Inc <info@aporeto.com>

RUN apt-get update && apt-get -y install libelf-dev libnetfilter-queue1 iptables ca-certificates curl wget

ADD run.sh /run.sh

ENTRYPOINT ["/run.sh"]
