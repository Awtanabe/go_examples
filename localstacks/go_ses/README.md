
### 参考

https://zenn.dev/yuki_tu/scraps/2709030959b966

https://docs.localstack.cloud/user-guide/aws/ses/

- go aws

https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/ses-example-send-email.html

### エラー

- 最初は実行できなかった

EMAIL_ADDRESS_VERIFICATIONで認証を無効にしたけど、効かなかった

メールの認証をする

```
aws --endpoint-url=http://localhost:4566 ses verify-email-identity \
    --email-address sender@example.com \
    --region ap-northeast-1
```