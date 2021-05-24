FROM mariadb

RUN apt update && \
  apt-get update && \
  apt install -y apt-utils \
  locales && \
  rm -rf /var/lib/apt/lists/* && \
  locale-gen ja_JP.UTF-8
ENV LANG ja_JP.UTF-8
ADD ./docker/mysql/my.cnf /etc/mysql/conf.d/my.cnf
