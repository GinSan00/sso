package suite

import (
	"context"
	"sso-service-grpc/internal/config"
	"testing"

	ssov1 "github.com/GolangLessons/protos/gen/go/sso"
	"google.golang.org/grpc"
)

type Suite struct {
	*testing.T
	Cfg        *config.Config
	AuthClient ssov1.AuthClient
}

func New(t *testing.T) (context.Context, *Suite) {
	t.Helper()
	t.Parallel()

	cfg := config.MustLoadByPath("../../configs/local_tests.yaml")

	ctx, cancelCtx := context.WithTimeout(context.Background(), cfg.GRPC.Timeout)

	t.Cleanup(func() {
		t.Helper()
		cancelCtx()
	})

	cc, err := grpc.DialContext(context.Background()),
	
}
