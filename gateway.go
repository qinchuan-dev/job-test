package main

import (
	"context"
	"job-test/proto"
)

type MsgServer struct {
	proto.UnimplementedMsgServer
}

func (s *MsgServer) Deposit(ctx context.Context, m *proto.MsgDeposit) (*proto.MsgDepositResp, error) {
	err := DoDepositCmd(ctx, m.Userid, m.Amt, m.Denom, m.Memo)
	if err != nil {
		return nil, err
	} else {
		return &proto.MsgDepositResp{}, nil
	}
}

func (s *MsgServer) Withdraw(ctx context.Context, m *proto.MsgWithdraw) (*proto.MsgWithdrawResp, error) {
	err := DoWithdrawCmd(ctx, m.Userid, m.Amt, m.Denom, m.Memo)
	if err != nil {
		return nil, err
	} else {
		return &proto.MsgWithdrawResp{}, nil
	}
}

func (s *MsgServer) Send(ctx context.Context, m *proto.MsgSend) (*proto.MsgSendResp, error) {
	err := DoMsgSendCmd(ctx, m.From, m.To, m.Amt, m.Denom, m.Memo)
	if err != nil {
		return nil, err
	} else {
		return &proto.MsgSendResp{}, nil
	}
}

type QueryServer struct {
	proto.UnimplementedMsgServer
}

func (s *QueryServer) QueryBalance(ctx context.Context, m *proto.QueryBalance) (*proto.QueryBalanceResp, error) {
	_, err := DoQueryBalanceCmd(ctx, m.Userid)
	if err != nil {
		return nil, err
	}
	return &proto.QueryBalanceResp{}, nil
}

func (s *QueryServer) QueryDepositHistory(ctx context.Context, m *proto.QueryDepositHistory) (*proto.QueryDepositHistoryResp, error) {
	_, err := DoQueryDepositHistoryCmd(ctx, m.Userid)
	if err != nil {
		return nil, err
	} else {
		return &proto.QueryDepositHistoryResp{}, nil
	}
}

func (s *QueryServer) QuerySendHistory(ctx context.Context, m *proto.QuerySendHistory) (*proto.QuerySendHistoryResp, error) {
	_, err := DoQuerySendHistoryCmd(ctx, m.Userid)
	if err != nil {
		return nil, err
	} else {
		return &proto.QuerySendHistoryResp{}, nil
	}
}
