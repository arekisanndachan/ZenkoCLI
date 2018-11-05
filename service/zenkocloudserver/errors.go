// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package zenkocloudserver

const (

	// ErrCodeBucketAlreadyExists for service response error code
	// "BucketAlreadyExists".
	//
	// The requested bucket name is not available. The bucket namespace is shared
	// by all users of the system. Please select a different name and try again.
	ErrCodeBucketAlreadyExists = "BucketAlreadyExists"

	// ErrCodeBucketAlreadyOwnedByYou for service response error code
	// "BucketAlreadyOwnedByYou".
	ErrCodeBucketAlreadyOwnedByYou = "BucketAlreadyOwnedByYou"

	// ErrCodeNoSuchBucket for service response error code
	// "NoSuchBucket".
	//
	// The specified bucket does not exist.
	ErrCodeNoSuchBucket = "NoSuchBucket"

	// ErrCodeNoSuchKey for service response error code
	// "NoSuchKey".
	//
	// The specified key does not exist.
	ErrCodeNoSuchKey = "NoSuchKey"

	// ErrCodeNoSuchUpload for service response error code
	// "NoSuchUpload".
	//
	// The specified multipart upload does not exist.
	ErrCodeNoSuchUpload = "NoSuchUpload"

	// ErrCodeObjectAlreadyInActiveTierError for service response error code
	// "ObjectAlreadyInActiveTierError".
	//
	// This operation is not allowed against this storage tier
	ErrCodeObjectAlreadyInActiveTierError = "ObjectAlreadyInActiveTierError"

	// ErrCodeObjectNotInActiveTierError for service response error code
	// "ObjectNotInActiveTierError".
	//
	// The source object of the COPY operation is not in the active tier and is
	// only stored in Amazon Glacier.
	ErrCodeObjectNotInActiveTierError = "ObjectNotInActiveTierError"
)
