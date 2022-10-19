package database

import (
	"auth/app/queries"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Queries struct {
	*queries.UserQueries // load queries from User model
}

func MysqlConnection() (*Queries, error) {
	// set url connection
	mysqlConnectionURL := "root:12345678@tcp(localhost:3306)/auth?charset=utf8&parseTime=True&loc=Local"

	// instance the logger
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: 500 * time.Millisecond, // Slow SQL threshold
			LogLevel:      logger.Info,            // Log level
			Colorful:      true,                   // Disable color
		},
	)

	// connect to the database
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       mysqlConnectionURL, // data source name, refer https://github.com/go-sql-driver/mysql#dsn-data-source-name
		DefaultStringSize:         256,                // add default size for string fields, by default, will use db type `longtext` for fields without size, not a primary key, no index defined and don't have default values
		DisableDatetimePrecision:  true,               // disable datetime precision support, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,               // drop & create index when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,               // use change when rename column, rename rename not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false,              // smart configure based on used version
	}), &gorm.Config{
		Logger:                                   newLogger,
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	// handle database connection errors
	if err != nil {
		fmt.Println("failed to connect database")
		return nil, err
	}

	return &Queries{
		// Set queries from models:
		UserQueries: &queries.UserQueries{DB: db}, // from User model
	}, nil
}
