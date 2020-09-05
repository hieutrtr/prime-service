package store

type PrimeCache struct {
	Marks []uint32
	Limit uint32
}

// Generate all primes marked array using Sieve of Sundaram
func NewPrimeCache(max uint32) *PrimeCache{
	primeCache := new(PrimeCache)
	primeCache.Limit = max
	primeCache.Marks = make([]uint32, max/2)
	halfMax := max/2
	for i := uint32(1); i <= halfMax; i++ {
		j := i
		for i+j+i*j*2 < halfMax {
			primeCache.Marks[i+j+i*j*2] = 1
			j++
		}
	}
	return primeCache
}
