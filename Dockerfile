FROM nginx:1.23.3-alpine
WORKDIR /usr/share/nginx/html
COPY . .
EXPOSE 80