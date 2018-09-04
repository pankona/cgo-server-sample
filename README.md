# cgo-server-sample

Go のサーバがリクエストを受けたら C にコールバックし、C 側で色々処理してもらった結果を Go がレスポンスを返すサンプル

## ビルド

```bash
$ cd C
$ make
```

## 動かし方

上記ビルド手順でビルドすると `main` という実行バイナリが生成されるので、実行する
* :8080 で待ち受けるサーバが立ち上がります。`/api` という API が生えています。
* ブラウザ等で `http://localhost:8080/api` にアクセスします。
* 何やら文字が出たら動いています。

## LICENSE

* MIT

## Author

* Yosuke Akatsuka (a.k.a pankona)
