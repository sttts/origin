package testclient

import (
	authorizationapi "github.com/openshift/origin/pkg/authorization/api"
	"k8s.io/kubernetes/pkg/client/testing/core"
)

type FakeLocalResourceAccessReviews struct {
	Fake      *Fake
	Namespace string
}

func (c *FakeLocalResourceAccessReviews) Create(inObj *authorizationapi.LocalResourceAccessReview) (*authorizationapi.ResourceAccessReviewResponse, error) {
	obj, err := c.Fake.Invokes(core.NewCreateAction(authorizationapi.SchemeGroupVersion.WithResource("localresourceaccessreviews"), c.Namespace, inObj), &authorizationapi.ResourceAccessReviewResponse{})
	if cast, ok := obj.(*authorizationapi.ResourceAccessReviewResponse); ok {
		return cast, err
	}
	return nil, err
}
