package v2

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

// SmcRunResultType sets name of json field for result response.
const SmcRunResultType = "smc.runResult"

// SmcActiveElectionIDMethod sets name for getting active elecetion id.
const SmcActiveElectionIDMethod = "active_election_id"

// SmcWalletSeqnoMethod sets method name to get smart-contract seqno.
const SmcWalletSeqnoMethod = "seqno"

// SmcParticipantListMethod sets method name to get participant list.
const SmcParticipantListMethod = "participant_list"

// SmcParticipantListExtendedMethod sets method name to get participant_list_extended.
const SmcParticipantListExtendedMethod = "participant_list_extended"

// SmcParticipatesInMethod sets method name to get participates_in.
const SmcParticipatesInMethod = "participates_in"

// SmcComputeReturnedStakeMethod sets method name to get compute_returned_stake.
const SmcComputeReturnedStakeMethod = "compute_returned_stake"

// NoErrorCode sets code for successful response.
const NoErrorCode = 0

// ElectionParticipant struct for election participant.
type ElectionParticipant struct {
	Id                 string      `json:"id"`
	Stake              string      `json:"stake"`
	MaxFactor          string      `json:"max_factor"`
	ParticipantAddress string      `json:"participant_address"`
	AdnlAddress        string      `json:"adnl_address"`
	Raw                interface{} `json:"-"`
}

// GetActiveElectionID return active election id if elections are running.
func (client *Client) GetActiveElectionID(address string) (int64, error) {
	smcInfo, err := client.LoadContract(address)
	if err != nil {
		return 0, err
	}

	method := NewSmcMethodIdName(SmcActiveElectionIDMethod)
	params := []TvmStackEntry{}
	runMethodResult, err := client.SmcRunGetMethod(smcInfo.Id, method, params)
	if err != nil {
		return 0, err
	}

	if runMethodResult.Type != SmcRunResultType {
		return 0, fmt.Errorf("Unexpected response from tonlib with type:%s. %#v", runMethodResult.Type, *runMethodResult)
	}
	if len(runMethodResult.Stack) < 1 {
		return 0, fmt.Errorf("Empty stack response: %s. %#v", runMethodResult.Type, *runMethodResult)
	}

	// map response
	firstEntity, ok := runMethodResult.Stack[0].(map[string]interface{})
	if !ok {
		return 0, fmt.Errorf("Failed to map `%#v  to `map[string]interface{}`", runMethodResult.Stack[0])
	}
	firstNum, ok := firstEntity["number"]
	if !ok {
		return 0, fmt.Errorf("got response with empty 'number': %#v", firstEntity)
	}

	secondEntity, ok := firstNum.(map[string]interface{})
	if !ok {
		return 0, fmt.Errorf("Failed to map `%#v  to `map[string]interface{}`", secondEntity)
	}
	secondNum, ok := secondEntity["number"]
	if !ok {
		return 0, fmt.Errorf("got response with empty second 'number': %#v", secondEntity)
	}

	return strconv.ParseInt(secondNum.(string), 10, 64)
}

// LoadContract by smart-contract address.
func (client *Client) LoadContract(address string) (*SmcInfo, error) {
	contract := NewAccountAddress(address)
	smcInfo, err := client.SmcLoad(*contract)
	if err != nil {
		return nil, err
	}
	return smcInfo, nil
}

// GetWalletSeqno returns current seqno of smart-contract.
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
	if runMethodResult.Type != SmcRunResultType || runMethodResult.ExitCode != NoErrorCode {
		return 0, fmt.Errorf("got response with type %s and with exit_code: %d", runMethodResult.Type, runMethodResult.ExitCode)
	}
	if len(runMethodResult.Stack) < 1 {
		return 0, fmt.Errorf("empty stack response:%s. %#v", runMethodResult.Type, *runMethodResult)
	}

	// map response
	firstEntity, ok := runMethodResult.Stack[0].(map[string]interface{})
	if !ok {
		return 0, fmt.Errorf("Failed to map `%#v  to `map[string]interface{}`", runMethodResult.Stack[0])
	}
	firstNum, ok := firstEntity["number"]
	if !ok {
		return 0, fmt.Errorf("got response with empty 'number': %#v", firstEntity)
	}

	secondEntity, ok := firstNum.(map[string]interface{})
	if !ok {
		return 0, fmt.Errorf("Failed to map `%#v  to `map[string]interface{}`", secondEntity)
	}
	secondNum, ok := secondEntity["number"]
	if !ok {
		return 0, fmt.Errorf("got response with empty second 'number': %#v", secondEntity)
	}

	return strconv.ParseInt(secondNum.(string), 10, 64)
}

