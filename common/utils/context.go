package utils

//nolint:gochecknoglobals // context key's constants
var (
	TokenKey     = &ContextKey{"Token"}
	UserIDKey    = &ContextKey{"UserID"}
	UserRoleKey  = &ContextKey{"role"}
	EmailKey     = &ContextKey{"Email"}
)

type ContextKey struct {
	name string
}

func (k *ContextKey) String() string {
	return k.name
}
