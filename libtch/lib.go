package libtch

// #cgo CFLAGS: -I${SRCDIR} -O3 -Wall -Wno-unused-variable -Wno-deprecated-declarations -Wno-c++11-narrowing -g -Wno-sign-compare -Wno-unused-function
// #cgo CFLAGS: -I/usr/local/include
// #cgo CFLAGS: -D_GLIBCXX_USE_CXX11_ABI=0
// #cgo CFLAGS: -DANDROID_STL=c++_shared
// #cgo LDFLAGS: -lstdc++ -Wl,-Bstatic -ltorch -lc10 -ltorch_cpu
// #cgo CXXFLAGS: -std=c++14 -I${SRCDIR} -g -O3
// #cgo CFLAGS: -I/Users/yuanyuexiang/libtorch/lib -I/Users/yuanyuexiang/libtorch/include -I/Users/yuanyuexiang/libtorch/include/torch/csrc/api/include -I/Users/yuanyuexiang/libtorch/include/torch/csrc
// #cgo LDFLAGS: -L/Users/yuanyuexiang/libtorch/lib
// #cgo CXXFLAGS: -I/Users/yuanyuexiang/libtorch/lib -I/Users/yuanyuexiang/libtorch/include -I/Users/yuanyuexiang/libtorch/include/torch/csrc/api/include -I/Users/yuanyuexiang/libtorch/include/torch/csrc
import "C"
