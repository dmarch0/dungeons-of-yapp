FROM node:18.16.0-alpine

WORKDIR /app

COPY package.json /app/package.json
COPY package-lock.json /app/package-lock.json

RUN npm install --include=dev

ENTRYPOINT  ["npm", "run", "dev"]