package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/api"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/kubelet/envvars"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/master"
	ketcd "github.com/GoogleCloudPlatform/kubernetes/pkg/registry/etcd"
	"github.com/coreos/go-etcd/etcd"
	"log"
	"os"
	"strings"
)

var (
	flags       = flag.NewFlagSet("kubevars", flag.ExitOnError)
	etcdAddress string
	format      string
	formatter   func([]api.EnvVar) string
)

func main() {
	parseFlags()

	switch format {
	case "docker":
		formatter = func(vars []api.EnvVar) string {
			var r []string
			for _, v := range vars {
				r = append(r, fmt.Sprintf("-e %s=%v", v.Name, v.Value))
			}
			return strings.Join(r, "\n")
		}
	case "plain":
		formatter = func(vars []api.EnvVar) string {
			var r []string
			for _, v := range vars {
				r = append(r, fmt.Sprintf("%s=%v", v.Name, v.Value))
			}
			return strings.Join(r, "\n")
		}
	case "json":
		formatter = func(vars []api.EnvVar) string {
			bin, err := json.Marshal(vars)
			if err != nil {
				log.Fatal(err)
			}
			return string(bin)
		}
	default:
		flags.Usage()
		os.Exit(1)
	}

	client := etcd.NewClient([]string{etcdAddress})
	helper, err := master.NewEtcdHelper(client, "")
	if err != nil {
		log.Fatal(err)
	}
	registry := ketcd.NewRegistry(helper, nil, nil)
	services, err := registry.ListServices(api.NewContext())
	if err != nil {
		log.Fatal(err)
	}
	vars := envvars.FromServices(services)
	fmt.Println(formatter(vars))
}

func parseFlags() {
	flags.StringVar(&etcdAddress, "etcd", "http://localhost:4001", "The ETCD endpoint")
	flags.StringVar(&format, "format", "docker", "One of: docker, json, plain")
	flags.Usage = func() {
		fmt.Fprint(os.Stderr,
			`Usage: kubevars [-etcd http://host[:port]] [-format <docker|json|plain>]
Flags:
`)
		flags.PrintDefaults()
	}
	flags.Parse(os.Args[1:])
}
