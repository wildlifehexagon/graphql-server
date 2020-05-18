package genresolver

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"context"
	"time"

	"github.com/dictyBase/go-genproto/dictybaseapis/auth"
	"github.com/dictyBase/go-genproto/dictybaseapis/content"
	"github.com/dictyBase/go-genproto/dictybaseapis/order"
	"github.com/dictyBase/go-genproto/dictybaseapis/publication"
	"github.com/dictyBase/go-genproto/dictybaseapis/user"
	"github.com/dictyBase/graphql-server/internal/graphql/generated"
	"github.com/dictyBase/graphql-server/internal/graphql/models"
)

type Resolver struct{}

func (r *authResolver) Identity(ctx context.Context, obj *auth.Auth) (*models.Identity, error) {
	panic("not implemented")
}

func (r *authorResolver) Rank(ctx context.Context, obj *publication.Author) (*string, error) {
	panic("not implemented")
}

func (r *contentResolver) ID(ctx context.Context, obj *content.Content) (string, error) {
	panic("not implemented")
}

func (r *contentResolver) Name(ctx context.Context, obj *content.Content) (string, error) {
	panic("not implemented")
}

func (r *contentResolver) Slug(ctx context.Context, obj *content.Content) (string, error) {
	panic("not implemented")
}

func (r *contentResolver) CreatedBy(ctx context.Context, obj *content.Content) (*user.User, error) {
	panic("not implemented")
}

func (r *contentResolver) UpdatedBy(ctx context.Context, obj *content.Content) (*user.User, error) {
	panic("not implemented")
}

func (r *contentResolver) CreatedAt(ctx context.Context, obj *content.Content) (*time.Time, error) {
	panic("not implemented")
}

func (r *contentResolver) UpdatedAt(ctx context.Context, obj *content.Content) (*time.Time, error) {
	panic("not implemented")
}

func (r *contentResolver) Content(ctx context.Context, obj *content.Content) (string, error) {
	panic("not implemented")
}

func (r *contentResolver) Namespace(ctx context.Context, obj *content.Content) (string, error) {
	panic("not implemented")
}

func (r *mutationResolver) Login(ctx context.Context, input *models.LoginInput) (*auth.Auth, error) {
	panic("not implemented")
}

func (r *mutationResolver) Logout(ctx context.Context) (*models.Logout, error) {
	panic("not implemented")
}

func (r *mutationResolver) CreateContent(ctx context.Context, input *models.CreateContentInput) (*content.Content, error) {
	panic("not implemented")
}

func (r *mutationResolver) UpdateContent(ctx context.Context, input *models.UpdateContentInput) (*content.Content, error) {
	panic("not implemented")
}

func (r *mutationResolver) DeleteContent(ctx context.Context, id string) (*models.DeleteContent, error) {
	panic("not implemented")
}

func (r *mutationResolver) CreateOrder(ctx context.Context, input *models.CreateOrderInput) (*order.Order, error) {
	panic("not implemented")
}

func (r *mutationResolver) UpdateOrder(ctx context.Context, id string, input *models.UpdateOrderInput) (*order.Order, error) {
	panic("not implemented")
}

func (r *mutationResolver) CreateStrain(ctx context.Context, input *models.CreateStrainInput) (*models.Strain, error) {
	panic("not implemented")
}

func (r *mutationResolver) CreatePlasmid(ctx context.Context, input *models.CreatePlasmidInput) (*models.Plasmid, error) {
	panic("not implemented")
}

func (r *mutationResolver) UpdateStrain(ctx context.Context, id string, input *models.UpdateStrainInput) (*models.Strain, error) {
	panic("not implemented")
}

func (r *mutationResolver) UpdatePlasmid(ctx context.Context, id string, input *models.UpdatePlasmidInput) (*models.Plasmid, error) {
	panic("not implemented")
}

func (r *mutationResolver) DeleteStock(ctx context.Context, id string) (*models.DeleteStock, error) {
	panic("not implemented")
}

func (r *mutationResolver) CreateUser(ctx context.Context, input *models.CreateUserInput) (*user.User, error) {
	panic("not implemented")
}

