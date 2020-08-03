package gopetrinet

const (
	ElementTypeNet = iota
	ElementTypeNodeOmni
	ElementTypeNodeState
	ElementTypeNodeTransition
	ElementTypeArcNormal
	ElementTypeArcInhibit
	ElementTypeTokenNormal
)

type NodeIndex = int
