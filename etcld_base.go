package main

import (  
    "fmt"
    "context"
    "log"
    "github.com/coreos/etcd/clientv3"
    "time"
    "strconv"
)

var (  
    dialTimeout    = 2 * time.Second
    requestTimeout = 10 * time.Second
)


type key_value struct{
    key string
    value string
}

type connection interface {
	connect(URL string)
}
type insertion interface {
	insert(k_v key_value)
}
type updation interface {
	update(k_v key_value)
}
type deletion interface {
	delete(k_v key_value)
}
type retreival interface {
	retreive(k_v key_value)
}

func (e *etcd_struct) connect(URL string) {
	cli, _ := clientv3.New(clientv3.Config{
        DialTimeout: dialTimeout,
        Endpoints: []string{URL},
    })
    defer cli.Close()
	kv = clientv3.NewKV(cli)
	e.kv=kv

}

func(e *etcd_struct) insert(k_v key_value){
	ctx, _ := context.WithTimeout(context.Background(), requestTimeout)
	kv:=e.kv
	k:=k_v.key
	v:=k_v.value
    _,err:=kv.Put(ctx, k, v)
    
}

func(e *etcd_struct) update(k_v key_value){
	ctx, _ := context.WithTimeout(context.Background(), requestTimeout)
	kv:=e.kv
	k:=k_v.key
	v:=k_v.value
    _,err:=kv.Put(ctx, k, v)
    
}
func(e *etcd_struct) delete(k_v key_value){
	ctx, _ := context.WithTimeout(context.Background(), requestTimeout)
	kv:=e.kv
	k:=k_v.key
	k_v.value
    kv.Delete(ctx, k, clientv3.WithPrefix())

}

func(e *etcd_struct) retreive(k_v key_value){
	ctx, _ := context.WithTimeout(context.Background(), requestTimeout)
	kv:=e.kv
	k:=k_v.key
	gr,err:=kv.Get(ctx, key)
	fmt.Println("Value: ", string(gr.Kvs[0].Value), "Revision: ", gr.Header.Revision)
    //return string(gr.Kvs[0].Value)
    
}

type etcd_struct struct {
    kv clientv3.KV
}


func main() {
	//ash := Trainer{"Ash", 10, "Pallet Town"}
	kv := new(etcd_struct)
	//temp(&ash)
	connector(kv)
	ash:=key_value{"New key","main"}
	kv.insert(ash)
	a:=key_value{"New Key","base"}
	kv.update(a)
	b:=key_value{"New Key"}
	kv.retreive(b)
	c:=key_value{"New Key"}
	kv.delete(c)
}

/*func temp (arg1 interface { }){
    arg2 := (Trainer*)arg1
    fmt.Printf("%v", arg2)
}*/

func connector(kv connection) {
	URL:="127.0.0.1:2379"
	kv.connect(URL)
}