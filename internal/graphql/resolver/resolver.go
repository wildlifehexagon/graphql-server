package resolver

import (
	"github.com/dictyBase/graphql-server/internal/graphql/generated"
	"github.com/dictyBase/graphql-server/internal/registry"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	pb "github.com/dictyBase/go-genproto/dictybaseapis/user"
)

type Resolver struct {
	registry.Registry
	Logger *logrus.Entry
}

func NewResolver(nr registry.Registry, l *logrus.Entry) *Resolver {
	return &Resolver{Registry: nr, Logger: l}
}

func (r *Resolver) Mutation() generated.MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

type queryResolver struct{ *Resolver }

func (r *Resolver) GetPermissionClient() (permissionResolver, error) {
	client, ok := r.GetAPIClient("permission")
	if !ok {
		panic("could not get permission client")
	}
	_, err := grpc.Dial(
		client.(string),
		grpc.WithInsecure(),
	)
	if err != nil {
		r.Logger.Fatalf("cannot connect to grpc server for permission microservice")
		return nil, err
	}
	uc, _ := client.(pb.PermissionServiceClient)
	return permissionResolver{
		Client: uc,
		Logger: r.Logger,
	}, nil
}
