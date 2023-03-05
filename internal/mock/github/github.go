// Code generated by MockGen. DO NOT EDIT.
// Source: github.go

// Package github is a generated GoMock package.
package github

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	github "github.com/google/go-github/v50/github"
	githubactions "github.com/sethvargo/go-githubactions"
)

// MockAction is a mock of Action interface.
type MockAction struct {
	ctrl     *gomock.Controller
	recorder *MockActionMockRecorder
}

// MockActionMockRecorder is the mock recorder for MockAction.
type MockActionMockRecorder struct {
	mock *MockAction
}

// NewMockAction creates a new mock instance.
func NewMockAction(ctrl *gomock.Controller) *MockAction {
	mock := &MockAction{ctrl: ctrl}
	mock.recorder = &MockActionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAction) EXPECT() *MockActionMockRecorder {
	return m.recorder
}

// Context mocks base method.
func (m *MockAction) Context() (*githubactions.GitHubContext, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(*githubactions.GitHubContext)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Context indicates an expected call of Context.
func (mr *MockActionMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockAction)(nil).Context))
}

// GetInput mocks base method.
func (m *MockAction) GetInput(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInput", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetInput indicates an expected call of GetInput.
func (mr *MockActionMockRecorder) GetInput(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInput", reflect.TypeOf((*MockAction)(nil).GetInput), arg0)
}

// MockPullRequestService is a mock of PullRequestService interface.
type MockPullRequestService struct {
	ctrl     *gomock.Controller
	recorder *MockPullRequestServiceMockRecorder
}

// MockPullRequestServiceMockRecorder is the mock recorder for MockPullRequestService.
type MockPullRequestServiceMockRecorder struct {
	mock *MockPullRequestService
}

// NewMockPullRequestService creates a new mock instance.
func NewMockPullRequestService(ctrl *gomock.Controller) *MockPullRequestService {
	mock := &MockPullRequestService{ctrl: ctrl}
	mock.recorder = &MockPullRequestServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPullRequestService) EXPECT() *MockPullRequestServiceMockRecorder {
	return m.recorder
}

// CreateComment mocks base method.
func (m *MockPullRequestService) CreateComment(ctx context.Context, owner, repo string, number int, comment *github.PullRequestComment) (*github.PullRequestComment, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateComment", ctx, owner, repo, number, comment)
	ret0, _ := ret[0].(*github.PullRequestComment)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateComment indicates an expected call of CreateComment.
func (mr *MockPullRequestServiceMockRecorder) CreateComment(ctx, owner, repo, number, comment interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateComment", reflect.TypeOf((*MockPullRequestService)(nil).CreateComment), ctx, owner, repo, number, comment)
}

// MockIssueService is a mock of IssueService interface.
type MockIssueService struct {
	ctrl     *gomock.Controller
	recorder *MockIssueServiceMockRecorder
}

// MockIssueServiceMockRecorder is the mock recorder for MockIssueService.
type MockIssueServiceMockRecorder struct {
	mock *MockIssueService
}

// NewMockIssueService creates a new mock instance.
func NewMockIssueService(ctrl *gomock.Controller) *MockIssueService {
	mock := &MockIssueService{ctrl: ctrl}
	mock.recorder = &MockIssueServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIssueService) EXPECT() *MockIssueServiceMockRecorder {
	return m.recorder
}

// CreateComment mocks base method.
func (m *MockIssueService) CreateComment(ctx context.Context, owner, repo string, number int, comment *github.IssueComment) (*github.IssueComment, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateComment", ctx, owner, repo, number, comment)
	ret0, _ := ret[0].(*github.IssueComment)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateComment indicates an expected call of CreateComment.
func (mr *MockIssueServiceMockRecorder) CreateComment(ctx, owner, repo, number, comment interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateComment", reflect.TypeOf((*MockIssueService)(nil).CreateComment), ctx, owner, repo, number, comment)
}

// EditComment mocks base method.
func (m *MockIssueService) EditComment(ctx context.Context, owner, repo string, commentID int64, comment *github.IssueComment) (*github.IssueComment, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EditComment", ctx, owner, repo, commentID, comment)
	ret0, _ := ret[0].(*github.IssueComment)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// EditComment indicates an expected call of EditComment.
func (mr *MockIssueServiceMockRecorder) EditComment(ctx, owner, repo, commentID, comment interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditComment", reflect.TypeOf((*MockIssueService)(nil).EditComment), ctx, owner, repo, commentID, comment)
}

// ListComments mocks base method.
func (m *MockIssueService) ListComments(ctx context.Context, owner, repo string, number int, opts *github.IssueListCommentsOptions) ([]*github.IssueComment, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListComments", ctx, owner, repo, number, opts)
	ret0, _ := ret[0].([]*github.IssueComment)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListComments indicates an expected call of ListComments.
func (mr *MockIssueServiceMockRecorder) ListComments(ctx, owner, repo, number, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListComments", reflect.TypeOf((*MockIssueService)(nil).ListComments), ctx, owner, repo, number, opts)
}

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// CreateCommentFromOutput mocks base method.
func (m *MockClient) CreateCommentFromOutput(ctx context.Context, planOutput []string, path string) (*github.IssueComment, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCommentFromOutput", ctx, planOutput, path)
	ret0, _ := ret[0].(*github.IssueComment)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateCommentFromOutput indicates an expected call of CreateCommentFromOutput.
func (mr *MockClientMockRecorder) CreateCommentFromOutput(ctx, planOutput, path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCommentFromOutput", reflect.TypeOf((*MockClient)(nil).CreateCommentFromOutput), ctx, planOutput, path)
}
