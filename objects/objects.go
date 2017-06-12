package objects

import (
	"bytes"
	"io"
	"io/ioutil"

	uuid "github.com/satori/go.uuid"
)

type ObjectType string

const (
	InvalidObject ObjectType = ""
	BlobObject    ObjectType = "blob"
)

func (t ObjectType) Bytes() []byte {
	return []byte(string(t))
}

var _ EncodedObject = (*Object)(nil)

func NewObject(id uuid.UUID) *Object {
	return &Object{id: id}
}

type Object struct {
	id uuid.UUID
	t  ObjectType
	ct string // content type
	d  []byte
	sz int64
}

func (o *Object) ID() uuid.UUID {
	return o.id
}

func (o *Object) Size() int64 {
	return o.sz
}

func (o *Object) SetSize(s int64) {
	o.sz = s
}

func (o *Object) Type() ObjectType {
	return o.t
}

func (o *Object) ContentType() string {
	return o.ct
}

func (o *Object) SetContentType(ct string) {
	o.ct = ct
}

func (o *Object) SetType(t ObjectType) {
	o.t = t
}

func (o *Object) Reader() (io.ReadCloser, error) {
	return ioutil.NopCloser(bytes.NewBuffer(o.d)), nil
}

func (o *Object) Writer() (io.WriteCloser, error) {
	return o, nil
}

func (o *Object) Write(p []byte) (n int, err error) {
	o.d = append(o.d, p...)
	o.sz = int64(len(o.d))

	return len(p), nil
}

func (o *Object) Close() error {
	return nil
}