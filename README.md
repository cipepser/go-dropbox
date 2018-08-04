# go-dropbox

[Dropbox SDK for Go \[UNOFFICIAL\]](https://github.com/dropbox/dropbox-sdk-go-unofficial)でDropboxのAPI叩いてみる。

## インストール

```sh
❯ go get github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/...
```

![create_app](https://github.com/cipepser/go-dropbox/blob/master/img/create_app.png)

※`trans`はすでに取られていたので`trans_app`に変更した。

dropboxのwebページからアプリの画面に移動して、`Generated access token`でトークンを発行する。


TODO: oauth2にする

[この辺](https://github.com/dropbox/dbxcli/blob/bf5af4056e70b5ac5845319522ea39c8fd9ee2ee/cmd/root.go)の実装が参考になりそう。

## References
* [Dropbox SDK for Go \[UNOFFICIAL\]](https://github.com/dropbox/dropbox-sdk-go-unofficial)