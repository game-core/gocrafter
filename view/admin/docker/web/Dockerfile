FROM node:18.12.0-alpine AS dev

WORKDIR /web
RUN apk upgrade --update && apk add bash && apk --no-cache add git
COPY package.json yarn.lock ./
RUN yarn install
COPY . .
RUN yarn build

EXPOSE 3000

CMD ["yarn", "dev"]

FROM node:18.12.0-alpine AS prod

WORKDIR /web
COPY --from=development /web/package.json /web/yarn.lock ./
COPY --from=development /web/.nuxt ./.nuxt
COPY --from=development /web/nuxt.config.ts ./
RUN yarn install

EXPOSE 3000

CMD ["yarn", "start"]
