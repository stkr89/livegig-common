// Code generated by mocker. DO NOT EDIT.
// github.com/travisjeffery/mocker
// Source: pb/casesvc/v1/casesvc.pb.go

package mock

import (
	context "context"
	sync "sync"

	github_com_stkr89_livegig_common_pb_casesvc_v1 "github.com/stkr89/livegig-common/pb/casesvc/v1"
	google_golang_org_grpc "google.golang.org/grpc"
)

// MockCaseSvcClient is a mock of CaseSvcClient interface
type MockCaseSvcClient struct {
	lockCreate sync.Mutex
	CreateFunc func(ctx context.Context, in *github_com_stkr89_livegig_common_pb_casesvc_v1.CreateRequest, opts ...google_golang_org_grpc.CallOption) (*github_com_stkr89_livegig_common_pb_casesvc_v1.CaseResponse, error)

	calls struct {
		Create []struct {
			Ctx  context.Context
			In   *github_com_stkr89_livegig_common_pb_casesvc_v1.CreateRequest
			Opts []google_golang_org_grpc.CallOption
		}
	}
}

// Create mocks base method by wrapping the associated func.
func (m *MockCaseSvcClient) Create(ctx context.Context, in *github_com_stkr89_livegig_common_pb_casesvc_v1.CreateRequest, opts ...google_golang_org_grpc.CallOption) (*github_com_stkr89_livegig_common_pb_casesvc_v1.CaseResponse, error) {
	m.lockCreate.Lock()
	defer m.lockCreate.Unlock()

	if m.CreateFunc == nil {
		panic("mocker: MockCaseSvcClient.CreateFunc is nil but MockCaseSvcClient.Create was called.")
	}

	call := struct {
		Ctx  context.Context
		In   *github_com_stkr89_livegig_common_pb_casesvc_v1.CreateRequest
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
func (m *MockCaseSvcClient) CreateCalled() bool {
	m.lockCreate.Lock()
	defer m.lockCreate.Unlock()

	return len(m.calls.Create) > 0
}

// CreateCalls returns the calls made to Create.
func (m *MockCaseSvcClient) CreateCalls() []struct {
	Ctx  context.Context
	In   *github_com_stkr89_livegig_common_pb_casesvc_v1.CreateRequest
	Opts []google_golang_org_grpc.CallOption
} {
	m.lockCreate.Lock()
	defer m.lockCreate.Unlock()

	return m.calls.Create
}

// Reset resets the calls made to the mocked methods.
func (m *MockCaseSvcClient) Reset() {
	m.lockCreate.Lock()
	m.calls.Create = nil
	m.lockCreate.Unlock()
}
