FROM nginx:mainline-alpine3.22

RUN rm -rf /usr/share/nginx/html/*

COPY build/ /usr/share/nginx/html/

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]

# frontend/public 下的所有资源会被一起打包。
# - public/config/app.json → build后变成 build/config/app.json
# - public/index.html 、 public/gemini_sztu.png 等也都会被复制

# 从 config.ts 可以看到，代码通过 fetch('/config/app.json') 获取配置，这是一个相对于网站根目录的路径。
# 在 nginx 中，网站根目录默认是 /usr/share/nginx/html 。所以：
# 外部挂载路径应该是： /usr/share/nginx/html/config/app.json


# daemon off; = 禁止后台运行
# Docker 必须让程序前台运行，否则容器退出
# 这是 Docker + Nginx 的标准固定写法