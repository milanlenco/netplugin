package vpppolicy

import "fmt"

// PolicyRule rules the world
type PolicyRule struct {
	Rule *VppnetPolicyRule // rule definition
}

// VppPolicy should be me
type VppPolicy struct {
	Rules map[string]*PolicyRule // rules database
}

// AddRule  adds a policy to the vpp driver
func AddRule(rule *VppnetPolicyRule) error {
	fmt.Println(rule)
	return nil
}

// DelRule  deletes a policy to the vpp driver
func DelRule(rule *VppnetPolicyRule) error {
	fmt.Println(rule)
	return nil
}
