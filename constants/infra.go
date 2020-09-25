package constants

// message db connection.
const (
	ConnectDBReadSuccess string = "Connected to DB \"read\""

	ConnectDBFail string = "Could not connect database, error"

	ClosingDBWriteSuccess string = "Database \"write\" conn gracefully close"
	ClosingDBReadSuccess  string = "Database \"read\" conn gracefully close"
	ClosingDBReadFailed   string = "Error closing DB \"read\" connection"
	ClosingDBWriteFailed  string = "Error closing DB \"write\" connection"
)
