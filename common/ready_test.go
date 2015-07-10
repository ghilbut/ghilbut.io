package common_test

import (
	. "github.com/ghilbut/ygg.go/common"
	. "github.com/ghilbut/ygg.go/test/fake"
	"testing"
)

//const kJson = "{ \"id\": \"A\", \"endpoint\": \"B\" }"

func Test_has_connection_after_set_connection(t *testing.T) {

	var conn Connection = NewFakeConnection()

	ready := NewCtrlReady()

	if ready.HasConnection(conn) {
		t.Fail()
	}

	ready.SetConnection(conn)
	if !ready.HasConnection(conn) {
		t.Fail()
	}
}

func Test_clear_connection(t *testing.T) {

	var conn0 Connection = NewFakeConnection()
	var conn1 Connection = NewFakeConnection()
	var conn2 Connection = NewFakeConnection()

	ready := NewCtrlReady()
	ready.SetConnection(conn0)
	ready.SetConnection(conn1)
	ready.SetConnection(conn2)

	if !ready.HasConnection(conn0) ||
		!ready.HasConnection(conn1) ||
		!ready.HasConnection(conn2) {
		t.Fail()
	}

	ready.Clear()

	if ready.HasConnection(conn0) ||
		ready.HasConnection(conn1) ||
		ready.HasConnection(conn2) {
		t.Fail()
	}
}

func Test_remove_connection_when_it_is_closed(t *testing.T) {

	var conn Connection = NewFakeConnection()

	ready := NewCtrlReady()
	ready.SetConnection(conn)

	conn.Close()
	if ready.HasConnection(conn) {
		t.Fail()
	}
}

func Test_remove_connection_when_invalid_json_is_passed(t *testing.T) {

	var conn Connection = NewFakeConnection()

	ready := NewCtrlReady()
	ready.OnCtrlReadyProc = func(proxy *CtrlProxy) {
		t.Fail()
	}

	ready.SetConnection(conn)
	ready.OnText(conn, "")
	if ready.HasConnection(conn) {
		t.Fail()
	}

	ready.SetConnection(conn)
	ready.OnText(conn, "{ \"key\": \"value")
	if ready.HasConnection(conn) {
		t.Fail()
	}

	ready.SetConnection(conn)
	ready.OnText(conn, "{}")
	if ready.HasConnection(conn) {
		t.Fail()
	}
}

func Test_ok(t *testing.T) {

	var lhs Connection = NewFakeConnection()
	var rhs Connection = lhs.(*FakeConnection).Other()

	ready := NewCtrlReady()
	ready.OnCtrlReadyProc = func(proxy *CtrlProxy) {
		if proxy == nil {
			t.Fail()
		}
		if proxy.Desc.CtrlId != "A" {
			t.Fail()
		}
		if proxy.Desc.Endpoint != "B" {
			t.Fail()
		}
	}

	ready.SetConnection(rhs)
	lhs.SendText(kJson)
}
