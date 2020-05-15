#docker run -it -p8888:80 testpb3:latest /bin/bash
FROM debian:10

#set hostname
RUN echo "SRVPLANNING" >> /etc/hostname

#general install
RUN apt-get update && \
apt-get upgrade -y && \
apt-get install -y \
apache2 \
default-mysql-server \
default-mysql-client \
php7.3 \
git \
apt-utils \
apt-transport-https \
ca-certificates \
vim \
php-cli \
php-zip \
wget \
unzip &&\
apt-get clean

#this is amcmz sh*t
#ENV GIT_SSL_NO_VERIFY=true
COPY ./amcmz.crt /usr/local/share/ca-certificates/amcmz.crt
RUN update-ca-certificates

#composer to install dependencies
RUN php -r "copy('https://getcomposer.org/installer', 'composer-setup.php');"
RUN php composer-setup.php --install-dir=/usr/local/bin --filename=composer

RUN mkdir -p /var/www/planningBiblio && \
git clone https://github.com/PlanningBiblio/PlanningBiblio/ /var/www/planningBiblio && \
chown -R www-data /var/www/planningBiblio && \
chgrp -R www-data /var/www/planningBiblio

#https://getcomposer.org/doc/01-basic-usage.md //composer install not good here
#RUN cd /var/www/planningBiblio && composer update
WORKDIR /home/guillaume

#to boot apache at start
#https://stackoverflow.com/questions/49764989/cannot-start-apache-automatically-with-docker
#apache2 conf a2ensite?
#later on??
COPY 000-default.conf /etc/apache2/sites-available/

#CMD ["/usr/sbin/apachectl", "-D", "FOREGROUND"]
EXPOSE 80
