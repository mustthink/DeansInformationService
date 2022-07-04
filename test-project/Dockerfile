FROM node:lts-alpine

# устанавливаем простой HTTP-сервер для статики
RUN npm install -g http-server

RUN mkdir web
COPY package*.json ./web

# устанавливаем зависимости проекта
RUN cd web && npm install

# копируем файлы и каталоги проекта в текущий рабочий каталог (т.е. в каталог 'app')
COPY . .

CMD npm run serve --prefix ./web