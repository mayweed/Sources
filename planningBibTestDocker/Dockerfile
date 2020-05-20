#docker run -it -p8888:80 testpb3:latest /bin/bash
FROM debian:testing

#set hostname
RUN echo "SRVPLANNING" >> /etc/hostname

#general install
RUN apt-get update && \
apt-get upgrade -y && \
apt-get install -y \
apache2 \
default-mysql-server \
default-mysql-client \
php7.4 \
php7.4-curl \
php7.4-xml \
composer \
php-symfony-framework-bundle \
git \
apt-utils \
apt-transport-https \
ca-certificates \
vim \
wget \
unzip && \
apt-get clean

#this is amcmz sh*t
#ENV GIT_SSL_NO_VERIFY=true
COPY ./amcmz.crt /usr/local/share/ca-certificates/amcmz.crt
RUN update-ca-certificates

RUN mkdir -p /var/www/planningBiblio && \
git clone https://github.com/PlanningBiblio/PlanningBiblio/ -b 19.10.xx /var/www/planningBiblio && \
chown -R www-data /var/www/planningBiblio && \
chgrp -R www-data /var/www/planningBiblio && \
chmod -R 755 /var/www/planningBiblio

#RUN cd /var/www/planningBiblio && \
#cp .env .env.local && \
#composer update && \
#composer install ext-xml && \
#composer install

WORKDIR /home/guillaume

COPY 000-default.conf /etc/apache2/sites-available/

EXPOSE 80
