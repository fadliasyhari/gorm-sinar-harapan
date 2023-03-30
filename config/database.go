package config

import (
	"fmt"

	"enigmacamp.camp/gorm-sinar-harapan-makmur/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConnection interface {
	Conn() *gorm.DB
	Migrate(model ...any) error
}

type dbConnection struct {
	db  *gorm.DB
	cfg *Config
}

func (d *dbConnection) intDb() error {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", d.cfg.Host, d.cfg.Port, d.cfg.User, d.cfg.Password, d.cfg.Name)

	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// Open("postgres", dsn) -> yang masih menggunakan native
	if err != nil {
		return err
	}

	d.db = conn
	return nil
}

func (d *dbConnection) Conn() *gorm.DB {
	return d.db
}

func (d *dbConnection) Migrate(model ...any) error {
	err := d.Conn().AutoMigrate(model...)
	if err != nil {
		return err
	}
	return nil
}

func NewDbConnection(cfg *Config) (DBConnection, error) {
	conn := &dbConnection{
		cfg: cfg,
	}

	err := conn.intDb()
	if cfg.DbConfig.Env != "dev" {
		conn.Migrate(&model.Customer{}, &model.UserCredential{})
	}
	if err != nil {
		return nil, err
	}

	return conn, nil
}
