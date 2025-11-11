package main

import (
	"bank/handler"
	"bank/repository"
	"bank/service"
	"fmt"
	"net/http"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

func main() {

	initTimeZone()
	initConfig()

	db := initDatabase()
	_ = db

	// using db data
	// customerRepositoryDB := repository.NewCustomerRepositoryDB(db)
	// customerService := service.NewCustomerService(customerRepositoryDB)
	// customerHandler := handler.NewCustomerHandler(customerService)

	// using mock data
	customerRepositoryMock := repository.NewCustomerRepositoryMock()
	customerServiceMock := service.NewCustomerService(customerRepositoryMock)
	customerHandlerMock := handler.NewCustomerHandler(customerServiceMock)

	router := mux.NewRouter()

	// router.HandleFunc("/customers", customerHandler.GetCustomers).Methods(http.MethodGet)
	// router.HandleFunc("/customers/{customerId:[0-9]+}", customerHandler.GetCustomer).Methods(http.MethodGet)

	router.HandleFunc("/customers", customerHandlerMock.GetCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customerID:[0-9]+}", customerHandlerMock.GetCustomer).Methods(http.MethodGet)

	fmt.Printf("Banking service started at port %v", viper.GetInt("app.port"))
	http.ListenAndServe(fmt.Sprintf(":%v", viper.GetInt("app.port")), router)

	// customers, err := customerRepositoryDB.GetAll()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(customers)

	// customer, err :=customerRepositoryDB.GetById(20)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(customer)

	// using service layer
	// customers, err := customerService.GetCustomers()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(customers)

	// customer, err := customerService.GetCustomer(10)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(customer)

}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func initTimeZone() {

	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}
	time.Local = ict
}

func initDatabase() *sqlx.DB {

	dns := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.host"),
		viper.GetInt("db.port"),
		viper.GetString("db.database"),
	)

	db, err := sqlx.Open(viper.GetString("db.driver"), dns)
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(3 * time.Minute)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}
