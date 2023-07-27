package domain

import "github.com/c12s/magnetar/pkg/magnetar"

type Node struct {
	Id     NodeId
	Labels []magnetar.Label
}

type NodeId struct {
	Value string
}

// todo: support operators other than == (<, >, !=)
type QuerySelector []magnetar.Label

type NodeRepo interface {
	Put(node Node) error
	Get(nodeId NodeId) (*Node, error)
	Query(selector QuerySelector) ([]Node, error)
}