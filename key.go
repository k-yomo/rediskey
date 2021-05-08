package rediskey

import (
	"bytes"
)

type Marshaller interface {
	marshal(b *bytes.Buffer)
}

// Namespace represents the objectType of redis key.
type Namespace struct {
	Name   string
	Parent Marshaller
}

func NewNamespace(name string, parent Marshaller) *Namespace {
	return &Namespace{
		Name:   name,
		Parent: parent,
	}
}

func (n *Namespace) marshal(b *bytes.Buffer) {
	if n.Parent != nil {
		n.Parent.marshal(b)
		b.WriteByte(':')
	}
	b.WriteString(n.Name)
}

// Key represents the redis key.
type Key struct {
	// ObjectType of the key
	ObjectType string
	// ID of the value
	ID string
	// Parent can be nil.
	Parent Marshaller
}

func NewKey(objectType string, id string, parent Marshaller) *Key {
	return &Key{
		ObjectType: objectType,
		ID:         id,
		Parent:     parent,
	}
}

// String returns a string representation of the key.
func (k *Key) String() string {
	if k == nil {
		return ""
	}
	b := bytes.NewBuffer(make([]byte, 0, 512))
	k.marshal(b)
	return b.String()
}

func (k *Key) marshal(b *bytes.Buffer) {
	if k.Parent != nil {
		k.Parent.marshal(b)
		b.WriteByte(':')
	}
	if k.ObjectType != "" {
		b.WriteString(k.ObjectType + ":")
	}
	b.WriteString(k.ID)
}
