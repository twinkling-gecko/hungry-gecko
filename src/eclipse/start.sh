#!/bin/sh

cd /usr/src/tangerine && NUXT_PORT=3000 npm run serve&
cd /usr/src/macksnow && ./macksnow&

# build
/usr/bin/erb /usr/src/eclipse/nginx.conf.erb > /etc/nginx/nginx.conf

# start nginx
nginx
