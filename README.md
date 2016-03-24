codegangsta/cli/altsrc のサンプル
--------

### サンプル・アプリケーションのビルド

サンプルを構築するにあたり、`mattn/gom` を使っています。
gom がインストールされていない場合、予め `$ go get github.com/mattn/gom` でインストールしてください。

ビルドは次のように行います。
_vendor ディレクトリをインストールしなおすなどでなければ、`gom install` の実行は
一度だけ行えばよいです。

```
$ gom install && gom build 
```


サンプル・アプリケーションの実行（メインコマンド）
--------

なにも指定しない場合、ファイルから読み込んだ値が利用されます。
読み込む元なるファイルは、--load オプションによって指定される値が利用されていて、
メインコマンドの --load オプションは、デフォルト値として `"./.rc1 "` を持っている。
このファイルから

```
$ ./cli-example 
cli-example --test = test by .rc1
```

`--test` オプションを指定することで、オプションをコマンドラインから指定することもできます。
このとき、ファイルから入力された内容より優先されます。

```
$ ./cli-example --test "foo"
cli-example --test = foo
```

オプションを指定するキーを持たないファイルを読むと、デフォルト値が使われます。
ためしに、`/dev/null` あたりを読んでみると次のようになります。

```
$ ./cli-example --load /dev/null     
cli-example --test = test default value
```


サンプル・アプリケーションの実行（サブコマンド･サブサブコマンド）
--------

サブサブコマンドも同じようにオプションを指定したり、`--load` オプションや `--test` オプションが使えます。
それぞれのオプションはメインコマンドのものとは別個に指定されており、
デフォルト値も異なるように設定されています。

```
$ ./cli-example sub             
cli-example sub --test-sub = test-sub by .rc2
```

サブサブコマンドにおいても同様です。

```
$ ./cli-example sub subsub
cli-example sub subsub --test-sub-sub = test-sub-sub by .rc3
```

