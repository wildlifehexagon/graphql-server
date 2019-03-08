package resolver

import (
	"github.com/dictyBase/go-genproto/dictybaseapis/user"
	"github.com/dictyBase/graphql-server/internal/graphql/generated"
	"github.com/dictyBase/graphql-server/internal/registry"
	"github.com/sirupsen/logrus"
)

type Resolver struct {
	UserClient       user.UserServiceClient
	RoleClient       user.RoleServiceClient
	PermissionClient user.PermissionServiceClient
	registry.Registry
	Logger *logrus.Entry
}

// func NewResolver(l *logrus.Entry) *Resolver {
// 	return &Resolver{Registry: registry, Logger: l}
// }

func (r *Resolver) Mutation() generated.MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

type queryResolver struct{ *Resolver }
