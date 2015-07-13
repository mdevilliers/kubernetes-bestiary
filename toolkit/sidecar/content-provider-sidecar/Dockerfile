FROM ubuntu:14.04

RUN apt-get update > /dev/null \
	&& apt-get -y install git

ENTRYPOINT sleep 99999999 | git clone https://github.com/mdevilliers/kubernetes-bestiary.git /etc/static