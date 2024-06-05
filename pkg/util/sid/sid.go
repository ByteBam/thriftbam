package sid

import "github.com/bwmarrin/snowflake"

type Sid struct {
	sf *snowflake.Node
}

func NewSid() *Sid {
	node, err := snowflake.NewNode(1)
	if err != nil {
		panic("snowflake not created")
	}
	return &Sid{node}
}

func (s Sid) GenString() string {
	return s.sf.Generate().String()
}

func (s Sid) GenInt64() int64 {
	return s.sf.Generate().Int64()
}
