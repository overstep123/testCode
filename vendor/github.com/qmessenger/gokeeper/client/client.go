package client

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/johntech-o/gorpc"
	"github.com/qmessenger/gokeeper/model"
)

//
var (
	Debug            = false
	Stdout io.Writer = os.Stdout
	Stderr io.Writer = os.Stderr

	errUsage = errors.New("gokeeper usage: ./component/bin -n=nodeID -d=domain -k=keeper_address")
)

// Client ...
type Client struct {
	node *model.Node
	rpc  *gorpc.Client

	objContainer model.ObjContainer
	data         map[string]interface{}
	callback     []func()
}

// New return client struct
func New(keeperAddr, domain, nodeID, component string, rawSubscription []string, tags map[string]string) *Client {
	for k := range rawSubscription {
		rawSubscription[k] = filepath.Join("/", rawSubscription[k])
	}
	hostname, _ := os.Hostname()
	nodeInfo := model.NewNodeInfo(nodeID, hostname, keeperAddr, domain, component, rawSubscription, tags)
	c := &Client{
		node: model.NewNode(*nodeInfo),
		rpc:  gorpc.NewClient(gorpc.NewNetOptions(ConnectTimeout, ReadTimeout, WriteTimeout)),
		data: map[string]interface{}{},
	}
	return c
}

// LoadData register data
func (c *Client) LoadData(objContainer model.ObjContainer) *Client {
	c.objContainer = objContainer
	s := c.objContainer.GetStructs()
	for k, v := range s {
		c.data[k] = v
	}
	return c
}

// RegisterCallback event
func (c *Client) RegisterCallback(args ...func()) *Client {
	for _, v := range args {
		c.callback = append(c.callback, v)
	}
	return c
}

// Work get data from keeper service and listen data change
func (c *Client) Work() error {
	if len(c.data) == 0 {
		Stdout.Write([]byte("gokeeper did not load any data (forgotten LoadData?) \n"))
	}
	if c.node.GetKeeperAddr() == "" || c.node.GetID() == "" || c.node.GetDomain() == "" {
		return errUsage
	}

	// 第一次必须阻塞式加载数据
	evtReq := model.Event{
		EventType: model.EventNodeRegister,
		Data:      c.node.Info(),
	}

	if Debug {
		Stdout.Write([]byte(fmt.Sprintf("%s|gokeeper|Work|event request|%#v \n", time.Now().String(), evtReq)))
	}

	evtResp := model.NewEvent()
	if err := c.rpc.CallWithAddress(c.node.GetKeeperAddr(), "Server", "Sync", &evtReq, &evtResp); err != nil {
		return errors.New(err.Error())
	}

	if Debug {
		Stdout.Write([]byte(fmt.Sprintf("%s|gokeeper|Work|event response|%#v \n", time.Now().String(), evtResp)))
	}

	if err := c.eventParser(evtResp); err != nil {
		return err
	}

	go c.eventLoop()
	go SignalNotifyDeamon(c)
	return nil
}

func (c *Client) eventParser(evt model.Event) error {
	switch evt.EventType {
	case model.EventNone:
		return nil
	default:
		if err := eventCallback(evt.EventType, c, evt); err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) eventLoop() {
	var gerr *gorpc.Error
	var err error
	for {
		evtReq := model.Event{EventType: model.EventNone, Data: c.node.Info()}

		select {
		case evtReq = <-c.node.Event:
		default:
		}

		if Debug {
			Stdout.Write([]byte(fmt.Sprintf("%s|gokeeper|eventLoop|event request|%#v \n", time.Now().String(), evtReq)))
		}

		evtResp := model.NewEvent()
		if gerr = c.rpc.CallWithAddress(c.node.GetKeeperAddr(), "Server", "Sync", &evtReq, &evtResp); gerr != nil {
			Stderr.Write([]byte(fmt.Sprintf("%s|gokeeper|eventLoop|Server|Sync|%s \n", time.Now().String(), gerr.Error())))
			time.Sleep(EventInterval)
			continue
		}

		if Debug {
			Stdout.Write([]byte(fmt.Sprintf("%s|gokeeper|eventLoop|event response|%#v \n", time.Now().String(), evtResp)))
		}

		if err = c.eventParser(evtResp); err != nil {
			Stderr.Write([]byte(fmt.Sprintf("%s|gokeeper|eventLoop|eventParser|%s \n", time.Now().String(), err.Error())))
			continue
		}
	}
}
