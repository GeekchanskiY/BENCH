FROM node:20.11-alpine

WORKDIR /app/

COPY app/public/ /app/public
COPY app/src/ /app/src
COPY app/package.json /app/

# TODO: npm install does not work in docker container(???)
RUN npm install

ENV PORT=3002

CMD ["npm", "start"]