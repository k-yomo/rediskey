package rediskey

import (
	"reflect"
	"testing"
)

func TestNewNamespace(t *testing.T) {
	type args struct {
		name   string
		parent Marshaller
	}
	tests := []struct {
		name string
		args args
		want *Namespace
	}{
		{
			name: "initializes new objectType",
			args: args{
				name: "objectType",
				parent: &Key{
					ObjectType: "key objectType",
					ID:         "id",
					Parent:     nil,
				},
			},
			want: &Namespace{
				Name: "objectType",
				Parent: &Key{
					ObjectType: "key objectType",
					ID:         "id",
					Parent:     nil,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNamespace(tt.args.name, tt.args.parent); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNamespace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewKey(t *testing.T) {
	type args struct {
		objectType string
		id         string
		parent     Marshaller
	}
	tests := []struct {
		name string
		args args
		want *Key
	}{
		{
			name: "initialize new key",
			args: args{
				objectType: "objectType",
				id:         "id",
				parent: &Namespace{
					Name:   "parent objectType",
					Parent: nil,
				},
			},
			want: &Key{
				ObjectType: "objectType",
				ID:         "id",
				Parent: &Namespace{
					Name:   "parent objectType",
					Parent: nil,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewKey(tt.args.objectType, tt.args.id, tt.args.parent); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKey_String(t *testing.T) {
	type fields struct {
		ObjectType string
		ID         string
		Parent     Marshaller
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "without objectType",
			fields: fields{
				ID: "1",
			},
			want: "1",
		},
		{
			name: "with objectType",
			fields: fields{
				ObjectType: "session",
				ID:         "1",
			},
			want: "session:1",
		},
		{
			name: "with only parent namespace",
			fields: fields{
				ID: "1",
				Parent: &Namespace{
					Name: "session",
				},
			},
			want: "session:1",
		},
		{
			name: "with objectType and parent namespace",
			fields: fields{
				ObjectType: "session",
				ID:         "1",
				Parent: &Namespace{
					Name: "auth",
				},
			},
			want: "auth:session:1",
		},
		{
			name: "with grand parent namespace",
			fields: fields{
				ObjectType: "sessionkey",
				ID:         "1",
				Parent: &Namespace{
					Name: "session",
					Parent: &Namespace{
						Name:   "auth",
						Parent: nil,
					},
				},
			},
			want: "auth:session:sessionkey:1",
		},
		{
			name: "with objectType and parent key",
			fields: fields{
				ObjectType: "comments",
				ID:         "1",
				Parent: &Key{
					ObjectType: "users",
					ID:         "1",
				},
			},
			want: "users:1:comments:1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := &Key{
				ObjectType: tt.fields.ObjectType,
				ID:         tt.fields.ID,
				Parent:     tt.fields.Parent,
			}
			if got := k.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
