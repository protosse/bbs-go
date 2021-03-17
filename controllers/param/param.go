package param

import (
	"bbs-go/util"
	"strings"
)

// UserName...
type UserName struct {
	Value string
}

func (p *UserName) UnmarshalJSON(data []byte) error {
	v := strings.Trim(string(data), `"`)
	if err := util.IsUsername(v); err != nil {
		return err
	}
	p.Value = v
	return nil
}

func (p *UserName) MarshalJSON(data []byte, err error) {
	if p != nil {
		data = []byte(p.Value)
	}
	return
}

// Password...
type Password struct {
	Value string
}

func (p *Password) UnmarshalJSON(data []byte) error {
	v := strings.Trim(string(data), `"`)
	if err := util.IsPassword(v); err != nil {
		return err
	}
	p.Value = v
	return nil
}

func (p *Password) MarshalJSON(data []byte, err error) {
	if p != nil {
		data = []byte(p.Value)
	}
	return
}

// Email...
type Email struct {
	Value string
}

func (p *Email) UnmarshalJSON(data []byte) error {
	v := strings.Trim(string(data), `"`)
	if err := util.IsEmail(v); err != nil {
		return err
	}
	p.Value = v
	return nil
}

func (p *Email) MarshalJSON(data []byte, err error) {
	if p != nil {
		data = []byte(p.Value)
	}
	return
}
