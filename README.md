# base64

base64でエンコード、デコードするだけのプログラム

ファイルを処理するものと、
パイプで渡ってきたものを処理するものを用意しています。

## インストール

releaseからzipを落として使ってください。
zipには5つのバイナリが含まれています。
* base64
* base64-decode-file
* base64-decode-pipe
* base64-encode-file
* base64-encode-pipe

golangのコンパイラがあれば、

	go get -u -v github.com/xcd0/base64/...

とすれば、上記すべてのバイナリがインストールされます。

