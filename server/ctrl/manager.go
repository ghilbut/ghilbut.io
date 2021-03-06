package ctrl

import (
	. "github.com/ghilbut/ygg.go/common"
	"github.com/ghilbut/ygg.go/debug"
	"log"
)

type CtrlManager struct {
	CtrlBridgeDelegate
	CtrlReadyDelegate
	TargetReadyDelegate

	bridge      CtrlBridge
	ctrlReady   *CtrlReady
	targetReady *TargetReady
	adapters    map[Adapter]bool
}

func NewCtrlManager() *CtrlManager {
	log.Println("======== [CtrlManager][NewCtrlManager] ========")

	manager := &CtrlManager{
		bridge:      nil,
		ctrlReady:   NewCtrlReady(),
		targetReady: NewTargetReady(),
		adapters:    make(map[Adapter]bool),
	}

	manager.ctrlReady.Delegate = manager
	manager.targetReady.Delegate = manager
	return manager
}

func (self *CtrlManager) OnCtrlBridgeStarted(bridge CtrlBridge) {
	log.Println("======== [CtrlManager][OnCtrlBridgeStarted] ========")
	assert.True(bridge != nil)

	self.bridge = bridge
}

func (self *CtrlManager) OnCtrlBridgeStopped() {
	log.Println("======== [CtrlManager][OnCtrlBridgeStopped] ========")
	assert.True(self.bridge != nil)

	self.ctrlReady.Clear()
	self.targetReady.Clear()

	for adapter, _ := range self.adapters {
		adapter.Close()
	}

	assert.True(len(self.adapters) == 0)

	self.bridge = nil

}

func (self *CtrlManager) HasAdapter(adapter Adapter) bool {
	log.Println("======== [CtrlManager][HasAdapter] ========")
	assert.True(adapter != nil)

	_, ok := self.adapters[adapter]
	return ok
}

func (self *CtrlManager) OnCtrlConnected(conn Connection) {
	log.Println("======== [CtrlManager][OnCtrlConnected] ========")

	assert.True(conn != nil)
	self.ctrlReady.SetConnection(conn)
}

func (self *CtrlManager) OnCtrlProxy(proxy *CtrlProxy) {
	log.Println("======== [CtrlManager][OnCtrlProxy] ========")

	conn := self.bridge.Connect(proxy)
	if conn != nil {
		self.targetReady.SetConnection(conn, proxy)
	} else {
		proxy.Close()
	}
}

func (self *CtrlManager) OnTargetProxy(ctrl *CtrlProxy, target *TargetProxy) {
	log.Println("======== [CtrlManager][OnTargetProxy] ========")

	adapter := NewOneToOneAdapter(target)
	if adapter == nil {
		ctrl.Close()
		target.Close()
		return
	}

	adapter.SetCtrlProxy(ctrl)

	self.adapters[adapter] = true
	adapter.BindDelegate(self)
}

func (self *CtrlManager) OnAdapterClosed(adapter Adapter) {
	log.Println("======== [CtrlManager][OnAdapterClosed] ========")
	assert.True(self.HasAdapter(adapter))

	adapter.UnbindDelegate()
	delete(self.adapters, adapter)
}
