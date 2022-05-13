package models

import (
    "go-blog-demo/pkg/setting"
    "log"
    "time"

    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "gorm.io/gorm/schema"
)

var db *gorm.DB

type Model struct {
    ID        uint `gorm:"primaryKey"`
    CreatedAt time.Time
    UpdatedAt time.Time
}

func init() {
    var (
        err         error
        sqlite_path string
    )

    sec, err := setting.Cfg.GetSection("database")
    if err != nil {
        log.Fatal(2, "Fail to get section 'database': %v", err)
    }

    sqlite_path = sec.Key("SQLITE_PATH").String()

    db, err = gorm.Open(sqlite.Open(sqlite_path), &gorm.Config{
        NamingStrategy: schema.NamingStrategy{
            SingularTable: true,
        }})
    if err != nil {
        log.Println(err)
    }

    sqlDB, err := db.DB()
    if err != nil {
        log.Println(err)
    }

    sqlDB.SetMaxIdleConns(10)
    sqlDB.SetMaxOpenConns(100)
    db.AutoMigrate(&Staff{})

}
