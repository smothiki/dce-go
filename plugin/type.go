/*
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

//go:generate go-extpoints . ComposePlugin
package plugin

import (
	"context"

	"github.com/mesos/mesos-go/executor"
	mesos "github.com/mesos/mesos-go/mesosproto"
)

type ComposePlugin interface {
	PreLaunchTask(ctx *context.Context, composeFiles *[]string, executorId string, taskInfo *mesos.TaskInfo) error
	PostLaunchTask(ctx *context.Context, composeFiles []string, taskInfo *mesos.TaskInfo) (string, error)
	PreKillTask(taskInfo *mesos.TaskInfo) error
	PostKillTask(taskInfo *mesos.TaskInfo) error
	Shutdown(executor.ExecutorDriver) error
}
