FROM node:19-alpine

WORKDIR /frontend/jobapp
COPY  ./package*.json ./
RUN apk update && apk add git && apk add vim && npm install

