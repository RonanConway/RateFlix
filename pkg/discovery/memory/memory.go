package memory

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/RonanConway/RateFlix/pkg/discovery"
)

type serviceName string
type instanceID string

// Registry defines an in-memory service registry.
type Registry struct {
	sync.RWMutex
	serviceAddrs map[serviceName]map[instanceID]*serviceInstance
}

type serviceInstance struct {
	hostPort   string
	lastActive time.Time
}

// NewRegistry creates a new in-memory service
// registry instance.
func NewRegistry() *Registry {
	return &Registry{serviceAddrs: map[serviceName]map[instanceID]*serviceInstance{}}
}

// Register creates a service record in the registry.
func (r *Registry) Register(ctx context.Context, instID instanceID, svcName serviceName, hostPort string) error {
	r.Lock()
	defer r.Unlock()
	if _, ok := r.serviceAddrs[svcName]; !ok {
		r.serviceAddrs[svcName] = map[instanceID]*serviceInstance{}
	}

	r.serviceAddrs[svcName][instID] = &serviceInstance{hostPort: hostPort, lastActive: time.Now()}
	return nil
}

// Deregister removes a service record from the
// registry.
func (r *Registry) Deregister(ctx context.Context, instID instanceID, svcName serviceName) error {
	r.Lock()
	defer r.Unlock()
	if _, ok := r.serviceAddrs[svcName]; !ok {
		return nil
	}
	delete(r.serviceAddrs[svcName], instID)
	return nil
}

// ReportHealthyState is a push mechanism for
// reporting healthy state to the registry.
func (r *Registry) ReportHealthyState(instID instanceID, svcName serviceName) error {
	r.Lock()
	defer r.Unlock()
	if _, ok := r.serviceAddrs[svcName]; !ok {
		return errors.New("service is not registered yet")
	}
	if _, ok := r.serviceAddrs[svcName][instID]; !ok {
		return errors.New("service instance is not registered yet")
	}
	r.serviceAddrs[svcName][instID].lastActive = time.Now()
	return nil
}

// ServiceAddresses returns the list of addresses of
// active instances of the given service.
func (r *Registry) ServiceAddresses(ctx context.Context, svcName serviceName) ([]string, error) {
	r.RLock()
	defer r.RUnlock()
	// If there are no service addresses then return no service found error
	if len(r.serviceAddrs[svcName]) == 0 {
		return nil, discovery.ErrNotFound
	}

	var activeHosts []string

	// Go through each service and if their last active tim and if they have been active in the last 5 seconds then append the hostport to the slice.
	for _, i := range r.serviceAddrs[svcName] {
		if i.lastActive.Before(time.Now().Add(-5 * time.Second)) {
			continue
		}
		activeHosts = append(activeHosts, i.hostPort)
	}

	return activeHosts, nil
}
