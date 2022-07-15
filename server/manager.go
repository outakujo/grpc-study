package main

import "grpc-study/pb"

type End struct {
	Id string
	Ch chan *pb.Cmd
}

type Manager struct {
	m map[string]*End
}

func NewManager() *Manager {
	return &Manager{m: make(map[string]*End)}
}

func (r *Manager) Add(id string) *End {
	end, ok := r.m[id]
	if ok {
		return end
	}
	end = &End{Id: id, Ch: make(chan *pb.Cmd)}
	r.m[id] = end
	return end
}

func (r *Manager) Del(id string) *End {
	end, ok := r.m[id]
	if !ok {
		return nil
	}
	delete(r.m, id)
	return end
}
