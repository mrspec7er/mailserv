FROM node:22-bookworm

WORKDIR /server

COPY package.json ./

RUN npm install

COPY . .

EXPOSE 1025

CMD npm run dev