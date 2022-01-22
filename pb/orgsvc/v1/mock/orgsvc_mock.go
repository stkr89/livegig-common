// Code generated by mocker. DO NOT EDIT.
// github.com/travisjeffery/mocker
// Source: pb/orgsvc/v1/orgsvc.pb.go

package mock

import (
	context "context"
	sync "sync"

	github_com_stkr89_livegig_common_pb_orgsvc_v1 "github.com/stkr89/livegig-common/pb/orgsvc/v1"
	google_golang_org_grpc "google.golang.org/grpc"
)

// MockOrgSvcClient is a mock of OrgSvcClient interface
type MockOrgSvcClient struct {
	lockCreate sync.Mutex
	CreateFunc func(ctx context.Context, in *github_com_stkr89_livegig_common_pb_orgsvc_v1.CreateRequest, opts ...google_golang_org_grpc.CallOption) (*github_com_stkr89_livegig_common_pb_orgsvc_v1.OrgResponse, error)

	lockDelete sync.Mutex
	DeleteFunc func(ctx context.Context, in *github_com_stkr89_livegig_common_pb_orgsvc_v1.DeleteRequest, opts ...google_golang_org_grpc.CallOption) (*github_com_stkr89_livegig_common_pb_orgsvc_v1.DeleteResonse, error)

	lockList sync.Mutex
	ListFunc func(ctx context.Context, in *github_com_stkr89_livegig_common_pb_orgsvc_v1.ListRequest, opts ...google_golang_org_grpc.CallOption) (*github_com_stkr89_livegig_common_pb_orgsvc_v1.ListResponse, error)

	calls struct {
		Create []struct {
			Ctx  context.Context
			In   *github_com_stkr89_livegig_common_pb_orgsvc_v1.CreateRequest
			Opts []google_golang_org_grpc.CallOption
		}
		Delete []struct {
			Ctx  context.Context
			In   *github_com_stkr89_livegig_common_pb_orgsvc_v1.DeleteRequest
			Opts []google_golang_org_grpc.CallOption
		}
		List []struct {
			Ctx  context.Context
			In   *github_com_stkr89_livegig_common_pb_orgsvc_v1.ListRequest
			Opts []google_golang_org_grpc.CallOption
		}
	}
}

// Create mocks base method by wrapping the associated func.
func (m *MockOrgSvcClient) Create(ctx context.Context, in *github_com_stkr89_livegig_common_pb_orgsvc_v1.CreateRequest, opts ...google_golang_org_grpc.CallOption) (*github_com_stkr89_livegig_common_pb_orgsvc_v1.OrgResponse, error) {
	m.lockCreate.Lock()
	defer m.lockCreate.Unlock()

	if m.CreateFunc == nil {
		panic("mocker: MockOrgSvcClient.CreateFunc is nil but MockOrgSvcClient.Create was called.")
	}

	call := struct {
		Ctx  context.Context
		In   *github_com_stkr89_livegig_common_pb_orgsvc_v1.CreateRequest
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
func (m *MockOrgSvcClient) CreateCalled() bool {
	m.lockCreate.Lock()
	defer m.lockCreate.Unlock()

	return len(m.calls.Create) > 0
}

// CreateCalls returns the calls made to Create.
func (m *MockOrgSvcClient) CreateCalls() []struct {
	Ctx  context.Context
	In   *github_com_stkr89_livegig_common_pb_orgsvc_v1.CreateRequest
	Opts []google_golang_org_grpc.CallOption
} {
	m.lockCreate.Lock()
	defer m.lockCreate.Unlock()

	return m.calls.Create
}

// Delete mocks base method by wrapping the associated func.
func (m *MockOrgSvcClient) Delete(ctx context.Context, in *github_com_stkr89_livegig_common_pb_orgsvc_v1.DeleteRequest, opts ...google_golang_org_grpc.CallOption) (*github_com_stkr89_livegig_common_pb_orgsvc_v1.DeleteResonse, error) {
	m.lockDelete.Lock()
	defer m.lockDelete.Unlock()

	if m.DeleteFunc == nil {
		panic("mocker: MockOrgSvcClient.DeleteFunc is nil but MockOrgSvcClient.Delete was called.")
	}

	call := struct {
		Ctx  context.Context
		In   *github_com_stkr89_livegig_common_pb_orgsvc_v1.DeleteRequest
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
func (m *MockOrgSvcClient) DeleteCalled() bool {
	m.lockDelete.Lock()
	defer m.lockDelete.Unlock()

	return len(m.calls.Delete) > 0
}

// DeleteCalls returns the calls made to Delete.
func (m *MockOrgSvcClient) DeleteCalls() []struct {
	Ctx  context.Context
	In   *github_com_stkr89_livegig_common_pb_orgsvc_v1.DeleteRequest
	Opts []google_golang_org_grpc.CallOption
} {
	m.lockDelete.Lock()
	defer m.lockDelete.Unlock()

	return m.calls.Delete
}

// List mocks base method by wrapping the associated func.
func (m *MockOrgSvcClient) List(ctx context.Context, in *github_com_stkr89_livegig_common_pb_orgsvc_v1.ListRequest, opts ...google_golang_org_grpc.CallOption) (*github_com_stkr89_livegig_common_pb_orgsvc_v1.ListResponse, error) {
	m.lockList.Lock()
	defer m.lockList.Unlock()

	if m.ListFunc == nil {
		panic("mocker: MockOrgSvcClient.ListFunc is nil but MockOrgSvcClient.List was called.")
	}

	call := struct {
		Ctx  context.Context
		In   *github_com_stkr89_livegig_common_pb_orgsvc_v1.ListRequest
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
func (m *MockOrgSvcClient) ListCalled() bool {
	m.lockList.Lock()
	defer m.lockList.Unlock()

	return len(m.calls.List) > 0
}

// ListCalls returns the calls made to List.
func (m *MockOrgSvcClient) ListCalls() []struct {
	Ctx  context.Context
	In   *github_com_stkr89_livegig_common_pb_orgsvc_v1.ListRequest
	Opts []google_golang_org_grpc.CallOption
} {
	m.lockList.Lock()
	defer m.lockList.Unlock()

	return m.calls.List
}

// Reset resets the calls made to the mocked methods.
func (m *MockOrgSvcClient) Reset() {
	m.lockCreate.Lock()
	m.calls.Create = nil
	m.lockCreate.Unlock()
	m.lockDelete.Lock()
	m.calls.Delete = nil
	m.lockDelete.Unlock()
	m.lockList.Lock()
	m.calls.List = nil
	m.lockList.Unlock()
}
