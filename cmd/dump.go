package cmd

import (
	"bytes"
	"context"
	"encoding/gob"
	"errors"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/urfave/cli"
	"go.etcd.io/etcd/clientv3"
)

const (
	defaultPrefix      = "/"
	defaultEtcdAddress = "localhost:2379"
)

func dumpCmd() cli.Command {
	return cli.Command{
		Name:   "dump",
		Usage:  "dump K/V with prefix",
		Action: dumpAction,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:     "address, a",
				Usage:    "etcd address",
				Value:    defaultEtcdAddress,
				Required: false,
			},
			cli.StringFlag{
				Name:     "prefix, p",
				Usage:    "key prefix",
				Value:    defaultPrefix,
				Required: false,
			},
			cli.StringFlag{
				Name:     "output, o",
				Usage:    "Output to `FILE`",
				Required: false,
			},
			cli.BoolFlag{
				Name:     "silent, s",
				Usage:    "verbose mode",
				Required: false,
			},
		},
	}
}

type kvData = []byte
type dumpData = []kvData

func dumpAction(c *cli.Context) error {
	address := c.String("address")
	if len(address) == 0 {
		return errors.New("address shouldn't be empty")
	}

	prefix := c.String("prefix")
	if len(prefix) == 0 {
		return errors.New("prefix shouldn't be empty")
	}

	silent := c.Bool("silent")
	dd, err := dump(address, prefix, !silent)
	if err != nil {
		return err
	}

	out := c.String("output")
	if len(out) != 0 {
		err = writeDumpData(out, dd)
		if err != nil {
			return err
		}
	}

	return nil
}

func dump(addr, prefix string, print bool) (dumpData, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{addr},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		return nil, err
	}
	defer cli.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	rsp, err := cli.Get(ctx, prefix, clientv3.WithPrefix())
	cancel()
	if err != nil {
		return nil, err
	}

	dd := make(dumpData, 0)

	for _, kv := range rsp.Kvs {
		b, err := kv.Marshal()
		if err != nil {
			return nil, err
		}

		dd = append(dd, b)

		if print {
			fmt.Println(string(kv.Key))
		}

	}

	return dd, nil
}

func writeDumpData(filename string, d dumpData) error {
	var buffer bytes.Buffer
	enc := gob.NewEncoder(&buffer)
	err := enc.Encode(d)
	if err != nil {
		return err
	}

	b := buffer.Bytes()
	err = ioutil.WriteFile(filename, b, 0644)
	if err != nil {
		return err
	}

	return nil
}
