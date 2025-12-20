FROM node:24-alpine


RUN yarn set version berry

WORKDIR /app

COPY ./frontend/package.json ./
COPY ./frontend/yarn.lock ./


COPY ./frontend/ .
COPY ./frontend/.env .

RUN rm .yarnrc.yml

RUN yarn install
RUN yarn build

CMD ["yarn", "dlx", "serve", "-s", "dist", "-l", "tcp://0.0.0.0:10006" ]

