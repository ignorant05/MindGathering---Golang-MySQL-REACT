FROM node:20-alpine as builder

WORKDIR /app

COPY  package*.json ./
RUN npm install --fronzen-lockfile

COPY . .
RUN npm run build 


FROM nginx:alpine
COPY nginx/nginx.conf /etc/nginx/conf.d/default.conf

COPY --from=builder /app/dist /usr/share/nginx/html
EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
