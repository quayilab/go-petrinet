package gopetrinet

const (
	// ElementTypeNet :
	ElementTypeNet = iota
	//ElementTypeNodeOmni :
	ElementTypeNodeOmni
	// ElementTypeNodeState :
	ElementTypeNodeState
	// ElementTypeNodeTransition :
	ElementTypeNodeTransition
	// ElementTypeArcNormal :
	ElementTypeArcNormal
	// ElementTypeArcInhibit :
	ElementTypeArcInhibit
	// ElementTypeTokenNormal :
	ElementTypeTokenNormal
)

// NodeIndex :
type NodeIndex = int
