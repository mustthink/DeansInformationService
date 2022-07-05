FROM node:lts-alpine


COPY . ./web

RUN cd web && npm install


CMD npm run serve --prefix ./web