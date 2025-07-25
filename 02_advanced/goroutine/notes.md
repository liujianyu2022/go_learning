## 进程
进程是程序的一次执行实例。操作系统以进程为单位分配CPU时间、内存等资源


### 线程
线程是进程内的执行单元，是CPU调度的基本单位。多个线程共享一个进程中的资源
线程是CPU实际调度的对象，多核CPU可以并行执行多个线程

### 协程
协程是用户态的轻量级线程，由程序控制调度（而非操作系统）。
协程包括执行的程序以及执行的状态。因此协程的本质是一段包含了运行状态的程序。
一个线程上可以执行多个协程，只需要在线程上进行协程的切换即可，避免了cpu在多个线程之间的切换，而cpu在多个线程之间的切换是比较耗费时间和资源的

### 线程 VS 协程
1. 用户态 和 内核态：
协程：完全在用户态由程序自行调度，不涉及操作系统内核。切换时无需陷入内核态，没有系统调用的开销。
线程：由操作系统内核调度，切换时需要从用户态切换到内核态，涉及CPU寄存器保存、内核栈切换等操作，开销较大。

2. 资源占用更小
协程栈：通常是预分配的少量内存（如KB级别），且可以动态增长（如Go的协程初始栈仅2KB）。
线程栈：由操作系统分配，默认较大（如Linux下通常为2~8MB），大量线程会消耗大量内存。

3. 无锁编程与协作式调度
协程：通过协作式调度（Coroutine主动让出控制权），避免了线程的抢占式调度带来的竞争条件。减少了对锁的需求，降低了同步开销。
线程：内核的抢占式调度可能导致线程在任何时间点被中断，需要锁/原子操作保证安全，锁竞争会显著降低性能。

注意点：如果任务完全是计算密集型的（无I/O等待），协程的优势会消失，甚至可能因单线程无法利用多核而变慢（此时需要多线程+协程结合，如Go的GOMAXPROCS）