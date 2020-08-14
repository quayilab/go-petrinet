package gopetrinet

const (
	// ElementNet :
	ElementNet = iota
	// ElementNode :
	ElementNode
	// ElementArc :
	ElementArc
)

const (
	// NodeState :
	NodeState = iota
	// NodeTransition :
	NodeTransition
	// NodeOmni :
	NodeOmni
)

const (
	// ArcNormal :
	ArcNormal = iota
	// ArcInhibit :
	ArcInhibit
)

const (
	EventStateCreate            = `creating state "%s"`
	EventStateAdd               = `adding state "%s"`
	EventStateReadyTest         = `testing state "%s"`
	EventTransitionCreate       = `creating transition "%s"`
	EventTransitionAdd          = `adding transition "%s"`
	EventTransitionReadyTest    = `testing transition "%s"`
	EventTokenCreate            = `creating token "%s"`
	EventTokenAdd               = `adding token "%s" to state "%"`
	EventTokenDataAssign        = `assigning data to token "%s"`
	EventConnectNodes           = `connecting nodes "%s" to "%s"`
	EventDisconnectNodes        = `disconnecting node "%s" from "%s"`
	EventConnectStateTransition = `connecting state "%s" to transition "%s"`
	EventConnectTransitionState = `connecting transition "%s" to state "%s"`

	ResultSucceeded = "succeeded"
	ResultFailed    = "failed"
	ResultValue     = "finished with result: %v"

	ReasonObjectExists       = `object "%s" exists`
	ReasonConnectionExists   = `connection between node "%s" and "%s" exists`
	ReasonConnectionInexist  = `connection between node "%s" and "%s" not exist`
	ReasonStateNotReady      = `state "%s" not ready`
	ReasonTransitionNotReady = `transition "%s" not ready`
	ReasonNetDeadlock        = `net "%s" deadlock`
	ReasonNetNotRunning      = `net "%s" not running`
	ReasonNetRunning         = `net "%s" running`
)

// NodeIndex :
type NodeIndex = int
