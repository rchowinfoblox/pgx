package pgtype

import (
	"fmt"
	"io"
	"math"
	"strconv"

	"github.com/jackc/pgx/pgio"
)

type Int2 struct {
	Int    int16
	Status Status
}

func (dst *Int2) ConvertFrom(src interface{}) error {
	switch value := src.(type) {
	case Int2:
		*dst = value
	case int8:
		*dst = Int2{Int: int16(value), Status: Present}
	case uint8:
		*dst = Int2{Int: int16(value), Status: Present}
	case int16:
		*dst = Int2{Int: int16(value), Status: Present}
	case uint16:
		if value > math.MaxInt16 {
			return fmt.Errorf("%d is greater than maximum value for Int2", value)
		}
		*dst = Int2{Int: int16(value), Status: Present}
	case int32:
		if value < math.MinInt16 {
			return fmt.Errorf("%d is greater than maximum value for Int2", value)
		}
		if value > math.MaxInt16 {
			return fmt.Errorf("%d is greater than maximum value for Int2", value)
		}
		*dst = Int2{Int: int16(value), Status: Present}
	case uint32:
		if value > math.MaxInt16 {
			return fmt.Errorf("%d is greater than maximum value for Int2", value)
		}
		*dst = Int2{Int: int16(value), Status: Present}
	case int64:
		if value < math.MinInt16 {
			return fmt.Errorf("%d is greater than maximum value for Int2", value)
		}
		if value > math.MaxInt16 {
			return fmt.Errorf("%d is greater than maximum value for Int2", value)
		}
		*dst = Int2{Int: int16(value), Status: Present}
	case uint64:
		if value > math.MaxInt16 {
			return fmt.Errorf("%d is greater than maximum value for Int2", value)
		}
		*dst = Int2{Int: int16(value), Status: Present}
	case int:
		if value < math.MinInt16 {
			return fmt.Errorf("%d is greater than maximum value for Int2", value)
		}
		if value > math.MaxInt16 {
			return fmt.Errorf("%d is greater than maximum value for Int2", value)
		}
		*dst = Int2{Int: int16(value), Status: Present}
	case uint:
		if value > math.MaxInt16 {
			return fmt.Errorf("%d is greater than maximum value for Int2", value)
		}
		*dst = Int2{Int: int16(value), Status: Present}
	case string:
		num, err := strconv.ParseInt(value, 10, 16)
		if err != nil {
			return err
		}
		*dst = Int2{Int: int16(num), Status: Present}
	default:
		if originalSrc, ok := underlyingNumberType(src); ok {
			return dst.ConvertFrom(originalSrc)
		}
		return fmt.Errorf("cannot convert %v to Int2", value)
	}

	return nil
}

func (src *Int2) AssignTo(dst interface{}) error {
	return int64AssignTo(int64(src.Int), src.Status, dst)
}

func (dst *Int2) DecodeText(r io.Reader) error {
	size, err := pgio.ReadInt32(r)
	if err != nil {
		return err
	}

	if size == -1 {
		*dst = Int2{Status: Null}
		return nil
	}

	buf := make([]byte, int(size))
	_, err = r.Read(buf)
	if err != nil {
		return err
	}

	n, err := strconv.ParseInt(string(buf), 10, 16)
	if err != nil {
		return err
	}

	*dst = Int2{Int: int16(n), Status: Present}
	return nil
}

func (dst *Int2) DecodeBinary(r io.Reader) error {
	size, err := pgio.ReadInt32(r)
	if err != nil {
		return err
	}

	if size == -1 {
		*dst = Int2{Status: Null}
		return nil
	}

	if size != 2 {
		return fmt.Errorf("invalid length for int2: %v", size)
	}

	n, err := pgio.ReadInt16(r)
	if err != nil {
		return err
	}

	*dst = Int2{Int: int16(n), Status: Present}
	return nil
}

func (src Int2) EncodeText(w io.Writer) error {
	if done, err := encodeNotPresent(w, src.Status); done {
		return err
	}

	s := strconv.FormatInt(int64(src.Int), 10)
	_, err := pgio.WriteInt32(w, int32(len(s)))
	if err != nil {
		return nil
	}
	_, err = w.Write([]byte(s))
	return err
}

func (src Int2) EncodeBinary(w io.Writer) error {
	if done, err := encodeNotPresent(w, src.Status); done {
		return err
	}

	_, err := pgio.WriteInt32(w, 2)
	if err != nil {
		return err
	}

	_, err = pgio.WriteInt16(w, src.Int)
	return err
}