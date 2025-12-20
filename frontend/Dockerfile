FROM node:lts-alpine

RUN yarn global add serve

WORKDIR /app

COPY ./package.json ./
COPY ./yarn.lock ./

RUN yarn install

COPY ./ .
COPY ./.env .

RUN yarn build

CMD ["yarn", "dlx", "serve", "-s", "dist", "-l", "tcp://0.0.0.0:10006" ]

