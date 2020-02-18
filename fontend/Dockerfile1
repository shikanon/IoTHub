FROM node:12.13.0

RUN mkdir -p /app

COPY . /app

WORKDIR /app

#VOLUME ["/app/dist"]

RUN npm install -g cnpm --registry=https://registry.npm.taobao.org

RUN npm config set registry https://registry.npm.taobao.org

RUN npm install

EXPOSE 8080
CMD ["npm", "run", "dev"]


