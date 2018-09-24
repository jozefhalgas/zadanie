#
# TASK: Fix Dockerfile if anything is missing.
FROM nginx

WORKDIR /workdir
COPY index.html /usr/share/nginx/html
COPY default.conf /etc/nginx/conf.d

COPY start.sh start.sh
EXPOSE 80
ENV APP="http://localhost:8080"
RUN chmod +x /workdir/start.sh
CMD ["/workdir/start.sh"]
