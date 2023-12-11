package repository

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/boltdb/bolt"
	"go.uber.org/zap"
	"net"
	"os/exec"
	"runtime"
	"smartgw/api/domain"
	"smartgw/lib/database"
	"smartgw/lib/io"
	"strings"
)

type EthernetRepository interface {
	Save(ethernet *domain.Ethernet) error
	Delete(name string) error
	Find(name string) (domain.Ethernet, error)
	FindAll() ([]domain.Ethernet, error)
	Migrate() error
}

var _ EthernetRepository = (*ethernetRepository)(nil)

type ethernetRepository struct {
	db *bolt.DB
}

func NewEthernetRepository(db *bolt.DB) EthernetRepository {
	return &ethernetRepository{
		db: db,
	}
}

func (eth ethernetRepository) Migrate() error {

	inters, err := net.Interfaces()
	if err != nil {
		zap.S().Errorf("获取网卡信息错误 %v", err)
	}
	param := domain.Ethernet{}
	for _, v := range inters {
		param.Index = v.Index
		param.Name = v.Name
		param.MTU = v.MTU
		param.MAC = strings.ToUpper(hex.EncodeToString(v.HardwareAddr))
		param.Flags = v.Flags.String()
		param.IP, param.Netmask = GetIPAndMask(v.Name)
		param.Gateway = GetGateway(v.Name)
		//------------------------------------
		msg, err := eth.Find(v.Name)
		// 找不到
		if err != nil {
			param.ConfigEnabled = false
			param.DHCPEnabled = false
			param.ConfigIP = ""
			param.ConfigNetmask = ""
			param.ConfigGateway = ""
		} else {
			param.ConfigEnabled = msg.ConfigEnabled
			param.DHCPEnabled = msg.DHCPEnabled
			param.ConfigIP = msg.ConfigIP
			param.ConfigNetmask = msg.ConfigNetmask
			param.ConfigGateway = msg.ConfigGateway
		}
		err = eth.Save(&param)
	}

	eths, err := eth.FindAll()
	for _, v := range eths {
		if v.ConfigEnabled == false {
			continue
		}
		if v.DHCPEnabled == false {
			io.CmdSetDHCP(&v)
		} else {
			io.CmdSetStaticIP(&v)
		}
	}
	return err
}

func (eth ethernetRepository) Save(ethernet *domain.Ethernet) error {
	return eth.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(database.Ethernet))

		if data, err := json.Marshal(ethernet); err != nil {
			return err
		} else {
			return b.Put([]byte(ethernet.Name), data)
		}
	})
}

func (eth ethernetRepository) Delete(name string) error {
	return eth.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(database.Ethernet))

		return b.Delete([]byte(name))
	})
}

func (eth ethernetRepository) Find(name string) (domain.Ethernet, error) {
	ethernet := domain.Ethernet{}

	err := eth.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(database.Ethernet))
		data := b.Get([]byte(name))

		if data != nil {
			return json.Unmarshal(data, &ethernet)
		} else {
			return errors.New("没有找到相关数据")
		}
	})
	return ethernet, err
}

func (eth ethernetRepository) FindAll() ([]domain.Ethernet, error) {
	ethernets := make([]domain.Ethernet, 0)

	err := eth.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(database.Ethernet))

		return b.ForEach(func(k, v []byte) error {
			ethernet := domain.Ethernet{}
			err := json.Unmarshal(v, &ethernet)
			if err == nil {
				ethernets = append(ethernets, ethernet)
			}
			return err
		})
	})

	return ethernets, err
}

func GetIPAndMask(name string) (ip, netmask string) {

	inter, err := net.InterfaceByName(name)
	if err != nil {
		return "", ""
	}
	address, _ := inter.Addrs()
	for _, addr := range address {
		if ip, ok := addr.(*net.IPNet); ok && !ip.IP.IsLoopback() {
			if ip.IP.To4() != nil {
				return ip.IP.String(), net.IP(ip.Mask).String()
			}
		}
	}

	return "", ""
}

func GetGateway(name string) string {
	if runtime.GOOS == "linux" {
		out, err := exec.Command("/bin/sh", "-c",
			fmt.Sprintf("route -n | grep %s | grep UG | awk '{print $2}'", name)).Output()
		if err != nil {
			return ""
		}
		return strings.Trim(string(out), "\n")
	}

	return ""
}
