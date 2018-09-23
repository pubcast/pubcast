package activitypub

import "net/url"

type pubObject struct{}

func (p pubObject) Serialize() (m map[string]interface{}, e error) {
	panic("not implemented")
}

func (p pubObject) TypeLen() (l int) {
	panic("not implemented")
}

func (p pubObject) GetType(index int) (v interface{}) {
	panic("not implemented")
}

func (p pubObject) GetId() *url.URL {
	panic("not implemented")
}

func (p pubObject) SetId(*url.URL) {
	panic("not implemented")
}

func (p pubObject) HasId() bool {
	panic("not implemented")
}

func (p pubObject) AppendType(interface{}) {
	panic("not implemented")
}

func (p pubObject) RemoveType(int) {
	panic("not implemented")
}
