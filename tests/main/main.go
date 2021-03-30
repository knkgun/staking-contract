package main

import "Zilliqa/stake-test/transitions"

func main() {
	t := transitions.NewTesting()
	// tested
	// t.UpgradeTo()
	// t.ChangeProxyAdmin()
	// t.Pause()

	// test in progress
	t.Unpause()

	// haven't tested
	// t.UpdateAdmin()
	// t.UpdateVerifier()
	// t.AddFunds()
	// t.DrainContractBalance()
	// t.UpdateStakingParameters()
	// t.AddSSN()
	// t.AddDelegator()
	// t.DelegateStake()
	// t.UpdateReceiveAddr()
	// t.AssignStakeReward()
	// t.AssignStakeReward2()
	// t.WithDrawStakeAmount()
	// t.WithDrawStakeAmount2()
	// t.WithDrawStakeAmount3()
	// t.AddSSNAfterUpgrade()
	// t.UpdateComm()
	// t.WithdrawComm()
	// t.WithdrawStakeReward()
	// t.WithdrawStakeReward2()
	// t.WithdrawStakeReward3()
}
