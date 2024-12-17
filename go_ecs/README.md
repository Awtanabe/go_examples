## 参考記事

- ecsって何か

https://zenn.dev/mi_01_24fu/articles/aws-ecs-2024_03_18

- 構築
https://zenn.dev/pirotyyy/articles/fae69e25f9807a

## やりたいこと

- go docker ecs
- rds接続
- ci-cd
- ログ検知


## ecs 構成

- クラスタ
- タスク
- サービス
- vpc

### ecs

https://ap-northeast-1.console.aws.amazon.com/ecr/repositories/private/<amazon-id>/go-api-ecr-test?region=ap-northeast-1


```
aws ecr get-login-password | docker login --username AWS --password-stdin <account-id>.dkr.ecr.<reagin>.amazonaws.com

/// You must specify a region. You can also configure your region by running "aws configure".
設定が必要でアクセスキーなどの設置が必要だった


// イメージのpush ECRにイメージを
docker build . -t go-api-ecr-test
docker tag go-api-ecr-test:latest <amazon-id>.dkr.ecr.ap-northeast-1.amazonaws.com/go-api-ecr-test:1.0.0
docker push <amazon-id>.dkr.ecr.ap-northeast-1.amazonaws.com/go-api-ecr-test:1.0.0
```

### クラスター

https://ap-northeast-1.console.aws.amazon.com/ecs/v2/getStarted?region=ap-northeast-1

- 手動でクラスター作成
  - fargateで作成

### タスク作成

- ecs画面の再度バーに「タスク定義」がある
- imageのurlは ecrにある

### サービス作成

- 記事に沿って
- ⭐️インバウンド設定
  - サービスから設定とネットワーク
- curl http://{Public IP}:8080/

### デプロイ後

- パブリックipは タスクの詳細
  - ネットワーキング
- インバウンドルール修正
  - 全てのネットワークを8080で
  

### デバック

タスクのログで