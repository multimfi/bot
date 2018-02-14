// +build !devel

package static

import "log"

var index_html = []byte{0x3c, 0x21, 0x44, 0x4f, 0x43, 0x54, 0x59, 0x50, 0x45, 0x20, 0x68, 0x74, 0x6d, 0x6c, 0x3e, 0xa, 0x3c, 0x68, 0x74, 0x6d, 0x6c, 0x20, 0x6c, 0x61, 0x6e, 0x67, 0x3d, 0x22, 0x65, 0x6e, 0x22, 0x3e, 0xa, 0x3c, 0x68, 0x65, 0x61, 0x64, 0x3e, 0xa, 0x3c, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x3e, 0x62, 0x6f, 0x74, 0x3c, 0x2f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x3e, 0xa, 0x3c, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x20, 0x74, 0x79, 0x70, 0x65, 0x3d, 0x22, 0x74, 0x65, 0x78, 0x74, 0x2f, 0x6a, 0x61, 0x76, 0x61, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x22, 0x3e, 0xa, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x2e, 0x6f, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x20, 0x3d, 0x20, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x28, 0x29, 0x20, 0x7b, 0xa, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x20, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x52, 0x43, 0x52, 0x65, 0x61, 0x64, 0x79, 0x20, 0x3d, 0x20, 0x30, 0xa, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x20, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x52, 0x43, 0x44, 0x6f, 0x77, 0x6e, 0x20, 0x3d, 0x20, 0x31, 0xa, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x20, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x65, 0x74, 0x20, 0x3d, 0x20, 0x32, 0xa, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x20, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x3d, 0x20, 0x33, 0xa, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x20, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64, 0x65, 0x72, 0x20, 0x3d, 0x20, 0x34, 0xa, 0xa, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x20, 0x53, 0x74, 0x61, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x20, 0x3d, 0x20, 0x20, 0x30, 0x3b, 0xa, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x20, 0x53, 0x74, 0x61, 0x74, 0x65, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x20, 0x3d, 0x20, 0x20, 0x31, 0x3b, 0xa, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x20, 0x53, 0x74, 0x61, 0x74, 0x65, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x20, 0x3d, 0x20, 0x32, 0x3b, 0xa, 0xa, 0x76, 0x61, 0x72, 0x20, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x20, 0x3d, 0x20, 0x30, 0x3b, 0xa, 0x76, 0x61, 0x72, 0x20, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x20, 0x3d, 0x20, 0x7b, 0x7d, 0x3b, 0xa, 0x76, 0x61, 0x72, 0x20, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64, 0x65, 0x72, 0x73, 0x20, 0x3d, 0x20, 0x7b, 0x7d, 0x3b, 0xa, 0xa, 0x76, 0x61, 0x72, 0x20, 0x73, 0x74, 0x61, 0x74, 0x65, 0x20, 0x3d, 0x20, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x67, 0x65, 0x74, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x49, 0x64, 0x28, 0x22, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x29, 0x2e, 0x74, 0x42, 0x6f, 0x64, 0x69, 0x65, 0x73, 0x5b, 0x30, 0x5d, 0x3b, 0xa, 0x76, 0x61, 0x72, 0x20, 0x73, 0x74, 0x61, 0x74, 0x65, 0x73, 0x20, 0x3d, 0x20, 0x5b, 0x22, 0x66, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x22, 0x2c, 0x20, 0x22, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x22, 0x5d, 0x3b, 0xa, 0x76, 0x61, 0x72, 0x20, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x20, 0x3d, 0x20, 0x7b, 0xa, 0x9, 0x22, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x22, 0x3a, 0x20, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x28, 0x6f, 0x62, 0x6a, 0x2c, 0x20, 0x69, 0x74, 0x65, 0x6d, 0x29, 0x20, 0x7b, 0xa, 0x9, 0x9, 0x66, 0x6f, 0x72, 0x20, 0x28, 0x78, 0x20, 0x69, 0x6e, 0x20, 0x6f, 0x62, 0x6a, 0x5b, 0x22, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x22, 0x5d, 0x29, 0x20, 0x7b, 0xa, 0x9, 0x9, 0x9, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x28, 0x6b, 0x76, 0x28, 0x78, 0x2c, 0x20, 0x6f, 0x62, 0x6a, 0x5b, 0x22, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x22, 0x5d, 0x5b, 0x78, 0x5d, 0x29, 0x29, 0x3b, 0xa, 0x9, 0x9, 0x7d, 0xa, 0x9, 0x7d, 0x2c, 0xa, 0xa, 0x9, 0x22, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x3a, 0x20, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x28, 0x6f, 0x62, 0x6a, 0x2c, 0x20, 0x69, 0x74, 0x65, 0x6d, 0x29, 0x20, 0x7b, 0xa, 0x9, 0x9, 0x66, 0x6f, 0x72, 0x20, 0x28, 0x78, 0x20, 0x69, 0x6e, 0x20, 0x6f, 0x62, 0x6a, 0x5b, 0x22, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x5d, 0x29, 0x20, 0x7b, 0xa, 0x9, 0x9, 0x9, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x28, 0x6b, 0x76, 0x28, 0x78, 0x2c, 0x20, 0x6f, 0x62, 0x6a, 0x5b, 0x22, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x5d, 0x5b, 0x78, 0x5d, 0x29, 0x29, 0x3b, 0xa, 0x9, 0x9, 0x7d, 0xa, 0x9, 0x7d, 0x2c, 0xa, 0xa, 0x9, 0x22, 0x73, 0x74, 0x61, 0x72, 0x74, 0x73, 0x41, 0x74, 0x22, 0x3a, 0x20, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x28, 0x6f, 0x62, 0x6a, 0x2c, 0x20, 0x69, 0x74, 0x65, 0x6d, 0x29, 0x20, 0x7b, 0xa, 0x9, 0x9, 0x69, 0x66, 0x20, 0x28, 0x6f, 0x62, 0x6a, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x20, 0x21, 0x3d, 0x20, 0x22, 0x66, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x22, 0x29, 0x20, 0x7b, 0xa, 0x9, 0x9, 0x9, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x28, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x5b, 0x6f, 0x62, 0x6a, 0x2e, 0x68, 0x5d, 0x2e, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x72, 0x29, 0xa, 0x9, 0x9, 0x9, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x3b, 0xa, 0x9, 0x9, 0x7d, 0xa, 0x9, 0x9, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x5b, 0x6f, 0x62, 0x6a, 0x2e, 0x68, 0x5d, 0x2e, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x72, 0x20, 0x3d, 0x20, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x28, 0x74, 0x29, 0x20, 0x7b, 0xa, 0x9, 0x9, 0x9, 0x76, 0x61, 0x72, 0x20, 0x64, 0x20, 0x3d, 0x20, 0x6e, 0x65, 0x77, 0x20, 0x44, 0x61, 0x74, 0x65, 0x28, 0x6f, 0x62, 0x6a, 0x5b, 0x22, 0x73, 0x74, 0x61, 0x72, 0x74, 0x73, 0x41, 0x74, 0x22, 0x5d, 0x29, 0x3b, 0xa, 0x9, 0x9, 0x9, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x48, 0x54, 0x4d, 0x4c, 0x20, 0x3d, 0x20, 0x6f, 0x62, 0x6a, 0x5b, 0x22, 0x73, 0x74, 0x61, 0x72, 0x74, 0x73, 0x41, 0x74, 0x22, 0x5d, 0xa, 0x9, 0x9, 0x9, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x48, 0x54, 0x4d, 0x4c, 0x20, 0x2b, 0x3d, 0x20, 0x22, 0x3c, 0x62, 0x72, 0x3e, 0x3c, 0x62, 0x3e, 0x22, 0x20, 0x2b, 0x20, 0x73, 0x69, 0x6e, 0x63, 0x65, 0x28, 0x74, 0x2c, 0x20, 0x64, 0x29, 0x20, 0x2b, 0x20, 0x22, 0x3c, 0x2f, 0x62, 0x3e, 0x22, 0x3b, 0xa, 0x9, 0x9, 0x7d, 0xa, 0x9, 0x9, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x5b, 0x6f, 0x62, 0x6a, 0x2e, 0x68, 0x5d, 0x2e, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x72, 0x28, 0x44, 0x61, 0x74, 0x65, 0x2e, 0x6e, 0x6f, 0x77, 0x28, 0x29, 0x29, 0xa, 0x9, 0x7d, 0x2c, 0xa, 0xa, 0x9, 0x22, 0x65, 0x6e, 0x64, 0x73, 0x41, 0x74, 0x22, 0x3a, 0x20, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x28, 0x6f, 0x62, 0x6a, 0x2c, 0x20, 0x69, 0x74, 0x65, 0x6d, 0x29, 0x20, 0x7b, 0xa, 0x9, 0x9, 0x76, 0x61, 0x72, 0x20, 0x64, 0x31, 0x20, 0x3d, 0x20, 0x6e, 0x65, 0x77, 0x20, 0x44, 0x61, 0x74, 0x65, 0x28, 0x6f, 0x62, 0x6a, 0x5b, 0x22, 0x65, 0x6e, 0x64, 0x73, 0x41, 0x74, 0x22, 0x5d, 0x29, 0xa, 0x9, 0x9, 0x76, 0x61, 0x72, 0x20, 0x64, 0x32, 0x20, 0x3d, 0x20, 0x6e, 0x65, 0x77, 0x20, 0x44, 0x61, 0x74, 0x65, 0x28, 0x6f, 0x62, 0x6a, 0x5b, 0x22, 0x73, 0x74, 0x61, 0x72, 0x74, 0x73, 0x41, 0x74, 0x22, 0x5d, 0x29, 0xa, 0x9, 0x9, 0x69, 0x66, 0x20, 0x28, 0x64, 0x31, 0x20, 0x3c, 0x20, 0x31, 0x29, 0x20, 0x7b, 0xa, 0x9, 0x9, 0x9, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x48, 0x54, 0x4d, 0x4c, 0x20, 0x3d, 0x20, 0x22, 0x74, 0x62, 0x64, 0x22, 0x3b, 0xa, 0x9, 0x9, 0x9, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0xa, 0x9, 0x9, 0x7d, 0xa, 0x9, 0x9, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x48, 0x54, 0x4d, 0x4c, 0x20, 0x3d, 0x20, 0x6f, 0x62, 0x6a, 0x5b, 0x22, 0x65, 0x6e, 0x64, 0x73, 0x41, 0x74, 0x22, 0x5d, 0xa, 0x9, 0x9, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x48, 0x54, 0x4d, 0x4c, 0x20, 0x2b, 0x3d, 0x20, 0x22, 0x3c, 0x62, 0x72, 0x3e, 0x3c, 0x62, 0x3e, 0x22, 0x2b, 0x73, 0x69, 0x6e, 0x63, 0x65, 0x28, 0x64, 0x31, 0x2c, 0x20, 0x64, 0x32, 0x29, 0x2b, 0x22, 0x3c, 0x2f, 0x62, 0x3e, 0x22, 0x3b, 0xa, 0x9, 0x7d, 0x2c, 0xa, 0xa, 0x9, 0x22, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64, 0x65, 0x72, 0x73, 0x22, 0x3a, 0x20, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x28, 0x6f, 0x62, 0x6a, 0x2c, 0x20, 0x69, 0x74, 0x65, 0x6d, 0x29, 0x20, 0x7b, 0xa, 0x9, 0x9, 0x69, 0x66, 0x20, 0x28, 0x6f, 0x62, 0x6a, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x20, 0x21, 0x3d, 0x20, 0x22, 0x66, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x22, 0x29, 0x20, 0x7b, 0xa, 0x9, 0x9, 0x9, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x3b, 0xa, 0x9, 0x9, 0x7d, 0xa, 0xa, 0x9, 0x9, 0x76, 0x61, 0x72, 0x20, 0x72, 0x20, 0x3d, 0x20, 0x6f, 0x62, 0x6a, 0x5b, 0x22, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64, 0x65, 0x72, 0x73, 0x22, 0x5d, 0x3b, 0xa, 0x9, 0x9, 0x69, 0x66, 0x20, 0x28, 0x21, 0x72, 0x29, 0x20, 0x7b, 0xa, 0x9, 0x9, 0x9, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x48, 0x54, 0x4d, 0x4c, 0x20, 0x3d, 0x20, 0x22, 0x6e, 0x6f, 0x6e, 0x65, 0x22, 0x3b, 0xa, 0x9, 0x9, 0x9, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x3b, 0xa, 0x9, 0x9, 0x7d, 0xa, 0x9, 0x9, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x48, 0x54, 0x4d, 0x4c, 0x20, 0x3d, 0x20, 0x22, 0x3c, 0x62, 0x3e, 0x22, 0x20, 0x2b, 0x20, 0x72, 0x5b, 0x6f, 0x62, 0x6a, 0x2e, 0x63, 0x5d, 0x20, 0x2b, 0x20, 0x22, 0x3c, 0x2f, 0x62, 0x3e, 0x22, 0x3b, 0xa, 0x9, 0x9, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x48, 0x54, 0x4d, 0x4c, 0x20, 0x2b, 0x3d, 0x20, 0x22, 0x3c, 0x62, 0x72, 0x3e, 0x22, 0x20, 0x2b, 0x20, 0x72, 0x3b, 0xa, 0x9, 0x7d, 0xa, 0x7d, 0xa, 0xa, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x72, 0x28, 0x29, 0x20, 0x7b, 0xa, 0x9, 0x76, 0x61, 0x72, 0x20, 0x6e, 0x6f, 0x77, 0x20, 0x3d, 0x20, 0x44, 0x61, 0x74, 0x65, 0x2e, 0x6e, 0x6f, 0x77, 0x28, 0x29, 0x3b, 0xa, 0x9, 0x66, 0x6f, 0x72, 0x20, 0x28, 0x69, 0x20, 0x69, 0x6e, 0x20, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x29, 0x20, 0x7b, 0xa, 0x9, 0x9, 0x69, 0x66, 0x20, 0x28, 0x74, 0x79, 0x70, 0x65, 0x6f, 0x66, 0x20, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x5b, 0x69, 0x5d, 0x2e, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x72, 0x20, 0x21, 0x3d, 0x20, 0x22, 0x75, 0x6e, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x64, 0x22, 0x29, 0x20, 0x7b, 0xa, 0x9, 0x9, 0x9, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x5b, 0x69, 0x5d, 0x2e, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x72, 0x28, 0x6e, 0x6f, 0x77, 0x29, 0x3b, 0xa, 0x9, 0x9, 0x7d, 0xa, 0x9, 0x7d, 0xa, 0x7d, 0xa, 0xa, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x73, 0x69, 0x6e, 0x63, 0x65, 0x28, 0x78, 0x2c, 0x20, 0x79, 0x29, 0x20, 0x7b, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x76, 0x61, 0x72, 0x20, 0x74, 0x20, 0x3d, 0x20, 0x28, 0x78, 0x20, 0x2d, 0x20, 0x79, 0x29, 0x2f, 0x31, 0x65, 0x33, 0x3b, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x76, 0x61, 0x72, 0x20, 0x73, 0x20, 0x3d, 0x20, 0x74, 0x25, 0x36, 0x30, 0x7c, 0x30, 0x3b, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x74, 0x20, 0x2f, 0x3d, 0x20, 0x36, 0x30, 0x3b, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x76, 0x61, 0x72, 0x20, 0x6d, 0x20, 0x3d, 0x20, 0x74, 0x25, 0x36, 0x30, 0x7c, 0x30, 0x3b, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x74, 0x20, 0x2f, 0x3d, 0x20, 0x36, 0x30, 0x3b, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x76, 0x61, 0x72, 0x20, 0x68, 0x20, 0x3d, 0x20, 0x74, 0x25, 0x32, 0x34, 0x7c, 0x30, 0x3b, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x74, 0x20, 0x2f, 0x3d, 0x20, 0x32, 0x34, 0x3b, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x76, 0x61, 0x72, 0x20, 0x64, 0x20, 0x3d, 0x20, 0x74, 0x7c, 0x30, 0x3b, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x20, 0x28, 0x64, 0x20, 0x3d, 0x3d, 0x20, 0x30, 0x20, 0x3f, 0x20, 0x22, 0x22, 0x20, 0x3a, 0x20, 0x64, 0x2b, 0x22, 0x64, 0x20, 0x22, 0x29, 0x20, 0x2b, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x28, 0x68, 0x2b, 0x64, 0x20, 0x3d, 0x3d, 0x20, 0x30, 0x20, 0x3f, 0x20, 0x22, 0x22, 0x20, 0x3a, 0x20, 0x68, 0x2b, 0x22, 0x68, 0x20, 0x22, 0x29, 0x20, 0x2b, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x28, 0x6d, 0x2b, 0x68, 0x2b, 0x64, 0x20, 0x3d, 0x3d, 0x20, 0x30, 0x20, 0x3f, 0x20, 0x22, 0x22, 0x20, 0x3a, 0x20, 0x6d, 0x2b, 0x22, 0x6d, 0x20, 0x22, 0x29, 0x20, 0x2b, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x73, 0x2b, 0x22, 0x73, 0x22, 0x3b, 0xa, 0x7d, 0xa, 0xa, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x6b, 0x76, 0x28, 0x6b, 0x2c, 0x20, 0x76, 0x29, 0x20, 0x7b, 0xa, 0x9, 0x76, 0x61, 0x72, 0x20, 0x72, 0x65, 0x74, 0x20, 0x3d, 0x20, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x28, 0x22, 0x70, 0x22, 0x29, 0x3b, 0xa, 0x9, 0x76, 0x61, 0x72, 0x20, 0x62, 0x20, 0x3d, 0x20, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x28, 0x22, 0x62, 0x22, 0x29, 0x3b, 0xa, 0x9, 0x62, 0x2e, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x48, 0x54, 0x4d, 0x4c, 0x20, 0x3d, 0x20, 0x6b, 0x3b, 0xa, 0x9, 0x72, 0x65, 0x74, 0x2e, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x28, 0x62, 0x29, 0x3b, 0xa, 0x9, 0x72, 0x65, 0x74, 0x2e, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x48, 0x54, 0x4d, 0x4c, 0x20, 0x2b, 0x3d, 0x20, 0x22, 0x20, 0x22, 0x20, 0x2b, 0x20, 0x76, 0x3b, 0xa, 0x9, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x20, 0x72, 0x65, 0x74, 0x3b, 0xa, 0x7d, 0xa, 0xa, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x73, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x28, 0x69, 0x74, 0x65, 0x6d, 0x2c, 0x20, 0x73, 0x74, 0x61, 0x74, 0x65, 0x29, 0x20, 0x7b, 0xa, 0x9, 0x66, 0x6f, 0x72, 0x20, 0x28, 0x78, 0x20, 0x69, 0x6e, 0x20, 0x73, 0x74, 0x61, 0x74, 0x65, 0x73, 0x29, 0x20, 0x7b, 0xa, 0x9, 0x9, 0x69, 0x66, 0x20, 0x28, 0x73, 0x74, 0x61, 0x74, 0x65, 0x20, 0x21, 0x3d, 0x20, 0x73, 0x74, 0x61, 0x74, 0x65, 0x73, 0x5b, 0x78, 0x5d, 0x29, 0x20, 0x7b, 0xa, 0x9, 0x9, 0x9, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x28, 0x73, 0x74, 0x61, 0x74, 0x65, 0x73, 0x5b, 0x78, 0x5d, 0x29, 0x3b, 0xa, 0x9, 0x9, 0x9, 0x63, 0x6f, 0x6e, 0x74, 0x69, 0x6e, 0x75, 0x65, 0x3b, 0xa, 0x9, 0x9, 0x7d, 0xa, 0x9, 0x9, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x2e, 0x61, 0x64, 0x64, 0x28, 0x73, 0x74, 0x61, 0x74, 0x65, 0x29, 0x3b, 0xa, 0x9, 0x7d, 0xa, 0x7d, 0xa, 0xa, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x67, 0x65, 0x74, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x28, 0x6e, 0x61, 0x6d, 0x65, 0x2c, 0x20, 0x6f, 0x62, 0x6a, 0x29, 0x20, 0x7b, 0xa, 0x9, 0x66, 0x6f, 0x72, 0x20, 0x28, 0x78, 0x20, 0x69, 0x6e, 0x20, 0x6f, 0x62, 0x6a, 0x2e, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x29, 0x20, 0x7b, 0xa, 0x9, 0x9, 0x79, 0x20, 0x3d, 0x20, 0x6f, 0x62, 0x6a, 0x2e, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x5b, 0x78, 0x5d, 0xa, 0x9, 0x9, 0x69, 0x66, 0x20, 0x28, 0x79, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x20, 0x21, 0x3d, 0x20, 0x6e, 0x61, 0x6d, 0x65, 0x29, 0xa, 0x9, 0x9, 0x9, 0x63, 0x6f, 0x6e, 0x74, 0x69, 0x6e, 0x75, 0x65, 0xa, 0x9, 0x9, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x20, 0x79, 0xa, 0x9, 0x7d, 0xa, 0x7d, 0xa, 0xa, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x28, 0x6f, 0x62, 0x6a, 0x29, 0x20, 0x7b, 0xa, 0x9, 0x76, 0x61, 0x72, 0x20, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x3b, 0xa, 0x9, 0x69, 0x66, 0x20, 0x28, 0x21, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x5b, 0x6f, 0x62, 0x6a, 0x2e, 0x68, 0x5d, 0x29, 0x20, 0x7b, 0xa, 0x9, 0x9, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x20, 0x3d, 0x20, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x28, 0x22, 0x74, 0x72, 0x22, 0x29, 0x3b, 0xa, 0x9, 0x9, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x2e, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x2e, 0x61, 0x64, 0x64, 0x28, 0x22, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x22, 0x29, 0xa, 0x9, 0x9, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x5b, 0x6f, 0x62, 0x6a, 0x2e, 0x68, 0x5d, 0x20, 0x3d, 0x20, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x3b, 0xa, 0x9, 0x9, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x28, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x29, 0x3b, 0xa, 0x9, 0x7d, 0x20, 0x65, 0x6c, 0x73, 0x65, 0x20, 0x7b, 0xa, 0x9, 0x9, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x20, 0x3d, 0x20, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x5b, 0x6f, 0x62, 0x6a, 0x2e, 0x68, 0x5d, 0x3b, 0xa, 0x9, 0x7d, 0xa, 0xa, 0x9, 0x73, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x28, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x2c, 0x20, 0x6f, 0x62, 0x6a, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x29, 0x3b, 0xa, 0x9, 0x66, 0x6f, 0x72, 0x20, 0x28, 0x6e, 0x61, 0x6d, 0x65, 0x20, 0x69, 0x6e, 0x20, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x29, 0x20, 0x7b, 0xa, 0x9, 0x9, 0x76, 0x61, 0x72, 0x20, 0x69, 0x74, 0x65, 0x6d, 0x20, 0x3d, 0x20, 0x67, 0x65, 0x74, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x28, 0x6e, 0x61, 0x6d, 0x65, 0x2c, 0x20, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x29, 0x3b, 0xa, 0x9, 0x9, 0x69, 0x66, 0x20, 0x28, 0x21, 0x69, 0x74, 0x65, 0x6d, 0x29, 0x20, 0x7b, 0xa, 0x9, 0x9, 0x9, 0x69, 0x74, 0x65, 0x6d, 0x20, 0x3d, 0x20, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x28, 0x22, 0x74, 0x64, 0x22, 0x29, 0x3b, 0xa, 0x9, 0x9, 0x9, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x2e, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x28, 0x69, 0x74, 0x65, 0x6d, 0x29, 0xa, 0x9, 0x9, 0x9, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x20, 0x3d, 0x20, 0x6e, 0x61, 0x6d, 0x65, 0x3b, 0xa, 0x9, 0x9, 0x9, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x5b, 0x6e, 0x61, 0x6d, 0x65, 0x5d, 0x28, 0x6f, 0x62, 0x6a, 0x2c, 0x20, 0x69, 0x74, 0x65, 0x6d, 0x29, 0xa, 0x9, 0x9, 0x7d, 0xa, 0x9, 0x9, 0x73, 0x77, 0x69, 0x74, 0x63, 0x68, 0x28, 0x6e, 0x61, 0x6d, 0x65, 0x29, 0x20, 0x7b, 0xa, 0x9, 0x9, 0x63, 0x61, 0x73, 0x65, 0x20, 0x22, 0x65, 0x6e, 0x64, 0x73, 0x41, 0x74, 0x22, 0x3a, 0xa, 0x9, 0x9, 0x63, 0x61, 0x73, 0x65, 0x20, 0x22, 0x73, 0x74, 0x61, 0x72, 0x74, 0x73, 0x41, 0x74, 0x22, 0x3a, 0xa, 0x9, 0x9, 0x63, 0x61, 0x73, 0x65, 0x20, 0x22, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64, 0x65, 0x72, 0x73, 0x22, 0x3a, 0xa, 0x9, 0x9, 0x9, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x5b, 0x6e, 0x61, 0x6d, 0x65, 0x5d, 0x28, 0x6f, 0x62, 0x6a, 0x2c, 0x20, 0x69, 0x74, 0x65, 0x6d, 0x29, 0x3b, 0xa, 0x9, 0x9, 0x7d, 0xa, 0x9, 0x7d, 0xa, 0x7d, 0xa, 0xa, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x52, 0x43, 0x28, 0x6e, 0x29, 0x20, 0x7b, 0xa, 0x9, 0x76, 0x61, 0x72, 0x20, 0x65, 0x20, 0x3d, 0x20, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x67, 0x65, 0x74, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x49, 0x64, 0x28, 0x22, 0x69, 0x72, 0x63, 0x22, 0x29, 0xa, 0x9, 0x73, 0x77, 0x69, 0x74, 0x63, 0x68, 0x28, 0x6e, 0x29, 0x20, 0x7b, 0xa, 0x9, 0x63, 0x61, 0x73, 0x65, 0x20, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x52, 0x43, 0x52, 0x65, 0x61, 0x64, 0x79, 0x3a, 0xa, 0x9, 0x9, 0x65, 0x2e, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x20, 0x3d, 0x20, 0x22, 0x67, 0x72, 0x65, 0x65, 0x6e, 0x22, 0x3b, 0xa, 0x9, 0x9, 0x65, 0x2e, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x48, 0x54, 0x4d, 0x4c, 0x20, 0x3d, 0x20, 0x22, 0x4f, 0x4b, 0x21, 0x22, 0x3b, 0xa, 0x9, 0x9, 0x62, 0x72, 0x65, 0x61, 0x6b, 0x3b, 0xa, 0x9, 0x63, 0x61, 0x73, 0x65, 0x20, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x52, 0x43, 0x44, 0x6f, 0x77, 0x6e, 0x3a, 0xa, 0x9, 0x9, 0x65, 0x2e, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x20, 0x3d, 0x20, 0x22, 0x72, 0x65, 0x64, 0x22, 0x3b, 0xa, 0x9, 0x9, 0x65, 0x2e, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x48, 0x54, 0x4d, 0x4c, 0x20, 0x3d, 0x20, 0x22, 0x46, 0x41, 0x49, 0x4c, 0x21, 0x22, 0x3b, 0xa, 0x9, 0x9, 0x62, 0x72, 0x65, 0x61, 0x6b, 0x3b, 0xa, 0x9, 0x7d, 0xa, 0x7d, 0xa, 0xa, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x6d, 0x61, 0x70, 0x20, 0x3d, 0x20, 0x7b, 0xa, 0x9, 0x30, 0x3a, 0x20, 0x22, 0x67, 0x72, 0x65, 0x65, 0x6e, 0x22, 0x2c, 0xa, 0x9, 0x31, 0x3a, 0x20, 0x22, 0x72, 0x65, 0x64, 0x22, 0x2c, 0xa, 0x9, 0x32, 0x3a, 0x20, 0x22, 0x62, 0x6c, 0x61, 0x63, 0x6b, 0x22, 0x2c, 0xa, 0x7d, 0xa, 0xa, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64, 0x65, 0x72, 0x73, 0x28, 0x6f, 0x62, 0x6a, 0x29, 0x20, 0x7b, 0xa, 0x9, 0x76, 0x61, 0x72, 0x20, 0x65, 0x20, 0x3d, 0x20, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x67, 0x65, 0x74, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x49, 0x64, 0x28, 0x22, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64, 0x65, 0x72, 0x73, 0x22, 0x29, 0xa, 0x9, 0x66, 0x6f, 0x72, 0x20, 0x28, 0x6e, 0x20, 0x69, 0x6e, 0x20, 0x6f, 0x62, 0x6a, 0x29, 0x20, 0x7b, 0xa, 0x9, 0x9, 0x76, 0x61, 0x72, 0x20, 0x69, 0x74, 0x65, 0x6d, 0x20, 0x3d, 0x20, 0x67, 0x65, 0x74, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x28, 0x6e, 0x2c, 0x20, 0x65, 0x29, 0x3b, 0xa, 0x9, 0x9, 0x69, 0x66, 0x20, 0x28, 0x21, 0x69, 0x74, 0x65, 0x6d, 0x29, 0x20, 0x7b, 0xa, 0x9, 0x9, 0x9, 0x69, 0x74, 0x65, 0x6d, 0x20, 0x3d, 0x20, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x28, 0x22, 0x73, 0x70, 0x61, 0x6e, 0x22, 0x29, 0x3b, 0xa, 0x9, 0x9, 0x9, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x28, 0x69, 0x74, 0x65, 0x6d, 0x29, 0xa, 0x9, 0x9, 0x9, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x20, 0x3d, 0x20, 0x6e, 0x3b, 0xa, 0x9, 0x9, 0x9, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x48, 0x54, 0x4d, 0x4c, 0x20, 0x3d, 0x20, 0x6e, 0x3b, 0xa, 0x9, 0x9, 0x7d, 0xa, 0x9, 0x9, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x20, 0x3d, 0x20, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x6d, 0x61, 0x70, 0x5b, 0x6f, 0x62, 0x6a, 0x5b, 0x6e, 0x5d, 0x5d, 0x3b, 0xa, 0x9, 0x7d, 0xa, 0x7d, 0xa, 0xa, 0x76, 0x61, 0x72, 0x20, 0x63, 0x6f, 0x6e, 0x6e, 0x20, 0x3d, 0x20, 0x6e, 0x65, 0x77, 0x20, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x28, 0x22, 0x73, 0x73, 0x65, 0x22, 0x29, 0x3b, 0xa, 0xa, 0x63, 0x6f, 0x6e, 0x6e, 0x2e, 0x6f, 0x6e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x20, 0x3d, 0x20, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x28, 0x65, 0x76, 0x74, 0x29, 0x20, 0x7b, 0xa, 0x9, 0x76, 0x61, 0x72, 0x20, 0x69, 0x74, 0x65, 0x6d, 0x20, 0x3d, 0x20, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x67, 0x65, 0x74, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x49, 0x64, 0x28, 0x22, 0x65, 0x72, 0x72, 0x22, 0x29, 0x3b, 0xa, 0x9, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x20, 0x3d, 0x20, 0x22, 0x72, 0x65, 0x64, 0x22, 0x3b, 0xa, 0x9, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x48, 0x54, 0x4d, 0x4c, 0x20, 0x3d, 0x20, 0x22, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x3b, 0xa, 0x7d, 0xa, 0xa, 0x63, 0x6f, 0x6e, 0x6e, 0x2e, 0x6f, 0x6e, 0x6f, 0x70, 0x65, 0x6e, 0x20, 0x3d, 0x20, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x28, 0x65, 0x76, 0x74, 0x29, 0x20, 0x7b, 0xa, 0x9, 0x76, 0x61, 0x72, 0x20, 0x69, 0x74, 0x65, 0x6d, 0x20, 0x3d, 0x20, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x67, 0x65, 0x74, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x49, 0x64, 0x28, 0x22, 0x65, 0x72, 0x72, 0x22, 0x29, 0x3b, 0xa, 0x9, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x20, 0x3d, 0x20, 0x22, 0x67, 0x72, 0x65, 0x65, 0x6e, 0x22, 0x3b, 0xa, 0x9, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x48, 0x54, 0x4d, 0x4c, 0x20, 0x3d, 0x20, 0x22, 0x6f, 0x6b, 0x22, 0x3b, 0xa, 0x7d, 0xa, 0xa, 0x63, 0x6f, 0x6e, 0x6e, 0x2e, 0x61, 0x64, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x28, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x2c, 0x20, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x28, 0x65, 0x29, 0x20, 0x7b, 0xa, 0x9, 0x76, 0x61, 0x72, 0x20, 0x78, 0x20, 0x3d, 0x20, 0x4a, 0x53, 0x4f, 0x4e, 0x2e, 0x70, 0x61, 0x72, 0x73, 0x65, 0x28, 0x65, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x29, 0x3b, 0xa, 0x9, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x28, 0x78, 0x29, 0x3b, 0xa, 0xa, 0x9, 0x69, 0x66, 0x20, 0x28, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x67, 0x65, 0x74, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x79, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x28, 0x22, 0x66, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x22, 0x29, 0x2e, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x20, 0x3c, 0x20, 0x31, 0x29, 0x20, 0x7b, 0xa, 0x9, 0x9, 0x63, 0x6c, 0x65, 0x61, 0x72, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x28, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x29, 0x3b, 0xa, 0x9, 0x9, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x20, 0x3d, 0x20, 0x30, 0x3b, 0xa, 0x9, 0x9, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x3b, 0xa, 0x9, 0x7d, 0xa, 0xa, 0x9, 0x69, 0x66, 0x20, 0x28, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x20, 0x3e, 0x20, 0x30, 0x29, 0x20, 0x7b, 0xa, 0x9, 0x9, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x3b, 0xa, 0x9, 0x7d, 0xa, 0xa, 0x9, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x20, 0x3d, 0x20, 0x73, 0x65, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x28, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x72, 0x2c, 0x20, 0x31, 0x30, 0x30, 0x30, 0x29, 0x3b, 0xa, 0x7d, 0x29, 0xa, 0xa, 0x63, 0x6f, 0x6e, 0x6e, 0x2e, 0x61, 0x64, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x28, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x52, 0x43, 0x52, 0x65, 0x61, 0x64, 0x79, 0x2c, 0x20, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x28, 0x65, 0x29, 0x20, 0x7b, 0xa, 0x9, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x52, 0x43, 0x28, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x52, 0x43, 0x52, 0x65, 0x61, 0x64, 0x79, 0x29, 0x3b, 0xa, 0x7d, 0x29, 0xa, 0xa, 0x63, 0x6f, 0x6e, 0x6e, 0x2e, 0x61, 0x64, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x28, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x52, 0x43, 0x44, 0x6f, 0x77, 0x6e, 0x2c, 0x20, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x28, 0x65, 0x29, 0x20, 0x7b, 0xa, 0x9, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x52, 0x43, 0x28, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x52, 0x43, 0x44, 0x6f, 0x77, 0x6e, 0x29, 0x3b, 0xa, 0x7d, 0x29, 0xa, 0xa, 0x63, 0x6f, 0x6e, 0x6e, 0x2e, 0x61, 0x64, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x28, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64, 0x65, 0x72, 0x2c, 0x20, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x28, 0x65, 0x29, 0x20, 0x7b, 0xa, 0x9, 0x76, 0x61, 0x72, 0x20, 0x78, 0x20, 0x3d, 0x20, 0x4a, 0x53, 0x4f, 0x4e, 0x2e, 0x70, 0x61, 0x72, 0x73, 0x65, 0x28, 0x65, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x29, 0x3b, 0xa, 0x9, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64, 0x65, 0x72, 0x73, 0x28, 0x78, 0x29, 0x3b, 0xa, 0x7d, 0x29, 0xa, 0xa, 0x63, 0x6f, 0x6e, 0x6e, 0x2e, 0x61, 0x64, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x28, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x65, 0x74, 0x2c, 0x20, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x28, 0x65, 0x29, 0x20, 0x7b, 0xa, 0x9, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x20, 0x3d, 0x20, 0x7b, 0x7d, 0x3b, 0xa, 0x9, 0x66, 0x6f, 0x72, 0x20, 0x28, 0x78, 0x20, 0x3d, 0x20, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x72, 0x6f, 0x77, 0x73, 0x2e, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x2d, 0x31, 0x3b, 0x20, 0x78, 0x20, 0x3e, 0x20, 0x30, 0x3b, 0x20, 0x78, 0x2d, 0x2d, 0x29, 0x20, 0x7b, 0xa, 0x9, 0x9, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x6f, 0x77, 0x28, 0x78, 0x29, 0x3b, 0xa, 0x9, 0x7d, 0xa, 0x7d, 0x29, 0xa, 0x7d, 0x3b, 0xa, 0x3c, 0x2f, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x3e, 0xa, 0x3c, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x20, 0x74, 0x79, 0x70, 0x65, 0x3d, 0x22, 0x74, 0x65, 0x78, 0x74, 0x2f, 0x63, 0x73, 0x73, 0x22, 0x3e, 0xa, 0x62, 0x6f, 0x64, 0x79, 0x20, 0x7b, 0xa, 0x9, 0x66, 0x6f, 0x6e, 0x74, 0x2d, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x3a, 0x20, 0x73, 0x61, 0x6e, 0x73, 0x2d, 0x73, 0x65, 0x72, 0x69, 0x66, 0x3b, 0xa, 0x7d, 0xa, 0xa, 0x74, 0x64, 0x20, 0x70, 0x20, 0x7b, 0xa, 0x9, 0x6d, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x3a, 0x20, 0x30, 0x3b, 0xa, 0x9, 0x70, 0x61, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x3a, 0x20, 0x30, 0x3b, 0xa, 0x7d, 0xa, 0xa, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x20, 0x74, 0x72, 0x20, 0x74, 0x68, 0x20, 0x7b, 0xa, 0x9, 0x74, 0x65, 0x78, 0x74, 0x2d, 0x61, 0x6c, 0x69, 0x67, 0x6e, 0x3a, 0x20, 0x6c, 0x65, 0x66, 0x74, 0x3b, 0xa, 0x7d, 0xa, 0xa, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x20, 0x7b, 0xa, 0x9, 0x77, 0x69, 0x64, 0x74, 0x68, 0x3a, 0x20, 0x31, 0x30, 0x30, 0x25, 0x3b, 0xa, 0x9, 0x62, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2d, 0x73, 0x70, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x3a, 0x20, 0x30, 0x3b, 0xa, 0x7d, 0xa, 0xa, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x20, 0x74, 0x72, 0x20, 0x74, 0x64, 0x2c, 0x20, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x20, 0x74, 0x72, 0x20, 0x74, 0x68, 0x2c, 0x20, 0x68, 0x31, 0x20, 0x7b, 0xa, 0x9, 0x70, 0x61, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x2d, 0x6c, 0x65, 0x66, 0x74, 0x3a, 0x20, 0x31, 0x65, 0x6d, 0x3b, 0xa, 0x7d, 0xa, 0xa, 0x2e, 0x66, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x20, 0x7b, 0xa, 0x9, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x3a, 0x20, 0x72, 0x67, 0x62, 0x28, 0x32, 0x35, 0x35, 0x2c, 0x20, 0x31, 0x35, 0x36, 0x2c, 0x20, 0x31, 0x35, 0x36, 0x29, 0x3b, 0xa, 0x7d, 0xa, 0xa, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x20, 0x7b, 0xa, 0x9, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x3a, 0x20, 0x72, 0x67, 0x62, 0x28, 0x32, 0x33, 0x32, 0x2c, 0x20, 0x32, 0x35, 0x35, 0x2c, 0x20, 0x32, 0x34, 0x30, 0x29, 0x3b, 0xa, 0x7d, 0xa, 0xa, 0x74, 0x64, 0x20, 0x7b, 0xa, 0x9, 0x70, 0x61, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x3a, 0x20, 0x31, 0x65, 0x6d, 0x3b, 0xa, 0x9, 0x62, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2d, 0x74, 0x6f, 0x70, 0x3a, 0x20, 0x35, 0x70, 0x78, 0x20, 0x73, 0x6f, 0x6c, 0x69, 0x64, 0x20, 0x23, 0x66, 0x66, 0x66, 0x3b, 0xa, 0x7d, 0xa, 0xa, 0x23, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64, 0x65, 0x72, 0x73, 0x20, 0x73, 0x70, 0x61, 0x6e, 0x20, 0x7b, 0xa, 0x9, 0x66, 0x6f, 0x6e, 0x74, 0x2d, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x3a, 0x20, 0x62, 0x6f, 0x6c, 0x64, 0x3b, 0xa, 0x9, 0x70, 0x61, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x3a, 0x20, 0x31, 0x65, 0x6d, 0x3b, 0xa, 0x7d, 0xa, 0x3c, 0x2f, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x3e, 0xa, 0x3c, 0x2f, 0x68, 0x65, 0x61, 0x64, 0x3e, 0xa, 0x3c, 0x62, 0x6f, 0x64, 0x79, 0x3e, 0xa, 0x3c, 0x64, 0x69, 0x76, 0x20, 0x69, 0x64, 0x3d, 0x22, 0x65, 0x72, 0x72, 0x22, 0x3e, 0x3c, 0x2f, 0x64, 0x69, 0x76, 0x3e, 0xa, 0x3c, 0x70, 0x3e, 0x49, 0x52, 0x43, 0x20, 0x3c, 0x62, 0x3e, 0x3c, 0x73, 0x70, 0x61, 0x6e, 0x20, 0x69, 0x64, 0x3d, 0x22, 0x69, 0x72, 0x63, 0x22, 0x3e, 0x54, 0x42, 0x44, 0x3c, 0x2f, 0x73, 0x70, 0x61, 0x6e, 0x3e, 0x3c, 0x2f, 0x62, 0x3e, 0x3c, 0x2f, 0x70, 0x3e, 0xa, 0x3c, 0x68, 0x34, 0x3e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64, 0x65, 0x72, 0x73, 0x3c, 0x2f, 0x68, 0x34, 0x3e, 0xa, 0x3c, 0x64, 0x69, 0x76, 0x20, 0x69, 0x64, 0x3d, 0x22, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64, 0x65, 0x72, 0x73, 0x22, 0x3e, 0x3c, 0x2f, 0x64, 0x69, 0x76, 0x3e, 0xa, 0x3c, 0x68, 0x33, 0x3e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x3c, 0x2f, 0x68, 0x33, 0x3e, 0xa, 0x3c, 0x64, 0x69, 0x76, 0x3e, 0xa, 0x3c, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x20, 0x69, 0x64, 0x3d, 0x22, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x3e, 0xa, 0x20, 0x20, 0x3c, 0x74, 0x72, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x74, 0x68, 0x3e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x3c, 0x2f, 0x74, 0x68, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x74, 0x68, 0x3e, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x3c, 0x2f, 0x74, 0x68, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x74, 0x68, 0x3e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x3c, 0x2f, 0x74, 0x68, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x74, 0x68, 0x3e, 0x45, 0x6e, 0x64, 0x3c, 0x2f, 0x74, 0x68, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x74, 0x68, 0x3e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64, 0x65, 0x72, 0x3c, 0x2f, 0x74, 0x68, 0x3e, 0xa, 0x20, 0x20, 0x3c, 0x2f, 0x74, 0x72, 0x3e, 0xa, 0x3c, 0x2f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x3e, 0xa, 0x3c, 0x2f, 0x64, 0x69, 0x76, 0x3e, 0xa, 0x3c, 0x2f, 0x62, 0x6f, 0x64, 0x79, 0x3e, 0xa, 0x3c, 0x2f, 0x68, 0x74, 0x6d, 0x6c, 0x3e, 0xa}

func File(file string) []byte {
	switch file {

	case "index.html":
		return index_html

	default:
		log.Fatalln(file, "not found")
	}

	return *(*[]byte)(nil)
}
