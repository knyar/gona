package gona

import (
	"encoding/json"
	"strconv"
)

type BGPSessions struct {
	Sessions []BGPSession `json:"sessions"`
	Session  *BGPSession  `json:"session"`
	Modified bool         `json:"modified"`
	Success  bool         `json:"success"`
}

type BGPSession struct {
	ID             int    `json:"id,string"`
	MbID           int    `json:"mb_id,string"`
	Description    string `json:"description"`
	RoutesReceived string `json:"routes_received"`
	ConfigStatus   string `json:"config_status"`
	LastUpdate     string `json:"last_update"`
	Locked         string `json:"locked"`
	GroupID        int    `json:"group_id,string"`
	GroupName      string `json:"group_name"`
	LocationName   string `json:"location_name"`
	CustomerIP     string `json:"customer_ip"`
	CustomerPeerIP string `json:"customer_peer_ip"`
	ProviderPeerIP string `json:"provider_peer_ip"`
	ProviderIPType string `json:"provider_ip_type"`
	ProviderASN    int    `json:"provider_asn,string"`
	CustomerASN    int    `json:"customer_asn,string"`
	State          string `json:"state"`
}

func (s *BGPSession) IsLocked() bool {
	return "1" == s.Locked
}

func (s *BGPSession) IsProviderIPTypeV4() bool {
	return string(IPv4) == s.ProviderIPType
}

// GetBGPSession external method on Client to get your BGP session
func (c *Client) GetBGPSession(id int) (sessions BGPSessions, err error) {
	err = c.get("cloud/bgpsession2/"+strconv.Itoa(id), &sessions)
	if err != nil {
		return BGPSessions{}, err
	}

	return sessions, nil
}

// GetBGPSessions external method on Client to get BGP sessions
func (c *Client) GetBGPSessions(mbPkgID int) (*[]BGPSession, error) {
	var allSessions BGPSessions

	err := c.get("/cloud/bgpsessions2", &allSessions)
	if err != nil {
		return nil, err
	}
	if len(allSessions.Sessions) == 0 {
		return nil, nil
	}

	ips, err := c.GetIPs(mbPkgID)
	if err != nil {
		return nil, err
	}
	if len(ips.IPv4) == 0 && len(ips.IPv6) == 0 {
		return nil, err
	}

	ipsMap := *ips.GetIPsMap()

	var sessionIDs []int

	for _, session := range allSessions.Sessions {
		_, exists := ipsMap[session.CustomerIP]
		if exists {
			sessionIDs = append(sessionIDs, session.ID)
		}
	}

	var sessions []BGPSession

	for _, id := range sessionIDs {
		ss, err := c.GetBGPSession(id)
		if err != nil {
			return nil, err
		}

		sessions = append(sessions, *ss.Session)
	}

	return &sessions, nil
}

// CreateBGPSession external method on Client to create a BGP session.
func (c *Client) CreateBGPSessions(mbPkgID int, groupID int, isIPV6 bool, redundant bool) (sessions BGPSessions, err error) {
	values := map[string]string{
		"group_id": strconv.Itoa(groupID),
	}

	if isIPV6 {
		values["ipv6"] = "1"
	}
	if redundant {
		values["redundant"] = "1"
	}

	postData, _ := json.Marshal(values)

	if err := c.post("/cloud/bgpcreatesessions/"+strconv.Itoa(mbPkgID), postData, &sessions); err != nil {
		return BGPSessions{}, err
	}

	return sessions, nil
}
