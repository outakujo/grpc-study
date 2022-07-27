package main

import (
	"grpc-study/pb"
	"sync"
)

type End struct {
	Id string
	Ch chan *pb.Cmd
}

type Manager struct {
	m   map[string]*End
	mut sync.Mutex
}

func NewManager() *Manager {
	return &Manager{m: make(map[string]*End)}
}

func (r *Manager) Add(id string) *End {
	r.mut.Lock()
	defer r.mut.Unlock()
	end, ok := r.m[id]
	if ok {
		return end
	}
	end = &End{Id: id, Ch: make(chan *pb.Cmd)}
	r.m[id] = end
	return end
}

func (r *Manager) Del(id string) *End {
	r.mut.Lock()
	defer r.mut.Unlock()
	end, ok := r.m[id]
	if !ok {
		return nil
	}
	delete(r.m, id)
	return end
}

func (r *Manager) Get(id string) *End {
	r.mut.Lock()
	defer r.mut.Unlock()
	end, ok := r.m[id]
	if !ok {
		return nil
	}
	return end
}

var manager = NewManager()
