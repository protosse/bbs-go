package models

type RspBuilder struct {
	Data map[string]interface{}
}

func (builder *RspBuilder) Put(key string, value interface{}) *RspBuilder {
	builder.Data[key] = value
	return builder
}

func (builder *RspBuilder) Build() map[string]interface{} {
	return builder.Data
}

func EmptyRsp() *RspBuilder {
	return &RspBuilder{Data: make(map[string]interface{})}
}

func UserRsp(user User) *RspBuilder {
	data := map[string]interface{}{}
	return &RspBuilder{Data: data}
}
