package main

import ("fmt")

type Shard struct{
	data map[string]string
}

func NewShard() *Shard{
	return &Shard{data : make(map[string]string)}
}

func (s *Shard) Insert(key, value string){
	s.data[key] = value
}

func (s* Shard) Get(key string, val string)(string, bool){
	val, exists := s.data[key]
	return val, exists
}