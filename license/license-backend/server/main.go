package main

import (
	"log"
	"server/config"
	"server/delivery/httpserver"
	"server/logger"
	"server/pkg/jwt"
	"server/repository/sqlite"
	"server/repository/sqlite/customerProductRepo"
	"server/repository/sqlite/customerRepo"
	"server/repository/sqlite/logRepo"
	"server/repository/sqlite/productRepo"
	"server/repository/sqlite/restrictionRepo"
	"server/repository/sqlite/userRepo"

	"server/service/customer_product_service"
	"server/service/customer_service"
	"server/service/log_service"
	"server/service/product_service"
	"server/service/restriction_service"
	"server/service/user_service"
)

var cfg *config.MainConfig

func init() {
	// load configs
	cfg = config.Load()
}

func main() {

	// start logger
	logger.NewLogger(cfg.LogCfg)

	switch cfg.DBType {

	case "postgres":
		// start database
		//postgresqlDB := postgresql.PostgreSQLConnection(cfg.PostgresqlCfg)
		//customerPostgresqlDB := pgCustomerRepo.New(postgresqlDB)
		//customersProductPostgresqlDB := pgCustomerProductRepo.New(postgresqlDB)
		//productPostgresqlDB := pgProductRepo.New(postgresqlDB)
		//restrictionPostgresqlDB := pgRestrictionRepo.New(postgresqlDB)
		//userPostgresqlDB := pgUserRepo.New(postgresqlDB)
		//logsRepo := pgLogRepo.New(postgresqlDB)
		//
		//customerSvc := customer_service.New(customerPostgresqlDB)
		//customersProductSvc := customer_product_service.New(customersProductPostgresqlDB)
		//productSvc := product_service.New(productPostgresqlDB)
		//restrictionSvc := restriction_service.New(restrictionPostgresqlDB)
		//logSvc := log_service.New(logsRepo, cfg.ActivityLog)
		//userSvc := user_service.New(userPostgresqlDB)
		//
		//// start jwt
		//jwtSvc := jwt.New(cfg.JwtCfg)
		//
		//// start server
		//
		//server := httpserver.New(cfg.ServerCfg, customersProductSvc,
		//	restrictionSvc, customerSvc, productSvc, userSvc, jwtSvc, logSvc)
		//
		//server.Serve()
		log.Println("can not use postgresql")

	case "sqlite":

		sqlDB := sqlite.SQLiteConnection(cfg.SqliteCfg)
		customerSqliteDB := customerRepo.New(sqlDB)
		customersProductSqliteDB := customerProductRepo.New(sqlDB)
		productSqliteDB := productRepo.New(sqlDB)
		restrictionSqliteDB := restrictionRepo.New(sqlDB)
		userSqliteDB := userRepo.New(sqlDB)
		logsRepo := logRepo.New(sqlDB)

		customerSvc := customer_service.New(customerSqliteDB)
		customersProductSvc := customer_product_service.New(customersProductSqliteDB)
		productSvc := product_service.New(productSqliteDB)
		restrictionSvc := restriction_service.New(restrictionSqliteDB)
		logSvc := log_service.New(logsRepo, cfg.ActivityLog)
		userSvc := user_service.New(userSqliteDB)

		// start jwt
		jwtSvc := jwt.New(cfg.JwtCfg)

		// start server

		server := httpserver.New(cfg.ServerCfg, customersProductSvc,
			restrictionSvc, customerSvc, productSvc, userSvc, jwtSvc, logSvc)

		server.Serve()

	}

}
