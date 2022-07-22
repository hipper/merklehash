package merkle

import (
	"crypto/sha256"
	"errors"
	"fmt"
)

// Hasher is the interface implemented by hash functions.
type Hasher interface {
	Hash(data []byte) [sha256.Size]byte
}

// TreeHash computes the SHA-256 tree hash for the given input.
type TreeHash struct {
	hasher Hasher
}

// NewTreeHash creates a new instance of a hashed tree.
// Note: This tree implementation doesn't duplicate or hash orphan leaves multiple times.
// Those leaves are added to the next available level instead.
func NewTreeHash(hasher Hasher) *TreeHash {
	return &TreeHash{hasher: hasher}
}

// CalculateRoot computes SHA-256 hash and returns it in a human-readable format.
func (t *TreeHash) CalculateRoot(d [][]byte) (string, error) {
	ln := len(d)
	if ln == 0 {
		return "", errors.New("tree must have at least 1 leaf specified")
	}

	// Build initial set of leaves by hashing input data.
	nodes := make([][sha256.Size]byte, 0, ln)

	for _, l := range d {
		if len(l) == 0 {
			continue
		}

		nodes = append(nodes, t.hasher.Hash(l))
	}

	root := t.compute(nodes)

	return fmt.Sprintf("%x", root), nil
}

// This approach uses a pair of arrays to compute hash for each layer, by taking 2 elements and
// computing the SHA-256 hash of the pair. The value is then added to the next layer for further processing
// during the next iteration.
func (t *TreeHash) compute(prev [][sha256.Size]byte) []byte {
	for len(prev) > 1 {
		ln := len(prev)
		next := make([][sha256.Size]byte, 0, ln)

		for i := 0; i < ln; i = i + 2 {
			// Check if we have at least two nodes remaining.
			if ln-i > 1 {
				next = append(next, t.hasher.Hash(append(prev[i][:], prev[i+1][:]...)))
				continue
			}

			// Add one remaining odd node.
			next = append(next, prev[i])
		}
		prev = next
	}

	return prev[0][:]
}

// SHA256Hasher uses SHA256 to hash data.
type SHA256Hasher struct {
}

// Hash implements Hasher.
func (h *SHA256Hasher) Hash(data []byte) [sha256.Size]byte {
	return sha256.Sum256(data)
}
