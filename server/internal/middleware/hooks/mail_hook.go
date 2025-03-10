// Copyright © 2023 OpenIM open source community. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package hooks

import (
	"fmt"

	"github.com/gin-gonic/gin"

	urltrie "github.com/OpenIMSDK/OpenKF/server/internal/middleware/hooks/url_trie"
	"github.com/OpenIMSDK/OpenKF/server/pkg/log"
)

func init() {
	urltrie.RegisterHook(MailHook{})
	fmt.Println("RegisterHook", "Register Hook[MailHook] success...")
}

// MailHook implement urltrie.Hook.
type MailHook struct {
	urltrie.Hook
}

// Pattern return pattern.
func (h MailHook) Pattern() string {
	return "/api/v1/register/email/code"
}

// BeforeRun do something before controller run.
func (h MailHook) BeforeRun(c *gin.Context) {
	log.Debugf("GlobalHook", "path: %v", c.Request.URL.Path)
	c.Next()
}

// AfterRun do something after controller run.
func (h MailHook) AfterRun(c *gin.Context) {
	c.Next()
}
