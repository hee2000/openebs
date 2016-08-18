// Copyright 2016 CloudByte, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This file will deal with destroy and destroy related operations w.r.t
// OpenEBS storage. Appropriate structures in this file will be used to
// invoke VSM destroy operations and their variants.

// NOTE - There can be multiple variations of destroying a VSM.
// NOTE - A destroy variant is thought of as an orthogonal action
//        i.e. applying some aspect(s) over regular destroy operation.

// NOTE - There can be multiple variations of destroying a VSM.
// Below represents some samples of VSM destroy variants:
//   e.g. - Destroy a VSM.
//   e.g. - Destroy a VSM forcibly.
//   e.g. - Simulate destroying a VSM without actually destroying rather
//          verifying if destroy is appropriate without any side-effects.
//   e.g. - Destroy a VSM and return various stats before and after the
//          operation.

package daemon