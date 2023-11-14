FROM node:18.12.0-alpine AS dev

WORKDIR /swagger
RUN apk upgrade --update && apk add bash && apk --no-cache add git
RUN yarn global add @stoplight/prism-cli

EXPOSE 8000
