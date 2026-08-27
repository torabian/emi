package emigo

// Permission is a single node of a module's permission tree, as declared via the
// module's `permissions:` yaml block and compiled by the emi go compiler (see
// lib/golang/go-permissions.go). It lives here, in the shared runtime, instead of
// being redefined at the top of every generated Permissions.go.
//
// Key is the fully qualified, dot separated identifier (e.g. "post.create") - not
// the short per-node slug from yaml. A node that groups children and whose key was
// left for the compiler to derive gets a trailing "*" (e.g. "post*"), signaling it
// covers itself and everything below it; a key set explicitly in yaml is used as-is.
//
// Children is keyed by each child's Name (falling back to its Key when Name isn't
// set), so a permission tree can be walked by name - e.g. Permissions["managePosts"]
// .Children["createPost"] - instead of scanning a slice by position.
type Permission struct {
	Key         string
	Name        string
	Title       map[string]string
	Description map[string]string
	Children    map[string]*Permission
}

func (x Permission) GetTitle() map[string]string {
	return x.Title
}

func (x Permission) GetDescription() map[string]string {
	return x.Description
}

func (x Permission) GetKey() string {
	return x.Key
}

func (x Permission) GetName() string {
	return x.Name
}

func (x Permission) GetChildren() map[string]*Permission {
	return x.Children
}
