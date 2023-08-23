package ibclient

import "fmt"

func (objMgr *ObjectManager) GetDNSView(name string) (*DNSView, error) {
	var res []DNSView
	if name == "" {
		return nil, fmt.Errorf(
			"DNS view's name is required to retreive DNS view object")
	}
	queryParams := NewQueryParams(false, map[string]string{"name": name})
	err := objMgr.connector.GetObject(NewEmptyDNSView(), "", queryParams, &res)
	if err != nil {
		return nil, err
	} else if res == nil || len(res) == 0 {
		return nil, NewNotFoundError(fmt.Sprintf("DNS view with name '%s' not found", name))
	}

	return &res[0], nil
}

func (objMgr *ObjectManager) CreateDnsView(name string, network_view string, comment string) (ref string, err error) {
	if name == "" || network_view == "" {
		return "", fmt.Errorf("dns View Name or Network View is null")
	}
	newDnsView := NewDNSView(name, network_view, comment)
	ref, err = objMgr.connector.CreateObject(newDnsView)
	return ref, err
}
