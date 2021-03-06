// Code generated by mocker. DO NOT EDIT.
// github.com/travisjeffery/mocker
// Source: pb/usersvc/v1/usersvc.pb.go

package mock

import (
	context "context"
	sync "sync"

	github_com_stkr89_livegig_common_pb_usersvc_v1 "github.com/stkr89/livegig-common/pb/usersvc/v1"
	google_golang_org_grpc "google.golang.org/grpc"
)

// MockUserSvcClient is a mock of UserSvcClient interface
type MockUserSvcClient struct {
	lockCreate sync.Mutex
	CreateFunc func(ctx context.Context, in *github_com_stkr89_livegig_common_pb_usersvc_v1.CreateRequest, opts ...google_golang_org_grpc.CallOption) (*github_com_stkr89_livegig_common_pb_usersvc_v1.UserResponse, error)

	lockList sync.Mutex
	ListFunc func(ctx context.Context, in *github_com_stkr89_livegig_common_pb_usersvc_v1.ListRequest, opts ...google_golang_org_grpc.CallOption) (*github_com_stkr89_livegig_common_pb_usersvc_v1.ListResponse, error)

	lockDelete sync.Mutex
	DeleteFunc func(ctx context.Context, in *github_com_stkr89_livegig_common_pb_usersvc_v1.DeleteRequest, opts ...google_golang_org_grpc.CallOption) (*github_com_stkr89_livegig_common_pb_usersvc_v1.DeleteResonse, error)

	lockUpdate sync.Mutex
	UpdateFunc func(ctx context.Context, in *github_com_stkr89_livegig_common_pb_usersvc_v1.UpdateRequest, opts ...google_golang_org_grpc.CallOption) (*github_com_stkr89_livegig_common_pb_usersvc_v1.UserResponse, error)

	calls struct {
		Create []struct {
			Ctx  context.Context
			In   *github_com_stkr89_livegig_common_pb_usersvc_v1.CreateRequest
			Opts []google_golang_org_grpc.CallOption
		}
		List []struct {
			Ctx  context.Context
			In   *github_com_stkr89_livegig_common_pb_usersvc_v1.ListRequest
			Opts []google_golang_org_grpc.CallOption
		}
		Delete []struct {
			Ctx  context.Context
			In   *github_com_stkr89_livegig_common_pb_usersvc_v1.DeleteRequest
			Opts []google_golang_org_grpc.CallOption
		}
		Update []struct {
			Ctx  context.Context
			In   *github_com_stkr89_livegig_common_pb_usersvc_v1.UpdateRequest
			Opts []google_golang_org_grpc.CallOption
		}
	}
}

// Create mocks base method by wrapping the associated func.
func (m *MockUserSvcClient) Create(ctx context.Context, in *github_com_stkr89_livegig_common_pb_usersvc_v1.CreateRequest, opts ...google_golang_org_grpc.CallOption) (*github_com_stkr89_livegig_common_pb_usersvc_v1.UserResponse, error) {
	m.lockCreate.Lock()
	defer m.lockCreate.Unlock()

	if m.CreateFunc == nil {
		panic("mocker: MockUserSvcClient.CreateFunc is nil but MockUserSvcClient.Create was called.")
	}

	call := struct {
		Ctx  context.Context
		In   *github_com_stkr89_livegig_common_pb_usersvc_v1.CreateRequest
		Opts []google_golang_org_grpc.CallOption
	}{
		Ctx:  ctx,
		In:   in,
		Opts: opts,
	}

	m.calls.Create = append(m.calls.Create, call)

	return m.CreateFunc(ctx, in, opts...)
}

// CreateCalled returns true if Create was called at least once.
func (m *MockUserSvcClient) CreateCalled() bool {
	m.lockCreate.Lock()
	defer m.lockCreate.Unlock()

	return len(m.calls.Create) > 0
}

// CreateCalls returns the calls made to Create.
func (m *MockUserSvcClient) CreateCalls() []struct {
	Ctx  context.Context
	In   *github_com_stkr89_livegig_common_pb_usersvc_v1.CreateRequest
	Opts []google_golang_org_grpc.CallOption
} {
	m.lockCreate.Lock()
	defer m.lockCreate.Unlock()

	return m.calls.Create
}

// List mocks base method by wrapping the associated func.
func (m *MockUserSvcClient) List(ctx context.Context, in *github_com_stkr89_livegig_common_pb_usersvc_v1.ListRequest, opts ...google_golang_org_grpc.CallOption) (*github_com_stkr89_livegig_common_pb_usersvc_v1.ListResponse, error) {
	m.lockList.Lock()
	defer m.lockList.Unlock()

	if m.ListFunc == nil {
		panic("mocker: MockUserSvcClient.ListFunc is nil but MockUserSvcClient.List was called.")
	}

	call := struct {
		Ctx  context.Context
		In   *github_com_stkr89_livegig_common_pb_usersvc_v1.ListRequest
		Opts []google_golang_org_grpc.CallOption
	}{
		Ctx:  ctx,
		In:   in,
		Opts: opts,
	}

	m.calls.List = append(m.calls.List, call)

	return m.ListFunc(ctx, in, opts...)
}

