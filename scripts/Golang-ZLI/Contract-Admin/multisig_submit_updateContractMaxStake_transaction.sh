#!/usr/bin/env bash
zli contract call -a ${MULTISIG_WALLET_ADDRESS} -t SubmitCustomUpdateContractMaxStakeTransaction -k ${OWNER_KEY} -r "[{\"vname\":\"proxyContract\",\"type\":\"ByStr20\",\"value\":\"${PROXY_CONTRACT}\"},{\"vname\":\"max_stake\",\"type\":\"Uint128\",\"value\":\"${MAX_STAKE}\"}]" -f true
