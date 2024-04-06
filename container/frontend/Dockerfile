FROM node:lts-alpine

# ワークディレクトリを指定
WORKDIR /src

# 環境変数を設定
ENV WATCHPACK_POLLING=true

# ソースコードをコンテナにコピー
COPY ./src /src

# 各パッケージをインストール
RUN apk update && \
    apk upgrade && \
    apk add git && \
    npm install -g npm@latest nuxi nuxt3

#CMD ["npm", "run", "dev"]