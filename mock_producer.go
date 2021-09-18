package easy_topic_producer

type mockTopicProducer struct {
}

//NewMockProducer instance
func NewMockProducer() (TopicProducer, error) {
	return &mockTopicProducer{}, nil
}

func (p *mockTopicProducer) SendMessage(_, _ []byte) error {
	return nil
}
