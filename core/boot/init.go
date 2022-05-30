package boot

import (
	"log"
	"os"

	"test-privy/core/db"
	helper "test-privy/helper"
	"test-privy/middleware/access"
	"test-privy/middleware/customcfg"
	routes "test-privy/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	// IsDebugMode is a variable status
	IsDebugMode bool

	// Env is an environment type
	Env string
)

// Init is a boot function to manage all of the project that will be run.
func Init() {
	if loadEnv() && envCheck() {
		Env = os.Getenv("ENVIRONMENT")
		if os.Getenv("APP_DEBUG") == "1" {
			IsDebugMode = true
		}
	} else {
		// helper.Print("err", "wrong environment variable, check again.", false)
		os.Exit(0)
	}

	if err := dbCheck(); err != nil {
		os.Exit(0)
	}

	createFolder()
}

// Run is used to boot the server after checking
func Run() {
	e := routes.Routes(echo.New())
	e.HideBanner = true
	e.HTTPErrorHandler = customcfg.DefaultErr()
	e.Use(access.GenerateRoutes(e))
	e.Use(
		middleware.Recover(),
		customcfg.Secure(),
		middleware.GzipWithConfig(middleware.GzipConfig{
			Level: 5,
		}),
		// middleware.BodyDump(beforeafter.Dump),
		// echo.WrapMiddleware(beforeafter.Latency),
		// em.RequestID(),
	)
	if IsDebugMode {
		e.Use(customcfg.LoggerCfg())
	}

	if err := e.Start(":" + os.Getenv("PORT")); err != nil {
		log.Println(err)
	}
}

func dbCheck() error {
	user := os.Getenv("MYSQL_DB_USERNAME")
	port := os.Getenv("MYSQL_DB_PORT")
	addr := os.Getenv("MYSQL_DB_ADDR")
	pswd := os.Getenv("MYSQL_DB_PASSWORD")
	dbname := os.Getenv("MYSQL_DB_NAME")

	// assigning a main connection databases variable with db env structure
	db.InitDBMain = user + ":" + pswd + "@(" + addr + ":" + port + ")/" + dbname + "?parseTime=true"

	if os.Getenv("APP_DB_DEBUG") == "1" {
		db.DebugMode = true
	}

	return nil
}
func createFolder() {
	paths := helper.Paths
	for _, s := range paths {
		err := os.MkdirAll(s, 0755)
		if err != nil {
			log.Println("Create Folder failed" + err.Error())
			continue
		}
	}
}
