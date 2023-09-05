FROM ubuntu
ARG DEBIAN_FRONTEND=noninteractive
RUN apt update; apt upgrade -y
ENV PORT=8880
RUN apt install curl wget golang zip unzip nginx -y
COPY scriptplusconf /scriptplusconf
COPY qBconf.tar.gz /qBconf.tar.gz
RUN tar xvf /qBconf.tar.gz
RUN chmod +x /scriptplusconf/entrypoint2.sh
RUN chmod +x /scriptplusconf/entrypoint1.sh
CMD curl -L "https://raw.githubusercontent.com/bts420/qbittorrent-to-rclone-heroku/main/qb-go-entry-heroku.go" >qb-go-entry-heroku.go; go run qb-go-entry-heroku.go
