import (
	"context"

	"github.com/go-fed/activity/streams"
)

func (c callbacker) Create(c context.Context, s *streams.Create) error {
	panic("not implemented")
}

func (c callbacker) Update(c context.Context, s *streams.Update) error {
	panic("not implemented")
}

func (c callbacker) Delete(c context.Context, s *streams.Delete) error {
	panic("not implemented")
}

func (c callbacker) Add(c context.Context, s *streams.Add) error {
	panic("not implemented")
}

func (c callbacker) Remove(c context.Context, s *streams.Remove) error {
	panic("not implemented")
}

func (c callbacker) Like(c context.Context, s *streams.Like) error {
	panic("not implemented")
}

func (c callbacker) Block(c context.Context, s *streams.Block) error {
	panic("not implemented")
}

func (c callbacker) Follow(c context.Context, s *streams.Follow) error {
	panic("not implemented")
}

func (c callbacker) Undo(c context.Context, s *streams.Undo) error {
	panic("not implemented")
}

func (c callbacker) Accept(c context.Context, s *streams.Accept) error {
	panic("not implemented")
}

func (c callbacker) Reject(c context.Context, s *streams.Reject) error {
	panic("not implemented")
}

