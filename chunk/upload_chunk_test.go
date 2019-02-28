// Copyright 2018-present The Yumcoder Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// 
// Author: yumcoder (omid.jn@gmail.com)
//
package chunk

import (
	"testing"
)

func Test_UploadChunkSize(t *testing.T) {
	// uploadChunkSize = (int) Math.max(slowNetwork ? minUploadChunkSlowNetworkSize : minUploadChunkSize, (totalFileSize + 1024 * 3000 - 1) / (1024 * 3000));
	// if (1024 % uploadChunkSize != 0) {
	//   int chunkSize = 64;
	//   while (uploadChunkSize > chunkSize) {
	//	   chunkSize *= 2;
	//   }
	//   uploadChunkSize = chunkSize;
	// }
	// maxRequestsCount = Math.max(1, (slowNetwork ? maxUploadingSlowNetworkKBytes : maxUploadingKBytes) / uploadChunkSize);

	minUploadChunkSize := 128
	minUploadChunkSlowNetworkSize := 32

	//slowNetwork := true

	totalFileSize := 20*1024*1024
	m2 := (totalFileSize + 1024 * 3000 - 1) / (1024 * 3000)

	t.Log("minUploadChunkSize: " , minUploadChunkSize)
	t.Log("minUploadChunkSlowNetworkSize: ", minUploadChunkSlowNetworkSize)
	t.Log(m2, "", 1024%m2)

	uploadChunkSize := minUploadChunkSlowNetworkSize
	if uploadChunkSize < m2{
		uploadChunkSize = m2
	}
	if 1024%m2 != 0 {
		chunkSize := 64
		for uploadChunkSize > chunkSize {
			chunkSize *= 2
		}
		uploadChunkSize = chunkSize
	}
	t.Log(uploadChunkSize)
}
