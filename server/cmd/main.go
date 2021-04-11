package main

import (
	"context"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"

	grpcAPI "github.com/lrotermund/port-challenge/server/pkg/api/grpc"
	protos "github.com/lrotermund/port-challenge/server/pkg/api/portservice/v1"
)

const (
	address = ":9090"
)

func main() {
	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	nl, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal().Err(err).Msgf("failed to listen: %s", address)
	}

	gs := grpc.NewServer()
	pds := grpcAPI.NewPortDomainService()

	protos.RegisterPortDomainServer(gs, pds)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		s := <-sigCh
		log.Info().Msgf("got signal %+v, attempting graceful shutdown", s)
		cancel()
		gs.GracefulStop()
		wg.Done()
	}()

	log.Info().Msg("starting grpc PortDomainService")

	err = gs.Serve(nl)
	if err != nil {
		log.Fatal().Err(err).Msgf("could not serve grpc")
	}

	wg.Wait()

	log.Info().Msg("clean shutdown")
}
