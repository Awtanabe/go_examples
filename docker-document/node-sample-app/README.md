
### その他

shコマンド
-c 引数の文字列をコマンド実行

- macなどは、zsh bashがすでに起動している
  sh を本来は実行しないと、だめ
  確かに docker exec -it ubuntu bash とかもbash起動してる


- シェルのネスト

https://chatgpt.com/c/677160d3-3ca8-8004-abcb-eb10224f88ab

docker exec -it ubuntu bash 
はコンテナでシェルを起動させるコマンド
=> シェルを起動させることでコマンドを受け付けて実行してくれる

### イメージ

- image
  - 環境設定が詰まった物(Dockerfileで定義される)
    - WORKDIRの作業ディレクトリ
    - image(ここにもライブラリがある)
    - apy-get install ライブラリ
    - CMD実行
- dockerfile
  - sh -c
    - シェルを起動して、コマンドを実行する
- 環境変数
  - 渡し方
- コンテナ
  - dockerホスト内で稼働する、プロセスの環境
  - up, stop のステータスがある
  - rmなどのステータス

- ファイルシステム
  - 基本はチ領域に保存される
  - バインド
  - volumeのいずれかで保存される
- ネットワーク
  - -pでポートマッピング



### 作業
- イメージ作成
Dockerfile
docker build -t getting-started .

- コンテナ起動

// d はdetach
// p publish ポートマッピング。ホストとコンテナ間でポートの関連付
docker run -dp 127.0.0.1:3000:3000 getting-started

- コード
https://github.com/docker/getting-started.git

## 手順

- git cloneする
- cd getting-started/app移動
- dockerfile作成
- ビルドとrun
- port を300で開けて、-pでマッピングする


```
FROM node:18-alpine

WORKDIR /app

COPY . .

RUN yarn install --production
## コンテナが起動した時にデフォルトで実行するコマンド
CMD ["node", "src/index.js"]
EXPOSE 3000
```

### 更新

https://docs.docker.jp/get-started/03_updating_app.html

- コードを変更する
- イメージの更新
docker build -t getting-started .

- docker stop

停止状態でプロセスを終了
データや設定はそのまま
docker startで再開できる

- docker rm


### アプリケーションの共有

https://docs.docker.jp/get-started/04_sharing_app.html

共有

### データベースの保持

https://docs.docker.jp/get-started/05_persisting_data.html


- コンテナのファイルシステム
  - 何もしないと、作成したデータは「スクラッチ領域scratch space 」、というのがある(一時的なデータの保存領域)

https://docs.docker.jp/get-started/06_bind_mounts.html

- バインド
  - ホスト側に保存されてコンテナを起動するときに共有される
  - ファイル変更を監視して、アプリケーションを自動更新する
  - volumeとの違い
    - volume
      - volume名とコンテナ内の場所
      - volumeドライバがある
    - バインド
      - src と targetの指定
  - 作業
    - docker run -it --mount type=bind,src="$(pwd)",target=/src ubuntu bash
  
- volume
  - dockerホストに保存される
  - ボリュームはディレクトリごと保存できる
    - docker run -dp 127.0.0.1:3000:3000 --mount type=volume,src=todo-db,target=/etc/todos getting-started
      - /etc/todosはコンテナの中に存在する
      - /src はボリューム名
      - bashに入る
        - docker exec -it 3ed8f3562a91 sh
    - やってること
      - docker volume create todo-db 保存領域
      - 紐付け(volumeとコンテナの中の場所)
        - docker run -dp 127.0.0.1:3000:3000 --mount type=volume,src=todo-db,target=/etc/todos getting-started


```
// data.txtを作成しコンテナに入れる
docker run -d ubuntu bash -c "shuf -i 1-10000 -n 1 -o /data.txt && tail -f /dev/null"


docker exec 413fe8269c1eb462f5a24cc356953e4dce20a49e7de169a63c0b5adfa6400706 cat /data.txt
7973
```


- コンテナの起動
  - pでポートのマッピング

```
// -w /app 作業ディレクトリ
// shでコマンドを十個おうしてる

docker run -dp 127.0.0.1:3000:3000 \
    -w /app --mount type=bind,src="$(pwd)",target=/app \
    node:18-alpine \
    sh -c "yarn install && yarn run dev"
```

### ファイルなどの共有

- ボリューム
  - どこに保存されているのか？
    - ホストマシンの /var/lib/docker/volumes/

docker desktopを使っていると仮想環境に保存されているから

