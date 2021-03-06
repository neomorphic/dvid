package dvid

import "fmt"

// StoreCloser stores can be closed.
type StoreCloser interface {
	Close()
}

// StoreIdentifiable stores can say whether they are identified by a given store configuration.
type StoreIdentifiable interface {
	// Equal returns true if this store matches the given store configuration.
	Equal(StoreConfig) bool
}

// Store allows polyglot persistence of data.  The Store implementation
// could be an ordered key-value database, graph store, etc.
type Store interface {
	fmt.Stringer
	StoreCloser
	StoreIdentifiable
}

// StoreConfig is a store-specific configuration where each store implementation
// defines the types of parameters it accepts.
type StoreConfig struct {
	Config

	// Engine is a simple name describing the engine, e.g., "basholeveldb"
	Engine string
}

// DataSpecifier is either a TypeString or a (InstanceName, UUID) tuple, allowing
// an immutable specification of a datatype or data instance.  The DataSpecifier
// can be used as the key of a map, and is useful when configuring which data is
// stored using different storage engines.
type DataSpecifier string

// GetDataSpecifier returns an DataSpecifier given an instance name and UUID.
func GetDataSpecifier(name InstanceName, uuid UUID) DataSpecifier {
	return DataSpecifier(name) + DataSpecifier(uuid)
}
