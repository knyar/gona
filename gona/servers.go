package gona

import (
	"net/url"
	"strconv"

	"github.com/google/go-querystring/query"
)

// Server struct defines what a VPS looks like
type Server struct {
	Name                     string `json:"fqdn"`
	ID                       int    `json:"mbpkgid"`
	OS                       string `json:"os"`
	OSID                     int    `json:"os_id"`
	PrimaryIPv4              string `json:"ip"`
	PrimaryIPv6              string `json:"ipv6"`
	PlanID                   int    `json:"plan_id"`
	Package                  string `json:"package"`
	PackageBilling           string `json:"package_billing"`
	PackageBillingContractId string `json:"package_billing_contract_id"`
	Location                 string `json:"city"`
	LocationID               int    `json:"location_id"`
	ServerStatus             string `json:"status"`
	PowerStatus              string `json:"state"`
	Installed                int    `json:"installed"`
}

// ServerOptions struct defines some extra options including SSH Auth
type ServerOptions struct {
	SSHKeyID    int
	SSHKey      string
	Password    string
	CloudConfig string
	UserData    string
	UserData64  string
}

// GetServers external method on Client to list your instances
func (c *Client) GetServers() ([]Server, error) {

	var serverList []Server

	if err := c.get("cloud/servers", &serverList); err != nil {
		return nil, err
	}

	return serverList, nil
}

// GetServer external method on Client to get an instance
func (c *Client) GetServer(id int) (server Server, err error) {
	if err := c.get("cloud/server?mbpkgid="+strconv.Itoa(id), &server); err != nil {
		return server, err
	}
	return server, nil
}

// StartServer external method on Client to boot up an instance
func (c *Client) StartServer(id int) error {

	if err := c.post("cloud/server/start/"+strconv.Itoa(id), nil, nil); err != nil {
		return err
	}

	return nil
}

// StopServer external method on Client to shut down an instance
func (c *Client) StopServer(id int) error {

	if err := c.post("cloud/server/shutdown/"+strconv.Itoa(id), nil, nil); err != nil {
		return err
	}

	return nil
}

// RebootServer external method on Client to reboot an instance
func (c *Client) RebootServer(id int) error {

	if err := c.post("cloud/server/reboot/"+strconv.Itoa(id), nil, nil); err != nil {
		return err
	}

	return nil
}

type CreateServerRequest struct {
	Plan                     string `url:"plan,omitempty"`
	Location                 int    `url:"location,omitempty"`
	Image                    int    `url:"image,omitempty"`
	FQDN                     string `url:"fqdn,omitempty"`
	SSHKey                   string `url:"ssh_key,omitempty"`
	SSHKeyID                 int    `url:"ssh_key_id,omitempty"`
	Password                 string `url:"password,omitempty"`
	PackageBilling           string `url:"package_billing,omitempty"`
	PackageBillingContractId string `url:"package_billing_contract_id,omitempty"`
	CloudConfig              string `url:"cloud_config,omitempty"`
	ScriptContent            string `url:"script_content,omitempty"`
}

type ServerBuild struct {
	ServerID int    `json:"mbpkgid"`
	Status   string `json:"status"`
	Build    int    `json:"build"`
}

// CreateServer external method on Client to buy and build a new instance.
func (c *Client) CreateServer(r *CreateServerRequest) (b ServerBuild, err error) {

	values, err := query.Values(r)
	if err != nil {
		return b, err
	}
	if values.Has("script_content") {
		values.Add("script_type", "user-data")
	}

	if err := c.post("cloud/server/buy_build", []byte(values.Encode()), &b); err != nil {
		return b, err
	}

	return b, nil
}

// CancelServer external method on Client to cancel/remove from billing an instance.
// this method completely removes an instance, it cannot be rebuilt afterward.
// billing should be prorated to the day or something like that.
// This method requires apikey_allow_cancel to be checked on the account.
func (c *Client) CancelServer(id int) error {

	if err := c.post("cloud/server/cancel/"+strconv.Itoa(id), nil, nil); err != nil {
		return err
	}

	return nil
}

type BuildServerRequest struct {
	Location      int    `url:"location,omitempty"`
	Image         int    `url:"image,omitempty"`
	FQDN          string `url:"fqdn,omitempty"`
	SSHKey        string `url:"ssh_key,omitempty"`
	SSHKeyID      int    `url:"ssh_key_id,omitempty"`
	Password      string `url:"password,omitempty"`
	CloudConfig   string `url:"cloud_config,omitempty"`
	ScriptContent string `url:"script_content,omitempty"`
}

// BuildServer external method on Client to re-build an instance
func (c *Client) BuildServer(id int, r *BuildServerRequest) (b ServerBuild, err error) {
	values, err := query.Values(r)
	if err != nil {
		return b, err
	}
	if values.Has("script_content") {
		values.Add("script_type", "user-data")
	}

	if err := c.post("cloud/server/build/"+strconv.Itoa(id), []byte(values.Encode()), &b); err != nil {
		return b, err
	}

	return b, nil
}

// DeleteServer external method on Client to destroy an instance.
// This method requires apikey_allow_delete to be checked on the account
func (c *Client) DeleteServer(id int, cancelBilling bool) error {
	values := url.Values{}
	if cancelBilling {
		values.Add("cancel_billing", "1")
	}
	return c.post("cloud/server/delete?mbpkgid="+strconv.Itoa(id), []byte(values.Encode()), nil)
}

// UnlinkServer external method on Client to unlink a billing package from a location
func (c *Client) UnlinkServer(id int) error {
	return c.post("cloud/server/unlink/"+strconv.Itoa(id), nil, nil)
}
