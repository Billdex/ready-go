package unique_id

import "github.com/bwmarrin/snowflake"

var node, _ = snowflake.NewNode(1)

func SnowflakeSetup(n int64) error {
	var err error
	node, err = snowflake.NewNode(n)
	return err
}

func SnowflakeNodeID() snowflake.ID {
	return node.Generate()
}

func SnowflakeIntID() int64 {
	return node.Generate().Int64()
}

func SnowflakeStrID() string {
	return node.Generate().String()
}
