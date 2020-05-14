#docker run -it -p8888:80 testpb3:latest /bin/bash
FROM debian:10
RUN apt-get update && apt-get install -y apache2 default-mysql-server default-mysql-client php7.3 git apt-transport-https ca-certificates
#this is amcmz sh*t
#ENV GIT_SSL_NO_VERIFY=true
COPY ./amcmz.crt /usr/local/share/ca-certificates/amcmz.crt
RUN update-ca-certificates
RUN mkdir -p /var/www/planningBiblio && \
git clone https://github.com/PlanningBiblio/PlanningBiblio/ /var/www/planningBiblio 
WORKDIR /home/guillaume
EXPOSE 80
