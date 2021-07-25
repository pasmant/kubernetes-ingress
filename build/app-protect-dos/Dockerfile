# For Alpine:
FROM alpine:latest

# Set Workdir
WORKDIR /app

# Install prerequisite packages:
RUN apk update && apk add wget

# Download NAP DOS Arbitrator
RUN --mount=type=secret,id=nginx-repo.crt,dst=/etc/ssl/nginx/nginx-repo.crt,mode=0644 \
    --mount=type=secret,id=nginx-repo.key,dst=/etc/ssl/nginx/nginx-repo.key,mode=0644 \
    wget --certificate=/etc/ssl/nginx/nginx-repo.crt --private-key=/etc/ssl/nginx/nginx-repo.key https://pkgs.nginx.com/app-protect-dos/app-protect-dos-arb

# Remove nginx repository key/cert from docker
RUN rm -rf /etc/ssl/nginx


RUN chmod 755 /app/app-protect-dos-arb
CMD ["/app/app-protect-dos-arb"]