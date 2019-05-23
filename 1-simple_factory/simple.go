/**
  * Author: JeffreyBool
  * Date: 2019/5/24
  * Time: 00:34
  * Software: GoLand
*/

package __simple_factory

import (
	"fmt"
)

//API is interface
type Api interface {
	Say(string) string
}

//NewAPI return Api instance by type
func NewApi(t int) Api {
	if t == 1 {
		return &hiAPI{}
	} else if t == 2 {
		return &helloAPI{}
	}
	return nil
}

//hiAPI is one of API implement
type hiAPI struct{}

func (hiAPI) Say(name string) string {
	return fmt.Sprintf("Hi, %s", name)
}

//HelloAPI is another API implement
type helloAPI struct{}

func (helloAPI) Say(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}
