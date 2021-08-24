package assets

import "fmt"

func GetMysqlDB(projectName string) (string, string) {
	return fmt.Sprintf(`package db

import (
	"%[1]s/source/config"

	"context"
	"log"
	"os"
	"time"

	"go.uber.org/fx"
	"go.uber.org/zap"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

//Param is all dependencies needed to create database connection
type Param struct {
	fx.In
	Config config.Config
	Logger *zap.Logger
}

// Module is all di needed for db connection
var Module = fx.Options(
	fx.Provide(NewConn),
	fx.Invoke(testConn),
)

type configuration struct {
	User       string
	Password   string
	Network    string
	Address    string
	Port       string
	DbName     string
	Parameters string
}

func (cfg configuration) formatDSN() string {
	return cfg.User + ":" + cfg.Password + "@" + cfg.Network +
		"(" + cfg.Address + ")/" + cfg.DbName + "?" + cfg.Parameters
}

//NewConn is provider for type *Pool
func NewConn(p Param) (db *gorm.DB, err error) {

	cfg := configuration{
		User:       p.Config.GetString("db_username"),
		Password:   p.Config.GetString("db_password"),
		Network:    p.Config.GetString("db_protocol"),
		Address:    p.Config.GetString("db_address"),
		Port:       p.Config.GetString("db_port"),
		DbName:     p.Config.GetString("db_name"),
		Parameters: p.Config.GetString("db_params"),
	}

	URL := cfg.formatDSN()

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      false,       // Disable color
		},
	)
	db, err = gorm.Open(mysql.Open(URL), &gorm.Config{Logger: newLogger})
	if err != nil {
		return nil, err
	}

	return db, nil
}

type param struct {
	fx.In
	DB       *gorm.DB
	Logger   *zap.Logger
}

func testConn(par param) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	sqlDB, err := par.DB.DB()
	if err != nil {
		par.Logger.Fatal("Unable to get database: %v\n", zap.Error(err))
		return
	}

	err = sqlDB.PingContext(ctx)
	if err != nil {
		par.Logger.Fatal("Unable to connect to database: %v\n", zap.Error(err))
		return
	}

	par.Logger.Info("Successfully connected to database.")
}`, projectName), "./" + projectName + "/source/db/db.go"
}
