package common

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

var mux sync.Mutex

type Backend struct {
	URL          *url.URL
	mux          sync.RWMutex
	ReverseProxy *httputil.ReverseProxy
}

type NodePool struct {
	backends []*Backend
	curInd  uint64
}

func (s *NodePool) AddContainer(nodeip *string, port *string) {
	containerUrl, err := url.Parse(node_ip+":"+port)
	if err != nil {
		log.Fatal(err)
	}
	proxy := httputil.NewSingleHostReverseProxy(containerUrl)
	s.backends = append(s.backends, &Backend{
			URL: containerUrl,
			ReverseProxy: proxy,
		})
}

func (s *NodePool) DeleteContainer(nodeip *string, port *string) {
	//Remove from Backend
}

func isBackendAlive(u *url.URL) bool {
	conn, err := net.DialTimeout("tcp", u.Host, 10 * time.Second)
	if err != nil {
		log.Println("Site unreachable, error: ", err)
		return false
	}
	defer conn.Close()
	return true
}

func (s *NodePool) HealthCheck() {
	for i, b := range s.backends {
		if !isBackendAlive(b.URL) {
			mux.Lock()
			s.backends = append(s.backends[:i], s.backends[i+1:]...)
			mux.Unlock()
		}
	}
}

func (s *NodePool) GetNextPeer() *Backend {
	if len(s.backends) != 0 
	{
		next := int(atomic.AddUint64(&s.curInd, uint64(1)) % uint64(len(s.backends)))

		atomic.StoreUint64(&s.curInd, next)

		return s.backends[next]
	}
	return nil
}

func lb(w http.ResponseWriter, r *http.Request) {

	peer := NodePool.GetNextPeer()
	if peer != nil {
		peer.ReverseProxy.ServeHTTP(w, r)
		return
	}
	http.Error(w, "Service not available", http.StatusServiceUnavailable)
}

func healthCheck() {
	t := time.NewTicker(time.Minute * 2)
	for {
		select {
		case <-t.C:
			NodePool.HealthCheck()
			log.Println("Health check completed")
		}
	}
}

