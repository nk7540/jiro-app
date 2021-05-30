# Artics API

## 動作環境

* Docker 20.10.5
* Docker Compose 1.28.5

## Setup

```
# 初回のみ(コンテナ起動)
docker-compose up -d

# サーバ立ち上げ
docker-compose exec app bash
[artics-api]$ go run src/cmd/main.go
```

### アクセス

```
curl localhost:8080/api/v1/user/1
```

### DB更新

* マイグレーションファイル

  `db/migrations/20210530063731-create_tables.sql`

```
# マイグレーションの作成
[artics-api]$ sql-migrate new add_new_column

# マイグレーションの実行
[artics-api]$ sql-migrate up

# ORMの更新(スキーマ更新後)
[artics-api]$ sqlboiler mysql
```

