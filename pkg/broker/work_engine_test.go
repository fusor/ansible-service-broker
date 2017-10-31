package broker

import (
	"sync"
	"testing"

	ft "github.com/openshift/ansible-service-broker/pkg/fusortest"
)

var engine *WorkEngine

func init() {
	engine = NewWorkEngine(10)
}

type mockSubscriber struct {
	buffer <-chan WorkMsg
	called bool
}

func (ms *mockSubscriber) Subscribe(buffer <-chan WorkMsg) {
	ms.buffer = buffer
	ms.called = true
}

type mockMsg struct {
	msg string
}

func (mm mockMsg) Render() string {
	return mm.msg
}

type mockWorker struct {
	called bool
	wg     *sync.WaitGroup
}

func (mw *mockWorker) Run(token string, buffer chan<- WorkMsg) {
	mw.called = true
	buffer <- mockMsg{msg: "hello"}
	mw.wg.Done()
}

func TestNewWorkEngine(t *testing.T) {
	we := NewWorkEngine(10)
	ft.AssertNotNil(t, we)
	ft.AssertEqual(t, we.bufsz, 10)
}

func TestGetActiveTopics(t *testing.T) {
	topics := engine.GetActiveTopics()
	ft.AssertEqual(t, 0, len(topics))
	dasub := mockSubscriber{}
	engine.AttachSubscriber(&dasub, ProvisionTopic)

	// ensure topic is added and buffer passed to subscriber
	topics = engine.GetActiveTopics()
	ft.AssertEqual(t, 1, len(topics))
	_, exists := topics[ProvisionTopic]
	ft.AssertTrue(t, exists, "topic does not exist")
}

func TestAttachSubscriber(t *testing.T) {
	dasub := mockSubscriber{}
	err := engine.AttachSubscriber(&dasub, ProvisionTopic)
	if err != nil {
		t.Fatal(err)
	}
	topics := engine.GetActiveTopics()
	_, exists := topics[ProvisionTopic]
	ft.AssertTrue(t, exists, "topic does not exist")
	ft.AssertTrue(t, dasub.called, "subscribe never called")
}

func TestInvalidWorkTopic(t *testing.T) {
	var faketopic WorkTopic
	faketopic = "fake"
	dasub := mockSubscriber{}
	err := engine.AttachSubscriber(&dasub, faketopic)
	if err == nil {
		t.Fail()
	}
	ft.AssertEqual(t, "invalid work topic", err.Error(), "invalid error")
}

func TestStartNewJob(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1) // we're launching 1 goroutine

	// work around to get pointer receivers to update the worker
	var work Work
	worker := &mockWorker{wg: &wg}
	work = worker

	token, err := engine.StartNewJob("testtoken", work, ProvisionTopic)
	ft.AssertNil(t, err)
	ft.AssertEqual(t, "testtoken", token, "token doesn't match")

	// let's wait until it's done
	wg.Wait()

	// verify we actually called the run method
	ft.AssertTrue(t, worker.called, "run not called")
}
