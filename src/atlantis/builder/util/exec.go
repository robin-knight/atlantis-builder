/* Copyright 2014 Ooyala, Inc. All rights reserved.
 *
 * This file is licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
 * except in compliance with the License. You may obtain a copy of the License at
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License is
 * distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and limitations under the License.
 */

package util

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
)

func EchoExec(cmd *exec.Cmd) []byte {
     return EchoExecCanSkipError(cmd, false)
}

func EchoExecCanSkipError(cmd *exec.Cmd, skipErr bool) []byte {
	// make streaming copies of stdout
	var buf bytes.Buffer
	outWriter := io.MultiWriter(&buf, os.Stdout)

	cmd.Stderr = os.Stderr
	cmd.Stdout = outWriter

	if err := cmd.Start(); err != nil {
		panic(err)
	}
	if err := cmd.Wait(); err != nil {
		if skipErr {
			fmt.Printf("cmd execution failed but error ignored\n")
		} else {
			panic(err)
		}

	}

	return buf.Bytes()
}
