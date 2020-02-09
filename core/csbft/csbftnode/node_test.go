package csbftnode

import (
	"testing"
	"github.com/golang/mock/gomock"
	"github.com/dipperin/dipperin-core/core/csbft/statemachine"
	"github.com/dipperin/dipperin-core/common"
	"github.com/dipperin/dipperin-core/core/model"
)

func testReady(t *testing.T) *CsBft  {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	chain := NewMockChainReader(ctrl)
	fetcher := NewMockFetcher(ctrl)
	signer := NewMockMsgSigner(ctrl)
	sender := NewMockMsgSender(ctrl)
	validator := NewMockValidator(ctrl)
	block := NewMockAbstractBlock(ctrl)
	verification :=NewMockAbstractVerification(ctrl)
	var config = &statemachine.BftConfig{
		ChainReader: chain,
		Fetcher:     fetcher,
		Signer:      signer,
		Sender:      sender,
		Validator:   validator,
	}
	var height uint64
	height = 1
	block.EXPECT().Number().Return(height).AnyTimes()
	chain.EXPECT().CurrentBlock().Return(block).AnyTimes()
	chain.EXPECT().GetSeenCommit(height).Return([]model.AbstractVerification{verification}).AnyTimes()
	verification.EXPECT().GetRound().Return(uint64(2333)).AnyTimes()
	chain.EXPECT().IsChangePoint(gomock.Any(), gomock.Any()).Return(true).AnyTimes()
	chain.EXPECT().GetNextVerifiers().Return([]common.Address{}).AnyTimes()

	return NewCsBft(config)
}

func TestNewCsBft(t *testing.T) {
	if c :=testReady(t);c==nil{
		t.Error("fail to NewCsBft")
	}
}


