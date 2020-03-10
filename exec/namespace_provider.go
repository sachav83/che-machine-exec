//
// Copyright (c) 2012-2019 Red Hat, Inc.
// This program and the accompanying materials are made
// available under the terms of the Eclipse Public License 2.0
// which is available at https://www.eclipse.org/legal/epl-2.0/
//
// SPDX-License-Identifier: EPL-2.0
//
// Contributors:
//   Red Hat, Inc. - initial API and implementation
//

package exec

import (
	"github.com/sirupsen/logrus"
	"io/ioutil"
)

const (
	NameSpaceFile = "/var/run/secrets/kubernetes.io/serviceaccount/namespace"
)

var namespace string

func readNameSpace() string {
	nsBytes, err := ioutil.ReadFile(NameSpaceFile)
	if err != nil {
		logrus.Fatal("Failed to get NameSpace", err)
	}
	return string(nsBytes)
}

// Get namespace for current Eclipse CHE workspace
func GetNameSpace() string {
	if namespace == "" {
		namespace = readNameSpace()
	}
	return namespace
}