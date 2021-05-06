// Code generated by "stringer -type=ChatMemberKind,ChatType -output stringer.go -linecomment"; DO NOT EDIT.

package domain

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Creator-1]
	_ = x[Administrator-2]
	_ = x[Member-3]
	_ = x[Restricted-4]
	_ = x[Left-5]
	_ = x[Kicked-6]
}

const _ChatMemberKind_name = "creatoradministratormemberrestrictedleftkicked"

var _ChatMemberKind_index = [...]uint8{0, 7, 20, 26, 36, 40, 46}

func (i ChatMemberKind) String() string {
	i -= 1
	if i >= ChatMemberKind(len(_ChatMemberKind_index)-1) {
		return "ChatMemberKind(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _ChatMemberKind_name[_ChatMemberKind_index[i]:_ChatMemberKind_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Sender-1]
	_ = x[Private-2]
	_ = x[Group-3]
	_ = x[Supergroup-4]
	_ = x[Channel-5]
}

const _ChatType_name = "senderprivategroupsupergroupchannel"

var _ChatType_index = [...]uint8{0, 6, 13, 18, 28, 35}

func (i ChatType) String() string {
	i -= 1
	if i >= ChatType(len(_ChatType_index)-1) {
		return "ChatType(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _ChatType_name[_ChatType_index[i]:_ChatType_index[i+1]]
}
