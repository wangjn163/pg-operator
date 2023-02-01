package v1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Postgresql struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PostgresSpec   `json:"Spec"`
	Status            PostgresStatus `json:"status"`
	Error             string         `json:"-"`
	DockerImage       string         `json:"dockerImage,omitempty"`
	ClusterName       string         `json:"-"`
}

type PostgresSpec struct {
	DockerImage string `json:"dockerImage,omitempty"`
	ClusterName string `json:"-"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type PostgresqlList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []Postgresql `json:"items"`
}

type PostgresStatus struct {
	PostgresClusterStatus string `json:"PostgresClusterStatus"`
}
