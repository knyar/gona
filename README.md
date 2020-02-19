

# gona
`import "github.com/netactuate/gona"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>
Package gona provides a simple golang interface to the NetActuate
Rest API at <a href="https://vapi.netactuate.com/">https://vapi.netactuate.com/</a>




## <a name="pkg-index">Index</a>
* [Constants](#pkg-constants)
* [func GetKeyFromEnv() string](#GetKeyFromEnv)
* [type Client](#Client)
  * [func NewClient(apikey string) *Client](#NewClient)
  * [func (c *Client) CancelServer(id int) error](#Client.CancelServer)
  * [func (c *Client) CreateSSHKey(name, key string) (sshkey SSHKey, err error)](#Client.CreateSSHKey)
  * [func (c *Client) CreateServer(name, plan string, locationID, osID int, options *ServerOptions) (server Server, err error)](#Client.CreateServer)
  * [func (c *Client) DeleteSSHKey(id int) error](#Client.DeleteSSHKey)
  * [func (c *Client) DeleteServer(id int) error](#Client.DeleteServer)
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
  * [func (c *Client) UpdateSSHKey(id int, name, key string) (SSHKey, error)](#Client.UpdateSSHKey)
* [type JobID](#JobID)
* [type Location](#Location)
* [type OS](#OS)
* [type Package](#Package)
* [type Plan](#Plan)
* [type SSHKey](#SSHKey)
* [type Server](#Server)
* [type ServerOptions](#ServerOptions)


#### <a name="pkg-files">Package files</a>
[client.go](/src/github.com/netactuate/gona/client.go) [locations.go](/src/github.com/netactuate/gona/locations.go) [os.go](/src/github.com/netactuate/gona/os.go) [packages.go](/src/github.com/netactuate/gona/packages.go) [plans.go](/src/github.com/netactuate/gona/plans.go) [servers.go](/src/github.com/netactuate/gona/servers.go) [sshkeys.go](/src/github.com/netactuate/gona/sshkeys.go) 


## <a name="pkg-constants">Constants</a>
``` go
const (
    Version      = "0.0.1"
    BaseEndpoint = "https://vapi.netactuate.com/"
    ContentType  = "application/json"
)
```



## <a name="GetKeyFromEnv">func</a> [GetKeyFromEnv](/src/target/client.go?s=727:754#L27)
``` go
func GetKeyFromEnv() string
```
GetKeyFromEnv is a simple function to try to yank the value for "VR_API_KEY" from
the environment




## <a name="Client">type</a> [Client](/src/target/client.go?s=521:623#L18)
``` go
type Client struct {
    // contains filtered or unexported fields
}
```
Client is the main object (struct) to which we attach most methods/functions.
It has the following fields: (client, userAgent, endPoint, apiKey)







### <a name="NewClient">func</a> [NewClient](/src/target/client.go?s=963:1000#L33)
``` go
func NewClient(apikey string) *Client
```
NewClient is the main entrypoint for instantiating a Client struct. It takes
your API Key as it's sole argument and returns the Client struct ready to talk to the API





### <a name="Client.CancelServer">func</a> (\*Client) [CancelServer](/src/target/servers.go?s=4379:4422#L148)
``` go
func (c *Client) CancelServer(id int) error
```
CancelServer external method on Client to cancel/remove from billing an instance.
this method completely removes an instance, it cannot be rebuilt afterward.
billing should be prorated to the day or something like that.
This method requires apikey_allow_cancel to be checked on the account.




### <a name="Client.CreateSSHKey">func</a> (\*Client) [CreateSSHKey](/src/target/sshkeys.go?s=621:695#L21)
``` go
func (c *Client) CreateSSHKey(name, key string) (sshkey SSHKey, err error)
```



### <a name="Client.CreateServer">func</a> (\*Client) [CreateServer](/src/target/servers.go?s=3352:3473#L119)
``` go
func (c *Client) CreateServer(name, plan string, locationID, osID int, options *ServerOptions) (server Server, err error)
```
CreateServer external method on Client to buy and build a new instance.




### <a name="Client.DeleteSSHKey">func</a> (\*Client) [DeleteSSHKey](/src/target/sshkeys.go?s=1241:1284#L47)
``` go
func (c *Client) DeleteSSHKey(id int) error
```



### <a name="Client.DeleteServer">func</a> (\*Client) [DeleteServer](/src/target/servers.go?s=3115:3158#L109)
``` go
func (c *Client) DeleteServer(id int) error
```
DeleteServer external method on Client to destroy an instance.
This should not be used in Terraform as we will use CancelServer instead.
This method requires apikey_allow_delete to be checked on the account




### <a name="Client.GetLocations">func</a> (\*Client) [GetLocations](/src/target/locations.go?s=233:284#L1)
``` go
func (c *Client) GetLocations() ([]Location, error)
```
GetLocations public method on Client to get a list of locations




### <a name="Client.GetOSs">func</a> (\*Client) [GetOSs](/src/target/os.go?s=358:397#L5)
``` go
func (c *Client) GetOSs() ([]OS, error)
```
GetOSs returns a list of OS objects from the api




### <a name="Client.GetPackage">func</a> (\*Client) [GetPackage](/src/target/packages.go?s=714:774#L18)
``` go
func (c *Client) GetPackage(id int) (pkg Package, err error)
```
GetPackage external method on Client that takes an id (int) as it's sole
argument and returns a single Package object




### <a name="Client.GetPackages">func</a> (\*Client) [GetPackages](/src/target/packages.go?s=398:447#L5)
``` go
func (c *Client) GetPackages() ([]Package, error)
```
GetPackages external method on Client that returns a list of Package object from the API




### <a name="Client.GetPlans">func</a> (\*Client) [GetPlans](/src/target/plans.go?s=397:440#L5)
``` go
func (c *Client) GetPlans() ([]Plan, error)
```
GetPlans external method on Client to list available Plans




### <a name="Client.GetSSHKey">func</a> (\*Client) [GetSSHKey](/src/target/sshkeys.go?s=431:492#L14)
``` go
func (c *Client) GetSSHKey(id int) (sshkey SSHKey, err error)
```



### <a name="Client.GetSSHKeys">func</a> (\*Client) [GetSSHKeys](/src/target/sshkeys.go?s=233:289#L3)
``` go
func (c *Client) GetSSHKeys() (keys []SSHKey, err error)
```



### <a name="Client.GetServer">func</a> (\*Client) [GetServer](/src/target/servers.go?s=1161:1222#L39)
``` go
func (c *Client) GetServer(id int) (server Server, err error)
```
GetServer external method on Client to get an instance




### <a name="Client.GetServers">func</a> (\*Client) [GetServers](/src/target/servers.go?s=917:964#L27)
``` go
func (c *Client) GetServers() ([]Server, error)
```
GetServers external method on Client to list your instances




### <a name="Client.ProvisionServer">func</a> (\*Client) [ProvisionServer](/src/target/servers.go?s=2169:2279#L78)
``` go
func (c *Client) ProvisionServer(name string, id, locationID, osID int, options *ServerOptions) (JobID, error)
```
ProvisionServer external method on Client to re-build an instance
This should not be used in Terraform as we will use CreateServer instead




### <a name="Client.RebootServer">func</a> (\*Client) [RebootServer](/src/target/servers.go?s=1863:1906#L67)
``` go
func (c *Client) RebootServer(id int) error
```
RebootServer external method on Client to reboot an instance




### <a name="Client.StartServer">func</a> (\*Client) [StartServer](/src/target/servers.go?s=1412:1454#L47)
``` go
func (c *Client) StartServer(id int) error
```
StartServer external method on Client to boot up an instance




### <a name="Client.StopServer">func</a> (\*Client) [StopServer](/src/target/servers.go?s=1637:1678#L57)
``` go
func (c *Client) StopServer(id int) error
```
StopServer external method on Client to shut down an instance




### <a name="Client.UpdateSSHKey">func</a> (\*Client) [UpdateSSHKey](/src/target/sshkeys.go?s=917:988#L34)
``` go
func (c *Client) UpdateSSHKey(id int, name, key string) (SSHKey, error)
```



## <a name="JobID">type</a> [JobID](/src/target/servers.go?s=805:853#L22)
``` go
type JobID struct {
    ID int `json:"id,string"`
}
```
JobID struct holds the current Job Id for what's being processed










## <a name="Location">type</a> [Location](/src/target/locations.go?s=82:165#L1)
``` go
type Location struct {
    ID   int    `json:"id,string"`
    Name string `json:"name"`
}
```
Location is a struct for storing the id and name of a location










## <a name="OS">type</a> [OS](/src/target/os.go?s=71:305#L1)
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










## <a name="Package">type</a> [Package](/src/target/packages.go?s=88:305#L1)
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










## <a name="Plan">type</a> [Plan](/src/target/plans.go?s=71:334#L1)
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










## <a name="SSHKey">type</a> [SSHKey](/src/target/sshkeys.go?s=58:231#L1)
``` go
type SSHKey struct {
    ID          int    `json:"id,string"`
    Name        string `json:"name"`
    Key         string `json:"ssh_key"`
    Fingerprint string `json:"fingerprint"`
}
```









## <a name="Server">type</a> [Server](/src/target/servers.go?s=122:579#L1)
``` go
type Server struct {
    Name         string `json:"fqdn"`
    ID           int    `json:"mbpkgid,string"`
    OS           string `json:"os"`
    PrimaryIPv4  string `json:"ip"`
    PrimaryIPv6  string `json:"ipv6"`
    PlanID       int    `json:"plan_id,string"`
    PkgID        int    `json:"pkg_id,string"`
    LocationID   int    `json:"location_id,string"`
    OSID         int    `json:"os_id,string"`
    ServerStatus string `json:"status"`
    PowerStatus  string `json:"state"`
}
```
Server struct defines what a VPS looks like










## <a name="ServerOptions">type</a> [ServerOptions](/src/target/servers.go?s=650:736#L15)
``` go
type ServerOptions struct {
    SSHKeyID    int
    Password    string
    CloudConfig string
}
```
ServerOptions struct defines some extra options including SSH Auth














- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)
