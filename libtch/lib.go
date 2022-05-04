package libtch

// #cgo CFLAGS: -I${SRCDIR} -O3 -Wall -Wno-unused-variable -Wno-deprecated-declarations -Wno-c++11-narrowing -g -Wno-sign-compare -Wno-unused-function
// #cgo CFLAGS: -I/usr/local/include
// #cgo CFLAGS: -D_GLIBCXX_USE_CXX11_ABI=0
// #cgo CFLAGS: -DANDROID_STL=c++_shared
// #cgo LDFLAGS: -lstdc++ -ltorch -lc10 -ltorch_cpu
// #cgo CXXFLAGS: -std=c++14 -I${SRCDIR} -g -O3 -frtti -fexceptions
// #cgo CFLAGS: -I/root/libtorch/lib -I/root/libtorch/include -I/root/libtorch/include/torch/csrc/api/include -I/root/libtorch/include/torch/csrc
// #cgo LDFLAGS: -L/root/libtorch/lib
// #cgo CXXFLAGS: -I/root/libtorch/lib -I/root/libtorch/include -I/root/libtorch/include/torch/csrc/api/include -I/root/libtorch/include/torch/csrc
import "C"
