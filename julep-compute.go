package main

import (
	"fmt"
	"github.com/chalupaul/viper"
	types "github.com/chalupaul/julep/types"
	log "github.com/Sirupsen/logrus"
	"os"
	cli "github.com/codegangsta/cli"
	"math/big"
)

func init() {
	BootstrapLogging()
}

var cfg *viper.Viper
var InfraTree types.HostGroup

func startup(c *cli.Context) {
	if c.Bool("verbose") {
		log.SetLevel(log.DebugLevel)
	}
	url := c.String("etcd")
	key := c.String("keyfile")
	// Load up the site config
	if c, err := LoadCfg(OptionUrl(url), OptionKey(key)); err != nil {
		log.Fatal(err)
		os.Exit(1)
	} else {
		log.Debug("Config loaded during startup.")
		cfg = c
	}
	// Load up the infrastructure tree representation
	if t, err := LoadCfg(OptionUrl(url), OptionKey(key), OptionCfgUrl(DefaultTreeUrl)); err != nil {
		log.Fatal(err)
		os.Exit(1)
	} else {
		log.Debug("Loaded infrastructure tree representation.")
		err := t.Marshal(&InfraTree); if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		InfraTree.OrderHostIds()
//		fmt.Println(InfraTree)
//		sort.Sort(types.HostById(InfraTree.Hosts))
		fmt.Println(InfraTree)
		fmt.Println(InfraTree.ChildGroup.ChildGroup.Hosts[0].Name)
	}
	
	//fmt.Println("::",cfg.GetString("hi"),"::")
}

func schedule(i types.Instance) (error) {
	return nil
}

func main() {
	// Set up cli framework
	app := cli.NewApp()
	app.Name = "julep"
	app.Usage = "simple. golang. cloud."
	app.Version = "0.1.0"
	app.Action = startup
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name: "verbose",
			Usage: "verbose mode",
		},
		cli.StringFlag{
			Name: "etcd, e",
			Value: DefaultEtcdUrl,
			Usage: "etcd URL",
			EnvVar: "JULEP_ETCD_URL",
		},
		cli.StringFlag{
			Name: "keyfile, k",
			Value: os.ExpandEnv(DefaultKeyFile),
			Usage: "private gpg key to decrypt etcd data",
			EnvVar: "JULEP_PRIVATE_KEY",
		},
	}
	app.Run(os.Args)
	id := big.NewInt(0)
	//h := md5.New()
	//h.Write([]byte(uuid.New()))
	//idHex := hex.EncodeToString(h.Sum(nil))
	if _, ok := id.SetString("340282366920938463463374607431768211456", 10); ok {
		var z big.Int
		x := big.NewInt(5)
		fmt.Println(z.Add(id, x))
		fmt.Printf("number = %v\n", id)
	} else {
		fmt.Printf("instance id %#v too large\n", id)
	}
	

	
	
	/*
	i := types.Instance{}
	id :=i.GenID()
	fmt.Println("instance id: ", id, i.Id)
	log.WithFields(log.Fields{
		"id": i.Id,
		"hash": id,
	}).Info("Instance created.")
	hosts := make([]types.Host, 3)
	h1 := types.Host{Name: "one.this.thing", HashStart: "000000000000000000000000000000000000000", HashEnd: "113427455640312821154458202477256070485"}
	h1.GenID()
	h2 := types.Host{Name: "one.this.thing", HashStart: "113427455640312821154458202477256070486", HashEnd: "226854911280625642308916404954512140970"}
	h2.GenID()
	h3 := types.Host{Name: "one.this.thing", HashStart: "226854911280625642308916404954512140971", HashEnd: "340282366920938463463374607431768211456"}
	h3.GenID()
	hosts = append(hosts, h1)
	hosts = append(hosts, h2)
	hosts = append(hosts, h3)
	*/
	
	
}