func (r *mutationResolver) CreateUserRoleRelationship(ctx context.Context, userID string, roleID string) (*user.User, error) {
	panic("not implemented")
}

func (r *mutationResolver) UpdateUser(ctx context.Context, id string, input *models.UpdateUserInput) (*user.User, error) {
	panic("not implemented")
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (*models.DeleteUser, error) {
	panic("not implemented")
}

func (r *mutationResolver) CreateRole(ctx context.Context, input *models.CreateRoleInput) (*user.Role, error) {
	panic("not implemented")
}

func (r *mutationResolver) CreateRolePermissionRelationship(ctx context.Context, roleID string, permissionID string) (*user.Role, error) {
	panic("not implemented")
}

func (r *mutationResolver) UpdateRole(ctx context.Context, id string, input *models.UpdateRoleInput) (*user.Role, error) {
	panic("not implemented")
}

func (r *mutationResolver) DeleteRole(ctx context.Context, id string) (*models.DeleteRole, error) {
	panic("not implemented")
}

func (r *mutationResolver) CreatePermission(ctx context.Context, input *models.CreatePermissionInput) (*user.Permission, error) {
	panic("not implemented")
}

func (r *mutationResolver) UpdatePermission(ctx context.Context, id string, input *models.UpdatePermissionInput) (*user.Permission, error) {
	panic("not implemented")
}

func (r *mutationResolver) DeletePermission(ctx context.Context, id string) (*models.DeletePermission, error) {
	panic("not implemented")
}

func (r *orderResolver) ID(ctx context.Context, obj *order.Order) (string, error) {
	panic("not implemented")
}

func (r *orderResolver) CreatedAt(ctx context.Context, obj *order.Order) (*time.Time, error) {
	panic("not implemented")
}

func (r *orderResolver) UpdatedAt(ctx context.Context, obj *order.Order) (*time.Time, error) {
	panic("not implemented")
}

func (r *orderResolver) Courier(ctx context.Context, obj *order.Order) (*string, error) {
	panic("not implemented")
}

func (r *orderResolver) CourierAccount(ctx context.Context, obj *order.Order) (*string, error) {
	panic("not implemented")
}

func (r *orderResolver) Comments(ctx context.Context, obj *order.Order) (*string, error) {
	panic("not implemented")
}

func (r *orderResolver) Payment(ctx context.Context, obj *order.Order) (*string, error) {
	panic("not implemented")
}

func (r *orderResolver) PurchaseOrderNum(ctx context.Context, obj *order.Order) (*string, error) {
	panic("not implemented")
}

func (r *orderResolver) Status(ctx context.Context, obj *order.Order) (*models.StatusEnum, error) {
	panic("not implemented")
}

func (r *orderResolver) Consumer(ctx context.Context, obj *order.Order) (*user.User, error) {
	panic("not implemented")
}

func (r *orderResolver) Payer(ctx context.Context, obj *order.Order) (*user.User, error) {
	panic("not implemented")
}

func (r *orderResolver) Purchaser(ctx context.Context, obj *order.Order) (*user.User, error) {
	panic("not implemented")
}

func (r *orderResolver) Items(ctx context.Context, obj *order.Order) ([]models.Stock, error) {
	panic("not implemented")
}

func (r *permissionResolver) ID(ctx context.Context, obj *user.Permission) (string, error) {
	panic("not implemented")
}

func (r *permissionResolver) Permission(ctx context.Context, obj *user.Permission) (string, error) {
	panic("not implemented")
}

func (r *permissionResolver) Description(ctx context.Context, obj *user.Permission) (string, error) {
	panic("not implemented")
}

func (r *permissionResolver) CreatedAt(ctx context.Context, obj *user.Permission) (*time.Time, error) {
	panic("not implemented")
}

func (r *permissionResolver) UpdatedAt(ctx context.Context, obj *user.Permission) (*time.Time, error) {
	panic("not implemented")
}

func (r *permissionResolver) Resource(ctx context.Context, obj *user.Permission) (*string, error) {
	panic("not implemented")
}

func (r *plasmidResolver) ID(ctx context.Context, obj *models.Plasmid) (string, error) {
	panic("not implemented")
}

func (r *plasmidResolver) CreatedAt(ctx context.Context, obj *models.Plasmid) (*time.Time, error) {
	panic("not implemented")
}

func (r *plasmidResolver) UpdatedAt(ctx context.Context, obj *models.Plasmid) (*time.Time, error) {
	panic("not implemented")
}

func (r *plasmidResolver) CreatedBy(ctx context.Context, obj *models.Plasmid) (*user.User, error) {
	panic("not implemented")
}

func (r *plasmidResolver) UpdatedBy(ctx context.Context, obj *models.Plasmid) (*user.User, error) {
	panic("not implemented")
}

func (r *plasmidResolver) Summary(ctx context.Context, obj *models.Plasmid) (*string, error) {
	panic("not implemented")
}

func (r *plasmidResolver) EditableSummary(ctx context.Context, obj *models.Plasmid) (*string, error) {
	panic("not implemented")
}

func (r *plasmidResolver) Depositor(ctx context.Context, obj *models.Plasmid) (string, error) {
	panic("not implemented")
}

func (r *plasmidResolver) Genes(ctx context.Context, obj *models.Plasmid) ([]*string, error) {
	panic("not implemented")
}

func (r *plasmidResolver) Dbxrefs(ctx context.Context, obj *models.Plasmid) ([]*string, error) {
	panic("not implemented")
}

func (r *plasmidResolver) Publications(ctx context.Context, obj *models.Plasmid) ([]*publication.Publication, error) {
	panic("not implemented")
}

func (r *plasmidResolver) Name(ctx context.Context, obj *models.Plasmid) (string, error) {
	panic("not implemented")
}

func (r *plasmidResolver) ImageMap(ctx context.Context, obj *models.Plasmid) (*string, error) {
	panic("not implemented")
}

func (r *plasmidResolver) Sequence(ctx context.Context, obj *models.Plasmid) (*string, error) {
	panic("not implemented")
}

func (r *plasmidResolver) InStock(ctx context.Context, obj *models.Plasmid) (bool, error) {
	panic("not implemented")
}

func (r *plasmidResolver) Keywords(ctx context.Context, obj *models.Plasmid) ([]*string, error) {
	panic("not implemented")
}

func (r *plasmidResolver) GenbankAccession(ctx context.Context, obj *models.Plasmid) (*string, error) {
	panic("not implemented")
}

func (r *publicationResolver) ID(ctx context.Context, obj *publication.Publication) (string, error) {
	panic("not implemented")
}

func (r *publicationResolver) Doi(ctx context.Context, obj *publication.Publication) (*string, error) {
	panic("not implemented")
}

func (r *publicationResolver) Title(ctx context.Context, obj *publication.Publication) (*string, error) {
	panic("not implemented")
}

func (r *publicationResolver) Abstract(ctx context.Context, obj *publication.Publication) (*string, error) {
	panic("not implemented")
}

func (r *publicationResolver) Journal(ctx context.Context, obj *publication.Publication) (*string, error) {
	panic("not implemented")
}

func (r *publicationResolver) PubDate(ctx context.Context, obj *publication.Publication) (*time.Time, error) {
	panic("not implemented")
}

func (r *publicationResolver) Volume(ctx context.Context, obj *publication.Publication) (*string, error) {
	panic("not implemented")
}

func (r *publicationResolver) Pages(ctx context.Context, obj *publication.Publication) (*string, error) {
	panic("not implemented")
}

func (r *publicationResolver) Issn(ctx context.Context, obj *publication.Publication) (*string, error) {
	panic("not implemented")
}

func (r *publicationResolver) PubType(ctx context.Context, obj *publication.Publication) (*string, error) {
	panic("not implemented")
}

func (r *publicationResolver) Source(ctx context.Context, obj *publication.Publication) (*string, error) {
	panic("not implemented")
}

func (r *publicationResolver) Issue(ctx context.Context, obj *publication.Publication) (*string, error) {
	panic("not implemented")
}

func (r *publicationResolver) Status(ctx context.Context, obj *publication.Publication) (*string, error) {
	panic("not implemented")
}

func (r *publicationResolver) Authors(ctx context.Context, obj *publication.Publication) ([]*publication.Author, error) {
	panic("not implemented")
}

func (r *queryResolver) GetRefreshToken(ctx context.Context, token string) (*auth.Auth, error) {
	panic("not implemented")
}

func (r *queryResolver) Content(ctx context.Context, id string) (*content.Content, error) {
	panic("not implemented")
}

func (r *queryResolver) ContentBySlug(ctx context.Context, slug string) (*content.Content, error) {
	panic("not implemented")
}

func (r *queryResolver) Order(ctx context.Context, id string) (*order.Order, error) {
	panic("not implemented")
}

func (r *queryResolver) ListOrders(ctx context.Context, input *models.ListOrderInput) (*models.OrderListWithCursor, error) {
	panic("not implemented")
}

func (r *queryResolver) Publication(ctx context.Context, id string) (*publication.Publication, error) {
	panic("not implemented")
}

func (r *queryResolver) Plasmid(ctx context.Context, id string) (*models.Plasmid, error) {
	panic("not implemented")
}

func (r *queryResolver) Strain(ctx context.Context, id string) (*models.Strain, error) {
	panic("not implemented")
}

func (r *queryResolver) ListStrains(ctx context.Context, input *models.ListStockInput) (*models.StrainListWithCursor, error) {
	panic("not implemented")
}

func (r *queryResolver) ListPlasmids(ctx context.Context, input *models.ListStockInput) (*models.PlasmidListWithCursor, error) {
	panic("not implemented")
}

func (r *queryResolver) User(ctx context.Context, id string) (*user.User, error) {
	panic("not implemented")
}

func (r *queryResolver) UserByEmail(ctx context.Context, email string) (*user.User, error) {
	panic("not implemented")
}

func (r *queryResolver) ListUsers(ctx context.Context, pagenum string, pagesize string, filter string) (*models.UserList, error) {
	panic("not implemented")
}

func (r *queryResolver) Role(ctx context.Context, id string) (*user.Role, error) {
	panic("not implemented")
}

func (r *queryResolver) ListRoles(ctx context.Context) ([]*user.Role, error) {
	panic("not implemented")
}

func (r *queryResolver) Permission(ctx context.Context, id string) (*user.Permission, error) {
	panic("not implemented")
}

func (r *queryResolver) ListPermissions(ctx context.Context) ([]*user.Permission, error) {
	panic("not implemented")
}

func (r *roleResolver) ID(ctx context.Context, obj *user.Role) (string, error) {
	panic("not implemented")
}

func (r *roleResolver) Role(ctx context.Context, obj *user.Role) (string, error) {
	panic("not implemented")
}

func (r *roleResolver) Description(ctx context.Context, obj *user.Role) (string, error) {
	panic("not implemented")
}

func (r *roleResolver) CreatedAt(ctx context.Context, obj *user.Role) (*time.Time, error) {
	panic("not implemented")
}

func (r *roleResolver) UpdatedAt(ctx context.Context, obj *user.Role) (*time.Time, error) {
	panic("not implemented")
}

func (r *roleResolver) Permissions(ctx context.Context, obj *user.Role) ([]*user.Permission, error) {
	panic("not implemented")
}

func (r *strainResolver) CreatedBy(ctx context.Context, obj *models.Strain) (*user.User, error) {
	panic("not implemented")
}

func (r *strainResolver) UpdatedBy(ctx context.Context, obj *models.Strain) (*user.User, error) {
	panic("not implemented")
}

func (r *strainResolver) Publications(ctx context.Context, obj *models.Strain) ([]*publication.Publication, error) {
	panic("not implemented")
}

func (r *strainResolver) SystematicName(ctx context.Context, obj *models.Strain) (string, error) {
	panic("not implemented")
}

func (r *strainResolver) Parent(ctx context.Context, obj *models.Strain) (*models.Strain, error) {
	panic("not implemented")
}

func (r *strainResolver) Names(ctx context.Context, obj *models.Strain) ([]*string, error) {
	panic("not implemented")
}

func (r *strainResolver) InStock(ctx context.Context, obj *models.Strain) (bool, error) {
	panic("not implemented")
}

func (r *strainResolver) Phenotypes(ctx context.Context, obj *models.Strain) ([]*models.Phenotype, error) {
	panic("not implemented")
}

func (r *strainResolver) GeneticModification(ctx context.Context, obj *models.Strain) (*string, error) {
	panic("not implemented")
}

func (r *strainResolver) MutagenesisMethod(ctx context.Context, obj *models.Strain) (*string, error) {
	panic("not implemented")
}

func (r *strainResolver) Characteristics(ctx context.Context, obj *models.Strain) ([]*string, error) {
	panic("not implemented")
}

func (r *strainResolver) Genotypes(ctx context.Context, obj *models.Strain) ([]*string, error) {
	panic("not implemented")
}

func (r *userResolver) ID(ctx context.Context, obj *user.User) (string, error) {
	panic("not implemented")
}

func (r *userResolver) FirstName(ctx context.Context, obj *user.User) (string, error) {
	panic("not implemented")
}

func (r *userResolver) LastName(ctx context.Context, obj *user.User) (string, error) {
	panic("not implemented")
}

func (r *userResolver) Email(ctx context.Context, obj *user.User) (string, error) {
	panic("not implemented")
}

func (r *userResolver) Organization(ctx context.Context, obj *user.User) (*string, error) {
	panic("not implemented")
}

func (r *userResolver) GroupName(ctx context.Context, obj *user.User) (*string, error) {
	panic("not implemented")
}

func (r *userResolver) FirstAddress(ctx context.Context, obj *user.User) (*string, error) {
	panic("not implemented")
}

func (r *userResolver) SecondAddress(ctx context.Context, obj *user.User) (*string, error) {
	panic("not implemented")
}

func (r *userResolver) City(ctx context.Context, obj *user.User) (*string, error) {
	panic("not implemented")
}

func (r *userResolver) State(ctx context.Context, obj *user.User) (*string, error) {
	panic("not implemented")
}

func (r *userResolver) Zipcode(ctx context.Context, obj *user.User) (*string, error) {
	panic("not implemented")
}

func (r *userResolver) Country(ctx context.Context, obj *user.User) (*string, error) {
	panic("not implemented")
}

func (r *userResolver) Phone(ctx context.Context, obj *user.User) (*string, error) {
	panic("not implemented")
}

func (r *userResolver) IsActive(ctx context.Context, obj *user.User) (bool, error) {
	panic("not implemented")
}

func (r *userResolver) CreatedAt(ctx context.Context, obj *user.User) (*time.Time, error) {
	panic("not implemented")
}

func (r *userResolver) UpdatedAt(ctx context.Context, obj *user.User) (*time.Time, error) {
	panic("not implemented")
}

func (r *userResolver) Roles(ctx context.Context, obj *user.User) ([]*user.Role, error) {
	panic("not implemented")
}

// Auth returns generated.AuthResolver implementation.
func (r *Resolver) Auth() generated.AuthResolver { return &authResolver{r} }

// Author returns generated.AuthorResolver implementation.
func (r *Resolver) Author() generated.AuthorResolver { return &authorResolver{r} }

// Content returns generated.ContentResolver implementation.
func (r *Resolver) Content() generated.ContentResolver { return &contentResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Order returns generated.OrderResolver implementation.
func (r *Resolver) Order() generated.OrderResolver { return &orderResolver{r} }

// Permission returns generated.PermissionResolver implementation.
func (r *Resolver) Permission() generated.PermissionResolver { return &permissionResolver{r} }

// Plasmid returns generated.PlasmidResolver implementation.
func (r *Resolver) Plasmid() generated.PlasmidResolver { return &plasmidResolver{r} }

// Publication returns generated.PublicationResolver implementation.
func (r *Resolver) Publication() generated.PublicationResolver { return &publicationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Role returns generated.RoleResolver implementation.
func (r *Resolver) Role() generated.RoleResolver { return &roleResolver{r} }

// Strain returns generated.StrainResolver implementation.
func (r *Resolver) Strain() generated.StrainResolver { return &strainResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type authResolver struct{ *Resolver }
type authorResolver struct{ *Resolver }
type contentResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type orderResolver struct{ *Resolver }
type permissionResolver struct{ *Resolver }
type plasmidResolver struct{ *Resolver }
type publicationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type roleResolver struct{ *Resolver }
type strainResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
