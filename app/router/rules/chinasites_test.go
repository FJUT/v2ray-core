package rules_test

import (
	"testing"

	. "v2ray.com/core/app/router/rules"
	v2net "v2ray.com/core/common/net"
	"v2ray.com/core/testing/assert"
)

func makeDomainDestination(domain string) v2net.Destination {
	return v2net.TCPDestination(v2net.DomainAddress(domain), 80)
}

func TestChinaSites(t *testing.T) {
	assert := assert.On(t)

	rule := NewChinaSitesRule("tag")
	assert.Bool(rule.Apply(makeDomainDestination("v.qq.com"))).IsTrue()
	assert.Bool(rule.Apply(makeDomainDestination("www.163.com"))).IsTrue()
	assert.Bool(rule.Apply(makeDomainDestination("ngacn.cc"))).IsTrue()
	assert.Bool(rule.Apply(makeDomainDestination("12306.cn"))).IsTrue()

	assert.Bool(rule.Apply(makeDomainDestination("v2ray.com"))).IsFalse()
}
