package resolver

import (
	"github.com/dictyBase/graphql-server/internal/graphql/generated"
	"github.com/dictyBase/graphql-server/internal/graphql/resolver/user"
	"github.com/dictyBase/graphql-server/internal/registry"
	"github.com/sirupsen/logrus"
)

type Resolver struct {
	registry.Registry
	Logger *logrus.Entry
}

type MutationResolver struct {
	registry.Registry
	Logger *logrus.Entry
}

type QueryResolver struct {
	registry.Registry
	Logger *logrus.Entry
}

func NewResolver(nr registry.Registry, l *logrus.Entry) *Resolver {
	return &Resolver{Registry: nr, Logger: l}
}

func (r *Resolver) Mutation() generated.MutationResolver {
	return &MutationResolver{
		Registry: r.Registry,
		Logger:   r.Logger,
	}
}
func (r *Resolver) Query() generated.QueryResolver {
	return &QueryResolver{
		Registry: r.Registry,
		Logger:   r.Logger,
	}
}

func (r *Resolver) User() generated.UserResolver {
	return &user.UserResolver{
		Client: r.GetUserClient(registry.USER),
		Logger: r.Logger,
	}
}

func (r *Resolver) Role() generated.RoleResolver {
	return &user.RoleResolver{
		Client: r.GetRoleClient(registry.ROLE),
		Logger: r.Logger,
	}
}

func (r *Resolver) Permission() generated.PermissionResolver {
	return &user.PermissionResolver{
		Client: r.GetPermissionClient(registry.PERMISSION),
		Logger: r.Logger,
	}
}
