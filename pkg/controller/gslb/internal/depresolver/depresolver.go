// Package depresolver abstracts and implements ohmyglb dependencies resolver.
// depresolver responsibilities
// - abstracts multiple configurations into single point of access
// - provides predefined values when configuration is missing
// - validates configuration
// - executes once
// TODO: Add the rest of configuration to be resolved
package depresolver

import (
	"context"
	"sync"

	"sigs.k8s.io/controller-runtime/pkg/client"
)

//Config holds operator configuration
type Config struct {
	//Reschedule of Reconcile loop to pickup external Gslb targets
	ReconcileRequeueSeconds int
}

//DependencyResolver resolves configuration for GSLB
type DependencyResolver struct {
	client      client.Client
	config      *Config
	context     context.Context
	onceConfig  sync.Once
	onceSpec    sync.Once
	errorConfig error
	errorSpec   error
}

const (
	lessOrEqualToZeroErrorMessage = "\"%s is less or equal to zero\""
	lessThanZeroErrorMessage      = "\"%s is less than zero\""
)

//NewDependencyResolver returns a new depresolver.DependencyResolver
func NewDependencyResolver(context context.Context, client client.Client) *DependencyResolver {
	resolver := new(DependencyResolver)
	resolver.client = client
	resolver.context = context
	return resolver
}
