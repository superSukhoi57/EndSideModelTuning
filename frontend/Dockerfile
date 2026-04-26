FROM nginx:mainline-alpine3.22

RUN rm -rf /usr/share/nginx/html/*

COPY build/ /usr/share/nginx/html/

# 添加 nginx 配置，支持 SPA 路由回退
RUN echo 'server { \
    listen 80; \
    server_name localhost; \
    root /usr/share/nginx/html; \
    index index.html; \
    location / { \
        try_files $uri $uri/ /index.html; \
    } \
}' > /etc/nginx/conf.d/default.conf

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
