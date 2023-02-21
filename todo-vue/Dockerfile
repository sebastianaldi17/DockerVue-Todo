FROM node:18.14.1

RUN mkdir -p /web
WORKDIR /web

COPY . .
RUN npm install

EXPOSE 8080

CMD npm run serve