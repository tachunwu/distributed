package main

// interface has an interface inside

type Server interface {
	store()
	setStorager(Storager)
}

type Storager interface {
	store()
}

type PebbleServer struct {
	Storager Storager
}
type BadgerServer struct {
	Storager Storager
}
