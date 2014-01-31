// Package plaintext implements an Instrumentation on an io.Writer.
package plaintext

import (
	"fmt"
	"io"
	"time"

	"github.com/soundcloud/roshi/instrumentation"
)

type plaintextInstrumentation struct{ io.Writer }

// New returns a new Instrumentation that prints metrics to the passed
// io.Writer. All metrics are prefixed with an appropriate bucket name, and
// take the form e.g. "insert.record.count 10".
func New(w io.Writer) instrumentation.Instrumentation {
	return plaintextInstrumentation{w}
}

func (i plaintextInstrumentation) InsertCall() {
	fmt.Fprintf(i, "insert.call.count 1")
}

func (i plaintextInstrumentation) InsertRecordCount(n int) {
	fmt.Fprintf(i, "insert.record.count %d", n)
}

func (i plaintextInstrumentation) InsertCallDuration(d time.Duration) {
	fmt.Fprintf(i, "insert.call.duration_ms %d", d.Nanoseconds()/1e6)
}

func (i plaintextInstrumentation) InsertRecordDuration(d time.Duration) {
	fmt.Fprintf(i, "insert.record.duration_ms %d", d.Nanoseconds()/1e6)
}

func (i plaintextInstrumentation) InsertQuorumFailure() {
	fmt.Fprintf(i, "insert.quorum_failure.count 1")
}

func (i plaintextInstrumentation) SelectCall() {
	fmt.Fprintf(i, "select.call.count 1")
}

func (i plaintextInstrumentation) SelectFirstResponseDuration(d time.Duration) {
	fmt.Fprintf(i, "select.first_response.duration_ms %d", d.Nanoseconds()/1e6)
}

func (i plaintextInstrumentation) SelectPartialError() {
	fmt.Fprintf(i, "select.partial_error.count 1")
}

func (i plaintextInstrumentation) SelectBlockingDuration(d time.Duration) {
	fmt.Fprintf(i, "select.blocking.duration_ms %d", d.Nanoseconds()/1e6)
}

func (i plaintextInstrumentation) SelectOverheadDuration(d time.Duration) {
	fmt.Fprintf(i, "select.overhead.duration_ms %d", d.Nanoseconds()/1e6)
}

func (i plaintextInstrumentation) SelectDuration(d time.Duration) {
	fmt.Fprintf(i, "select.duration_ms %d", d.Nanoseconds()/1e6)
}

func (i plaintextInstrumentation) SelectSendAllPromotion() {
	fmt.Fprintf(i, "select.send_all_promotion.count 1")
}

func (i plaintextInstrumentation) DeleteCall() {
	fmt.Fprintf(i, "delete.call.count 1")
}

func (i plaintextInstrumentation) DeleteRecordCount(n int) {
	fmt.Fprintf(i, "delete.record.count %d", n)
}

func (i plaintextInstrumentation) DeleteCallDuration(d time.Duration) {
	fmt.Fprintf(i, "delete.call.duration_ms %d", d.Nanoseconds()/1e6)
}

func (i plaintextInstrumentation) DeleteRecordDuration(d time.Duration) {
	fmt.Fprintf(i, "delete.record.duration_ms %d", d.Nanoseconds()/1e6)
}

func (i plaintextInstrumentation) DeleteQuorumFailure() {
	fmt.Fprintf(i, "delete.quorum_failure.count 1")
}

func (i plaintextInstrumentation) KeysFailure() {
	fmt.Fprintf(i, "keys.failure.count 1")
}

func (i plaintextInstrumentation) KeysInstanceCompleted() {
	fmt.Fprintf(i, "keys.instance_completed.count 1")
}

func (i plaintextInstrumentation) KeysClusterCompleted() {
	fmt.Fprintf(i, "keys.cluster_completed.count 1")
}

func (i plaintextInstrumentation) KeysFarmCompleted() {
	fmt.Fprintf(i, "keys.farm_completed.count 1")
}

func (i plaintextInstrumentation) RepairCall() {
	fmt.Fprintf(i, "repair.call.count 1")
}

func (i plaintextInstrumentation) RepairRequestCount(n int) {
	fmt.Fprintf(i, "repair.request.count %d", n)
}

func (i plaintextInstrumentation) RepairDiscarded(n int) {
	fmt.Fprintf(i, "repair.discarded.count %d", n)
}

func (i plaintextInstrumentation) RepairCheckPartialFailure() {
	fmt.Fprintf(i, "repair.check.partial_failure.count 1")
}

func (i plaintextInstrumentation) RepairCheckCompleteFailure() {
	fmt.Fprintf(i, "repair.check.complete_failure.count 1")
}

func (i plaintextInstrumentation) RepairCheckDuration(d time.Duration) {
	fmt.Fprintf(i, "repair.check.duration_ms %d", d.Nanoseconds()/1e6)
}

func (i plaintextInstrumentation) RepairCheckRedundant() {
	fmt.Fprintf(i, "repair.check.redundant.count 1")
}

func (i plaintextInstrumentation) RepairWriteCount() {
	fmt.Fprintf(i, "repair.write.count 1")
}

func (i plaintextInstrumentation) RepairWriteSuccess() {
	fmt.Fprintf(i, "repair.write.success.count 1")
}

func (i plaintextInstrumentation) RepairWriteFailure() {
	fmt.Fprintf(i, "repair.write.failure.count 1")
}

func (i plaintextInstrumentation) RepairWriteDuration(d time.Duration) {
	fmt.Fprintf(i, "repair.write.duration_ms %d", d.Nanoseconds()/1e6)
}
