package proto

//go:generate protoc --proto_path=. --micro_out=. --go_out=. user.proto

const (
	QualifiedName = "srv.user"
)
