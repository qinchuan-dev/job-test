package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"job-test/pg"
	"job-test/proto"
	"job-test/types"
	"log"
	"math/big"
	"net"
	"net/http"
)

const (
	FlagGrpcPort    = "grpcPort"
	FlagGatewayPort = "gatewayPort"
)

func DoDepositCmd(ctx context.Context, userId, amount, denom, memo string) error {
	db, err := pg.NewPG(ctx, DefaultPostgresConStr)
	if err != nil {
		return err
	}

	amt, ok := new(big.Int).SetString(amount, 0)
	if !ok {
		return fmt.Errorf("invalid old amount %s", amount)
	}
	err = db.Deposit(ctx, userId, denom, *amt)
	if err != nil {
		return err
	}

	err = db.InsertDepositHistory(ctx, userId, denom, *amt, types.DEPOSIT, memo)
	if err != nil {
		return err
	}

	return nil
}

func DoWithdrawCmd(ctx context.Context, userId string, amount, denom, memo string) error {
	db, err := pg.NewPG(ctx, DefaultPostgresConStr)
	if err != nil {
		return err
	}

	amt, ok := new(big.Int).SetString(amount, 0)
	if !ok {
		return fmt.Errorf("invalid old amount %s", amount)
	}
	err = db.Withdraw(ctx, userId, denom, *amt)
	if err != nil {
		return err
	}

	err = db.InsertDepositHistory(ctx, userId, denom, *amt, types.WITHDRAW, memo)
	if err != nil {
		return err
	}

	return nil
}

func DoMsgSendCmd(ctx context.Context, sender, receiver, amount, denom, memo string) error {
	db, err := pg.NewPG(ctx, DefaultPostgresConStr)
	if err != nil {
		return err
	}

	amt, ok := new(big.Int).SetString(amount, 0)
	if !ok {
		return fmt.Errorf("invalid old amount %s", amount)
	}
	err = db.Withdraw(ctx, sender, denom, *amt)
	if err != nil {
		return err
	}

	err = db.Deposit(ctx, receiver, denom, *amt)
	if err != nil {
		return err
	}

	err = db.InsertSendHistory(ctx, sender, receiver, denom, *amt, memo)
	if err != nil {
		return err
	}

	return nil
}

func DoQueryBalanceCmd(ctx context.Context, userId string) ([]types.DepositItem, error) {
	db, err := pg.NewPG(ctx, DefaultPostgresConStr)
	if err != nil {
		return nil, err
	}

	return db.GetDepositByCustomer(ctx, userId)
}

func DoQueryHistoryCmd(ctx context.Context, userId string) ([]types.DepositHistoryItem, error) {
	db, err := pg.NewPG(ctx, DefaultPostgresConStr)
	if err != nil {
		return nil, err
	}

	return db.GetDepositHistoryByCustomer(ctx, userId)
}

func StartCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "Run the job-test process",
		Long:  `Run the job-test process`,
		RunE: func(cmd *cobra.Command, _ []string) error {
			grpcPort, _ := cmd.Flags().GetString(FlagGrpcPort)
			gatewayPort, _ := cmd.Flags().GetString(FlagGatewayPort)

			grpcEndpoint := "localhost" + grpcPort

			// grpc server
			lis, err := net.Listen("tcp", grpcPort)
			if err != nil {
				log.Fatalln("Failed to listen:", err)
			}

			s := grpc.NewServer()
			proto.RegisterMsgServer(s, &MsgServer{})
			reflection.Register(s)
			log.Println("Serving gRPC on localhost" + grpcPort)
			go func() {
				log.Fatalln(s.Serve(lis))
			}()

			//gateway
			ctx := context.Background()
			ctx, cancel := context.WithCancel(ctx)
			defer cancel()

			mux := runtime.NewServeMux()
			opts := []grpc.DialOption{grpc.WithInsecure()}
			err = proto.RegisterMsgHandlerFromEndpoint(ctx, mux, grpcEndpoint, opts)
			if err != nil {
				log.Fatalln("Failed to register gateway:", err)
			}

			log.Println("Serving gRPC-Gateway on http://0.0.0.0" + gatewayPort)
			if err := http.ListenAndServe(gatewayPort, mux); err != nil {
				log.Fatalf("Could not setup HTTP endpoint: %v", err)
			}
			return nil
		},
	}
	fs := pflag.NewFlagSet("", pflag.ContinueOnError)

	fs.String(FlagGrpcPort, ":50051", "The grpc port")
	fs.String(FlagGatewayPort, ":8081", "the http gateway port")
	cmd.Flags().AddFlagSet(fs)

	return cmd
}
