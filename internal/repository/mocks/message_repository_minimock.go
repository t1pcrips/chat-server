// Code generated by http://github.com/gojuno/minimock (v3.4.3). DO NOT EDIT.

package mocks

//go:generate minimock -i chat-server/internal/repository.MessageRepository -o message_repository_minimock.go -n MessageRepositoryMock -p mocks

import (
	"github.com/t1pcrips/chat-service/internal/model"
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// MessageRepositoryMock implements mm_repository.MessageRepository
type MessageRepositoryMock struct {
	t          minimock.Tester
	finishOnce sync.Once

	funcCreateMessage          func(ctx context.Context, info *model.CreateMessageRequest) (err error)
	funcCreateMessageOrigin    string
	inspectFuncCreateMessage   func(ctx context.Context, info *model.CreateMessageRequest)
	afterCreateMessageCounter  uint64
	beforeCreateMessageCounter uint64
	CreateMessageMock          mMessageRepositoryMockCreateMessage
}

// NewMessageRepositoryMock returns a mock for mm_repository.MessageRepository
func NewMessageRepositoryMock(t minimock.Tester) *MessageRepositoryMock {
	m := &MessageRepositoryMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.CreateMessageMock = mMessageRepositoryMockCreateMessage{mock: m}
	m.CreateMessageMock.callArgs = []*MessageRepositoryMockCreateMessageParams{}

	t.Cleanup(m.MinimockFinish)

	return m
}

type mMessageRepositoryMockCreateMessage struct {
	optional           bool
	mock               *MessageRepositoryMock
	defaultExpectation *MessageRepositoryMockCreateMessageExpectation
	expectations       []*MessageRepositoryMockCreateMessageExpectation

	callArgs []*MessageRepositoryMockCreateMessageParams
	mutex    sync.RWMutex

	expectedInvocations       uint64
	expectedInvocationsOrigin string
}

// MessageRepositoryMockCreateMessageExpectation specifies expectation struct of the MessageRepository.CreateMessage
type MessageRepositoryMockCreateMessageExpectation struct {
	mock               *MessageRepositoryMock
	params             *MessageRepositoryMockCreateMessageParams
	paramPtrs          *MessageRepositoryMockCreateMessageParamPtrs
	expectationOrigins MessageRepositoryMockCreateMessageExpectationOrigins
	results            *MessageRepositoryMockCreateMessageResults
	returnOrigin       string
	Counter            uint64
}

// MessageRepositoryMockCreateMessageParams contains parameters of the MessageRepository.CreateMessage
type MessageRepositoryMockCreateMessageParams struct {
	ctx  context.Context
	info *model.CreateMessageRequest
}

// MessageRepositoryMockCreateMessageParamPtrs contains pointers to parameters of the MessageRepository.CreateMessage
type MessageRepositoryMockCreateMessageParamPtrs struct {
	ctx  *context.Context
	info **model.CreateMessageRequest
}

// MessageRepositoryMockCreateMessageResults contains results of the MessageRepository.CreateMessage
type MessageRepositoryMockCreateMessageResults struct {
	err error
}

// MessageRepositoryMockCreateMessageOrigins contains origins of expectations of the MessageRepository.CreateMessage
type MessageRepositoryMockCreateMessageExpectationOrigins struct {
	origin     string
	originCtx  string
	originInfo string
}

// Marks this method to be optional. The default behavior of any method with Return() is '1 or more', meaning
// the test will fail minimock's automatic final call check if the mocked method was not called at least once.
// Optional() makes method check to work in '0 or more' mode.
// It is NOT RECOMMENDED to use this option unless you really need it, as default behaviour helps to
// catch the problems when the expected method call is totally skipped during test run.
func (mmCreateMessage *mMessageRepositoryMockCreateMessage) Optional() *mMessageRepositoryMockCreateMessage {
	mmCreateMessage.optional = true
	return mmCreateMessage
}

// Expect sets up expected params for MessageRepository.CreateMessage
func (mmCreateMessage *mMessageRepositoryMockCreateMessage) Expect(ctx context.Context, info *model.CreateMessageRequest) *mMessageRepositoryMockCreateMessage {
	if mmCreateMessage.mock.funcCreateMessage != nil {
		mmCreateMessage.mock.t.Fatalf("MessageRepositoryMock.CreateMessage mock is already set by Set")
	}

	if mmCreateMessage.defaultExpectation == nil {
		mmCreateMessage.defaultExpectation = &MessageRepositoryMockCreateMessageExpectation{}
	}

	if mmCreateMessage.defaultExpectation.paramPtrs != nil {
		mmCreateMessage.mock.t.Fatalf("MessageRepositoryMock.CreateMessage mock is already set by ExpectParams functions")
	}

	mmCreateMessage.defaultExpectation.params = &MessageRepositoryMockCreateMessageParams{ctx, info}
	mmCreateMessage.defaultExpectation.expectationOrigins.origin = minimock.CallerInfo(1)
	for _, e := range mmCreateMessage.expectations {
		if minimock.Equal(e.params, mmCreateMessage.defaultExpectation.params) {
			mmCreateMessage.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmCreateMessage.defaultExpectation.params)
		}
	}

	return mmCreateMessage
}

