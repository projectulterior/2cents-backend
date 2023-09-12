package resolver

import "context"

type ChannelsGetter interface {
	Channels(context.Context) ([]*Channel, error)
	Next(context.Context) (*string, error)
}

type Channels struct {
	ChannelsGetter
}

func NewChannels(getter ChannelsGetter) *Channels {
	return &Channels{ChannelsGetter: getter}
}
