package testclient

import (
	authorizationapi "github.com/openshift/origin/pkg/authorization/api"
	"k8s.io/kubernetes/pkg/client/testing/core"
)

type FakeClusterResourceAccessReviews struct {
	Fake *Fake
}

func (c *FakeClusterResourceAccessReviews) Create(inObj *authorizationapi.ResourceAccessReview) (*authorizationapi.ResourceAccessReviewResponse, error) {
	obj, err := c.Fake.Invokes(core.NewRootCreateAction(authorizationapi.SchemeGroupVersion.WithResource("resourceaccessreviews"), inObj), &authorizationapi.ResourceAccessReviewResponse{})
	if cast, ok := obj.(*authorizationapi.ResourceAccessReviewResponse); ok {
		return cast, err
	}
	return nil, err
}
