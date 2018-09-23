import "net/url"

func (d deliverer) Do(b []byte, to *url.URL, toDo func(b []byte, u *url.URL) error) {
	panic("not implemented")
}

