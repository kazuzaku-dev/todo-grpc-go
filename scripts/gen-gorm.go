package main

import (
	"os"
	"strings"
	"unicode"

	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	dbHost := os.Getenv("DATABASE_HOST")
	dbPort := os.Getenv("DATABASE_PORT")
	dbUser := os.Getenv("DATABASE_USER")
	dbPasswored := os.Getenv("DATABASE_PASSWORD")
	dbName := os.Getenv("DATABASE_NAME")

	dsn := dbUser + ":" + dbPasswored + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"

	// GORMを使ってデータベースに接続
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// ジェネレータを作成
	g := gen.NewGenerator(gen.Config{
		OutPath: "internal/generated/gorm/model", // 生成されたファイルを出力するパス
	})

	// 接続されたデータベースに基づいてモデルを自動生成
	g.UseDB(db)

	// モデルの名前をカスタマイズする
	// ドメイン層のモデルと区別するために、モデル名の先頭に「G」を付ける
	g.WithModelNameStrategy(func(tableName string) string {
		return "G" + toCamelCase(tableName)
	})

	// すべてのテーブルに対応するモデルを生成
	g.GenerateAllTable()

	// ファイルを出力
	g.Execute()
}

// スネークケースの文字列をキャメルケースに変換する
func toCamelCase(s string) string {
	parts := strings.Split(s, "_")
	for i, part := range parts {
		if len(part) > 0 {
			parts[i] = string(unicode.ToUpper(rune(part[0]))) + part[1:]
		}
	}
	return strings.Join(parts, "")
}
