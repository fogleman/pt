package pt

import "math/rand"

type BRDF interface {
	Sample(hit Hit, rnd *rand.Rand)
	Evaluate(hit Hit, ray Ray)
}
