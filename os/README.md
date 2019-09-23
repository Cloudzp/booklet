# 操作系统
是一个大型的系统程序
- 提供用户接口，方便用户控制程序；
- 负责为应用程序分配及调度资源，并且要控制与协调应用程序的并发运行，帮助用户存储及保护信息；

## 操作系统功能

1. 进程管理功能（核心）
2. 内存管理功能（核心）
3. 设备管理功能
4. 文件管理功能
5. 网络管理功能
6. 分布式管理功能
7. 图形界面管理功能


## 操作系统发展史
- 用户需求不断提升
- 硬件技术进步

### 四个典型阶段
- 电子管时代
- 晶体管时代
- 集成电路时代
- 大规模集成电路时代

- 手工操作
- 单批处理系统
- 多到批处理系统
- 分时系统

### MINIXOS
// TODO 看看这本书 www.mini3.org

不同版本的linux公用同一个内核

## 2.1 操作系统的逻辑结构
### 1. 逻辑结构
- os的设计和实现思路

### 2. cpu的态（Mode）
- CPU的工作状态
- 对资源和指令使用权限的描述

态的分类
- 核态
  能够访问所有的资源和指令
- 用户态
  权限受限

用户态和核态之间的转换

存储器
按照读取方式分类
- RAM ：随机存储器 Random-Access Memory,也叫主存，是与CPU中介交换数据的内部存储器，它可以随时读写，速度很快
- ROM ：只读存储器 Read-Only Memory，是一种只能读取事先所存数据的固态半导体存储器。其特征是一单存储资料就无法再降至改变或删除1.

按照存储元的材料
- 半导体存储器
- 磁存储器
- 光存储器

按照Cpu的联系
- 主存：直接与cpu交换信息
- 辅存： 不能与cpu直接交换信息

存储体系
实际存储体系：
- 寄存器
- 告诉缓存
- 高速存
- 主存
- 辅存

Cpu读取指令或者数据的访问顺序
1. 访问缓存（命中）
2. 访问内存（没有命中）
3. 访问辅存（缺页）

2.3 中断机制
- 指CPU对突发的外部事件的反应过程或者机制。
- CPU收到外部信号后，停止当前工作，专区处理该外部事件，处理完外部后回到原来的工作的中断处继续原来的工作。 

2.3.1 引入中断的目的
- 实现并发活动
- 实现实时处理
- 故障自动处理

2.3.2 中断源和中断类型
- 引起系统中断的事件成为中断源。
- 中断类型
- 强迫性中断和资源中断
  - 强迫性中断： 程序没有预期；
  - 资源中断： 程序有预期的；
  - 外中断和内终端
    - 外中断： 由Cpu外部事件引起，例如：I/O,外部事件
    - 内部中断： 由CPU内部事件引起，例： 访问中断，程序中断
- 断点：
  - 程序中断的地方，将要执行的下一指令的地址
  - CS.IP
- 现场
  - 程序正确运行所依赖的·信息集合。
    - 相关寄存器
- 现场的 两个处理过程
  - 现场的保护： 进入中断服务程序之前，栈
  - 现场数据的恢复： 退出中断程序之后，栈
- 中断响应的过程
  - 1） 识别中断源
  - 2） 保护中断和现场
  - 3） 装入·1中断程序的入口地址
  - 4） 进入中断程序
  - 5） 恢复现场和断点
  - 6） 中断返回

- 中断响应的实质
  - 交换指令执行地址
  - 交换CPU的态
  - 工作

## 第三章 操作系统用户界面
### 3.1 操作系统启动过程
#### 3.1.1 BIOS和主引导记录
- 实模式（实地址模式，REAL MODE）
  - 程序按照8086寻址方式访问0h-FFFFFh

- 保护模式（内存保护模式，Protcet mode）
  - 寻址方式：
  - 端也是寻址机制
  - 虚拟地址

- 系统BIOS
  - basic I/O System(Firmware, 固件)
    - 基本输入/输出系统
    - 位置：F0000-FFFFF
    - 功能：
      - 系统启动配置
      - 基本的设备I/O服务
      - 系统的家电自检和启动
-        