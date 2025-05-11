package qerror

import (
	"errors"
)

// Định nghĩa các constants error
var (
	ErrReceiverAlreadyHasInvitation = errors.New("receiver already has a pending or accepted invitation")
	ErrNotClassroomOwner            = errors.New("user is not the owner of the classroom")
	ErrInvitationNotFound           = errors.New("invitation not found")
	ErrInvitationNotPending         = errors.New("invitation is not in pending status")
	ErrNotInvitationReceiver        = errors.New("user is not the receiver of this invitation")
	ErrUserAlreadyClassroomMember   = errors.New("user is already a member of this classroom")
	ErrCannotInviteYourself         = errors.New("cannot invite yourself")
	ErrEmailNotFound                = errors.New("email not found")
)
