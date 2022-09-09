

# gona
`import "github.com/netactuate/gona/gona"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>
Package gona provides a simple golang interface to the NetActuate
Rest API at <a href="https://vapi.netactuate.com/">https://vapi.netactuate.com/</a>




## <a name="pkg-index">Index</a>
* [Constants](#pkg-constants)
* [func GetKeyFromEnv() string](#GetKeyFromEnv)
* [type BGPSession](#BGPSession)
  * [func (s *BGPSession) IsLocked() bool](#BGPSession.IsLocked)
  * [func (s *BGPSession) IsProviderIPTypeV4() bool](#BGPSession.IsProviderIPTypeV4)
* [type BGPSessions](#BGPSessions)
* [type Client](#Client)
  * [func NewClient(apikey string) *Client](#NewClient)
  * [func NewClientCustom(apikey string, apiurl string) *Client](#NewClientCustom)
  * [func (c *Client) CancelServer(id int) error](#Client.CancelServer)
  * [func (c *Client) CreateBGPSessions(mbPkgID int, groupID int, isIPV6 bool, redundant bool) (sessions BGPSessions, err error)](#Client.CreateBGPSessions)
  * [func (c *Client) CreateSSHKey(name, key string) (sshkey SSHKey, err error)](#Client.CreateSSHKey)
  * [func (c *Client) CreateServer(s *Server, options *ServerOptions) (server Server, err error)](#Client.CreateServer)
  * [func (c *Client) DeleteSSHKey(id int) error](#Client.DeleteSSHKey)
  * [func (c *Client) DeleteServer(id int) error](#Client.DeleteServer)
  * [func (c *Client) GetBGPSession(id int) (sessions BGPSessions, err error)](#Client.GetBGPSession)
  * [func (c *Client) GetBGPSessions(mbPkgID int) (*[]BGPSession, error)](#Client.GetBGPSessions)
  * [func (c *Client) GetIPs(mbPkgID int) (ips IPs, err error)](#Client.GetIPs)
  * [func (c *Client) GetLocations() ([]Location, error)](#Client.GetLocations)
  * [func (c *Client) GetOSs() ([]OS, error)](#Client.GetOSs)
  * [func (c *Client) GetPackage(id int) (pkg Package, err error)](#Client.GetPackage)
  * [func (c *Client) GetPackages() ([]Package, error)](#Client.GetPackages)
  * [func (c *Client) GetPlans() ([]Plan, error)](#Client.GetPlans)
  * [func (c *Client) GetSSHKey(id int) (sshkey SSHKey, err error)](#Client.GetSSHKey)
  * [func (c *Client) GetSSHKeys() (keys []SSHKey, err error)](#Client.GetSSHKeys)
  * [func (c *Client) GetServer(id int) (server Server, err error)](#Client.GetServer)
  * [func (c *Client) GetServers() ([]Server, error)](#Client.GetServers)
  * [func (c *Client) ProvisionServer(name string, id, locationID, osID int, options *ServerOptions) (JobID, error)](#Client.ProvisionServer)
  * [func (c *Client) RebootServer(id int) error](#Client.RebootServer)
  * [func (c *Client) StartServer(id int) error](#Client.StartServer)
  * [func (c *Client) StopServer(id int) error](#Client.StopServer)
  * [func (c *Client) UnlinkServer(id int) error](#Client.UnlinkServer)
  * [func (c *Client) UpdateSSHKey(id int, name, key string) (SSHKey, error)](#Client.UpdateSSHKey)
* [type IP](#IP)
* [type IPType](#IPType)
* [type IPs](#IPs)
  * [func (ips *IPs) GetIPsMap() *map[string]IPType](#IPs.GetIPsMap)
* [type JobID](#JobID)
* [type Location](#Location)
* [type OS](#OS)
* [type Package](#Package)
* [type Plan](#Plan)
* [type SSHKey](#SSHKey)
* [type Server](#Server)
* [type ServerOptions](#ServerOptions)


#### <a name="pkg-files">Package files</a>
[bgp.go](/src/github.com/netactuate/gona/gona/bgp.go) [client.go](/src/github.com/netactuate/gona/gona/client.go) [ip.go](/src/github.com/netactuate/gona/gona/ip.go) [locations.go](/src/github.com/netactuate/gona/gona/locations.go) [os.go](/src/github.com/netactuate/gona/gona/os.go) [packages.go](/src/github.com/netactuate/gona/gona/packages.go) [plans.go](/src/github.com/netactuate/gona/gona/plans.go) [servers.go](/src/github.com/netactuate/gona/gona/servers.go) [sshkeys.go](/src/github.com/netactuate/gona/gona/sshkeys.go) 


## <a name="pkg-constants">Constants</a>
``` go
const (
    Version      = "0.1.3"
    BaseEndpoint = "https://vapi.netactuate.com/"
    ContentType  = "application/json"
)
```
Version, BaseEndpoint, ContentType constants




## <a name="GetKeyFromEnv">func</a> [GetKeyFromEnv](/src/target/client.go?s=797:824#L40)
``` go
func GetKeyFromEnv() string
```
GetKeyFromEnv is a simple function to try to yank the value for
"NA_API_KEY" from the environment




## <a name="BGPSession">type</a> [BGPSession](/src/target/bgp.go?s=245:1049#L15)
``` go
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

