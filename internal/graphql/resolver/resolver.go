//go:generate go run ../../../scripts/gqlgen.go
package resolver

import (
	"github.com/dictyBase/graphql-server/internal/graphql/dataloader"
	"github.com/dictyBase/graphql-server/internal/graphql/generated"
	"github.com/dictyBase/graphql-server/internal/graphql/resolver/auth"
	"github.com/dictyBase/graphql-server/internal/graphql/resolver/content"
	"github.com/dictyBase/graphql-server/internal/graphql/resolver/gene"
	"github.com/dictyBase/graphql-server/internal/graphql/resolver/order"
	"github.com/dictyBase/graphql-server/internal/graphql/resolver/organism"
	"github.com/dictyBase/graphql-server/internal/graphql/resolver/publication"
	"github.com/dictyBase/graphql-server/internal/graphql/resolver/stock"
	"github.com/dictyBase/graphql-server/internal/graphql/resolver/user"
	"github.com/dictyBase/graphql-server/internal/registry"
	"github.com/sirupsen/logrus"
)

type Resolver struct {
	registry.Registry
	Dataloaders dataloader.Retriever
	Logger      *logrus.Entry
}

type MutationResolver struct {
	registry.Registry
	Logger *logrus.Entry
}

type QueryResolver struct {
	registry.Registry
	Dataloaders dataloader.Retriever
	Logger      *logrus.Entry
}

func NewResolver(nr registry.Registry, dl dataloader.Retriever, l *logrus.Entry) *Resolver {
	return &Resolver{Registry: nr, Dataloaders: dl, Logger: l}
}

func (r *Resolver) Mutation() generated.MutationResolver {
	return &MutationResolver{
		Registry: r.Registry,
		Logger:   r.Logger,
	}
}
func (r *Resolver) Query() generated.QueryResolver {
	return &QueryResolver{
		Registry:    r.Registry,
		Dataloaders: r.Dataloaders,
		Logger:      r.Logger,
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
func (r *Resolver) Publication() generated.PublicationResolver {
	return &publication.PublicationResolver{
		Logger: r.Logger,
	}
}
func (r *Resolver) Author() generated.AuthorResolver {
	return &publication.AuthorResolver{
		Logger: r.Logger,
	}
}
func (r *Resolver) Strain() generated.StrainResolver {
	return &stock.StrainResolver{
		Client:           r.GetStockClient(registry.STOCK),
		UserClient:       r.GetUserClient(registry.USER),
		AnnotationClient: r.GetAnnotationClient(registry.ANNOTATION),
		Registry:         r.Registry,
		Logger:           r.Logger,
	}
}
func (r *Resolver) Plasmid() generated.PlasmidResolver {
	return &stock.PlasmidResolver{
		Client:           r.GetStockClient(registry.STOCK),
		UserClient:       r.GetUserClient(registry.USER),
		AnnotationClient: r.GetAnnotationClient(registry.ANNOTATION),
		Registry:         r.Registry,
		Logger:           r.Logger,
	}
}

func (r *Resolver) Order() generated.OrderResolver {
	return &order.OrderResolver{
		Client:      r.GetOrderClient(registry.ORDER),
		StockClient: r.GetStockClient(registry.STOCK),
		UserClient:  r.GetUserClient(registry.USER),
		Logger:      r.Logger,
	}
}

func (r *Resolver) Content() generated.ContentResolver {
	return &content.ContentResolver{
		Client:     r.GetContentClient(registry.CONTENT),
		UserClient: r.GetUserClient(registry.USER),
		Logger:     r.Logger,
	}
}

func (r *Resolver) Auth() generated.AuthResolver {
	return &auth.AuthResolver{
		Client:         r.GetAuthClient(registry.AUTH),
		UserClient:     r.GetUserClient(registry.USER),
		IdentityClient: r.GetIdentityClient(registry.IDENTITY),
		Logger:         r.Logger,
	}
}

func (r *Resolver) Gene() generated.GeneResolver {
	return &gene.GeneResolver{
		Redis:  r.GetRedisRepository("redis"),
		Logger: r.Logger,
	}
}

func (r *Resolver) Organism() generated.OrganismResolver {
	return &organism.OrganismResolver{
		Logger: r.Logger,
	}
}
