package main

type constraint_checker[V comparable, D any] func(map[V]D) bool

type CSP[V any, D any] struct {
	variables   []V
	domains     []D
	constraints []constraint_checker
}
