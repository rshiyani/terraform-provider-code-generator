package models


type Contract struct {
                    TenantDn string `json:tenant_dn",omitempty"`
                    Name string `json:name",omitempty"`
                    MyMap map `json:my_map",omitempty"`
                    Prio string `json:prio",omitempty"`
                        Filters []Filter `json:filter",omitempty"`
}

func (contract *Contract) ToMap() (map[string]interface{}, error) {
    contractMap := make(map[string]interface{})
            contractMap["tenant_dn"] = contract.TenantDn
            contractMap["name"] = contract.Name
            contractMap["my_map"] = contract.MyMap
            contractMap["prio"] = contract.Prio
            filterSet := make([]interface{}, 0, 1)
            for _,filter := range contract.Filters {
                filterMap, err := filter.ToMap()
                if err != nil {
                    return nil, err
                }
                filterSet = append(filterSet, filterMap)
            }
            contractMap["filter"] = filterSet
    return contractMap, nil
}



                    type Filter struct {
                        FilterName string `json:filter_name",omitempty"`
                        Id string `json:id",omitempty"`
                        Description string `json:description",omitempty"`
                            FilterEntrys []FilterEntry `json:filter_entry",omitempty"`
    }

    func (filter *Filter) ToMap() (map[string]interface{}, error) {
        filterMap := make(map[string]interface{})
                filterMap["filter_name"] = filter.FilterName
                filterMap["id"] = filter.Id
                filterMap["description"] = filter.Description
                    filterEntrySet := make([]interface{}, 0, 1)
                    for _,filterEntry := range filter.FilterEntrys {
                        filterEntryMap, err := filterEntry.ToMap()
                        if err != nil {
                            return nil, err
                        }
                        filterEntrySet = append(filterEntrySet, filterEntryMap)
                    }
                    filterMap["filter_entry"] = filterEntrySet
        return filterMap, nil
    }


                                    type FilterEntry struct {
                        FilterEntryName string `json:filter_entry_name",omitempty"`
                        Id string `json:id",omitempty"`
                        ApplyToFrag string `json:apply_to_frag",omitempty"`
    }

    func (filterEntry *FilterEntry) ToMap() (map[string]interface{}, error) {
        filterEntryMap := make(map[string]interface{})
                filterEntryMap["filter_entry_name"] = filterEntry.FilterEntryName
                filterEntryMap["id"] = filterEntry.Id
                filterEntryMap["apply_to_frag"] = filterEntry.ApplyToFrag
        return filterEntryMap, nil
    }


                

