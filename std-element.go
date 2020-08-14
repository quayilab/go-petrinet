package gopetrinet

// Element struct ...
type Element struct {
	id      string
	label   string
	desc    string
	typ     int
	net     INet
	enabled bool
}

// ID :
func (e *Element) ID(values ...string) (result string) {
	if len(values) > 0 {
		e.id = values[0]
	}
	result = e.id
	return
}

// Label :
func (e *Element) Label(values ...string) (result string) {
	if len(values) > 0 {
		e.label = values[0]
	}
	result = e.label
	return
}

// Desc :
func (e *Element) Desc(values ...string) (result string) {
	if len(values) > 0 {
		e.desc = values[0]
	}
	result = e.desc
	return
}

// Type :
func (e *Element) Type(values ...int) (result int) {
	if len(values) > 0 {
		e.typ = values[0]
	}
	result = e.typ
	return
}

// Net :
func (e *Element) Net(values ...INet) (result INet) {
	if len(values) > 0 {
		e.net = values[0]
	}
	result = e.net
	return
}

// Enabled :
func (e *Element) Enabled(values ...bool) (result bool) {
	if len(values) > 0 {
		e.enabled = values[0]
	}
	result = e.enabled
	return
}

// IdenticWith :
func (e *Element) IdenticWith(e1 IElement) (result bool, reason string) {
	result = false
	if e1.Type() != e.typ {
		reason = "type not equal"
	} else if e1.ID() != e.id {
		reason = "id not equal"
	} else if e1.Label() != e.label {
		reason = "label not equal"
	} else if e1.Desc() != e.desc {
		reason = "desc not equal"
	} else if e1.Enabled() != e.enabled {
		reason = "status not equal"
	} else {
		result = true
	}
	return
}

// NewElement function
func NewElement(net INet, id, label, desc string, typ int) (result IElement) {
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
