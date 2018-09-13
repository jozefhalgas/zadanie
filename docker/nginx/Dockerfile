#
# TASK: Fix Dockerfile if anything is missing.
FROM nginx

COPY index.html /usr/share/nginx/html
COPY default.conf /etc/nginx/conf.d

COPY start.sh start.sh

EXPOSE 80

CMD ["./start.sh"]
