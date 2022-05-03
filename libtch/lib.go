package libtch

// #cgo CFLAGS: -I${SRCDIR} -O3 -Wall -Wno-unused-variable -Wno-deprecated-declarations -Wno-c++11-narrowing -g -Wno-sign-compare -Wno-unused-function
// #cgo CFLAGS: -I/usr/local/include
// #cgo CFLAGS: -D_GLIBCXX_USE_CXX11_ABI=0
// #cgo LDFLAGS: -lstdc++ -ltorch -lc10 -ltorch_cpu -L/lib64
// #cgo CXXFLAGS: -std=c++17 -I${SRCDIR} -g -O3
// #cgo CFLAGS: -I/root/libtorch/lib -I/root/libtorch/include -I/root/libtorch/include/torch/csrc/api/include -I/root/libtorch/include/torch/csrc
// #cgo LDFLAGS: -L/root/libtorch/lib
// #cgo CXXFLAGS: -I/root/libtorch/lib -I/root/libtorch/include -I/root/libtorch/include/torch/csrc/api/include -I/root/libtorch/include/torch/csrc
import "C"
