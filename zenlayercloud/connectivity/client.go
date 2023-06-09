package connectivity

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

import (
	bmc "github.com/zenlayer/zenlayercloud-sdk-go/zenlayercloud/bmc20221120"
	"github.com/zenlayer/zenlayercloud-sdk-go/zenlayercloud/common"
)

type ZenlayerCloudClient struct {
	SecretKeyId       string
	SecretKeyPassword string
	Domain            string
	Scheme            string
	Timeout           int
	BmcConn           *bmc.Client
}

func (client *ZenlayerCloudClient) WithBmcClient() *bmc.Client {
	if client.BmcConn != nil {
		return client.BmcConn
	}
	config := client.NewConfig()
	client.BmcConn, _ = bmc.NewClient(config, client.SecretKeyId, client.SecretKeyPassword)
	return client.BmcConn
}

func (client *ZenlayerCloudClient) NewConfig() *common.Config {
	config := common.NewConfig()
	config.Timeout = client.Timeout
	config.Scheme = client.Scheme
	config.Domain = client.Domain
	return config
}
