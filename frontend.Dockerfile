FROM node:lts-alpine

RUN yarn global add serve

WORKDIR /app

COPY ./frontend/doobidoo/package.json ./
COPY ./frontend/doobidoo/yarn.lock ./

RUN yarn install

COPY ./frontend/doobidoo/ .
COPY ./frontend/doobidoo/.env .

RUN yarn build

CMD ["serve", "-s", "dist", "-l", "tcp://0.0.0.0:10006" ]

