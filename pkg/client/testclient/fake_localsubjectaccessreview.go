package testclient

import (
	authorizationapi "github.com/openshift/origin/pkg/authorization/api"
	"k8s.io/kubernetes/pkg/client/testing/core"
)

type FakeLocalSubjectAccessReviews struct {
	Fake      *Fake
	Namespace string
}

func (c *FakeLocalSubjectAccessReviews) Create(inObj *authorizationapi.LocalSubjectAccessReview) (*authorizationapi.SubjectAccessReviewResponse, error) {
	obj, err := c.Fake.Invokes(core.NewCreateAction(authorizationapi.SchemeGroupVersion.WithResource("localsubjectaccessreviews"), c.Namespace, inObj), &authorizationapi.SubjectAccessReviewResponse{})
	if cast, ok := obj.(*authorizationapi.SubjectAccessReviewResponse); ok {
		return cast, err
	}
	return nil, err
}
