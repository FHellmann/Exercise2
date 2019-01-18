# Mutex and Channel basics

### What is an atomic operation?
> Atomic operations in concurrent programming are program operations that run completely independently of any other processes.

### What is a semaphore?
> A semaphore is a variable or abstract data type used to control access to a common resource by multiple processes in a concurrent system such as a multitasking operating system.

### What is a mutex?
> A lock or mutex is a synchronization mechanism for enforcing limits on access to a resource in an environment where there are many threads of execution. A lock is designed to enforce a mutual exclusion concurrency control policy.

### What is the difference between a mutex and a binary semaphore?
> Different processes can execute wait or signal operation on a semaphore. Mutexes have ownership, unlike semaphores. Although any thread, within the scope of a mutex, can get an unlocked mutex and lock access to the same critical section of code, only the thread that locked a mutex should unlock it.

### What is a critical section?
> A Critical Section is the part of a program that accesses shared resources. Only when a process is in its Critical Section can it be in a position to disrupt other processes. We can avoid race conditions by making sure that no two processes enter their Critical Sections at the same time.

### What is the difference between race conditions and data races?
 > A data race occurs when 2 instructions from different threads access the same memory location, at least one of these accesses is a write and there is no synchronization that is mandating any particular order among these accesses. A race condition is a semantic error.

### List some advantages of using message passing over lock-based synchronization primitives.
> - The hardware can be simpler
> - Communication explicit -> simpler to understand; in shared memory it can be hard to know when communicating and when not, and how costly it is
> - Explicit communication focuses attention on costly aspect of parallel computation, sometimes leading to improved structure in multiprocesor program
> - Synchronization is naturally associated with sending messages, reducing the possibility for errors introduced by incorrect synchronization
> - Easier to use sender-initiated communication, which may have some advantages in performance

### List some advantages of using lock-based synchronization primitives over message passing.
> - If it is done right, lock-based synchronization primitives can be faster than message passing because the messaging process consumes a lot of time.
> - Necessary to access same memory locations.
