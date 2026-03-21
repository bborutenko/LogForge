package config

import (
	"context"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
)

var DBPool *pgxpool.Pool

func checkIfSchemaExists(pool *pgxpool.Pool, schemaName string) bool {
	var exists bool
	res, err := pool.Exec(context.Background(),
		`SELECT schema_name FROM information_schema.schemata WHERE schema_name = $1;
`, schemaName)

	if err != nil {
		log.Fatal().Err(err).Msgf("Failed to check if schema %s exists", schemaName)
		panic(err)
	}

	if res.RowsAffected() > 0 {
		exists = true
	}

	return exists
}

func createSchema(pool *pgxpool.Pool, schemaName string) {
	_, err := pool.Exec(context.Background(), "CREATE SCHEMA IF NOT EXISTS "+schemaName+";")
	if err != nil {
		log.Fatal().Err(err).Msgf("Failed to create schema %s", schemaName)
		panic(err)
	}
	log.Info().Msgf("Schema %s created successfully", schemaName)
}

func createTable(
	pool *pgxpool.Pool,
	tableName string,
	rowNames []string,
	rowValues []string,
	partitionBy string,
) {
	if len(rowNames) != len(rowValues) {
		log.Fatal().Msgf("Row names and values length mismatch for table %s", tableName)
		panic("Row names and values length mismatch")
	}

	rows := make([]string, len(rowNames))
	for i := range rowNames {
		row := rowNames[i] + " " + rowValues[i]
		log.Debug().Msgf("Constructed row definition: %s", row)
		rows[i] = row
	}

	query := `
	CREATE TABLE IF NOT EXISTS log_forge.` + tableName + ` (` + strings.Join(rows, ", ") + `) ` + partitionBy + `;`
	_, err := pool.Exec(context.Background(), query)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create table endpoint_metrics")
		panic(err)
	}
	log.Info().Msg("Table endpoint_metrics created successfully (or already exists)")
}

func installExtension(pool *pgxpool.Pool, extension string, schema string) {
	_, err := pool.Exec(context.Background(), "CREATE EXTENSION IF NOT EXISTS "+extension+" WITH SCHEMA "+schema+";")
	if err != nil {
		log.Fatal().Err(err).Msgf("Failed to install extension %s", extension)
		panic(err)
	}
	log.Info().Msgf("Extension %s installed successfully (or already exists)", extension)
}

func createPartition(
	pool *pgxpool.Pool,
	p_parent_table string,
	p_control string,
	p_interval string) {

	fullTableName := "log_forge." + p_parent_table

	_, err := pool.Exec(context.Background(),
		`SELECT log_forge.create_partition(
			p_parent_table := $1::text,
			p_control := $2::text,
			p_interval := $3::text
		);`,
		fullTableName, p_control, p_interval,
	)
	if err != nil {
		log.Fatal().Err(err).Msgf("Failed to create partitions for %s table", p_parent_table)
		panic(err)
	}
	log.Info().Msgf("Partitions created successfully for %s table (or already exist)", p_parent_table)
}

func grantUsage(pool *pgxpool.Pool, schema string, user string) {
	_, err := pool.Exec(context.Background(), "GRANT USAGE ON SCHEMA "+schema+" TO "+user+";")
	if err != nil {
		log.Fatal().Err(err).Msgf("Failed to grant usage on schema %s to user %s", schema, user)
		panic(err)
	}
	log.Info().Msgf("Granted usage on schema %s to user %s", schema, user)
}

func scheduleCron(pool *pgxpool.Pool, jobName string, schedule string, command string) {
	_, err := pool.Exec(context.Background(),
		`SELECT cron.schedule($1, $2, $3);`,
		jobName, schedule, command,
	)
	if err != nil {
		log.Fatal().Err(err).Msgf("Failed to schedule cron job %s", jobName)
		panic(err)
	}
	log.Info().Msgf("Scheduled cron job %s successfully", jobName)
}

func ConnectDatabase() (*pgxpool.Pool, *pgxpool.Config) {
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
	return pool, config
}

func InitDatabase() {
	dbPool, config := ConnectDatabase()

	if !checkIfSchemaExists(dbPool, "log_forge") {
		createSchema(dbPool, "log_forge")

		installExtension(dbPool, "pg_partman", "log_forge")
		installExtension(dbPool, "pg_cron", "pg_catalog")

		createTable(dbPool, "endpoint_metrics", endpointMetricsRowNames, endpointMetricsRowValues, "")
		createTable(dbPool, "logs", logsRowNames, logsRowValues, "PARTITION BY RANGE (timestamp)")
		createTable(dbPool, "user_actions", userActionsRowNames, userActionsRowValues, "PARTITION BY RANGE (timestamp)")

		// createPartition(dbPool, "endpoint_metrics", "endpoint", "day")
		createPartition(dbPool, "logs", "timestamp", "1 day")
		createPartition(dbPool, "user_actions", "timestamp", "1 day")

		grantUsage(dbPool, "pg_catalog", config.ConnConfig.User)

		// scheduleCron(dbPool, "maintain-endpoint-metrics-partitions", "*/30 * * * *", `SELECT public.run_maintenance(p_parent_table := ''log_forge.endpoint_metrics'');`)
		scheduleCron(dbPool, "maintain-logs-partitions", "*/30 * * * *", `SELECT public.run_maintenance(p_parent_table := ''log_forge.logs'');`)
		scheduleCron(dbPool, "maintain-user-actions-partitions", "*/30 * * * *", `SELECT public.run_maintenance(p_parent_table := ''log_forge.user_actions'');`)
	}
	DBPool = dbPool
}
