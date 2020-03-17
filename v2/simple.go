package v2

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

const SmcRunResultType  = "smc.runResult"
const SmcElectionIdMethod  = "active_election_id"
const SmcWalletSeqnoMethod  =  "seqno"
const SmcParicipiantListMethod  = "participant_list"
const SmcParicipiantListExtendedMethod = "participant_list_extended"
const SmcParticipatesInMethod  = "participates_in"
const SmcComputeReturnedStakeMethod = "compute_returned_stake"

func (client *Client) GetActiveElectionID(address string) (int64, error) {
	smcInfo, err := client.LoadContract(address)
	if err != nil {
		return 0, err
	}

	method := NewSmcMethodIdName(SmcElectionIdMethod)
	params := []TvmStackEntry{}
	runMethodResult, err := client.SmcRunGetMethod(smcInfo.Id, method, params)
	if err != nil {
		return 0, err
	}

	if runMethodResult.Type != SmcRunResultType {
		return 0, fmt.Errorf("Unexpected response from tonlib with type:%s. %#v", runMethodResult.Type, *runMethodResult)
	}

	return strconv.ParseInt(runMethodResult.Stack[0].(map[string]interface{})["number"].(map[string]interface{})["number"].(string), 10, 64)
}

func (client *Client) LoadContract(address string) (*SmcInfo, error) {
	contract := NewAccountAddress(address)
	smcInfo, err := client.SmcLoad(*contract)
	if err != nil {
		return nil, err
	}
	return smcInfo, nil
}

func (client *Client) GetWalletSeqno(address string) (int64, error) {
	smcInfo, err := client.LoadContract(address)
	if err != nil {
		return 0, err
	}
	method := NewSmcMethodIdName(SmcWalletSeqnoMethod)
	params := []TvmStackEntry{}
	runMethodResult, err := client.SmcRunGetMethod(smcInfo.Id, method, params)
	if err != nil {
		return 0, fmt.Errorf("runMethodResult failed. %v", err)
	}
	if runMethodResult.Type != SmcRunResultType || runMethodResult.ExitCode != 0 {
		return 0, fmt.Errorf("Got response with type %s and with exit_code: %d.", runMethodResult.Type, runMethodResult.ExitCode)
	}
	return strconv.ParseInt(runMethodResult.Stack[0].(map[string]interface{})["number"].(map[string]interface{})["number"].(string), 10, 64)
}

func (client *Client) GetParticipantList(address string) (*[]TvmStackEntry, error) {
	smcInfo, err := client.LoadContract(address)
	if err != nil {
		return nil, err
	}
	method := NewSmcMethodIdName(SmcParicipiantListMethod)
	params := []TvmStackEntry{}
	runMethodResult, err := client.SmcRunGetMethod(smcInfo.Id, method, params)
	if err != nil {
		return nil, err
	}
	if runMethodResult.Type != SmcRunResultType {
		return nil, fmt.Errorf("Got response with type `%s` instead of `%s`", runMethodResult.Type, SmcRunResultType)
	}
	return &runMethodResult.Stack, nil
}

func (client *Client) GetParticipantListExtended(electorAddress string) (*[]TvmStackEntry, error) {
	smcInfo, err := client.LoadContract(electorAddress)
	if err != nil {
		return nil, err
	}
	method := NewSmcMethodIdName(SmcParicipiantListExtendedMethod)
	params := []TvmStackEntry{}
	runMethodResult, err := client.SmcRunGetMethod(smcInfo.Id, method, params)
	if err != nil {
		return nil, err
	}
	if runMethodResult.Type != SmcRunResultType {
		return nil, fmt.Errorf("Got response with type `%s` instead of `%s`", runMethodResult.Type, SmcRunResultType)
	}
	return &runMethodResult.Stack, nil
}

func (client *Client) CheckParticipatesIn(pubKey, address string) (int64, error) {
	smcInfo, err := client.LoadContract(address)
	if err != nil {
		return 0, err
	}

	method := NewSmcMethodIdName(SmcParticipatesInMethod)
	params := []TvmStackEntry{}
	valAddress := NewTvmNumberDecimal(hex2int(pubKey).String())
	stackAddress := NewTvmStackEntryNumber(valAddress)
	params = append(params, stackAddress)
	runMethodResult, err := client.SmcRunGetMethod(smcInfo.Id, method, params)
	if err != nil {
		return 0, err
	}

	if runMethodResult.Type != SmcRunResultType || runMethodResult.ExitCode != 0 {
		return 0, fmt.Errorf("got response with type %s and with exit_code: %d.", runMethodResult.Type, runMethodResult.ExitCode)
	}
	if len(runMethodResult.Stack) < 1 {
		return 0, fmt.Errorf("got an empty Stack in the response")
	}

	return strconv.ParseInt(runMethodResult.Stack[0].(map[string]interface{})["number"].(map[string]interface{})["number"].(string), 10, 64)

}

func (client *Client) CheckReward(address, electorAddress string) (int64, error) {
	smcInfo, err := client.LoadContract(electorAddress)
	if err != nil {
		return 0, err
	}

	method := NewSmcMethodIdName(SmcComputeReturnedStakeMethod)
	params := []TvmStackEntry{}
	tvmnum := NewTvmNumberDecimal(hex2int(address).String())
	stnum := NewTvmStackEntryNumber(tvmnum)
	params = append(params, stnum)
	runMethodResult, err := client.SmcRunGetMethod(smcInfo.Id, method, params)
	if err != nil {
		return 0, fmt.Errorf("runMethodResult failed with params %#v. error: %v", params, err)
	}
	if runMethodResult.Type != SmcRunResultType || runMethodResult.ExitCode != 0 {
		return 0, fmt.Errorf("got response with type %s and with exit_code: %d.", runMethodResult.Type, runMethodResult.ExitCode)
	}
	return strconv.ParseInt(runMethodResult.Stack[0].(map[string]interface{})["number"].(map[string]interface{})["number"].(string), 10, 64)
}

func (client *Client) GetAccountStateSimple(address string) (*FullAccountState, error) {
	accountAddress := NewAccountAddress(address)
	return client.GetAccountState(*accountAddress)
}

func (client *Client) GetLastBlock() (string, error) {
	return client.Sync(SyncState(SyncState{}))

}

func (client *Client) TonlibSendFile(bocFilePath string) error {
	if !fileExists(bocFilePath) {
		return fmt.Errorf("file does not exist")
	}
	data, err := ioutil.ReadFile(bocFilePath)
	if err != nil {
		return err
	}

	_, err = client.RawSendMessage(data)
	return err
}
