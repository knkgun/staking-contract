package transitions

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Zilliqa/gozilliqa-sdk/bech32"
	contract2 "github.com/Zilliqa/gozilliqa-sdk/contract"
	"math/rand"
	"strings"
)

func (p *Proxy) UpdateContractMaxStake(valid, invalid string) {
	err := p.updateContractMaxStake(invalid)
	if err == nil {
		panic("update contract max stake with invalid key failed")
	}

	fmt.Println("update contract max stake with invalid key succeed")

	err2 := p.updateContractMaxStake(valid)
	if err2 != nil {
		panic("update contract max stake with valid key failed")
	}

	fmt.Println("update contract max stake with valid key succeed")

}

func (p *Proxy) updateContractMaxStake(private string) error {
	proxy, _ := bech32.ToBech32Address(p.Addr)
	stakeNum := fmt.Sprintf("%d", rand.Int())
	parameters := []contract2.Value{
		{
			VName: "max_stake",
			Type:  "Uint128",
			Value: stakeNum,
		},
	}
	args, _ := json.Marshal(parameters)
	if err2, output := ExecZli("contract", "call",
		"-k", private,
		"-a", proxy,
		"-t", "update_contractmaxstake",
		"-r", string(args)); err2 != nil {
		return errors.New("call transition error: " + err2.Error())
	} else {
		tx := strings.TrimSpace(strings.Split(output, "confirmed!")[1])
		payload := p.Provider.GetTransaction(tx).Result.(map[string]interface{})
		receipt := payload["receipt"].(map[string]interface{})
		success := receipt["success"].(bool)
		if success {
			res := p.Provider.GetSmartContractState(p.ImplAddress).Result.(map[string]interface{})
			minstake := res["contractmaxstake"].(string)
			if minstake == stakeNum {
				return nil
			} else {
				return errors.New("state failed")
			}

		} else {
			return errors.New("transaction failed")
		}
	}
}
