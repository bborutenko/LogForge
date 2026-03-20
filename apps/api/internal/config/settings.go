package config

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

var St *Settings

type Settings struct {
	DBUrl                 string
	MaxConnLifetime       time.Duration
	MaxConnLifetimeJitter time.Duration
	MaxConnIdleTime       time.Duration
	PingTimeout           time.Duration
	HealthCheckPeriod     time.Duration
	MaxConns              int32
	MinConns              int32
	MinIdleConns          int32
}

func getEnvAsDuration(key string, fallback time.Duration) time.Duration {
	if val, ok := os.LookupEnv(key); ok {
		parsed, err := time.ParseDuration(val)
		if err == nil {
			log.Debug().Msgf("Parsed duration for %s: %s", key, parsed)
			return parsed
		}
		log.Warn().Err(err).Msgf("Failed to parse duration for %s, using fallback", key)
	}
	log.Warn().Msgf("Environment variable %s not set, using fallback", key)
	return fallback
}

func getEnvAsInt32(key string, fallback int32) int32 {
	if val, ok := os.LookupEnv(key); ok {
		parsed, err := strconv.ParseInt(val, 10, 32)
		if err == nil {
			log.Debug().Msgf("Parsed int32 for %s: %d", key, parsed)
			return int32(parsed)
		}
		log.Warn().Err(err).Msgf("Failed to parse int32 for %s, using fallback", key)
	}
	log.Warn().Msgf("Environment variable %s not set, using fallback", key)
	return fallback
}

func buildDBUrl(base string, settings *Settings) string {
	u, err := url.Parse(base)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to parse DATABASE_URL")
		panic(err)
	}

	q := u.Query()
	q.Set("pool_max_conns", fmt.Sprintf("%d", settings.MaxConns))
	q.Set("pool_min_conns", fmt.Sprintf("%d", settings.MinConns))
	q.Set("pool_min_idle_conns", fmt.Sprintf("%d", settings.MinIdleConns))
	q.Set("pool_max_conn_lifetime", settings.MaxConnLifetime.String())
	q.Set("pool_max_conn_idle_time", settings.MaxConnIdleTime.String())
	q.Set("pool_health_check_period", settings.HealthCheckPeriod.String())
	q.Set("pool_max_conn_lifetime_jitter", settings.MaxConnLifetimeJitter.String())

	u.RawQuery = q.Encode()
	return u.String()
}

func InitSettings() {
	if err := godotenv.Load(); err != nil {
		log.Info().Msg("No .env file found")
	}

	baseDBUrl := os.Getenv("DATABASE_URL")
	if baseDBUrl == "" {
		log.Fatal().Msg("DATABASE_URL environment variable is required")
		panic("DATABASE_URL environment variable is required")
	}

	settings := &Settings{
		MaxConnLifetime:       getEnvAsDuration("DB_MAX_CONN_LIFETIME", time.Hour),
		MaxConnLifetimeJitter: getEnvAsDuration("DB_MAX_CONN_LIFETIME_JITTER", 0),
		MaxConnIdleTime:       getEnvAsDuration("DB_MAX_CONN_IDLE_TIME", time.Minute*30),
		PingTimeout:           getEnvAsDuration("DB_PING_TIMEOUT", 0),
		HealthCheckPeriod:     getEnvAsDuration("DB_HEALTH_CHECK_PERIOD", time.Minute),
		MaxConns:              getEnvAsInt32("DB_MAX_CONNS", 4), // fallback is usually max(4, runtime.NumCPU())
		MinConns:              getEnvAsInt32("DB_MIN_CONNS", 0),
		MinIdleConns:          getEnvAsInt32("DB_MIN_IDLE_CONNS", 0),
	}

	settings.DBUrl = buildDBUrl(baseDBUrl, settings)
	St = settings
}