```
➜  ~ docker volume ls
DRIVER    VOLUME NAME
local     ac5b2deb810c5573c7647dea4e35807071d96febbc58175604bbebca829f62d5
local     autocall_postgres_volume
local     autocall_redis
local     gopracticetodo_db-data
➜  ~ 

➜  ~ docker volume inspect ac5b2deb810c5573c7647dea4e35807071d96febbc58175604bbebca829f62d5
[
    {
        "CreatedAt": "2024-12-29T09:45:50Z",
        "Driver": "local",
        "Labels": null,
        "Mountpoint": "/var/lib/docker/volumes/ac5b2deb810c5573c7647dea4e35807071d96febbc58175604bbebca829f62d5/_data",
        "Name": "ac5b2deb810c5573c7647dea4e35807071d96febbc58175604bbebca829f62d5",
        "Options": null,
        "Scope": "local"
    }
]


// アクセス
docker run --rm -it -v ac5b2deb810c5573c7647dea4e35807071d96febbc58175604bbebca829f62d5:/data alpine sh

// postgresのデータを確認できる

ls /data
PG_VERSION            pg_dynshmem           pg_multixact          pg_snapshots          pg_tblspc             postgresql.auto.conf
base                  pg_hba.conf           pg_notify             pg_stat               pg_twophase           postgresql.conf
global                pg_ident.conf         pg_replslot           pg_stat_tmp           pg_wal                postmaster.opts
pg_commit_ts          pg_logical            pg_serial             pg_subtrans           pg_xact               postmaster.pid
/
```

- バインドマウント

```
// マウント自体を理解するためにgptに
// mount時権限不要
docker run -it --rm --privileged ubuntu bash

// bindコマンド

apt update && apt install -y mount

// 準備
mkdir /source_dir
mkdir /target_dir
echo "Hello from source_dir" > /source_dir/example.txt

// bind
mount --bind /source_dir /target_dir

```

### 複数コンテナ

https://docs.docker.jp/get-started/07_multi_container.html

- メモ
  - todo アプリとmysqlなど
  - コンテナは独立してる

- やること
  - docker network create todo-app
    - ネットワーク構築
  - mysqlでデータ確認
  - todoアプリからdbに接続

```
// ネットワーク追加
docker run -d \
     --network todo-app --network-alias mysql \
     -v todo-mysql-data:/var/lib/mysql \
     -e MYSQL_ROOT_PASSWORD=secret \
     -e MYSQL_DATABASE=todos \
     mysql:8.0

//todoアプリからdbに接続

docker run -it --network todo-app nicolaka/netshoot

// ipにリクエストできる
dig mysql


// todo とmysqlの
$ docker run -dp 127.0.0.1:3000:3000 \
   -w /app -v "$(pwd):/app" \
   --network todo-app \
   -e MYSQL_HOST=mysql \
   -e MYSQL_USER=root \
   -e MYSQL_PASSWORD=secret \
   -e MYSQL_DB=todos \
   node:18-alpine \
   sh -c "yarn install && yarn run dev"
```

### docker composeを使う

- dockerだけで作成してたものをdocker-compose
  - dockerと githubから落としてきたsrcを利用して

```
$ docker run -dp 127.0.0.1:3000:3000 \
  -w /app -v "$(pwd):/app" \
  --network todo-app \
  -e MYSQL_HOST=mysql \
  -e MYSQL_USER=root \
  -e MYSQL_PASSWORD=secret \
  -e MYSQL_DB=todos \
  node:18-alpine \
  sh -c "yarn install && yarn run dev"
```

### イメージのベストプラクティス


- イメージの階層化

docker image history 

- キャッシュ

- マルチステージビルド

### ネットワーク

https://docs.docker.jp/engine/userguide/networking/dockernetworks.html

- ネットワーク一覧を表示


- コンテナとホストの関係
  - NATで繋がれている
    - プライベートIPをパブリックIPアドレスに変換
      - iptablesが利用 ファイアーウォールとルーティングを兼ね備えている「
- bridge
  - デフォルトのネットワーク空間
    - todo apiと mysql の独立したコンテナを作ってもお互い bridgeに属していれば通信できる
      - 今までにあったのが todo api が自作 networkで mysqlがbridgeネットワークだったのかもしれない

```
// host はdockerを稼働してるパソコン環境のホスト
$ docker network ls
NETWORK ID     NAME      DRIVER    SCOPE
cc1741c9feca   bridge    bridge    local
d1b5ab168f8c   host      host      local
ccade9920c46   none      null      local

// bridge
$ docker run -p 3000:3000 nginx

// host だから -pのポートマッピングが不要になる
$docker run --network=host nginx
```

- ecsでのネットワークは別途学習

https://zenn.dev/fdnsy/articles/43b7f4d745ed1f

- 関連記事


bridgeについて
https://qiita.com/kenny_J_7/items/77de780d7193b75444c3

iptables
https://qiita.com/dtanimoto00/items/1d0a9b02867add646ea5

### docker compose


https://docs.docker.jp/compose/toc.html



