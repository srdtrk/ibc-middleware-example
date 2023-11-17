package example

// NewCallbackCounter creates a new CallbackCounter
func NewCallbackCounter(channelID string) CallbackCounter {
	return CallbackCounter{
		OnRecvPacket:            0,
		OnAcknowledgementPacket: 0,
		OnTimeoutPacket:         0,
		SendPacket:              0,
		ChannelId:               channelID,
	}
}
