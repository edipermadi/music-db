package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/edipermadi/music-db/internal/theory"
	"github.com/edipermadi/music-db/pkg/midi"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/sinshu/go-meltysynth/meltysynth"
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
		return
	}

	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/api").Subrouter()

	apiV1Router := apiRouter.PathPrefix("/v1").Subrouter()
	theoryRouter := apiV1Router.PathPrefix("/theory").Subrouter()

	apiRouter.PathPrefix("/docs").Handler(http.StripPrefix("/api/docs/", http.FileServer(http.Dir("./docs/dist/"))))

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
		return
	}

	// connect to database
	db, err := sqlx.Open("pgx", stdlib.RegisterConnConfig(connConfig))
	if err != nil {
		logger.With(zap.Error(err)).Fatal("failed to connect to database")
		return
	}

	// open soundfont
	sf2, err := os.Open("./soundfonts/FluidR3_GM.sf2")
	if err != nil {
		logger.With(zap.Error(err)).Fatal("failed to open soundfont file")
		return
	}

	defer func() { _ = sf2.Close() }()

	// load soundfont
	soundFont, err := meltysynth.NewSoundFont(sf2)
	if err != nil {
		logger.With(zap.Error(err)).Fatal("failed to parse soundfont")
		return
	}

	theoryRepository := theory.NewRepository(logger, db)
	theoryService := theory.NewService(logger, theoryRepository)
	handlerHandler := theory.NewHandler(logger, theoryService, synthFactory{soundFont: soundFont})
	handlerHandler.InstallEndpoints(theoryRouter)

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

type synthFactory struct {
	soundFont *meltysynth.SoundFont
}

func (f synthFactory) Instantiate(sampleRate int32) (midi.Synthesizer, error) {
	midiSynthSettings := meltysynth.NewSynthesizerSettings(sampleRate)
	return meltysynth.NewSynthesizer(f.soundFont, midiSynthSettings)
}