// ListCalled returns true if List was called at least once.
func (m *MockUserSvcClient) ListCalled() bool {
	m.lockList.Lock()
	defer m.lockList.Unlock()

	return len(m.calls.List) > 0
}

// ListCalls returns the calls made to List.
func (m *MockUserSvcClient) ListCalls() []struct {
	Ctx  context.Context
	In   *github_com_stkr89_livegig_common_pb_usersvc_v1.ListRequest
	Opts []google_golang_org_grpc.CallOption
} {
	m.lockList.Lock()
	defer m.lockList.Unlock()

	return m.calls.List
}

// Delete mocks base method by wrapping the associated func.
func (m *MockUserSvcClient) Delete(ctx context.Context, in *github_com_stkr89_livegig_common_pb_usersvc_v1.DeleteRequest, opts ...google_golang_org_grpc.CallOption) (*github_com_stkr89_livegig_common_pb_usersvc_v1.DeleteResonse, error) {
	m.lockDelete.Lock()
	defer m.lockDelete.Unlock()

	if m.DeleteFunc == nil {
		panic("mocker: MockUserSvcClient.DeleteFunc is nil but MockUserSvcClient.Delete was called.")
	}

	call := struct {
		Ctx  context.Context
		In   *github_com_stkr89_livegig_common_pb_usersvc_v1.DeleteRequest
		Opts []google_golang_org_grpc.CallOption
	}{
		Ctx:  ctx,
		In:   in,
		Opts: opts,
	}

	m.calls.Delete = append(m.calls.Delete, call)

	return m.DeleteFunc(ctx, in, opts...)
}

// DeleteCalled returns true if Delete was called at least once.
func (m *MockUserSvcClient) DeleteCalled() bool {
	m.lockDelete.Lock()
	defer m.lockDelete.Unlock()

	return len(m.calls.Delete) > 0
}

// DeleteCalls returns the calls made to Delete.
func (m *MockUserSvcClient) DeleteCalls() []struct {
	Ctx  context.Context
	In   *github_com_stkr89_livegig_common_pb_usersvc_v1.DeleteRequest
	Opts []google_golang_org_grpc.CallOption
} {
	m.lockDelete.Lock()
	defer m.lockDelete.Unlock()

	return m.calls.Delete
}

// Update mocks base method by wrapping the associated func.
func (m *MockUserSvcClient) Update(ctx context.Context, in *github_com_stkr89_livegig_common_pb_usersvc_v1.UpdateRequest, opts ...google_golang_org_grpc.CallOption) (*github_com_stkr89_livegig_common_pb_usersvc_v1.UserResponse, error) {
	m.lockUpdate.Lock()
	defer m.lockUpdate.Unlock()

	if m.UpdateFunc == nil {
		panic("mocker: MockUserSvcClient.UpdateFunc is nil but MockUserSvcClient.Update was called.")
	}

	call := struct {
		Ctx  context.Context
		In   *github_com_stkr89_livegig_common_pb_usersvc_v1.UpdateRequest
		Opts []google_golang_org_grpc.CallOption
	}{
		Ctx:  ctx,
		In:   in,
		Opts: opts,
	}

	m.calls.Update = append(m.calls.Update, call)

	return m.UpdateFunc(ctx, in, opts...)
}

// UpdateCalled returns true if Update was called at least once.
func (m *MockUserSvcClient) UpdateCalled() bool {
	m.lockUpdate.Lock()
	defer m.lockUpdate.Unlock()

	return len(m.calls.Update) > 0
}

// UpdateCalls returns the calls made to Update.
func (m *MockUserSvcClient) UpdateCalls() []struct {
	Ctx  context.Context
	In   *github_com_stkr89_livegig_common_pb_usersvc_v1.UpdateRequest
	Opts []google_golang_org_grpc.CallOption
} {
	m.lockUpdate.Lock()
	defer m.lockUpdate.Unlock()

	return m.calls.Update
}

// Reset resets the calls made to the mocked methods.
func (m *MockUserSvcClient) Reset() {
	m.lockCreate.Lock()
	m.calls.Create = nil
	m.lockCreate.Unlock()
	m.lockList.Lock()
	m.calls.List = nil
	m.lockList.Unlock()
	m.lockDelete.Lock()
	m.calls.Delete = nil
	m.lockDelete.Unlock()
	m.lockUpdate.Lock()
	m.calls.Update = nil
	m.lockUpdate.Unlock()
}
