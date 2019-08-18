# goa v3 tutorial log

## Just looking

Take a look at the commit messages of these pull requests.

このPull Requestのコミットメッセージを読んでみてください。

- https://github.com/akm/goa_v3_tutorial_log/pull/1
- https://github.com/akm/goa_v3_tutorial_log/pull/2


## Try

### Setup

Clone this repository outside of $GOPATH .

GOPATHに含まれないディレクトリにこのリポジトリをcloneします。

```
git clone git@github.com:akm/goa_v3_tutorial_log.git
cd goa_v3_tutorial_log
```

### Run server

See https://github.com/goadesign/goa#3-run

https://github.com/goadesign/goa#3-run の内容を実行して、サーバを起動します。

```
cd cmd/calc
go build
./calc
```

### Call APIs

See https://github.com/goadesign/goa#3-run

Open another terminal.

もう一つ別のターミナルを起動します。

```
cd calcsvc/cmd/calc-cli
go build
./calc-cli calc add -a 1 -b 2
./calc-cli calc add -a 1 -b foo
```

You can see how to call APIs with `--help`.

APIの使い方は `--help` で確認できます。

```
./calc-cli --help
./calc-cli calc add --help
```

Get static files like /swagger.json with curl.

/swagger.json のような静的ファイルをcurlで取得できます。

```
curl localhost:8088/swagger.json
```
