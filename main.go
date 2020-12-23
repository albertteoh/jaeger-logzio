package main

import (
	"flag"
	"github.com/jaegertracing/jaeger/plugin/storage/grpc/shared"
	"github.com/jaegertracing/jaeger/storage/spanstore"
	"github.com/logzio/jaeger-logzio/store"
	"os"

	"github.com/hashicorp/go-hclog"
	"github.com/jaegertracing/jaeger/plugin/storage/grpc"
)

const (
	loggerName = "jaeger-logzio"
)

func main() {
	logger := hclog.New(&hclog.LoggerOptions{
		Level:      hclog.Debug,
		Name:       loggerName,
		JSONFormat: true,
	})
	logger.Info("Initializing logz.io storage")
	var configPath string
	flag.StringVar(&configPath, "config", "", "The absolute path to the logz.io plugin's configuration file")
	flag.Parse()

	logzioConfig, err := store.ParseConfig(configPath, logger)
	if err != nil {
		logger.Error("can't parse config: ", err.Error())
		os.Exit(0)
	}

	logger.Info(logzioConfig.String())
	logzioStore := store.NewLogzioStore(*logzioConfig, logger)
	ps := &shared.PluginServices{Store: logzioStore, ArchiveStore: &archiveStore{}}
	grpc.Serve(ps)
	logzioStore.Close()
}

type archiveStore struct {
}

func (as *archiveStore) ArchiveSpanReader() spanstore.Reader {
	return nil
}


func (as *archiveStore) ArchiveSpanWriter() spanstore.Writer {
	return nil
}