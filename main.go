/* Copyright 2021 Jan Tytgat

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"fmt"
	"github.com/jantytgat/cnsbackup-config/cmd"
)

var banner = "   ___ _ _       _         _   ___   ___   ___          _             \n  / __(_) |_ _ _(_)_ __   /_\\ |   \\ / __| | _ ) __ _ __| |___  _ _ __ \n | (__| |  _| '_| \\ \\ /  / _ \\| |) | (__  | _ \\/ _` / _| / / || | '_ \\\n  \\___|_|\\__|_| |_/_\\_\\ /_/ \\_\\___/ \\___| |___/\\__,_\\__|_\\_\\\\_,_| .__/\n                                                                |_|   "

func main() {
	fmt.Println(banner)
	cmd.Execute()
}
