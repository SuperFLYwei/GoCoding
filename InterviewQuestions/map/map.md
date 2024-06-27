1、map取一个key，然后修改这个值，原map数据的值会不会变化? map属于引用类型，取一个key，然后修改这个值，原map数据的值会发生变化。

2、如何实现一个线程安全的map? 三种方式。 1）加读写锁RWMutex 对查找和遍历加读锁，对增删改加写锁。

2) 分片加锁

3) sync.map 将读写操作分来，来减少锁对性能的影响。 两个推荐场景： a) when the entry for a given key is only ever written once but read many times,
   as in caches that only grow. b) when multiple goroutines read, write, and overwrite entries for disjoint sets of
   keys.

读写分离。读（更新）相关的操作尽量通过不加锁的 read 实现，写（新增）相关的操作通过 dirty 加锁实现。 动态调整。新写入的 key 都只存在 dirty 中，如果 dirty 中的 key 被多次读取，dirty
就会上升成不需要加锁的 read。 延迟删除。Delete 只是把被删除的 key 标记成 nil，新增 key-value 的时候，标记成 enpunged；dirty 上升成 read 的时候，标记删除的 key 被批量移出
map。这样的好处是 dirty 变成 read 之前，这些 key 都会命中 read，而 read 不需要加锁，无论是读还是更新，性能都很高。

3、 Go map的底层原理 线程不安全 安全的map(sync.map) 特性: 1. 无序. 2. 长度不固定. 3. 引用类型. 底层实现: 1.hmap 2.bmap(bucket) hmap中含有n个bmap，是一个数组.
每个bucket又以链表的形式向下连接新的bucket. bucket关注三个字段: 1. 高位哈希值 2. 存储key和value的数组 3. 指向扩容bucket的指针 高位哈希值: 用于寻找bucket中的哪个key. 低位哈希值:
用于寻找当前key属于hmap中的哪个bucket. map的扩容: 当map中的元素增长的时候，Go语言会将bucket数组的数量扩充一倍，产生一个新的bucket数组，并将旧数组的数据迁移至新数组。 加载因子
判断扩充的条件，就是哈希表中的加载因子(即loadFactor)。 加载因子是一个阈值，一般表示为：散列包含的元素数 除以
位置总数。是一种“产生冲突机会”和“空间使用”的平衡与折中：加载因子越小，说明空间空置率高，空间使用率小，但是加载因子越大，说明空间利用率上去了，但是“产生冲突机会”高了。 每种哈希表的都会有一个加载因子，数值超过加载因子就会为哈希表扩容。
Golang的map的加载因子的公式是：map长度 / 2^B(这是代表bmap数组的长度，B是取的低位的位数)阈值是6.5。其中B可以理解为已扩容的次数。
当Go的map长度增长到大于加载因子所需的map长度时，Go语言就会将产生一个新的bucket数组，然后把旧的bucket数组移到一个属性字段oldbucket中。注意：并不是立刻把旧的数组中的元素转义到新的bucket当中，而是，只有当访问到具体的某个bucket的时候，会把bucket中的数据转移到新的bucket中。
map删除: 并不会直接删除旧的bucket，而是把原来的引用去掉，利用GC清除内存。

4、map的key可以是哪些类型？可以嵌套map吗？ golang中的map的 key 可以是很多种类型， 比如 bool, 数字，string, 指针, channel , 还有包含前面几个类型的 interface types,
structs, arrays； map是可以进行嵌套的。

5、map遍历的时候顺序是无序的 在初始化完成之后，调用了runtime.mapiterinit()方法 通过对 mapiterinit 方法阅读，可得知其主要用途是在 map
进行遍历迭代时进行初始化动作。共有三个形参，用于读取当前哈希表的类型信息、当前哈希表的存储信息和当前遍历迭代的数据。 咱们关注到源码中 fastrand 的部分，这个方法名，是不是迷之眼熟。 没错，它是一个生成随机数的方法。再看看上下文.
在这段代码中，它生成了随机数。用于决定从哪里开始循环迭代。更具体的话就是根据随机数，选择一个桶位置作为起始点进行遍历迭代 因此每次重新 for range map，你见到的结果都是不一样的。那是因为它的起始位置根本就不固定！

6、sync.Map sync.Map 采用读写分离和用空间换时间的策略保证 Map 的读写安全
尽量使用原子操作，最大程度上减少了锁的使用，从而接近了“lock free”的效果。
读取、插入和删除的时间复杂度都是 O(1)。

read
read 使用 map[any]*entry 存储数据，本身支持无锁的并发读
read 可以在无锁的状态下支持 CAS 更新，但如果更新的值是之前已经删除过的 entry 则需要加锁操作
由于 read 只负责读取，dirty 负责写入，因此使用 amended 来标记 dirty 中是否包含 read 没有的字段

dirty
dirty 本身就是一个原生 map，需要加锁保证并发写入

entry
read 和 dirty 都是用到 entry 结构
entry 内部只有一个 unsafe.Pointer 指针 p 指向 entry 实际存储的值
指针 p 有三种状态
e.p==nil：entry已经被标记删除，不过此时还未经过read=>dirty重塑，此时可能仍然属于dirty（如果dirty非nil）

e.p==expunged：entry已经被标记删除，经过read=>dirty重塑，不属于dirty，仅仅属于read，下一次dirty=>read升级，会被彻底清理（因为升级的操作是直接覆盖，read中的expunged会被自动释放回收）

e.p==普通指针：此时entry是一个普通的存在状态，属于read，如果dirty非nil，也属于dirty。对应架构图中的normal状态。

1）删除操作的细节，e.p到底是设置成了nil还是expunged？
如果key不在read中，但是在dirty中，则直接delete。
如果key在read中，则逻辑删除，e.p赋值为nil(后续在重塑的时候，nil会变成expunged)



7、golang中两个map对象如何比较？
使用reflect.DeepEqual 这个函数进行比较。使用 reflect.DeepEqual 有一点注意：由于使用了反射，所以有性能的损失。
如果你多做一些测试，那么你会发现 reflect.DeepEqual 会比 == 慢 100 倍以上。
8、map如何顺序读取?
go中map如果要实现顺序读取的话，可以先把map中的key,通过sort包排序。