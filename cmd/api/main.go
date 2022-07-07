package main

import (
	"errors"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"net/http"
	"strings"
	"time"

	"github.com/edipermadi/music-db/internal/theory"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	defer func() { _ = logger.Sync() }()

	vpr, err := loadConfig()
	if err != nil {
		logger.With(zap.Error(err)).Fatal("failed to load configuration")
	}

	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/api/").Subrouter()
	apiV1Router := apiRouter.PathPrefix("/v1").Subrouter()
	theoryV2Router := apiV1Router.PathPrefix("/theory").Subrouter()

	user := vpr.GetString("db.user")
	password := vpr.GetString("db.password")
	database := vpr.GetString("db.database")
	host := vpr.GetString("db.host")
	port := vpr.GetInt("db.port")
	apiPort := vpr.GetInt("api.port")

	connConfig, err := pgx.ParseConfig(fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=disable",
		user, password, database, host, port,
	))
	if err != nil {
		logger.With(zap.Error(err)).Fatal("failed to parse database connection config")
	}

	db, err := sqlx.Open("pgx", stdlib.RegisterConnConfig(connConfig))

	theoryRepository := theory.NewRepository(logger, db)
	theoryService := theory.NewService(logger, theoryRepository)
	handlerHandler := theory.NewHandler(logger, theoryService)

	theoryV2Router.HandleFunc("/pitches", handlerHandler.ListPitches).
		Methods(http.MethodGet).Name("LIST_PITCHES")

	theoryV2Router.HandleFunc("/chords", handlerHandler.ListChords).
		Methods(http.MethodGet).Name("LIST_CHORDS")

	theoryV2Router.HandleFunc("/scales", handlerHandler.ListScales).
		Methods(http.MethodGet).Name("LIST_SCALES")
	theoryV2Router.HandleFunc("/scales/{id:[0-9]+}", handlerHandler.GetScale).
		Methods(http.MethodGet).Name("GET_SCALE")
	theoryV2Router.HandleFunc("/scales/{id:[0-9]+}/keys", handlerHandler.ListScaleKeys).
		Methods(http.MethodGet).Name("LIST_KEYS_FROM_SCALE")

	theoryV2Router.HandleFunc("/keys", handlerHandler.ListKeys).
		Methods(http.MethodGet).Name("LIST_KEYS")
	theoryV2Router.HandleFunc("/keys/{id:[0-9]+}/modes", handlerHandler.ListKeyModes).
		Methods(http.MethodGet).Name("LIST_KEY_MODES")
	theoryV2Router.HandleFunc("/keys/{id:[0-9]+}/chords", handlerHandler.ListKeyChords).
		Methods(http.MethodGet).Name("LIST_KEY_CHORDS")
	theoryV2Router.HandleFunc("/keys/{id:[0-9]+}/pitches", handlerHandler.ListKeyPitches).
		Methods(http.MethodGet).Name("LIST_KEY_PITCHES")
	theoryV2Router.HandleFunc("/keys/{id:[0-9]+}", handlerHandler.GetKey).
		Methods(http.MethodGet).Name("GET_KEY")

	srv := &http.Server{
		Handler:           handlers.RecoveryHandler(handlers.PrintRecoveryStack(true))(router),
		Addr:              fmt.Sprintf(":%d", apiPort),
		WriteTimeout:      15 * time.Second,
		ReadTimeout:       15 * time.Second,
		ReadHeaderTimeout: 15 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		logger.With(zap.Error(err)).Error("failed to start server")
	}

	logger.Info("exiting")
}

func loadConfig() (*viper.Viper, error) {
	vpr := viper.New()
	vpr.SetDefault("mode", "docker")
	vpr.AutomaticEnv()
	vpr.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	vpr.SetConfigType("yaml")
	vpr.AddConfigPath("./config") //for go run main.go

	mode := vpr.GetString("mode")
	vpr.SetConfigName(mode)
	if err := vpr.ReadInConfig(); err != nil {
		return nil, err
	}

	return vpr, nil
}
