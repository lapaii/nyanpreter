package registers

func (r *Registers) SetEqualFlag(value bool) {
	if value {
		r.flags = r.flags | Equal // sets equal bit to 1
	} else {
		r.flags = r.flags & ^Equal // sets equal bit to 0
	}
}

func (r *Registers) GetEqualFlag() bool {
	return r.flags&Equal == Equal
}

func (r *Registers) SetGreaterThanFlag(value bool) {
	if value {
		r.flags = r.flags | GreaterThan // sets gt bit to 1
	} else {
		r.flags = r.flags & ^GreaterThan // sets gt bit to 0
	}
}

func (r *Registers) GetGreaterThanFlag() bool {
	return r.flags&GreaterThan == GreaterThan
}
