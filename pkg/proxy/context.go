/*
 * Copyright 2019 storyicon@foxmail.com
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package proxy

import (
    "github.com/gin-gonic/gin"
    "github.com/storyicon/gos/pkg/proxy/module"
)

// Context is used to pass variables between processors
type Context struct {
    engine *Engine
    *module.Path
    *gin.Context
}

func (c *Context) reset() {
    c.Path = nil
    c.Context = nil
}
