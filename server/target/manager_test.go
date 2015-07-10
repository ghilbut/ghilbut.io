package target_test

import (
	. "github.com/ghilbut/ygg.go/common"
	. "github.com/ghilbut/ygg.go/server/target"
	. "github.com/ghilbut/ygg.go/test/fake"
	. "github.com/ghilbut/ygg.go/test/mock/common"
	"github.com/golang/mock/gomock"
	"log"
	"testing"
)

// const kCtrlA0Json = "{ \"id\": \"A0\", \"endpoint\": \"B\" }"
// const kCtrlA1Json = "{ \"id\": \"A1\", \"endpoint\": \"B\" }"
// const kTargetJson = "{ \"endpoint\": \"B\" }"

// const kText = "Message"

// var bytes = []byte{0x01, 0x02}

func Test_Manager_notify_text(t *testing.T) {
	log.Println("======== [Test_Manager_notify_text] ========")

	var ctrlA0 Connection = NewFakeConnection()
	var ctrlA1 Connection = NewFakeConnection()
	var target Connection = NewFakeConnection()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDelegate := NewMockConnectionDelegate(mockCtrl)
	mockDelegate.EXPECT().OnText(ctrlA0, kText)
	mockDelegate.EXPECT().OnText(ctrlA1, kText)
	ctrlA0.BindDelegate(mockDelegate)
	ctrlA1.BindDelegate(mockDelegate)

	manager := NewManager()
	manager.SetTargetConnection(target.(*FakeConnection).Other())
	target.SendText(kTargetJson)

	manager.SetCtrlConnection(ctrlA0.(*FakeConnection).Other())
	ctrlA0.SendText(kCtrlA0Json)
	manager.SetCtrlConnection(ctrlA1.(*FakeConnection).Other())
	ctrlA1.SendText(kCtrlA1Json)

	target.SendText(kText)
}

func Test_Manager_recv_text(t *testing.T) {
	log.Println("======== [Test_Manager_recv_text] ========")

	var ctrlA0 Connection = NewFakeConnection()
	var ctrlA1 Connection = NewFakeConnection()
	var target Connection = NewFakeConnection()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDelegate := NewMockConnectionDelegate(mockCtrl)
	mockDelegate.EXPECT().OnText(target, kText).Times(2)
	target.BindDelegate(mockDelegate)

	manager := NewManager()
	manager.SetTargetConnection(target.(*FakeConnection).Other())
	target.SendText(kTargetJson)

	manager.SetCtrlConnection(ctrlA0.(*FakeConnection).Other())
	ctrlA0.SendText(kCtrlA0Json)
	manager.SetCtrlConnection(ctrlA1.(*FakeConnection).Other())
	ctrlA1.SendText(kCtrlA1Json)

	ctrlA0.SendText(kText)
	ctrlA1.SendText(kText)
}