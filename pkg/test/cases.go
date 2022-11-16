package test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type Suite interface {
	Run(name string, subtest func()) bool
	SetupTest()
	T() *testing.T
}

type Case[RQ, RS any] struct {
	Name        string
	Prepare     func()
	Req         RQ
	ExpectedErr error
	ValidateRes func(res RS)
}

func (c Case[RQ, RS]) Validate(s Suite, res RS, err error) {
	require.Equal(s.T(), c.ExpectedErr, err)

	if c.ValidateRes == nil {
		require.Nil(s.T(), res)
	} else {
		require.NotNil(s.T(), res)
		c.ValidateRes(res)
	}
}

type Cases[RQ, RS any] []*Case[RQ, RS]

func (cs Cases[RQ, RS]) Test(s Suite, process func(c *Case[RQ, RS])) {
	for i, c := range cs {
		c := c

		if i > 0 {
			s.SetupTest()
		}

		s.Run(c.Name, func() {
			if c.Prepare != nil {
				c.Prepare()
			}

			process(c)
		})
	}
}
