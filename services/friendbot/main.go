package main

import (
	"database/sql"
	"fmt"
	stdhttp "net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/spf13/cobra"
	"github.com/test/go/services/friendbot/internal"
	"github.com/test/go/support/app"
	"github.com/test/go/support/config"
	"github.com/test/go/support/errors"
	"github.com/test/go/support/http"
	"github.com/test/go/support/log"
	"github.com/test/go/support/render/problem"
)

// Config represents the configuration of a friendbot server
type Config struct {
	Port              int         `toml:"port" valid:"required"`
	FriendbotSecret   string      `toml:"friendbot_secret" valid:"required"`
	NetworkPassphrase string      `toml:"network_passphrase" valid:"required"`
	TestHorizonURL        string      `toml:"testhorizon_url" valid:"required"`
	StartingBalance   string      `toml:"starting_balance" valid:"required"`
	TLS               *config.TLS `valid:"optional"`
}

func main() {

	rootCmd := &cobra.Command{
		Use:   "friendbot",
		Short: "friendbot for the Test Test Network",
		Long:  "client-facing api server for the friendbot service on the Test Test Network",
		Run:   run,
	}

	rootCmd.PersistentFlags().String("conf", "./friendbot.cfg", "config file path")
	rootCmd.Execute()
}

func run(cmd *cobra.Command, args []string) {
	var (
		cfg     Config
		cfgPath = cmd.PersistentFlags().Lookup("conf").Value.String()
	)
	log.SetLevel(log.InfoLevel)

	err := config.Read(cfgPath, &cfg)
	if err != nil {
		switch cause := errors.Cause(err).(type) {
		case *config.InvalidConfigError:
			log.Error("config file: ", cause)
		default:
			log.Error(err)
		}
		os.Exit(1)
	}

	fb := initFriendbot(cfg.FriendbotSecret, cfg.NetworkPassphrase, cfg.TestHorizonURL, cfg.StartingBalance)
	router := initRouter(fb)
	registerProblems()

	addr := fmt.Sprintf("0.0.0.0:%d", cfg.Port)

	http.Run(http.Config{
		ListenAddr: addr,
		Handler:    router,
		TLS:        cfg.TLS,
		OnStarting: func() {
			log.Infof("starting friendbot server - %s", app.Version())
			log.Infof("listening on %s", addr)
		},
	})
}

func initRouter(fb *internal.Bot) *chi.Mux {
	mux := http.NewAPIMux(false)

	handler := &internal.FriendbotHandler{Friendbot: fb}
	mux.Get("/", handler.Handle)
	mux.Post("/", handler.Handle)
	mux.NotFound(stdhttp.HandlerFunc(func(w stdhttp.ResponseWriter, r *stdhttp.Request) {
		problem.Render(r.Context(), w, problem.NotFound)
	}))

	return mux
}

func registerProblems() {
	problem.RegisterError(sql.ErrNoRows, problem.NotFound)
}