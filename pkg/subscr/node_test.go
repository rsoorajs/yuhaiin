package subscr

import (
	"context"
	"encoding/json"
	"log"
	"testing"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestMarshalJson(t *testing.T) {
	s := &Point{
		NHash:   "n_hash",
		NName:   "n_name",
		NGroup:  "n_group",
		NOrigin: Point_manual,
		Node: &Point_Shadowsocksr{
			Shadowsocksr: &Shadowsocksr{
				Server: "server",
			},
		},
	}

	ss, err := protojson.Marshal(s)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	t.Log(string(ss))
	zz := &Point{}
	err = protojson.Unmarshal(ss, zz)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Log(zz)

	s.Node = &Point_Vmess{
		Vmess: &Vmess{
			Address: "address",
		},
	}

	ss, err = protojson.Marshal(s)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	t.Log(string(ss))
	err = protojson.Unmarshal(ss, zz)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Log(zz)
}

func TestNodeManager(t *testing.T) {
	n, err := NewNodeManager("/tmp/yuhaiin/nodeManagerTest/config.json")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	_, err = n.AddLink(context.TODO(), &NodeLink{
		Name: "test",
		Type: "test",
		Url:  "test",
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	// _, err = n.RefreshSubscr(context.TODO(), &emptypb.Empty{})
	// if err != nil {
	// t.Error(err)
	// t.FailNow()
	// }
	hash := "db084f1d4f90140540e47a13ca77204d1f597e933481d58dfe2e5860f76f75ff"
	t.Log(n.GetNode(context.TODO(), &wrapperspb.StringValue{Value: hash}))
	t.Log(n.Latency(context.TODO(), &wrapperspb.StringValue{Value: hash}))
	// t.Log(n.node)
}

func TestDelete(t *testing.T) {
	a := []string{"a", "b", "c"}

	for i := range a {
		if a[i] != "b" {
			continue
		}

		log.Println(i, a[:i], a[i:])
		a = append(a[:i], a[i+1:]...)
		break
	}

	t.Log(a)
}

func TestMarshalMap(t *testing.T) {
	s := &Point{
		NHash:   "n_hash",
		NName:   "n_name",
		NGroup:  "n_group",
		NOrigin: Point_manual,
		Node: &Point_Shadowsocksr{
			Shadowsocksr: &Shadowsocksr{
				Server: "server",
			},
		},
	}

	data, _ := protojson.MarshalOptions{UseProtoNames: true, EmitUnpopulated: true}.Marshal(s)

	var z map[string]interface{}

	err := json.Unmarshal(data, &z)
	if err != nil {
		t.Error(err)
	}

	t.Log(z)

	for k, v := range z {
		t.Log(k)
		switch x := v.(type) {
		case string:
			t.Log("string", x)
		case map[string]interface{}:
			t.Log("map[string]interface{}", x)
			x["server"] = "server2"
		}
	}

	b, err := json.Marshal(z)
	if err != nil {
		t.Error(err)
	}

	t.Log(string(b))

	err = protojson.Unmarshal(b, s)
	if err != nil {
		t.Error(err)
	}

	t.Log(s)
}

func TestReflect(t *testing.T) {
	z := &Point{Node: &Point_Shadowsocksr{Shadowsocksr: &Shadowsocksr{}}}

	v, err := z.Conn()
	if err != nil {
		t.Error(err)
	}

	t.Logf("%#v", v)

	z = &Point{Node: &Point_Shadowsocksr{}}

	v, err = z.Conn()
	if err != nil {
		t.Error(err)
	}

	t.Logf("%#v", v)

	z = &Point{}

	v, err = z.Conn()
	if err != nil {
		t.Error(err)
	}

	t.Logf("%#v", v)
}