```









### <a name="BGPSession.IsLocked">func</a> (\*BGPSession) [IsLocked](/src/target/bgp.go?s=1051:1087#L35)
``` go
func (s *BGPSession) IsLocked() bool
```



### <a name="BGPSession.IsProviderIPTypeV4">func</a> (\*BGPSession) [IsProviderIPTypeV4](/src/target/bgp.go?s=1117:1163#L39)
``` go
func (s *BGPSession) IsProviderIPTypeV4() bool
```



## <a name="BGPSessions">type</a> [BGPSessions](/src/target/bgp.go?s=54:243#L8)
``` go
type BGPSessions struct {
    Sessions []BGPSession `json:"sessions"`
    Session  *BGPSession  `json:"session"`
    Modified bool         `json:"modified"`
    Success  bool         `json:"success"`
}

```









## <a name="Client">type</a> [Client](/src/target/client.go?s=589:691#L31)
``` go
type Client struct {
    // contains filtered or unexported fields
}

```
Client is the main object (struct) to which we attach most
methods/functions.
It has the following fields:
(client, userAgent, endPoint, apiKey)







### <a name="NewClient">func</a> [NewClient](/src/target/client.go?s=1569:1606#L68)
``` go
func NewClient(apikey string) *Client
```
NewClient takes an apikey and calls NewClientCustom with the hardcoded
BaseEndpoint constant API URL


### <a name="NewClientCustom">func</a> [NewClientCustom](/src/target/client.go?s=1044:1102#L47)
``` go
func NewClientCustom(apikey string, apiurl string) *Client
```
NewClientCustom is the main entrypoint for instantiating a Client struct.
It takes your API Key as it's sole argument
and returns the Client struct ready to talk to the API





### <a name="Client.CancelServer">func</a> (\*Client) [CancelServer](/src/target/servers.go?s=4090:4133#L140)
``` go
func (c *Client) CancelServer(id int) error
```
CancelServer external method on Client to cancel/remove from billing an instance.
this method completely removes an instance, it cannot be rebuilt afterward.
billing should be prorated to the day or something like that.
This method requires apikey_allow_cancel to be checked on the account.




### <a name="Client.CreateBGPSessions">func</a> (\*Client) [CreateBGPSessions](/src/target/bgp.go?s=2430:2553#L99)
``` go
func (c *Client) CreateBGPSessions(mbPkgID int, groupID int, isIPV6 bool, redundant bool) (sessions BGPSessions, err error)
```
CreateBGPSession external method on Client to create a BGP session.




### <a name="Client.CreateSSHKey">func</a> (\*Client) [CreateSSHKey](/src/target/sshkeys.go?s=728:802#L37)
``` go
func (c *Client) CreateSSHKey(name, key string) (sshkey SSHKey, err error)
```
CreateSSHKey creates a key




### <a name="Client.CreateServer">func</a> (\*Client) [CreateServer](/src/target/servers.go?s=2510:2601#L94)
``` go
func (c *Client) CreateServer(s *Server, options *ServerOptions) (server Server, err error)
```
CreateServer external method on Client to buy and build a new instance.




### <a name="Client.DeleteSSHKey">func</a> (\*Client) [DeleteSSHKey](/src/target/sshkeys.go?s=1414:1457#L65)
``` go
func (c *Client) DeleteSSHKey(id int) error
```
DeleteSSHKey deletes a key




### <a name="Client.DeleteServer">func</a> (\*Client) [DeleteServer](/src/target/servers.go?s=5552:5595#L189)
``` go
func (c *Client) DeleteServer(id int) error
```
DeleteServer external method on Client to destroy an instance.
This method requires apikey_allow_delete to be checked on the account




### <a name="Client.GetBGPSession">func</a> (\*Client) [GetBGPSession](/src/target/bgp.go?s=1277:1349#L44)
``` go
func (c *Client) GetBGPSession(id int) (sessions BGPSessions, err error)
```
GetBGPSession external method on Client to get your BGP session




### <a name="Client.GetBGPSessions">func</a> (\*Client) [GetBGPSessions](/src/target/bgp.go?s=1553:1620#L54)
``` go
func (c *Client) GetBGPSessions(mbPkgID int) (*[]BGPSession, error)
```
GetBGPSessions external method on Client to get BGP sessions




### <a name="Client.GetIPs">func</a> (\*Client) [GetIPs](/src/target/ip.go?s=837:894#L50)
``` go
func (c *Client) GetIPs(mbPkgID int) (ips IPs, err error)
```
GetIPs returns a list of IPs for the selected mbPkgID from the API




### <a name="Client.GetLocations">func</a> (\*Client) [GetLocations](/src/target/locations.go?s=232:283#L10)
``` go
func (c *Client) GetLocations() ([]Location, error)
```
GetLocations public method on Client to get a list of locations




### <a name="Client.GetOSs">func</a> (\*Client) [GetOSs](/src/target/os.go?s=357:396#L15)
``` go
func (c *Client) GetOSs() ([]OS, error)
```
GetOSs returns a list of OS objects from the api




### <a name="Client.GetPackage">func</a> (\*Client) [GetPackage](/src/target/packages.go?s=718:778#L29)
``` go
func (c *Client) GetPackage(id int) (pkg Package, err error)
```
GetPackage external method on Client that takes an id (int) as it's sole
argument and returns a single Package object




### <a name="Client.GetPackages">func</a> (\*Client) [GetPackages](/src/target/packages.go?s=400:449#L16)
``` go
func (c *Client) GetPackages() ([]Package, error)
```
GetPackages external method on Client that returns a
list of Package object from the API




### <a name="Client.GetPlans">func</a> (\*Client) [GetPlans](/src/target/plans.go?s=396:439#L15)
``` go
func (c *Client) GetPlans() ([]Plan, error)
```
GetPlans external method on Client to list available Plans




### <a name="Client.GetSSHKey">func</a> (\*Client) [GetSSHKey](/src/target/sshkeys.go?s=508:569#L29)
``` go
func (c *Client) GetSSHKey(id int) (sshkey SSHKey, err error)
```
GetSSHKey as in one key




### <a name="Client.GetSSHKeys">func</a> (\*Client) [GetSSHKeys](/src/target/sshkeys.go?s=283:339#L17)
``` go
func (c *Client) GetSSHKeys() (keys []SSHKey, err error)
```
GetSSHKeys as in many keys




### <a name="Client.GetServer">func</a> (\*Client) [GetServer](/src/target/servers.go?s=1567:1628#L56)
``` go
func (c *Client) GetServer(id int) (server Server, err error)
```
GetServer external method on Client to get an instance




### <a name="Client.GetServers">func</a> (\*Client) [GetServers](/src/target/servers.go?s=1322:1369#L44)
``` go
func (c *Client) GetServers() ([]Server, error)
```
GetServers external method on Client to list your instances




### <a name="Client.ProvisionServer">func</a> (\*Client) [ProvisionServer](/src/target/servers.go?s=4315:4425#L150)
``` go
func (c *Client) ProvisionServer(name string, id, locationID, osID int, options *ServerOptions) (JobID, error)
```
ProvisionServer external method on Client to re-build an instance




### <a name="Client.RebootServer">func</a> (\*Client) [RebootServer](/src/target/servers.go?s=2272:2315#L84)
``` go
func (c *Client) RebootServer(id int) error
```
RebootServer external method on Client to reboot an instance




### <a name="Client.StartServer">func</a> (\*Client) [StartServer](/src/target/servers.go?s=1819:1861#L64)
``` go
func (c *Client) StartServer(id int) error
```
StartServer external method on Client to boot up an instance




### <a name="Client.StopServer">func</a> (\*Client) [StopServer](/src/target/servers.go?s=2045:2086#L74)
``` go
func (c *Client) StopServer(id int) error
```
StopServer external method on Client to shut down an instance




### <a name="Client.UnlinkServer">func</a> (\*Client) [UnlinkServer](/src/target/servers.go?s=5801:5844#L199)
``` go
func (c *Client) UnlinkServer(id int) error
```
UnlinkServer external method on Client to unlink a billing package from a location




### <a name="Client.UpdateSSHKey">func</a> (\*Client) [UpdateSSHKey](/src/target/sshkeys.go?s=1059:1130#L51)
``` go
func (c *Client) UpdateSSHKey(id int, name, key string) (SSHKey, error)
```
UpdateSSHKey updates it I guess




## <a name="IP">type</a> [IP](/src/target/ip.go?s=203:472#L21)
``` go
type IP struct {
    ID        int    `json:"id,string"`
    Primary   int    `json:"primary,string"`
    Reverse   string `json:"reverse"`
    IP        string `json:"ip"`
    Gateway   string `json:"gateway"`
    Netmask   string `json:"netmask"`
    Broadcast string `json:"broadcast"`
}

