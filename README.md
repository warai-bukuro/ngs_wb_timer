# NGSウィークバレットタイマー

PSO2NGSでWBの時間管理をするためのタイマー<br>
Windows10、マウスを利用した場合でのみ実行を確認しています

## 実行内容

1. 15からカウントダウンを始めます
2. マウスの追加ボタン1と2を同時に押した場合にカウントを13に戻します<br>
(感知からリセットまで2秒ほどかかるためです)
3. カウントが0になった場合にもカウントを13に戻します
4. Ctrl + Cで停止できます

## 起動

[Release](https://github.com/warai-bukuro/ngs_wb_timer/releases)から実行ファイルをダウンロードすれば使えると思います<br>
PSO2NGSは管理者権限で実行されているため、このツールも管理者権限で実行する必要があります<br>
(管理者権限でなくても起動はしますが、NGSの操作中にマウスの入力を感知できません)

## ビルド

GoはWindows版ですが、コマンドラインはBash(WSL)を使っています

```
$ git clone git@github.com:warai-bukuro/ngs_wb_timer.git
$ cd ngs_wb_timer
$ go.exe mod tidy
$ go.exe build main.go
```

## 改変を行いたい場合

GPLライセンスが適用できていると思うので、Fork後はソースコードを公開してください<br>
オーバーレイ化やより高速・柔軟にリセットを感知する方法などを追加していただけると幸いです

以上
