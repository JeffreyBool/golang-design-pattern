/**
  * Author: JeffreyBool
  * Date: 2019/5/24
  * Time: 00:41
  * Software: GoLand
*/

package __simple_factory

import (
	"testing"
)

func TestType1(t *testing.T) {
	api := NewApi(1)
	t.Log(api.Say("Tom"))
}

func TestType2(t *testing.T)  {
	api := NewApi(2)
	t.Log(api.Say("Tom"))
}
