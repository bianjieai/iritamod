package types

const (
	QueryNode  = "node"  // query node
	QueryNodes = "nodes" // query nodes
)

// QueryNodeParams defines the params to query a node
type QueryNodeParams struct {
	ID string
}

// QueryNodesParams defines the params to query nodes
type QueryNodesParams struct {
	Page  int
	Limit int
}
