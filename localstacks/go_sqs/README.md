### cli install

```
brew install localstack/tap/localstack-cli
```

### 

- キューの作成

```
aws --endpoint-url=http://localhost:4566 sqs create-queue --queue-name test-queue
```


- キューの一覧

```
aws --endpoint-url=http://localhost:4566 sqs list-queues 
```

- message送信

```
aws --endpoint-url=http://localhost:4566 sqs send-message --queue-url http://localhost:4566/000000000000/test-queue --message-body "Hello, LocalStack!"
```

- 受信

```
aws --endpoint-url=http://localhost:4566 sqs receive-message --queue-url http://localhost:4566/000000000000/test-queue
```

- 削除

receipthandleのidを指定するみたい


```
aws --endpoint-url=http://localhost:4566 sqs delete-message \
--queue-url http://localhost:4566/000000000000/test-queue \
--receipt-handle "xxxxxx"
```

- キューのメッセージの状態確認


```
aws --endpoint-url=http://localhost:4566 sqs get-queue-attributes \
  --queue-url http://localhost:4566/000000000000/test-queue \
  --attribute-names ApproximateNumberOfMessages ApproximateNumberOfMessagesNotVisible
```