// ExpectCtxParam1 sets up expected param ctx for MessageRepository.CreateMessage
func (mmCreateMessage *mMessageRepositoryMockCreateMessage) ExpectCtxParam1(ctx context.Context) *mMessageRepositoryMockCreateMessage {
	if mmCreateMessage.mock.funcCreateMessage != nil {
		mmCreateMessage.mock.t.Fatalf("MessageRepositoryMock.CreateMessage mock is already set by Set")
	}

	if mmCreateMessage.defaultExpectation == nil {
		mmCreateMessage.defaultExpectation = &MessageRepositoryMockCreateMessageExpectation{}
	}

	if mmCreateMessage.defaultExpectation.params != nil {
		mmCreateMessage.mock.t.Fatalf("MessageRepositoryMock.CreateMessage mock is already set by Expect")
	}

	if mmCreateMessage.defaultExpectation.paramPtrs == nil {
		mmCreateMessage.defaultExpectation.paramPtrs = &MessageRepositoryMockCreateMessageParamPtrs{}
	}
	mmCreateMessage.defaultExpectation.paramPtrs.ctx = &ctx
	mmCreateMessage.defaultExpectation.expectationOrigins.originCtx = minimock.CallerInfo(1)

	return mmCreateMessage
}

// ExpectInfoParam2 sets up expected param info for MessageRepository.CreateMessage
func (mmCreateMessage *mMessageRepositoryMockCreateMessage) ExpectInfoParam2(info *model.CreateMessageRequest) *mMessageRepositoryMockCreateMessage {
	if mmCreateMessage.mock.funcCreateMessage != nil {
		mmCreateMessage.mock.t.Fatalf("MessageRepositoryMock.CreateMessage mock is already set by Set")
	}

	if mmCreateMessage.defaultExpectation == nil {
		mmCreateMessage.defaultExpectation = &MessageRepositoryMockCreateMessageExpectation{}
	}

	if mmCreateMessage.defaultExpectation.params != nil {
		mmCreateMessage.mock.t.Fatalf("MessageRepositoryMock.CreateMessage mock is already set by Expect")
	}

	if mmCreateMessage.defaultExpectation.paramPtrs == nil {
		mmCreateMessage.defaultExpectation.paramPtrs = &MessageRepositoryMockCreateMessageParamPtrs{}
	}
	mmCreateMessage.defaultExpectation.paramPtrs.info = &info
	mmCreateMessage.defaultExpectation.expectationOrigins.originInfo = minimock.CallerInfo(1)

	return mmCreateMessage
}

// Inspect accepts an inspector function that has same arguments as the MessageRepository.CreateMessage
func (mmCreateMessage *mMessageRepositoryMockCreateMessage) Inspect(f func(ctx context.Context, info *model.CreateMessageRequest)) *mMessageRepositoryMockCreateMessage {
	if mmCreateMessage.mock.inspectFuncCreateMessage != nil {
		mmCreateMessage.mock.t.Fatalf("Inspect function is already set for MessageRepositoryMock.CreateMessage")
	}

	mmCreateMessage.mock.inspectFuncCreateMessage = f

	return mmCreateMessage
}

// Return sets up results that will be returned by MessageRepository.CreateMessage
func (mmCreateMessage *mMessageRepositoryMockCreateMessage) Return(err error) *MessageRepositoryMock {
	if mmCreateMessage.mock.funcCreateMessage != nil {
		mmCreateMessage.mock.t.Fatalf("MessageRepositoryMock.CreateMessage mock is already set by Set")
	}

	if mmCreateMessage.defaultExpectation == nil {
		mmCreateMessage.defaultExpectation = &MessageRepositoryMockCreateMessageExpectation{mock: mmCreateMessage.mock}
	}
	mmCreateMessage.defaultExpectation.results = &MessageRepositoryMockCreateMessageResults{err}
	mmCreateMessage.defaultExpectation.returnOrigin = minimock.CallerInfo(1)
	return mmCreateMessage.mock
}

