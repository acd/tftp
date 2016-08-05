package tftp

const (
	ErrCodeNotDefined       = uint16(0)
	ErrCodeFileNotFound     = uint16(1)
	ErrCodeAccessViolation  = uint16(2)
	ErrCodeDiskFull         = uint16(3)
	ErrCodeIllegalOperation = uint16(4)
	ErrCodeUnknownTID       = uint16(5)
	ErrCodeFileExists       = uint16(6)
	ErrCodeNoSuchUser       = uint16(7)
)

var (
	ErrNotDefined       = &TFTPError{ErrCodeNotDefined, "Not defined."}
	ErrFileNotFound     = &TFTPError{ErrCodeFileNotFound, "File not found."}
	ErrAccessViolation  = &TFTPError{ErrCodeAccessViolation, "Access violation."}
	ErrDiskFull         = &TFTPError{ErrCodeDiskFull, "Disk full or allocation exceeded."}
	ErrIllegalOperation = &TFTPError{ErrCodeIllegalOperation, "Illegal TFTP operation."}
	ErrUnknownTID       = &TFTPError{ErrCodeUnknownTID, "Unknown transfer ID."}
	ErrFileExists       = &TFTPError{ErrCodeFileExists, "File already exists."}
	ErrNoSuchUser       = &TFTPError{ErrCodeNoSuchUser, "No such user."}
)

type TFTPError struct {
	code uint16
	msg  string
}

func (e *TFTPError) Error() string {
	return e.msg
}

func Error(code uint16, msg string) *TFTPError {
	return &TFTPError{code, msg}
}
