FROM node:20.11-alpine

WORKDIR /app/

COPY app/public/ /app/public
COPY app/src/ /app/src
COPY app/package.json /app/

RUN npm install

CMD ["npm", "start"]