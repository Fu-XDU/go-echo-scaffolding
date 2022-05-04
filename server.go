package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"gopkg.in/urfave/cli.v1"
	"os"
	"server/controller"
	"server/database"
	"server/utils/flags"
)

const (
	clientIdentifier = "ServerApp" // Client identifier to advertise over the network
	clientVersion    = "1.0.0"
	clientUsage      = "A golang web server scaffolding using Echo framework"
)

var (
	app       = cli.NewApp()
	baseFlags = []cli.Flag{
		flags.PortFlag,
	}
	redisFlags = []cli.Flag{
		flags.RedisHostFlag,
		flags.RedisPortFlag,
		flags.RedisDbFlag,
		flags.RedisUserFlag,
		flags.RedisPasswdFlag,
	}
	mysqlFlags = []cli.Flag{
		flags.MysqlHostFlag,
		flags.MysqlPortFlag,
		flags.MysqlUserFlag,
		flags.MysqlPasswdFlag,
		flags.MysqlDBFlag,
	}
)

func init() {
	app.Action = ServerApp
	app.Name = clientIdentifier
	app.Version = clientVersion
	app.Usage = clientUsage
	app.Commands = []cli.Command{}
	app.Flags = append(app.Flags, baseFlags...)
	app.Flags = append(app.Flags, mysqlFlags...)
	app.Flags = append(app.Flags, redisFlags...)
}

func ServerApp(ctx *cli.Context) error {
	if args := ctx.Args(); len(args) > 0 {
		return fmt.Errorf("invalid command: %q", args[0])
	}
	err := prepare(ctx)
	if err != nil {
		log.Error(err)
	}
	return err
}

func prepare(ctx *cli.Context) (err error) {
	if err = database.Setup(); err != nil {
		return
	}

	p := ctx.String("port")
	startNetwork(p)
	return
}

func main() {
	if err := app.Run(os.Args); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func startNetwork(port string) {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/info", controller.Info)

	e.HideBanner = true
	e.Logger.Fatal(e.Start(":" + port))
}
