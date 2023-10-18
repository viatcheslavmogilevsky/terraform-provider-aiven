package main

import (
	"context"
	"flag"
	"log"

	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6/tf6server"

	"github.com/aiven/terraform-provider-aiven/internal/server"
)

//go:generate go test -tags userconfig ./internal/schemautil/userconfig
//go:generate go run ./ucgenerator/... --integrations clickhouse_kafka,clickhouse_postgresql,datadog,external_aws_cloudwatch_metrics,kafka_connect,kafka_logs,kafka_mirrormaker,logs,metrics
//go:generate go run ./ucgenerator/... --services pg,cassandra,elasticsearch,opensearch,grafana,influxdb,redis,mysql,kafka,kafka_connect,kafka_mirrormaker,m3db,m3aggregator,flink,clickhouse

// version is the version of the provider.
var version = "dev"

func main() {
	debugFlag := flag.Bool("debug", false, "Start provider in debug mode.")

	flag.Parse()
	ctx := context.Background()
	muxServer, err := server.NewMuxServer(ctx, version)
	if err != nil {
		log.Fatal(err)
	}

	var serveOpts []tf6server.ServeOpt
	if *debugFlag {
		serveOpts = append(serveOpts, tf6server.WithManagedDebug())
	}

	err = tf6server.Serve(
		"registry.terraform.io/aiven/aiven",
		func() tfprotov6.ProviderServer {
			return muxServer
		},
		serveOpts...,
	)

	if err != nil {
		log.Fatal(err)
	}
}
