FROM docker.elastic.co/elasticsearch/elasticsearch:5.6.5

USER root

# pre-stop-hook.sh and dependencies
RUN yum -y install jq curl

ENV GOSU_VERSION 1.10
RUN set -ex; \
  \
  yum -y install epel-release; \
  yum -y install wget dpkg; \
  \
  dpkgArch="$(dpkg --print-architecture | awk -F- '{ print $NF }')"; \
  wget -O /usr/bin/gosu "https://github.com/tianon/gosu/releases/download/$GOSU_VERSION/gosu-$dpkgArch"; \
  wget -O /tmp/gosu.asc "https://github.com/tianon/gosu/releases/download/$GOSU_VERSION/gosu-$dpkgArch.asc"; \
  \
  # verify the signature
  export GNUPGHOME="$(mktemp -d)"; \
  gpg --keyserver ha.pool.sks-keyservers.net --recv-keys B42F6819007F00F88E364FD4036A9C25BF357DD4; \
  gpg --batch --verify /tmp/gosu.asc /usr/bin/gosu; \
  rm -r "$GNUPGHOME" /tmp/gosu.asc; \
  \
  chmod +x /usr/bin/gosu; \
  # verify that the binary works
  gosu nobody true; \
  \
  yum -y remove wget dpkg; \
  yum clean all

COPY entrypoint.sh /entrypoint.sh
COPY pre-stop-hook.sh /pre-stop-hook.sh
COPY elasticsearch.yml /usr/share/elasticsearch/config/elasticsearch.yml
COPY log4j2.properties /usr/share/elasticsearch/config/log4j2.properties

RUN bin/elasticsearch-plugin install repository-s3

ENTRYPOINT ["/entrypoint.sh"]
