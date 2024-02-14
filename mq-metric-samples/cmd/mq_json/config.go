package main

/*
  Copyright (c) IBM Corporation 2016, 2022

  Licensed under the Apache License, Version 2.0 (the "License");
  you may not use this file except in compliance with the License.
  You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

  Unless required by applicable law or agreed to in writing, software
  distributed under the License is distributed on an "AS IS" BASIS,
  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  See the License for the specific

   Contributors:
     Mark Taylor - Initial Contribution
*/

import (
	cf "github.com/ibm-messaging/mq-metric-samples/v5/pkg/config"
)

type mqTTYConfig struct {
	cf        cf.Config
	interval  string
	oneline   bool
	recordmax int
}

type ConfigYJson struct {
	Interval  string
	OneLine   bool `yaml:"oneline"`
	RecordMax int  `yaml:"recordmax"`
}

type mqExporterConfigYaml struct {
	Global     cf.ConfigYGlobal
	Connection cf.ConfigYConnection
	Objects    cf.ConfigYObjects
	Filters    cf.ConfigYFilters
	JSON       ConfigYJson `yaml:"json"`
}

var config mqTTYConfig
var cfy mqExporterConfigYaml

/*
initConfig parses the command line parameters.
*/
func initConfig() error {
	var err error

	cf.InitConfig(&config.cf)

	cf.AddParm(&config.interval, "10s", cf.CP_STR, "ibmmq.interval", "json", "interval", "How long between each collection")
	cf.AddParm(&config.oneline, false, cf.CP_BOOL, "ibmmq.oneline", "json", "oneline", "JSON output on a single line")
	cf.AddParm(&config.recordmax, 100, cf.CP_INT, "ibmmq.recordmax", "json", "recordmax", "Max records in a single JSON array")

	err = cf.ParseParms()

	if err == nil {
		if config.cf.ConfigFile != "" {
			err = cf.ReadConfigFile(config.cf.ConfigFile, &cfy)
			if err == nil {
				cf.CopyYamlConfig(&config.cf, cfy.Global, cfy.Connection, cfy.Objects, cfy.Filters)
				config.interval = cf.CopyParmIfNotSetStr("json", "interval", cfy.JSON.Interval)
				config.oneline = cf.CopyParmIfNotSetBool("json", "oneline", cfy.JSON.OneLine)
				config.recordmax = cf.CopyParmIfNotSetInt("json", "recordmax", cfy.JSON.RecordMax)
			}
		}
	}

	if err == nil {
		cf.InitLog(config.cf)
	}

	if err == nil {
		err = cf.VerifyConfig(&config.cf, config)
	}

	if err == nil {
		if config.cf.CC.UserId != "" && config.cf.CC.Password == "" {
			config.cf.CC.Password = cf.GetPasswordFromStdin("Enter password for MQ: ")
		}
	}

	return err

}
