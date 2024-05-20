package domain

type PayeeStatus struct {
	value   string
	display string
}

func (ps PayeeStatus) Value() string {
	return ps.value
}

func (ps PayeeStatus) Display() string {
	return ps.display
}

var (
	PayeeDraftStatus = PayeeStatus{"DRAFT", "Rascunho"}
	PayeeValidStatus = PayeeStatus{"VALID", "Validado"}
)

func restorePayeeStatus(status string) (PayeeStatus, error) {
	switch status {
	case PayeeDraftStatus.Value():
		return PayeeDraftStatus, nil
	case PayeeValidStatus.Value():
		return PayeeValidStatus, nil
	default:
		return PayeeStatus{}, ErrTemperedValue
	}
}
