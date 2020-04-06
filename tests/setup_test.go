// Copyright 2020 CrowdStrike Holdings, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package tests

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/reporters"
	"github.com/onsi/gomega"
)

// mockController is the global mock controller used by tests
var mockController *gomock.Controller

var _ = ginkgo.AfterEach(func() {
	mockController.Finish()
})

var _ = ginkgo.BeforeEach(func() {
	reporter := testReporter{}
	mockController = gomock.NewController(reporter)
})

func TestUtils(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	junitReporter := reporters.NewJUnitReporter("junit.xml")
	ginkgo.RunSpecsWithDefaultAndCustomReporters(t, "Kafka replicator tests suite", []ginkgo.Reporter{junitReporter})
}

type testReporter struct{}

func (r testReporter) Errorf(format string, args ...interface{}) {
	ginkgo.Fail(fmt.Sprintf(format, args...))
}

func (r testReporter) Fatalf(format string, args ...interface{}) {
	ginkgo.Fail(fmt.Sprintf(format, args...))
}
