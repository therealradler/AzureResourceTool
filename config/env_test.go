package config

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestParseEnvironment(t *testing.T) {
	//os.Setenv("AZURE_SUBSCRIPTION_ID", "f135b611-b863-4f6a-a5ec-507b1111641b")
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0])
	}

	ParseEnvironment()
}
