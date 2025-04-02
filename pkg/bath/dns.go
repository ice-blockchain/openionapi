package bath

import (
	"github.com/ice-blockchain/iongo/abi"
	"github.com/ice-blockchain/iongo/tlb"
	"github.com/ice-blockchain/iongo/ton"
)

type BubbleDnsItemRenew struct {
	DnsRenewAction
	Success bool
}

type DnsRenewAction struct {
	Item    ton.AccountID
	Renewer ton.AccountID
}

func (b BubbleDnsItemRenew) ToAction() *Action {
	return &Action{Success: b.Success, Type: DomainRenew, DnsRenew: &b.DnsRenewAction}
}

func (a DnsRenewAction) SubjectAccounts() []ton.AccountID {
	return []ton.AccountID{a.Renewer, a.Item}
}

var DNSRenewStraw = Straw[BubbleDnsItemRenew]{
	CheckFuncs: []bubbleCheck{IsTx, HasOperation(abi.DeleteDnsRecordMsgOp), HasInterface(abi.NftItem), func(bubble *Bubble) bool {
		return bubble.Info.(BubbleTx).decodedBody.Value.(abi.DeleteDnsRecordMsgBody).Key.Equal(tlb.Bits256{})
	}},
	Builder: func(newAction *BubbleDnsItemRenew, bubble *Bubble) error {
		tx := bubble.Info.(BubbleTx)
		newAction.Renewer = tx.inputFrom.Address
		newAction.Item = tx.account.Address
		newAction.Success = tx.success
		return nil
	},
	SingleChild: &Straw[BubbleDnsItemRenew]{
		Optional:   true,
		CheckFuncs: []bubbleCheck{IsTx, HasOperation(abi.BounceMsgOp)},
	},
}
