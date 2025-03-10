package xpath

import (
	"bytes"
	"fmt"

	"github.com/antchfx/jsonquery"
	path "github.com/antchfx/xpath"
	"github.com/tinylib/msgp/msgp"
)

type msgpackDocument jsonDocument

func (*msgpackDocument) Parse(buf []byte) (dataNode, error) {
	var json bytes.Buffer

	// Unmarshal the message-pack binary message to JSON and proceed with the jsonquery class
	if _, err := msgp.UnmarshalAsJSON(&json, buf); err != nil {
		return nil, fmt.Errorf("unmarshalling to json failed: %w", err)
	}
	return jsonquery.Parse(&json)
}

func (d *msgpackDocument) QueryAll(node dataNode, expr string) ([]dataNode, error) {
	return (*jsonDocument)(d).QueryAll(node, expr)
}

func (d *msgpackDocument) CreateXPathNavigator(node dataNode) path.NodeNavigator {
	return (*jsonDocument)(d).CreateXPathNavigator(node)
}

func (d *msgpackDocument) GetNodePath(node, relativeTo dataNode, sep string) string {
	return (*jsonDocument)(d).GetNodePath(node, relativeTo, sep)
}
func (d *msgpackDocument) GetNodeName(node dataNode, sep string, withParent bool) string {
	return (*jsonDocument)(d).GetNodeName(node, sep, withParent)
}

func (d *msgpackDocument) OutputXML(node dataNode) string {
	return (*jsonDocument)(d).OutputXML(node)
}
