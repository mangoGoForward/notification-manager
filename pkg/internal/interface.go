package internal

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type Receiver interface {
	GetTenantID() string
	GetName() string
	Enabled() bool
	GetType() string
	GetLabels() map[string]string
	GetAlertSelector() *metav1.LabelSelector
	GetConfigSelector() *metav1.LabelSelector
	SetConfig(c Config)
	Validate() error
	Clone() Receiver
	GetHash() string
	SetHash(h string)
}

type Config interface {
	GetLabels() map[string]string
	GetPriority() int
	Validate() error
	Clone() Config
}
