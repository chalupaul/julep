package main

import (
	"fmt"
	cfg "github.com/chalupaul/viper"
	"github.com/chalupaul/julep/types"
	log "github.com/Sirupsen/logrus"
	"os"
	cli "github.com/codegangsta/cli"
)

func init() {
	BootstrapLogging()
}

func startup(c *cli.Context) {
	if c.Bool("verbose") {
		log.SetLevel(log.DebugLevel)
	}
	url := c.String("etcd")
	key := c.String("keyfile")
	if err := LoadCfg(url, key); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("::",cfg.GetString("hi"),"::")
}

func main() {
	// Set up cli framework
	app := cli.NewApp()
	app.Name = "julep"
	app.Usage = "simple. golang. cloud."
	app.Version = "0.1.0"
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
	app.Action = startup
	app.Run(os.Args)
	

	
	
	
	i := types.Instance{}
	id :=i.GenID()
	fmt.Println("instance id: ", id, i.Id)
	log.WithFields(log.Fields{
		"id": i.Id,
		"hash": id,
	}).Info("Instance created.")
	hosts := make([]types.Host, 3)
	h1 := types.Host{Hostname: "one.this.thing", HashStart: "000000000000000000000000000000000000000", HashEnd: "113427455640312821154458202477256070485"}
	h1.GenID()
	h2 := types.Host{Hostname: "one.this.thing", HashStart: "113427455640312821154458202477256070486", HashEnd: "226854911280625642308916404954512140970"}
	h2.GenID()
	h3 := types.Host{Hostname: "one.this.thing", HashStart: "226854911280625642308916404954512140971", HashEnd: "340282366920938463463374607431768211456"}
	h3.GenID()
	hosts = append(hosts, h1)
	hosts = append(hosts, h2)
	hosts = append(hosts, h3)
	
	
}
