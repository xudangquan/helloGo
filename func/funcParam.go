package main

import "fmt"

type Req struct {
	name string
}

type Resp struct {
	name string
}

func TransportBarcodePay(req *Req) (resp *Resp) {
	return &(Resp{req.name + "jiang"})
}

func TimeoutCtrlWrap(handle func(req *Req) (resp *Resp)) func(req *Req) (resp *Resp) {
	return func(req *Req) (resp *Resp) {
		result := handle(req)
		resp = &(Resp{result.name + req.name + "nan"})
		return
	}
}

func TimeoutCtrlWrap2(handle func(req *Req) (resp *Resp), req *Req) (resp *Resp) {
	result := handle(req)
	resp = &(Resp{result.name + req.name + "nan"})
	return
}

func main() {
	req := &(Req{"xu"})
	ret := TimeoutCtrlWrap(TransportBarcodePay)(req)
	ret2 := TimeoutCtrlWrap2(TransportBarcodePay, req)
	fmt.Println(ret.name)
	fmt.Println(ret2.name)
}
