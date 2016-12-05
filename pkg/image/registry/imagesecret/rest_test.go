package imagesecret

import (
	"testing"

	kapi "k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/fake"

	"github.com/openshift/origin/pkg/image/api"
)

func TestGetSecrets(t *testing.T) {
	fake := fake.NewSimpleClientset(&kapi.SecretList{
		Items: []kapi.Secret{
			{
				ObjectMeta: kapi.ObjectMeta{Name: "secret-1"},
				Type:       kapi.SecretTypeDockercfg,
			},
			{
				ObjectMeta: kapi.ObjectMeta{Name: "secret-2", Annotations: map[string]string{api.ExcludeImageSecretAnnotation: "true"}},
				Type:       kapi.SecretTypeDockercfg,
			},
			{
				ObjectMeta: kapi.ObjectMeta{Name: "secret-3"},
				Type:       kapi.SecretTypeOpaque,
			},
			{
				ObjectMeta: kapi.ObjectMeta{Name: "secret-4"},
				Type:       kapi.SecretTypeServiceAccountToken,
			},
			{
				ObjectMeta: kapi.ObjectMeta{Name: "secret-5"},
				Type:       kapi.SecretTypeDockerConfigJson,
			},
		},
	})
	rest := NewREST(fake.Core())
	opts, _, _ := rest.NewGetOptions()
	obj, err := rest.Get(kapi.NewDefaultContext(), "", opts)
	if err != nil {
		t.Fatal(err)
	}
	list := obj.(*kapi.SecretList)
	if len(list.Items) != 2 {
		t.Fatal(list)
	}
	if list.Items[0].Name != "secret-1" || list.Items[1].Name != "secret-5" {
		t.Fatal(list)
	}
}
