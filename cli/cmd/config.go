package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "View and modifies the config file",
}

func init() {
	RootCmd.AddCommand(configCmd)
}

type clusterConfig struct {
	Server string `yaml:"server"`
	Token  string `yaml:"token"`
}

type configFile struct {
	Version        string                   `yaml:"version"`
	Clusters       map[string]clusterConfig `yaml:"clusters"`
	CurrentCluster string                   `yaml:"current_cluster"`
}

// read the config file from disk
func readConfigFile(f string) (c *configFile, err error) {
	y, err := ioutil.ReadFile(f)
	if err != nil {
		if os.IsNotExist(err) {
			log.WithError(err).WithField("fileName", f).Debug("File not found when trying to read the config file")
		} else {
			log.WithError(err).Error("Error loading the config file")
		}
		return
	}
	conf := configFile{}
	if err = yaml.Unmarshal(y, &conf); err != nil {
		log.WithError(err).WithField("cfgFile", f).Error("Error trying to unmarshal the config file")
		return nil, err
	}
	return &conf, nil
}

// return the config file loaded from disk or creates a new one (empty with the base needs)
func readOrCreateConfigFile(f string) (c *configFile, err error) {
	if c, err = readConfigFile(cfgFile); err == nil || !os.IsNotExist(err) {
		return
	}

	// set defaults
	log.Debug("Config file not found... creating the base one")
	conf := configFile{Version: version, Clusters: make(map[string]clusterConfig)}
	return &conf, nil
}

// parse the config object to yaml (byte array pointer)
func marshalConfigFile(c *configFile) (b *[]byte, err error) {
	z, err := yaml.Marshal(&c)
	if err != nil {
		log.WithError(err).WithField("config", c).Error("Error marshaling the config file")
		return nil, err
	}
	return &z, nil
}

// write the config file to disk in yaml format
func writeConfigFile(f string, c *configFile) error {
	// TODO: implement validate before writing
	log.WithField("fileName", f).WithField("config", *c).Debug("Marshaling the config file to save")
	b, err := marshalConfigFile(c)
	if err != nil {
		return err
	}
	// get config basepath
	p := filepath.Dir(f)
	d, err := os.Stat(p)
	if err != nil {
		if !os.IsNotExist(err) {
			log.WithError(err).WithField("directory", p).Error("Failed to check if the directory exists")
			return err
		}
		log.WithField("directory", p).Debug("Config basepath not found... creating")
		if err = os.MkdirAll(p, 0755); err != nil {
			log.WithError(err).WithField("directory", p).Error("Failed to create the directory")
			return err
		}
	} else if !d.IsDir() {
		log.WithField("directory", p).Error("Path exists, but isn't a directory")
		return errors.New("Path exists, but isn't a directory")
	}
	if err = ioutil.WriteFile(f, *b, 0600); err != nil {
		log.WithError(err).Error("Error while writing the config file to disk")
	}
	return nil
}

// get the name of the current cluster in the config file
func getCurrentClusterName() (n string, err error) {
	n = viper.GetString("current_cluster")
	if n == "" {
		log.Debug("Cluster not set yet")
		err = newSysError("Set a cluster to use before continue")
	}
	return
}

// return the current cluster
func getCurrentCluster() (c *clusterConfig, err error) {
	n, err := getCurrentClusterName()
	if err != nil {
		return
	}
	k := fmt.Sprintf("clusters.%s", n)
	if err := viper.UnmarshalKey(k, &c); err != nil {
		log.WithError(err).Fatal("Erro trying to unmarshal the current cluster")
		return nil, errors.New("Erro trying to unmarshal the current cluster")
	}
	return
}

// return the config file yaml
func getConfigFileYaml(f string) (y string, err error) {
	c, err := readOrCreateConfigFile(f)
	if err != nil {
		return
	}
	b, err := marshalConfigFile(c)
	if err != nil {
		return
	}
	y = string((*b)[:])
	return
}
