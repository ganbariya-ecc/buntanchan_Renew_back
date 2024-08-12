# buntanchan_Renew_back

## 機能を追加する時
1. apps の template をコピーする
2. dockerfile の workdir を編集する
3. docker-compose.yaml に追記する (volumes にさっきのworkdirを指定)
4. nginx の default.conf を編集してリーバスプロキシを通す

## 機能を追加し終わったら
dockerfile の最後に cmd を追加してコンテナ起動時に起動するようにすることをお勧めします。

## 静的ファイル
静的ファイルは nginx の statics に放り込んでください

## アクセスする際は
docker compose で公開したnginx のポート番号に https://localhost:8447/statics/ とかでアクセスすればok

## ユーザーアイコン
- サイズは256 × 256 にする
- ファイル形式はjpeg にする