package util

import (
	"encoding/json"
	"flag"
	"os"
	goruntime "runtime"
	"strings"
	"testing"

	"golang.org/x/net/context"

	"github.com/coreos/pkg/capnslog"

	etcdclient "github.com/coreos/etcd/client"
	etcdclientv3 "github.com/coreos/etcd/clientv3"

	etcdtest "k8s.io/apiserver/pkg/storage/etcd/testing"
	kapi "k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/capabilities"

	serveretcd "github.com/openshift/origin/pkg/cmd/server/etcd"
)

func init() {
	capabilities.SetForTests(capabilities.Capabilities{
		AllowPrivileged: true,
	})
	flag.Set("v", "5")
	if len(os.Getenv("OS_TEST_VERBOSE_ETCD")) > 0 {
		capnslog.SetGlobalLogLevel(capnslog.DEBUG)
		capnslog.SetFormatter(capnslog.NewGlogFormatter(os.Stderr))
	} else {
		capnslog.SetGlobalLogLevel(capnslog.INFO)
		capnslog.SetFormatter(capnslog.NewGlogFormatter(os.Stderr))
	}
}

// url is the url for the launched etcd server
var url string

// RequireEtcd verifies if the etcd is running and accessible for testing
func RequireEtcd(t *testing.T) *etcdtest.EtcdTestServer {
	s := etcdtest.NewUnsecuredEtcdTestClientServer(t)
	url = s.Client.Endpoints()[0]
	return s
}

func RequireEtcd3(t *testing.T) *etcdtest.EtcdTestServer {
	s, _ := etcdtest.NewUnsecuredEtcd3TestClientServer(t, kapi.Scheme)
	url = s.V3Client.Endpoints()[0]
	return s
}

func NewEtcdClient() etcdclient.Client {
	client, _ := MakeNewEtcdClient()
	return client
}

func MakeNewEtcdClient() (etcdclient.Client, error) {
	etcdServers := []string{GetEtcdURL()}

	cfg := etcdclient.Config{
		Endpoints: etcdServers,
	}
	client, err := etcdclient.New(cfg)
	if err != nil {
		return nil, err
	}
	return client, serveretcd.TestEtcdClient(client)
}

func NewEtcd3Client() *etcdclientv3.Client {
	client, _ := MakeNewEtcd3Client()
	return client
}

func MakeNewEtcd3Client() (*etcdclientv3.Client, error) {
	etcdServers := []string{GetEtcdURL()}

	cfg := etcdclientv3.Config{
		Endpoints: etcdServers,
	}
	client, err := etcdclientv3.New(cfg)
	if err != nil {
		return nil, err
	}
	return client, serveretcd.TestEtcdClientV3(client)
}

func GetEtcdURL() string {
	if len(url) == 0 {
		panic("can't invoke GetEtcdURL prior to calling RequireEtcd")
	}
	return url
}

func DumpEtcdOnFailure(t *testing.T) {
	if !t.Failed() {
		return
	}

	client := NewEtcdClient()
	keyClient := etcdclient.NewKeysAPI(client)

	response, err := keyClient.Get(context.Background(), "/", &etcdclient.GetOptions{Recursive: true, Sort: true})
	if err != nil {
		t.Logf("error dumping etcd: %v", err)
		return
	}

	writeEtcdDump(t, response.Node)
}

func DumpEtcd3OnFailure(t *testing.T) {
	if !t.Failed() {
		return
	}

	client := NewEtcd3Client()

	response, err := client.KV.Get(context.Background(), "/", etcdclientv3.WithPrefix(), etcdclientv3.WithSort(etcdclientv3.SortByKey, etcdclientv3.SortDescend))
	if err != nil {
		t.Logf("error dumping etcd: %v", err)
		return
	}

	kvData := []etcd3kv{}
	for _, kvs := range response.Kvs {
		obj, _, err := kapi.Codecs.UniversalDeserializer().Decode(kvs.Value, nil, nil)
		if err != nil {
			t.Logf("error decoding value %s: %v", string(kvs.Value), err)
			continue
		}
		objJSON, err := json.Marshal(obj)
		if err != nil {
			t.Logf("error encoding object %#v as JSON: %v", obj, err)
			continue
		}
		kvData = append(kvData, etcd3kv{string(kvs.Key), string(objJSON)})
	}

	writeEtcdDump(t, kvData)
}

func writeEtcdDump(t *testing.T, data interface{}) {
	jsonResponse, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		t.Logf("error encoding etcd dump: %v", err)
		return
	}
	name := getCallingTestName()
	t.Logf("dumping etcd to %q", GetBaseDir()+"/etcd-dump-"+name+".json")
	dumpFile, err := os.OpenFile(GetBaseDir()+"/etcd-dump-"+name+".json", os.O_WRONLY|os.O_CREATE, 0444)
	if err != nil {
		t.Logf("error writing etcd dump: %v", err)
		return
	}
	defer dumpFile.Close()
	_, err = dumpFile.Write(jsonResponse)
	if err != nil {
		t.Logf("error writing etcd dump: %v", err)
		return
	}
}

func getCallingTestName() string {
	pc := make([]uintptr, 10)
	goruntime.Callers(4, pc)
	f := goruntime.FuncForPC(pc[0])
	last := strings.LastIndex(f.Name(), "Test")
	if last == -1 {
		last = 0
	}
	return f.Name()[last:]
}

type etcd3kv struct {
	Key, Value string
}
