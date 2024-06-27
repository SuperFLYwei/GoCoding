1、eface 不包含方法的接口
type eface struct {
    _type *_type
    data  unsafe.Pointer
}
data指向interface{}实例对象信息的存储地址
_type存放的是类型信息

2、iface 包含方法的接口

type iface struct {
    tab  *itab
    data unsafe.Pointer
}

type itab struct {
    inter  *interfacetype
    _type  *_type
    link   *itab
    hash   uint32 // copy of _type.hash. Used for type switches.
    bad    bool   // type does not implement interface
    inhash bool   // has this itab been added to hash?
    unused [2]byte
    fun    [1]uintptr // variable sized
}

type interfacetype struct {
    typ     _type
    pkgpath name
    mhdr    []imethod
}

type imethod struct {   //这里的 method 只是一种函数声明的抽象，比如  func Print() error
    name nameOff
    ityp typeOff
}

3、接口的特点：
Go的接口实现是非侵入式的，而是鸭子模式。
只要我们实现了接口对应的方法，也就实现了对应的接口。
Golang 的 interface 也是一种多态的体现。



