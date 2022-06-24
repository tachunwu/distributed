package main

type ServerBuilder interface {
	setRaftServer()
	setGrpcServer()
	setDatabaseClient()
}

func NewServerBuilder() ServerBuilder {
	return nil
}
