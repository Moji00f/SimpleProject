package limiter

import (
	"golang.org/x/time/rate"
	"sync"
)

type IPRateLimiter struct {
	ips map[string]*rate.Limiter
	mux *sync.RWMutex
	r   rate.Limit
	b   int
}

func NewIPRateLimiter(r rate.Limit, b int) *IPRateLimiter {
	return &IPRateLimiter{
		ips: make(map[string]*rate.Limiter),
		mux: &sync.RWMutex{},
		r:   r,
		b:   b,
	}
}

func (i *IPRateLimiter) AddIp(ip string) *rate.Limiter {
	i.mux.Lock()
	defer i.mux.Unlock()

	limiter := rate.NewLimiter(i.r, i.b)
	i.ips[ip] = limiter

	return limiter
}

func (i *IPRateLimiter) GetLimiter(ip string) *rate.Limiter {
	i.mux.Lock()
	limiter, exists := i.ips[ip]
	if !exists {
		i.mux.Unlock()
		return i.AddIp(ip)
	}
	i.mux.Unlock()
	return limiter
}
