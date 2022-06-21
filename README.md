# vethack

To be very clear,

# This package exists for mischief and unsafety #

https://pkg.go.dev/unsafe explains why turning a `uintptr` into an `unsafe.Pointer` is not valid.
However, in some fringe cases, like when doing a syscall that returns a pointer from the kernel, it
is needed to turn a `uintptr` into an `unsafe.Pointer`. That is ok, since the GC is not responsible
for the memory pointed to.

Those cases are pointed out by vet, but vet cannot be silenced on a per case basis. This package
provides a layer of indirection to hide the conversion from vet. This way vet can be run on the
original package that needs the conversion.
