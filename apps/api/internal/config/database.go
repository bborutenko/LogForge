package config

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
)

var DBPool *pgxpool.Pool

func isSchemaExists(pool *pgxpool.Pool) bool {
	rows, _ := pool.Query(context.Background(), "SELECT schema_name FROM information_schema.schemata WHERE schema_name = 'log_forge';")
	schema := rows.Next()
	if schema {
		log.Info().Msg("Schema log_forge already exists")
		return true
	}
	log.Info().Msg("Schema log_forge does not exist")
	return false
}

func createSchema(pool *pgxpool.Pool) {
	_, err := pool.Exec(context.Background(), "CREATE SCHEMA log_forge;")
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create schema log_forge")
		panic(err)
	}
	log.Info().Msg("Schema log_forge created successfully")
}

func createTables(pool *pgxpool.Pool) {
	queryEndpointMetrics := `
	CREATE TABLE IF NOT EXISTS log_forge.endpoint_metrics (
		id BIGSERIAL,
		user_id TEXT,
		endpoint TEXT NOT NULL,
		status_code INT NOT NULL,
		response_time_ms FLOAT NOT NULL,
		meta JSONB,
		PRIMARY KEY (id, endpoint)
	) PARTITION BY LIST (endpoint);`
	_, err := pool.Exec(context.Background(), queryEndpointMetrics)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create table endpoint_metrics")
		panic(err)
	}
	log.Info().Msg("Table endpoint_metrics created successfully (or already exists)")

	queryLogs := `
	CREATE TABLE IF NOT EXISTS log_forge.logs (
		id BIGSERIAL,
		timestamp TIMESTAMPTZ NOT NULL DEFAULT NOW(),
		level TEXT NOT NULL,
		message TEXT NOT NULL,
		meta JSONB,
		endpoint_metrics_id BIGINT,
		endpoint_metrics_name TEXT,
		PRIMARY KEY (id, timestamp),
		CONSTRAINT fk_endpoint_metrics FOREIGN KEY (endpoint_metrics_id, endpoint_metrics_name) REFERENCES log_forge.endpoint_metrics(id, endpoint) ON DELETE SET NULL
	) PARTITION BY RANGE (timestamp);`
	_, err = pool.Exec(context.Background(), queryLogs)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create table logs")
		panic(err)
	}
	log.Info().Msg("Table logs created successfully (or already exists)")

	queryUserActions := `
	CREATE TABLE IF NOT EXISTS log_forge.user_actions (
		id BIGSERIAL,
		timestamp TIMESTAMPTZ NOT NULL DEFAULT NOW(),
		user_id TEXT,
		action TEXT NOT NULL,
		session_id TEXT,
		endpoint_metrics_id BIGINT,
		endpoint_metrics_name TEXT,
		PRIMARY KEY (id, timestamp),
		CONSTRAINT fk_endpoint_metrics_actions FOREIGN KEY (endpoint_metrics_id, endpoint_metrics_name) REFERENCES log_forge.endpoint_metrics(id, endpoint) ON DELETE SET NULL
	) PARTITION BY RANGE (timestamp);`
	_, err = pool.Exec(context.Background(), queryUserActions)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create table user_actions")
		panic(err)
	}
	log.Info().Msg("Table user_actions created successfully (or already exists)")
}

func ConnectDatabase() *pgxpool.Pool {
	log.Info().Msg("Connecting to database")
	config, err := pgxpool.ParseConfig(St.DBUrl)

	if err != nil {
		log.Fatal().Err(err).Msg("Failed to parse database URL")
		panic(err)
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), config)

	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
		panic(err)
	}

	log.Info().Msg("Successfully connected to database")
	return pool
}

func InitDatabase() {
	DBPool = ConnectDatabase()
	if !isSchemaExists(DBPool) {
		log.Info().Msg("Creating schema log_forge")
		createSchema(DBPool)
	} else {
		log.Info().Msg("Schema log_forge already exists, skipping creation")
	}
	createTables(DBPool)
}