// GetParticipantList returns participant list for current election.
func (client *Client) GetParticipantList(address string) (*[]TvmStackEntry, error) {
	smcInfo, err := client.LoadContract(address)
	if err != nil {
		return nil, err
	}
	method := NewSmcMethodIdName(SmcParticipantListMethod)
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

// GetParticipantListExtended returns extended participant list for current election.
func (client *Client) GetParticipantListExtended(electorAddress string) (*[]ElectionParticipant, error) {
	smcInfo, err := client.LoadContract(electorAddress)
	if err != nil {
		return nil, err
	}
	method := NewSmcMethodIdName(SmcParticipantListExtendedMethod)
	params := []TvmStackEntry{}
	runMethodResult, err := client.SmcRunGetMethod(smcInfo.Id, method, params)
	if err != nil {
		return nil, err
	}
	if runMethodResult.Type != SmcRunResultType {
		return nil, fmt.Errorf("Got response with type `%s` instead of `%s`", runMethodResult.Type, SmcRunResultType)
	}
	if len(runMethodResult.Stack) != 1 {
		return nil, fmt.Errorf("expected length of Stack: 1, but got: %d. Resp: %#v", len(runMethodResult.Stack), runMethodResult.Stack)
	}
	stackEntryList, ok := runMethodResult.Stack[0].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("#1 failed to parse element as map[string]interface{}. element: %#v", runMethodResult.Stack[0])
	}
	listValueInterface, ok := stackEntryList["list"]
	if !ok {
		return nil, fmt.Errorf("#2 failed to find `list` in dict. element: %#v", stackEntryList)
	}
	tvmList, ok := listValueInterface.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("#3 failed to parse element as map[string]interface{}. element: %#v", listValueInterface)
	}
	elementsListInterface, ok := tvmList["elements"]
	if !ok {
		return nil, fmt.Errorf("#4 failed to find `elements` in dict. element: %#v", tvmList)
	}
	elementsList, ok := elementsListInterface.([]interface{})
	if !ok {
		return nil, fmt.Errorf("#5 failed to parse element as []interface{}. element: %#v", elementsListInterface)
	}

	// parse elements
	participants := []ElectionParticipant{}
	for _, el := range elementsList {
		tupleEl, ok := el.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("#6 failed to parse element as map[string]interface{}. element: %#v", el)
		}
		tupleElementsInterface, ok := tupleEl["tuple"]
		if !ok {
			return nil, fmt.Errorf("#7 failed to find `tuple` in dict. element: %#v", tupleEl)
		}
		tupleElementsDict, ok := tupleElementsInterface.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("#8 failed to parse element as map[string]interface{}. element: %#v", tupleElementsInterface)
		}
		elElementsInterface, ok := tupleElementsDict["elements"]
		if !ok {
			return nil, fmt.Errorf("#9 failed to find `elements` in dict. element: %#v", tupleElementsDict)
		}
		elElements, ok := elElementsInterface.([]interface{})
		if !ok {
			return nil, fmt.Errorf("#10 failed to parse element as []interface{}. element: %#v", elElementsInterface)
		}
		if len(elElements) != 2 {
			return nil, fmt.Errorf("#11 expected length of elElements: 2, but got: %d. elements: %#v", len(elElements), elElements)
		}
		// parse id
		idElementDict, ok := elElements[0].(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("#12 failed to parse element as map[string]interface{}. element: %#v", elElements[0])
		}
		idElementNumberInterface, ok := idElementDict["number"]
		if !ok {
			return nil, fmt.Errorf("#13 failed to find `number` in dict. element: %#v", idElementDict)
		}
		idElementNumberDict, ok := idElementNumberInterface.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("#14 failed to parse element as map[string]interface{}. element: %#v", idElementNumberInterface)
		}
		idElementNumberValueInterface, ok := idElementNumberDict["number"]
		if !ok {
			return nil, fmt.Errorf("#15 failed to find `number` in dict. element: %#v", idElementNumberDict)
		}
		idElementValue, ok := idElementNumberValueInterface.(string)
		if !ok {
			return nil, fmt.Errorf("#16 failed to parse element as string. element: %#v", idElementNumberValueInterface)
		}
		// parse other values struct
		valuesElementDict, ok := elElements[1].(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("#17 failed to parse element as map[string]interface{}. element: %#v", elElements[1])
		}
		valuesElementTupleInterface, ok := valuesElementDict["tuple"]
		if !ok {
			return nil, fmt.Errorf("#18 failed to find `tuple` in dict. element: %#v", valuesElementDict)
		}
		valuesElementTupleDict, ok := valuesElementTupleInterface.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("#19 failed to parse element as map[string]interface{}. element: %#v", valuesElementTupleInterface)
		}
		valuesElementTupleElementsInterface, ok := valuesElementTupleDict["elements"]
		if !ok {
			return nil, fmt.Errorf("#20 failed to find `elements` in dict. element: %#v", valuesElementTupleDict)
		}
		valuesElementTupleElements, ok := valuesElementTupleElementsInterface.([]interface{})
		if !ok {
			return nil, fmt.Errorf("#21 failed to parse element as []interface{}. element: %#v", valuesElementTupleElementsInterface)
		}
		if len(valuesElementTupleElements) != 4 {
			return nil, fmt.Errorf("#22 expected length of valuesElementTupleElements: 4, but got: %d. elements: %#v", len(valuesElementTupleElements), valuesElementTupleElements)
		}
		values := []string{"", "", "", ""}
		// parse value
		for i, valuesElementTupleElementInterface := range valuesElementTupleElements {
			valuesElementTupleElement, ok := valuesElementTupleElementInterface.(map[string]interface{})
			if !ok {
				return nil, fmt.Errorf("#23 failed to parse element as map[string]interface{}. element: %#v", valuesElementTupleElementInterface)
			}
			valuesElementTupleElementNumberInterface, ok := valuesElementTupleElement["number"]
			if !ok {
				return nil, fmt.Errorf("#24 failed to find `number` in dict. element: %#v", valuesElementTupleElement)
			}
			valuesElementTupleElementNumberDict, ok := valuesElementTupleElementNumberInterface.(map[string]interface{})
			if !ok {
				return nil, fmt.Errorf("#25 failed to parse element as map[string]interface{}. element: %#v", valuesElementTupleElementNumberInterface)
			}
			valuesElementTupleElementNumberValueInterface, ok := valuesElementTupleElementNumberDict["number"]
			if !ok {
				return nil, fmt.Errorf("#26 failed to find `number` in dict. element: %#v", valuesElementTupleElementNumberDict)
			}
			values[i], ok = valuesElementTupleElementNumberValueInterface.(string)
			if !ok {
				return nil, fmt.Errorf("#27 failed to parse element as string. element: %#v", valuesElementTupleElementNumberValueInterface)
			}
		}

		item := ElectionParticipant{
			Id:                 idElementValue,
			Stake:              values[0],
			MaxFactor:          values[1],
			ParticipantAddress: values[2],
			AdnlAddress:        values[3],
			Raw:                el,
		}
		participants = append(participants, item)
	}

	return &participants, nil
}

