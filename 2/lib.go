package lib

import (
	"fmt"

	"github.com/hashicorp/consul/acl"
)

func Foo() {
	fmt.Println(acl.DefaultPolicyEnforcementLevel)
}
