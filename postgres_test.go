package main

import (
	"testing"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestPostgresConnectionInCI(t *testing.T) {
	// CI sunucusunda ayağa kalkan PostgreSQL'e bağlanıyoruz
	dsn := "host=localhost user=postgres password=password dbname=featureflow_test port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("Github Actions üzerinde PostgreSQL'e ulaşılamadı! Hata: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		t.Fatalf("Veritabanı nesnesi alınamadı: %v", err)
	}
	defer sqlDB.Close()

	err = sqlDB.Ping()
	if err != nil {
		t.Fatalf("PostgreSQL'e Ping atılamadı: %v", err)
	}

	t.Log("Github Actions üzerinde PostgreSQL başarıyla ayağa kalktı ve yanıt veriyor!")
}
