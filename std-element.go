package gopetrinet

// Element struct ...
type Element struct {
	id      string
	label   string
	desc    string
	typ     int
	net     NetIntf
	enabled bool
}

// GetID :
func (e *Element) GetID() (result string) {
	result = e.id
	return
}

// GetLabel :
func (e *Element) GetLabel() (result string) {
	result = e.label
	return
}

// GetDesc :
func (e *Element) GetDesc() (result string) {
	result = e.desc
	return
}

// GetType :
func (e *Element) GetType() (result int) {
	result = e.typ
	return
}

// GetNet :
func (e *Element) GetNet() (result NetIntf) {
	result = e.net
	return
}

// GetStatus :
func (e *Element) GetStatus() (result bool) {
	result = e.enabled
	return
}

// Enable :
func (e *Element) Enable() {
	e.enabled = true
}

// Disable :
func (e *Element) Disable() {
	e.enabled = false
}

// IsIdentic :
func (e *Element) IsIdentic(e1 ElementIntf) (result bool, reason string) {
	result = false
	if e1.GetType() != e.typ {
		reason = "type not equal"
	} else if e1.GetID() != e.id {
		reason = "id not equal"
	} else if e1.GetLabel() != e.label {
		reason = "label not equal"
	} else if e1.GetDesc() != e.desc {
		reason = "desc not equal"
	} else if e1.GetStatus() != e.enabled {
		reason = "status not equal"
	} else {
		result = true
	}
	return
}

// NewElement function
func NewElement(net NetIntf, id, label, desc string, typ int) (result ElementIntf) {
	result = &Element{
		id:      id,
		label:   label,
		desc:    desc,
		typ:     typ,
		net:     net,
		enabled: true,
	}
	return
}
