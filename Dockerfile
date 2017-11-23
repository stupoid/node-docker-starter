FROM node:9-alpine

WORKDIR /usr/src/app

COPY ./app/package*.json /usr/src/app

RUN npm install