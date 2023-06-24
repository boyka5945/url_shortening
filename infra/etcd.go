package infra

import (
	"context"
	"fmt"
	"time"
	"url_shortening/config"

	"go.etcd.io/etcd/client/v3"
)

type ServiceRegistrar struct {
	etcdClient   *clientv3.Client
	leaseID      clientv3.LeaseID
	serviceName  string
	serviceAddr  string
	methodPaths  []string
	leaseTimeout time.Duration
	leaseKeepCh  <-chan *clientv3.LeaseKeepAliveResponse
}

func NewServiceRegistrar(serviceName, serviceAddr string, methodPaths []string, leaseTimeout time.Duration) (*ServiceRegistrar, error) {
	var etcdEndpoints []string
	for _, v := range config.GetConfig().Etcd {
		etcdEndpoints = append(etcdEndpoints, fmt.Sprintf("%s:%d", v.Host, v.Port))
	}

	etcdClient, err := clientv3.New(clientv3.Config{
		Endpoints: etcdEndpoints,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to etcd: %v", err)
	}

	resp, err := etcdClient.Grant(context.Background(), int64(leaseTimeout.Seconds()))
	if err != nil {
		etcdClient.Close()
		return nil, fmt.Errorf("failed to create lease in etcd: %v", err)
	}

	keepAliveCh, err := etcdClient.KeepAlive(context.Background(), resp.ID)
	if err != nil {
		etcdClient.Close()
		return nil, fmt.Errorf("failed to start lease keep-alive: %v", err)
	}

	registrar := &ServiceRegistrar{
		etcdClient:   etcdClient,
		leaseID:      resp.ID,
		serviceName:  serviceName,
		serviceAddr:  serviceAddr,
		methodPaths:  methodPaths,
		leaseTimeout: leaseTimeout,
		leaseKeepCh:  keepAliveCh,
	}

	return registrar, nil
}

func (r *ServiceRegistrar) RegisterService() error {
	for _, methodPath := range r.methodPaths {
		key := fmt.Sprintf("/service/%s/%s", r.serviceName, methodPath)
		value := r.serviceAddr
	
		_, err := r.etcdClient.Put(context.Background(), key, value, clientv3.WithLease(r.leaseID))
		if err != nil {
			return fmt.Errorf("failed to register service in etcd: %v", err)
		}
	}

	GetLogger().Printf("Registered service %s method %s with address %s in etcd", r.serviceName, r.methodPaths, r.serviceAddr)
	return nil
}

func (r *ServiceRegistrar) DeleteInactiveServices() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	leaseResp, err := r.etcdClient.Lease.TimeToLive(ctx, r.leaseID)
	if err != nil {
		return fmt.Errorf("failed to get lease TTL: %v", err)
	}

	if leaseResp.TTL <= 0 {
		_, err := r.etcdClient.Delete(context.Background(), r.serviceName)
		if err != nil {
			return fmt.Errorf("failed to delete inactive service: %v", err)
		}

		GetLogger().Printf("Deleted inactive service %s from etcd", r.serviceName)
	}

	return nil
}
