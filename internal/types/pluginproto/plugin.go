package pluginproto

import "google.golang.org/protobuf/proto"

func (p *PluginInfo) Marshal() ([]byte, error) {
	return proto.Marshal(p)
}

func (p *PluginInfo) Unmarshal(data []byte) error {
	return proto.Unmarshal(data, p)
}

func (s *SendPacket) Marshal() ([]byte, error) {
	return proto.Marshal(s)
}

func (s *SendPacket) Unmarshal(data []byte) error {
	return proto.Unmarshal(data, s)
}

func (m *Message) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *Message) Unmarshal(data []byte) error {
	return proto.Unmarshal(data, m)
}

func (m *MessageBatch) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *MessageBatch) Unmarshal(data []byte) error {
	return proto.Unmarshal(data, m)
}

func (h *HttpRequest) Marshal() ([]byte, error) {
	return proto.Marshal(h)
}

func (h *HttpRequest) Unmarshal(data []byte) error {
	return proto.Unmarshal(data, h)
}

func (h *HttpResponse) Marshal() ([]byte, error) {
	return proto.Marshal(h)
}

func (h *HttpResponse) Unmarshal(data []byte) error {
	return proto.Unmarshal(data, h)
}
