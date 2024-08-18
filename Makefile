SHELL := /bin/bash
.DEFAULT_GOAL := help
CURRENT_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

# 環境変数を読み込む
ifneq (,$(wildcard .env))
    include .env
    export $(shell sed 's/=.*//' .env)
endif

.PHONY: help
help: ## ヘルプ
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sed 's/Makefile://' | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

.PHONY: dev
dev: ## 開発サーバーを起動する
	go run cmd/main.go

.PHONY: gen-grpc
gen-grpc: ## protoファイルからgRPCのコード生成
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
	protoc -I=./proto --go_out=. --go-grpc_out=. ./proto/*.proto

.PHONY: db-start, db-stop, db-destroy
db-start: ## DB用Dockerコンテナを起動する
	docker-compose up -d
db-stop: ## DB用Dockerコンテナを停止する
	docker-compose down
db-destroy: ## DB用Dockerコンテナを破棄する
	docker-compose down --rmi all --volumes --remove-orphans

.PHONY: open-grpcui
open-grpcui: ## gRPC-UIを開く
	grpcui -plaintext localhost:50051

.PHONY: gen-gorm
gen-gorm: ## Gormのモデルを生成する
	go run scripts/gen-gorm.go
