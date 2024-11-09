package cgroup

// Controller defines the interface to a cgroup controller -- objects that
// model concrete cgroup controlers and their configuration options.
type Controller interface {
	// Controller defines the interface to a cgroup controller -- objects that
	Name() string

	// Apply applies this controller's configuration to the cgroup at the
	// given path.
	Apply(path string) error
}