// CheckParticipatesIn checks if pubKey participates in current election.
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

	if runMethodResult.Type != SmcRunResultType || runMethodResult.ExitCode != NoErrorCode {
		return 0, fmt.Errorf("got response with type %s and with exit_code: %d", runMethodResult.Type, runMethodResult.ExitCode)
	}
	if len(runMethodResult.Stack) < 1 {
		return 0, fmt.Errorf("got an empty Stack in the response")
	}

	// map response
	firstEntity, ok := runMethodResult.Stack[0].(map[string]interface{})
	if !ok {
		return 0, fmt.Errorf("Failed to map `%#v  to `map[string]interface{}`", runMethodResult.Stack[0])
	}
	firstNum, ok := firstEntity["number"]
	if !ok {
		return 0, fmt.Errorf("got response with empty 'number': %#v", firstEntity)
	}

	secondEntity, ok := firstNum.(map[string]interface{})
	if !ok {
		return 0, fmt.Errorf("Failed to map `%#v  to `map[string]interface{}`", secondEntity)
	}
	secondNum, ok := secondEntity["number"]
	if !ok {
		return 0, fmt.Errorf("got response with empty second 'number': %#v", secondEntity)
	}

	return strconv.ParseInt(secondNum.(string), 10, 64)
}

// CheckReward check if we have anything to recover from elections smart-contract.
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
	if runMethodResult.Type != SmcRunResultType || runMethodResult.ExitCode != NoErrorCode {
		return 0, fmt.Errorf("got response with type %s and with exit_code: %d", runMethodResult.Type, runMethodResult.ExitCode)
	}
	if len(runMethodResult.Stack) < 1 {
		return 0, fmt.Errorf("got response with empty stack: %#v", runMethodResult)
	}

	// map response
	firstEntity, ok := runMethodResult.Stack[0].(map[string]interface{})
	if !ok {
		return 0, fmt.Errorf("Failed to map `%#v  to `map[string]interface{}`", runMethodResult.Stack[0])
	}
	firstNum, ok := firstEntity["number"]
	if !ok {
		return 0, fmt.Errorf("got response with empty 'number': %#v", firstEntity)
	}

	secondEntity, ok := firstNum.(map[string]interface{})
	if !ok {
		return 0, fmt.Errorf("Failed to map `%#v  to `map[string]interface{}`", secondEntity)
	}
	secondNum, ok := secondEntity["number"]
	if !ok {
		return 0, fmt.Errorf("got response with empty second 'number': %#v", secondEntity)
	}

	return strconv.ParseInt(secondNum.(string), 10, 64)
}

// GetAccountStateSimple is wrapper around GetAccountState.
func (client *Client) GetAccountStateSimple(address string) (*FullAccountState, error) {
	accountAddress := NewAccountAddress(address)
	return client.GetAccountState(*accountAddress)
}

// GetLastBlock syncs with the node until last block.
func (client *Client) GetLastBlock() (string, error) {
	return client.Sync(SyncState(SyncState{}))
}

// TonlibSendFile sends BoC file to the network.
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
