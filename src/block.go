package main

import (
	"math/rand"
	"sync"
	"time"
)

type Block struct {
	index      int
	rand       *rand.Rand
	mu         sync.Mutex
	conditions []string
	block      []string
}

func blockNew(sizeMultiple int, conditions []string) *Block {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	block := make([]string, len(conditions)*sizeMultiple)
	index := 0
	for _, c := range conditions {
		for range 2 {
			block[index] = c
			index++
		}
	}

	r.Shuffle(len(block), func(i, j int) {
		block[i], block[j] = block[j], block[i]
	})

	b := Block{
		index:      0,
		rand:       r,
		mu:         sync.Mutex{},
		conditions: conditions,
		block:      block,
	}

	return &b
}

func blockGetCondition(block *Block) *string {
	block.mu.Lock()
	defer block.mu.Unlock()

	if block.index >= len(block.block) {
		block.index = 0
		block.rand.Shuffle(len(block.block), func(i, j int) {
			block.block[i], block.block[j] = block.block[j], block.block[i]
		})
	}

	condition := &block.block[block.index]
	block.index++

	return condition
}
