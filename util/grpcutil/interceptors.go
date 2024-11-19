package grpcutil

import (
	"context"
	"errors"

	"github.com/obaraelijah/secureproc/pkg/jobmanager"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
)

// userIDContext is the type of the key in a context.Context where the
// userID is stored.
type userIDContext struct{}

// userIDFromContext extracts and returns the userID from the given context.
// If the userID doesn't exist, returns an error ready to be returned by a
// gRPC API.
func GetUserIDFromContext(ctx context.Context) (string, error) {
	if userID, ok := ctx.Value(&userIDContext{}).(string); ok && userID != "" {
		return userID, nil
	}

	return "", jobmanager.Unauthenticated
}

// UnaryGetUserIDFromContextInterceptor extracts the CommonName from the client-
// supplied certificate, creates a new context.Context with that value using
// key UserIDContext, and invokes the given handler.
func UnaryGetUserIDFromContextInterceptor(
	ctx context.Context,
	req interface{},
	_ *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {

	ctx, err = getUserIDFromCommonName(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "jobmanager: unauthenticated")
	}

	resp, err = handler(ctx, req)
	if err != nil {
		err = status.Errorf(errorToGRPCErrorCode(err), err.Error())
	}

	return
}

// streamWrapper is a wrapper over a grpc.ServerStream that enables us to
// override Context() and add the userID to it.
type streamWrapper struct {
	grpc.ServerStream
}

// Context overrides the wrapped grpc.ServerStream's Context(), finds the
// userID, and adds it to the returned context.
func (s *streamWrapper) Context() context.Context {
	ctx := s.ServerStream.Context()

	newCtx, err := getUserIDFromCommonName(ctx)
	if err != nil {
		return ctx
	}

	return newCtx
}

// StreamGetUserIDFromContextInterceptor is a StreamInterceptor that
// extracts the CommonName from the client-supplied certificate, creates a
// new context.Context with that value using key UserIDContext, and invokes
// the given handler.
func StreamGetUserIDFromContextInterceptor(
	srv interface{},
	ss grpc.ServerStream,
	info *grpc.StreamServerInfo,
	handler grpc.StreamHandler,
) error {

	err := handler(srv, &streamWrapper{ServerStream: ss})
	if err != nil {
		return status.Errorf(errorToGRPCErrorCode(err), err.Error())
	}

	return nil
}

// getUserIDFromCommonName searches the client-supplied certificates associated
// with the given ctx, extracts the CommonName, creates a new context,
// attaches the CommonName as a new value, and returns the new context.
func getUserIDFromCommonName(ctx context.Context) (context.Context, error) {
	if p, ok := peer.FromContext(ctx); ok {
		if mtls, ok := p.AuthInfo.(credentials.TLSInfo); ok {
			for _, item := range mtls.State.PeerCertificates {
				if item.Subject.CommonName != "" {
					return AttachUserIDToContext(ctx, item.Subject.CommonName), nil
				}
			}
		}
	}

	return nil, status.Error(codes.Unauthenticated, "jobmanager: unauthenticated")
}

func AttachUserIDToContext(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, &userIDContext{}, userID)
}

// errorToGRPCErrorCode maps the given error to a suitable gRPC error code.
// If no mapping is found, it will return codes.Internal.
func errorToGRPCErrorCode(err error) codes.Code {
	code := codes.Internal

	if errors.Is(err, jobmanager.JobExistsError) {
		code = codes.AlreadyExists
	} else if errors.Is(err, jobmanager.JobNotFoundError) {
		code = codes.NotFound
	} else if errors.Is(err, jobmanager.InvalidJobIDError) {
		code = codes.InvalidArgument
	} else if errors.Is(err, jobmanager.InvalidArgument) {
		code = codes.InvalidArgument
	} else if errors.Is(err, jobmanager.Unauthenticated) {
		code = codes.Unauthenticated
	} else if errors.Is(err, context.DeadlineExceeded) {
		code = codes.DeadlineExceeded
	}

	return code
}
