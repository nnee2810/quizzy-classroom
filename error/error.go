package error

import (
	"errors"
)

// Định nghĩa các constants error
var (
	// Invitation errors
	ErrReceiverAlreadyHasInvitation = errors.New("receiver already has a pending or accepted invitation")
	ErrNotClassroomOwner            = errors.New("user is not the owner of the classroom")
)
