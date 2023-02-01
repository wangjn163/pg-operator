package v1

import (
	acidttdo "github.com/wangjianc/pg-operator/pkg/apis/acid.tt.do"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

const (
	APIVersion = "v1"
)

var (
	SchemeBuilder      runtime.SchemeBuilder
	localSchemebuild   = &SchemeBuilder
	AddToScheme        = localSchemebuild.AddToScheme
	SchemeGroupVersion = schema.GroupVersion{Group: acidttdo.GroupName, Version: APIVersion}
)

func init() {
	localSchemebuild.Register(addKnownTypes)
}

func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypeWithName(SchemeGroupVersion.WithKind("postgresql"), &Postgresql{})
	scheme.AddKnownTypeWithName(SchemeGroupVersion.WithKind("postgresqlList"), &PostgresqlList{})
	return nil
}