```









## <a name="IPType">type</a> [IPType](/src/target/ip.go?s=57:75#L9)
``` go
type IPType string
```

``` go
const (
    IPv4 IPType = "ipv4"
    IPv6 IPType = "ipv6"
)
```









## <a name="IPs">type</a> [IPs](/src/target/ip.go?s=132:201#L16)
``` go
type IPs struct {
    IPv4 []IP `json:"IPv4"`
    IPv6 []IP `json:"IPv6"`
}

```









### <a name="IPs.GetIPsMap">func</a> (\*IPs) [GetIPsMap](/src/target/ip.go?s=474:520#L31)
``` go
func (ips *IPs) GetIPsMap() *map[string]IPType
```



## <a name="JobID">type</a> [JobID](/src/target/servers.go?s=1209:1257#L39)
``` go
type JobID struct {
    ID int `json:"id,string"`
}

```
JobID struct holds the current Job Id for what's being processed










## <a name="Location">type</a> [Location](/src/target/locations.go?s=80:163#L4)
``` go
type Location struct {
    ID   int    `json:"id,string"`
    Name string `json:"name"`
}

```
Location is a struct for storing the id and name of a location










## <a name="OS">type</a> [OS](/src/target/os.go?s=69:303#L4)
``` go
type OS struct {
    ID      int    `json:"id,string"`
    Os      string `json:"os"`
    Type    string `json:"type"`
    Subtype string `json:"subtype"`
    Size    string `json:"size"`
    Bits    string `json:"bits"`
    Tech    string `json:"tech"`
}

