package testclient

import (
	authorizationapi "github.com/openshift/origin/pkg/authorization/api"
	"k8s.io/kubernetes/pkg/client/testing/core"
)

type FakeSelfSubjectRulesReviews struct {
	Fake      *Fake
	Namespace string
}

func (c *FakeSelfSubjectRulesReviews) Create(inObj *authorizationapi.SelfSubjectRulesReview) (*authorizationapi.SelfSubjectRulesReview, error) {
	obj, err := c.Fake.Invokes(core.NewCreateAction(authorizationapi.SchemeGroupVersion.WithResource("selfsubjectrulesreviews"), c.Namespace, inObj), &authorizationapi.SelfSubjectRulesReview{})
	if cast, ok := obj.(*authorizationapi.SelfSubjectRulesReview); ok {
		return cast, err
	}
	return nil, err
}
