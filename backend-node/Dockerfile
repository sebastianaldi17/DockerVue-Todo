FROM node:18.14.1

RUN mkdir -p /node-server
WORKDIR /node-server

COPY . .
RUN npm install

EXPOSE 8000

CMD npm run dev