```
OS is a struct for storing the attributes of and OS










## <a name="Package">type</a> [Package](/src/target/packages.go?s=86:303#L6)
``` go
type Package struct {
    ID        int    `json:"mbpkgid,string"`
    Status    string `json:"package_status"`
    Locked    string `json:"locked"`
    PlanName  string `json:"name"`
    Installed int    `json:"installed,string"`
}

```
Package struct stores the purchaced package values










## <a name="Plan">type</a> [Plan](/src/target/plans.go?s=69:332#L4)
``` go
type Plan struct {
    ID        int    `json:"plan_id,string"`
    Name      string `json:"plan"`
    RAM       string `json:"ram"`
    Disk      string `json:"disk"`
    Transfer  string `json:"transfer"`
    Price     string `json:"price"`
    Available string `json:"available"`
}

```
Plan struct defines the purchaceable plans/packages










## <a name="SSHKey">type</a> [SSHKey](/src/target/sshkeys.go?s=78:251#L9)
``` go
type SSHKey struct {
    ID          int    `json:"id,string"`
    Name        string `json:"name"`
    Key         string `json:"ssh_key"`
    Fingerprint string `json:"fingerprint"`
}

```
SSHKey is what it is










## <a name="Server">type</a> [Server](/src/target/servers.go?s=120:925#L10)
``` go
type Server struct {
    Name                     string `json:"fqdn"`
    ID                       int    `json:"mbpkgid,string"`
    OS                       string `json:"os"`
    OSID                     int    `json:"os_id,string"`
    PrimaryIPv4              string `json:"ip"`
    PrimaryIPv6              string `json:"ipv6"`
    Plan                     string `json:"plan"`
    PlanID                   int    `json:"plan_id,string"`
    Package                  string `json:"package"`
    PackageBilling           string `json:"package_billing"`
    PackageBillingContractId string `json:"package_billing_contract_id"`
    Location                 string `json:"city"`
    LocationID               int    `json:"location_id,string"`
    ServerStatus             string `json:"status"`
    PowerStatus              string `json:"state"`
}

```
Server struct defines what a VPS looks like










## <a name="ServerOptions">type</a> [ServerOptions](/src/target/servers.go?s=997:1139#L29)
``` go
type ServerOptions struct {
    SSHKeyID    int
    SSHKey      string
    Password    string
    CloudConfig string
    UserData    string
    UserData64  string
}

```
ServerOptions struct defines some extra options including SSH Auth














- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)
