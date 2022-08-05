package parsers

import (
	"io/ioutil"
	"testing"

	"github.com/linuxr/terragraph/models"
	"github.com/spf13/cast"
	"github.com/tidwall/gjson"
)

func TestParseResources(t *testing.T) {
	//fp := "/home/wangyubin/share/tfstate.json" // 2个资源
	fp := "/home/wangyubin/share/terrafrom-graph/cmp-signle/tfstate.json" // 47个资源
	content, err := ioutil.ReadFile(fp)
	if err != nil {
		t.Fatal(err)
	}

	module := gjson.Get(cast.ToString(content), "values.root_module")
	var nodes []models.Node
	nodes, err = ParseResources(module, nodes)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("nodes count: %v\n", len(nodes))
}