// Set uses given function f to mock the MessageRepository.CreateMessage method
func (mmCreateMessage *mMessageRepositoryMockCreateMessage) Set(f func(ctx context.Context, info *model.CreateMessageRequest) (err error)) *MessageRepositoryMock {
	if mmCreateMessage.defaultExpectation != nil {
		mmCreateMessage.mock.t.Fatalf("Default expectation is already set for the MessageRepository.CreateMessage method")
	}

	if len(mmCreateMessage.expectations) > 0 {
		mmCreateMessage.mock.t.Fatalf("Some expectations are already set for the MessageRepository.CreateMessage method")
	}

	mmCreateMessage.mock.funcCreateMessage = f
	mmCreateMessage.mock.funcCreateMessageOrigin = minimock.CallerInfo(1)
	return mmCreateMessage.mock
}

// When sets expectation for the MessageRepository.CreateMessage which will trigger the result defined by the following
// Then helper
func (mmCreateMessage *mMessageRepositoryMockCreateMessage) When(ctx context.Context, info *model.CreateMessageRequest) *MessageRepositoryMockCreateMessageExpectation {
	if mmCreateMessage.mock.funcCreateMessage != nil {
		mmCreateMessage.mock.t.Fatalf("MessageRepositoryMock.CreateMessage mock is already set by Set")
	}

	expectation := &MessageRepositoryMockCreateMessageExpectation{
		mock:               mmCreateMessage.mock,
		params:             &MessageRepositoryMockCreateMessageParams{ctx, info},
		expectationOrigins: MessageRepositoryMockCreateMessageExpectationOrigins{origin: minimock.CallerInfo(1)},
	}
	mmCreateMessage.expectations = append(mmCreateMessage.expectations, expectation)
	return expectation
}

// Then sets up MessageRepository.CreateMessage return parameters for the expectation previously defined by the When method
func (e *MessageRepositoryMockCreateMessageExpectation) Then(err error) *MessageRepositoryMock {
	e.results = &MessageRepositoryMockCreateMessageResults{err}
	return e.mock
}

// Times sets number of times MessageRepository.CreateMessage should be invoked
func (mmCreateMessage *mMessageRepositoryMockCreateMessage) Times(n uint64) *mMessageRepositoryMockCreateMessage {
	if n == 0 {
		mmCreateMessage.mock.t.Fatalf("Times of MessageRepositoryMock.CreateMessage mock can not be zero")
	}
	mm_atomic.StoreUint64(&mmCreateMessage.expectedInvocations, n)
	mmCreateMessage.expectedInvocationsOrigin = minimock.CallerInfo(1)
	return mmCreateMessage
}

func (mmCreateMessage *mMessageRepositoryMockCreateMessage) invocationsDone() bool {
	if len(mmCreateMessage.expectations) == 0 && mmCreateMessage.defaultExpectation == nil && mmCreateMessage.mock.funcCreateMessage == nil {
		return true
	}

	totalInvocations := mm_atomic.LoadUint64(&mmCreateMessage.mock.afterCreateMessageCounter)
	expectedInvocations := mm_atomic.LoadUint64(&mmCreateMessage.expectedInvocations)

	return totalInvocations > 0 && (expectedInvocations == 0 || expectedInvocations == totalInvocations)
}

// CreateMessage implements mm_repository.MessageRepository
func (mmCreateMessage *MessageRepositoryMock) CreateMessage(ctx context.Context, info *model.CreateMessageRequest) (err error) {
	mm_atomic.AddUint64(&mmCreateMessage.beforeCreateMessageCounter, 1)
	defer mm_atomic.AddUint64(&mmCreateMessage.afterCreateMessageCounter, 1)

	mmCreateMessage.t.Helper()

	if mmCreateMessage.inspectFuncCreateMessage != nil {
		mmCreateMessage.inspectFuncCreateMessage(ctx, info)
	}

	mm_params := MessageRepositoryMockCreateMessageParams{ctx, info}

	// Record call args
	mmCreateMessage.CreateMessageMock.mutex.Lock()
	mmCreateMessage.CreateMessageMock.callArgs = append(mmCreateMessage.CreateMessageMock.callArgs, &mm_params)
	mmCreateMessage.CreateMessageMock.mutex.Unlock()

	for _, e := range mmCreateMessage.CreateMessageMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmCreateMessage.CreateMessageMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmCreateMessage.CreateMessageMock.defaultExpectation.Counter, 1)
		mm_want := mmCreateMessage.CreateMessageMock.defaultExpectation.params
		mm_want_ptrs := mmCreateMessage.CreateMessageMock.defaultExpectation.paramPtrs

		mm_got := MessageRepositoryMockCreateMessageParams{ctx, info}

		if mm_want_ptrs != nil {

			if mm_want_ptrs.ctx != nil && !minimock.Equal(*mm_want_ptrs.ctx, mm_got.ctx) {
				mmCreateMessage.t.Errorf("MessageRepositoryMock.CreateMessage got unexpected parameter ctx, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmCreateMessage.CreateMessageMock.defaultExpectation.expectationOrigins.originCtx, *mm_want_ptrs.ctx, mm_got.ctx, minimock.Diff(*mm_want_ptrs.ctx, mm_got.ctx))
			}

			if mm_want_ptrs.info != nil && !minimock.Equal(*mm_want_ptrs.info, mm_got.info) {
				mmCreateMessage.t.Errorf("MessageRepositoryMock.CreateMessage got unexpected parameter info, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmCreateMessage.CreateMessageMock.defaultExpectation.expectationOrigins.originInfo, *mm_want_ptrs.info, mm_got.info, minimock.Diff(*mm_want_ptrs.info, mm_got.info))
			}

		} else if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmCreateMessage.t.Errorf("MessageRepositoryMock.CreateMessage got unexpected parameters, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
				mmCreateMessage.CreateMessageMock.defaultExpectation.expectationOrigins.origin, *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmCreateMessage.CreateMessageMock.defaultExpectation.results
		if mm_results == nil {
			mmCreateMessage.t.Fatal("No results are set for the MessageRepositoryMock.CreateMessage")
		}
		return (*mm_results).err
	}
	if mmCreateMessage.funcCreateMessage != nil {
		return mmCreateMessage.funcCreateMessage(ctx, info)
	}
	mmCreateMessage.t.Fatalf("Unexpected call to MessageRepositoryMock.CreateMessage. %v %v", ctx, info)
	return
}

// CreateMessageAfterCounter returns a count of finished MessageRepositoryMock.CreateMessage invocations
func (mmCreateMessage *MessageRepositoryMock) CreateMessageAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCreateMessage.afterCreateMessageCounter)
}

