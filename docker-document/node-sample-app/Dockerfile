FROM node:18-alpine

WORKDIR /app

COPY . .

RUN yarn install --production
## コンテナが起動した時にデフォルトで実行するコマンド
CMD ["node", "src/index.js"]
EXPOSE 3000