# goa v3 tutorial log

## Just looking

Take a look at the commit messages of these pull requests.

このPull Requestのコミットメッセージを読んでみてください。

- https://github.com/akm/goa_v3_tutorial_log/pull/1
- https://github.com/akm/goa_v3_tutorial_log/pull/2


## Try

### Setup

Use Go 1.12.x

Go 1.12.x を使ってください。


Clone this repository outside of $GOPATH .

GOPATHに含まれないディレクトリにこのリポジトリをcloneします。

```
git clone git@github.com:akm/goa_v3_tutorial_log.git
cd goa_v3_tutorial_log
```

### Setup datastore emulator

```
gcloud components update
gcloud components install cloud-datastore-emulator
```

### Run datastore emulator

```
gcloud beta emulators datastore start
```

Then datastore emulator shows the message about `DATASTORE_EMULATOR_HOST` like

Datastoreエミュレータは以下のようなメッセージを表示します。

```
[datastore] If you are using a library that supports the DATASTORE_EMULATOR_HOST environment variable, run:
[datastore] 
[datastore]   export DATASTORE_EMULATOR_HOST=localhost:8081
[datastore] 
```

### Run server

See https://github.com/goadesign/goa#3-run

https://github.com/goadesign/goa#3-run の内容を実行して、サーバを起動します。


Define environment variable `DATASTORE_EMULATOR_HOST` already shown before starting server. For Example:

サーバを起動する前に、前に表示された環境変数 `DATASTORE_EMULATOR_HOST` を定義します

```
export DATASTORE_EMULATOR_HOST=localhost:8081
```

Set a dummy GCP project ID to environment variable `DATASTORE_PROJECT_ID` .

またDatastoreのクライアントライブラリが接続するダミーのGCPプロジェクト名を指定します。

```
export DATASTORE_PROJECT_ID=dummy-project-id
```

Then start sever.

サーバを起動します。

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
