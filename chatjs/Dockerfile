FROM node:16.16.0-alpine
WORKDIR /usr/src
COPY ./ /usr/src/
RUN npm install -g serve
RUN npm install
RUN npm run build
EXPOSE 3000
CMD ["serve", "-s", "./build", "-l", "3000"]