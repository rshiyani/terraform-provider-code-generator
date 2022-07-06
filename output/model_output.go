package models

import "encoding/json"

type Contract struct {
	TenantDn string                 `json:"tenant_dn"`
	Name     string                 `json:"name"`
	MyMap    map[string]interface{} `json:"my_map"`
	Prio     string                 `json:"prio,omitempty"`
	Casts    []string               `json:"cast"`
	Filters  []Filter               `json:"filter"`
}

type Filter struct {
	FilterName   string        `json:"filter_name"`
	Id           string        `json:"id,omitempty"`
	Description  string        `json:"description,omitempty"`
	FilterEntrys []FilterEntry `json:"filter_entry"`
}

type FilterEntry struct {
	EntryNexts      []EntryNext `json:"entry_next"`
	Casts           []Cast      `json:"cast"`
	FilterEntryName string      `json:"filter_entry_name"`
	Id              string      `json:"id,omitempty"`
	ApplyToFrag     string      `json:"apply_to_frag,omitempty"`
}

type EntryNext struct {
	EntryNextName string `json:"entry_next_name"`
}

type Cast struct {
	Cast2s []string `json:"cast2"`
}

func (contract *Contract) ToMap() (map[string]interface{}, error) {
	contractJSON, err := json.Marshal(contract)
	if err != nil {
		return nil, err
	}
	contractMap := make(map[string]interface{})
	err = json.Unmarshal(contractJSON, &contractMap)
	if err != nil {
		return nil, err
	}
	return contractMap, nil
}
