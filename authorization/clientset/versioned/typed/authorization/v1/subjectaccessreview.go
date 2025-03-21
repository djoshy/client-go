// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	context "context"

	authorizationv1 "github.com/openshift/api/authorization/v1"
	scheme "github.com/openshift/client-go/authorization/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	gentype "k8s.io/client-go/gentype"
)

// SubjectAccessReviewsGetter has a method to return a SubjectAccessReviewInterface.
// A group's client should implement this interface.
type SubjectAccessReviewsGetter interface {
	SubjectAccessReviews() SubjectAccessReviewInterface
}

// SubjectAccessReviewInterface has methods to work with SubjectAccessReview resources.
type SubjectAccessReviewInterface interface {
	Create(ctx context.Context, subjectAccessReview *authorizationv1.SubjectAccessReview, opts metav1.CreateOptions) (*authorizationv1.SubjectAccessReviewResponse, error)

	SubjectAccessReviewExpansion
}

// subjectAccessReviews implements SubjectAccessReviewInterface
type subjectAccessReviews struct {
	*gentype.Client[*authorizationv1.SubjectAccessReview]
}

// newSubjectAccessReviews returns a SubjectAccessReviews
func newSubjectAccessReviews(c *AuthorizationV1Client) *subjectAccessReviews {
	return &subjectAccessReviews{
		gentype.NewClient[*authorizationv1.SubjectAccessReview](
			"subjectaccessreviews",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *authorizationv1.SubjectAccessReview { return &authorizationv1.SubjectAccessReview{} },
		),
	}
}

// Create takes the representation of a subjectAccessReview and creates it.  Returns the server's representation of the subjectAccessReviewResponse, and an error, if there is any.
func (c *subjectAccessReviews) Create(ctx context.Context, subjectAccessReview *authorizationv1.SubjectAccessReview, opts metav1.CreateOptions) (result *authorizationv1.SubjectAccessReviewResponse, err error) {
	result = &authorizationv1.SubjectAccessReviewResponse{}
	err = c.GetClient().Post().
		Resource("subjectaccessreviews").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(subjectAccessReview).
		Do(ctx).
		Into(result)
	return
}
