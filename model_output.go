package models





    type Contract struct {
                        TenantDn string `json:"tenant_dn"`
                        Name string `json:"name"`
                        Prio string `json:"prio,omitempty"`
                            Filters []Filter `json:"filter,omitempty"`
    }

                    type Filter struct {
                        FilterName string `json:"filter_name"`
                        Id string `json:"id,omitempty"`
                        Description string `json:"description,omitempty"`
                            FilterEntrys []FilterEntry `json:"filter_entry,omitempty"`
    }

                                    type FilterEntry struct {
                            Casts []string `json:"cast"`
                        FilterEntryName string `json:"filter_entry_name"`
                        Id string `json:"id,omitempty"`
                        ApplyToFrag string `json:"apply_to_frag,omitempty"`
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