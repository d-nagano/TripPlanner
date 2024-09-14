FROM node:22.6.0

# 作業ディレクトリの設定
WORKDIR /var/www/frontend

COPY package*.json ./
RUN npm install

COPY . .

CMD ["npm", "run", "serve"]