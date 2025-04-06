FROM node:22-bookworm AS base

COPY . .

RUN npm install && npm run build

FROM node:latest

COPY --from=base dist /usr/share/http

RUN npm install -g http-server

ENTRYPOINT [ "http-server", "/usr/share/http" ]