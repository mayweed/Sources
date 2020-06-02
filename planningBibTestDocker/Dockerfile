#docker run -it -p8888:80 testpb3:latest /bin/bash
FROM debian:testing

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
php7.4-mysql \
composer \
git \
apt-utils \
apt-transport-https \
ca-certificates \
vim \
wget \
unzip && \
apt-get clean

#this is amcmz sh*t
COPY ./amcmz.crt /usr/local/share/ca-certificates/amcmz.crt
RUN update-ca-certificates

RUN git clone https://github.com/PlanningBiblio/PlanningBiblio/ -b 19.10.xx /var/www/planningBiblio && \
chown -R www-data:www-data /var/www/planningBiblio/include && \
chown -R www-data:www-data /var/www/planningBiblio/data 

RUN cd /var/www/planningBiblio && \
composer update 

WORKDIR /home/guillaume

COPY 000-default.conf /etc/apache2/sites-available/

EXPOSE 80