// CreateMessageBeforeCounter returns a count of MessageRepositoryMock.CreateMessage invocations
func (mmCreateMessage *MessageRepositoryMock) CreateMessageBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCreateMessage.beforeCreateMessageCounter)
}

// Calls returns a list of arguments used in each call to MessageRepositoryMock.CreateMessage.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmCreateMessage *mMessageRepositoryMockCreateMessage) Calls() []*MessageRepositoryMockCreateMessageParams {
	mmCreateMessage.mutex.RLock()

	argCopy := make([]*MessageRepositoryMockCreateMessageParams, len(mmCreateMessage.callArgs))
	copy(argCopy, mmCreateMessage.callArgs)

	mmCreateMessage.mutex.RUnlock()

	return argCopy
}

// MinimockCreateMessageDone returns true if the count of the CreateMessage invocations corresponds
// the number of defined expectations
func (m *MessageRepositoryMock) MinimockCreateMessageDone() bool {
	if m.CreateMessageMock.optional {
		// Optional methods provide '0 or more' call count restriction.
		return true
	}

	for _, e := range m.CreateMessageMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	return m.CreateMessageMock.invocationsDone()
}

// MinimockCreateMessageInspect logs each unmet expectation
func (m *MessageRepositoryMock) MinimockCreateMessageInspect() {
	for _, e := range m.CreateMessageMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to MessageRepositoryMock.CreateMessage at\n%s with params: %#v", e.expectationOrigins.origin, *e.params)
		}
	}

	afterCreateMessageCounter := mm_atomic.LoadUint64(&m.afterCreateMessageCounter)
	// if default expectation was set then invocations count should be greater than zero
	if m.CreateMessageMock.defaultExpectation != nil && afterCreateMessageCounter < 1 {
		if m.CreateMessageMock.defaultExpectation.params == nil {
			m.t.Errorf("Expected call to MessageRepositoryMock.CreateMessage at\n%s", m.CreateMessageMock.defaultExpectation.returnOrigin)
		} else {
			m.t.Errorf("Expected call to MessageRepositoryMock.CreateMessage at\n%s with params: %#v", m.CreateMessageMock.defaultExpectation.expectationOrigins.origin, *m.CreateMessageMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCreateMessage != nil && afterCreateMessageCounter < 1 {
		m.t.Errorf("Expected call to MessageRepositoryMock.CreateMessage at\n%s", m.funcCreateMessageOrigin)
	}

	if !m.CreateMessageMock.invocationsDone() && afterCreateMessageCounter > 0 {
		m.t.Errorf("Expected %d calls to MessageRepositoryMock.CreateMessage at\n%s but found %d calls",
			mm_atomic.LoadUint64(&m.CreateMessageMock.expectedInvocations), m.CreateMessageMock.expectedInvocationsOrigin, afterCreateMessageCounter)
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *MessageRepositoryMock) MinimockFinish() {
	m.finishOnce.Do(func() {
		if !m.minimockDone() {
			m.MinimockCreateMessageInspect()
		}
	})
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *MessageRepositoryMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *MessageRepositoryMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockCreateMessageDone()
}
