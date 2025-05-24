package platform

import (
	"fmt"
	// "github.com/hewenyu/modelbridge/client" // For client.Logger, if used in constructors
)

// HandlerConstructor defines the function signature for creating a new platform handler.
// It takes a PlatformConfig and potentially a logger, and returns a PlatformHandler or an error.
type HandlerConstructor func(config *PlatformConfig /*, logger client.Logger */) (PlatformHandler, error)

var (
	handlerRegistry = make(map[string]HandlerConstructor)
)

// RegisterHandler registers a new platform handler constructor for a given platform name.
// It will overwrite any existing handler for the same platform name.
func RegisterHandler(platformName string, constructor HandlerConstructor) {
	if constructor == nil {
		// Potentially log a warning or panic if a nil constructor is provided.
		// For now, we'll allow it, but GetHandler will fail for this platformName.
	}
	handlerRegistry[platformName] = constructor
}

// GetHandler retrieves a new PlatformHandler instance for the specified platform name.
// It uses the registered constructor and the provided configuration.
func GetHandler(platformName string, config *PlatformConfig /*, logger client.Logger */) (PlatformHandler, error) {
	constructor, ok := handlerRegistry[platformName]
	if !ok {
		return nil, fmt.Errorf("platform %s not registered", platformName)
	}
	if constructor == nil { // Should ideally not happen if RegisterHandler prevents nil constructors
		return nil, fmt.Errorf("platform %s has a nil constructor registered", platformName)
	}
	return constructor(config /*, logger */)
}

// ListRegisteredPlatforms returns a list of all registered platform names.
func ListRegisteredPlatforms() []string {
	names := make([]string, 0, len(handlerRegistry))
	for name := range handlerRegistry {
		names = append(names, name)
	}
	return names
}
