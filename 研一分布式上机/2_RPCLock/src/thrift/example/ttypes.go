// Autogenerated by Thrift Compiler (0.9.1)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package example

import (
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"math"
)

// (needed to ensure safety because of naive import list construction.)
var _ = math.MinInt32
var _ = thrift.ZERO
var _ = fmt.Printf

var GoUnusedProtection__ int

type Rsp struct {
	CliID    int64  `thrift:"cliID,1"`
	Operator string `thrift:"operator,2"`
	Buffer   string `thrift:"buffer,3"`
}

func NewRsp() *Rsp {
	return &Rsp{}
}

func (p *Rsp) IsSetBuffer() bool {
	return p.Buffer != ""
}

func (p *Rsp) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error", p)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
		case 3:
			if err := p.readField3(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *Rsp) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return fmt.Errorf("error reading field 1: %s")
	} else {
		p.CliID = v
	}
	return nil
}

func (p *Rsp) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return fmt.Errorf("error reading field 2: %s")
	} else {
		p.Operator = v
	}
	return nil
}

func (p *Rsp) readField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return fmt.Errorf("error reading field 3: %s")
	} else {
		p.Buffer = v
	}
	return nil
}

func (p *Rsp) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("Rsp"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("%T write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("%T write struct stop error: %s", err)
	}
	return nil
}

func (p *Rsp) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("cliID", thrift.I64, 1); err != nil {
		return fmt.Errorf("%T write field begin error 1:cliID: %s", p, err)
	}
	if err := oprot.WriteI64(int64(p.CliID)); err != nil {
		return fmt.Errorf("%T.cliID (1) field write error: %s", p)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 1:cliID: %s", p, err)
	}
	return err
}

func (p *Rsp) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("operator", thrift.STRING, 2); err != nil {
		return fmt.Errorf("%T write field begin error 2:operator: %s", p, err)
	}
	if err := oprot.WriteString(string(p.Operator)); err != nil {
		return fmt.Errorf("%T.operator (2) field write error: %s", p)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 2:operator: %s", p, err)
	}
	return err
}

func (p *Rsp) writeField3(oprot thrift.TProtocol) (err error) {
	if p.IsSetBuffer() {
		if err := oprot.WriteFieldBegin("buffer", thrift.STRING, 3); err != nil {
			return fmt.Errorf("%T write field begin error 3:buffer: %s", p, err)
		}
		if err := oprot.WriteString(string(p.Buffer)); err != nil {
			return fmt.Errorf("%T.buffer (3) field write error: %s", p)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 3:buffer: %s", p, err)
		}
	}
	return err
}

func (p *Rsp) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Rsp(%+v)", *p)
}

type Req struct {
	CliID    int64  `thrift:"cliID,1"`
	Operator string `thrift:"operator,2"`
}

func NewReq() *Req {
	return &Req{}
}

func (p *Req) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error", p)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *Req) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return fmt.Errorf("error reading field 1: %s")
	} else {
		p.CliID = v
	}
	return nil
}

func (p *Req) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return fmt.Errorf("error reading field 2: %s")
	} else {
		p.Operator = v
	}
	return nil
}

func (p *Req) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("Req"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("%T write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("%T write struct stop error: %s", err)
	}
	return nil
}

func (p *Req) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("cliID", thrift.I64, 1); err != nil {
		return fmt.Errorf("%T write field begin error 1:cliID: %s", p, err)
	}
	if err := oprot.WriteI64(int64(p.CliID)); err != nil {
		return fmt.Errorf("%T.cliID (1) field write error: %s", p)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 1:cliID: %s", p, err)
	}
	return err
}

func (p *Req) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("operator", thrift.STRING, 2); err != nil {
		return fmt.Errorf("%T write field begin error 2:operator: %s", p, err)
	}
	if err := oprot.WriteString(string(p.Operator)); err != nil {
		return fmt.Errorf("%T.operator (2) field write error: %s", p)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 2:operator: %s", p, err)
	}
	return err
}

func (p *Req) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Req(%+v)", *p)
}
