package metrics

type ClientConfig struct {
	// Namespace is the prefix to be prepended to all client calls
	Namespace string
	// GlobalTags are tags that will be added to every metric
	GlobalTags []string
}

type ClientOption func(o *ClientConfig)

// WithNamespace configures the client to use the provided namespace.
func WithNamespace(namespace string) ClientOption {
	return func(o *ClientConfig) {
		o.Namespace = namespace
	}
}

// WithGlobalTags configures the client to include a set of tags to all submitted
// metrics.
func WithGlobalTags(tagOptions ...TagOption) ClientOption {
	tags := GetTags(tagOptions...)
	return func(o *ClientConfig) {
		o.GlobalTags = append(o.GlobalTags, tags...)
	}
}
