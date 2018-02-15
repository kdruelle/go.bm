////////////////////////////////////////////////////////////////////////////////
// 
// (C) 2011 Kevin Druelle <kevin@druelle.info>
//
// this software is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
// 
// This software is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
// 
// You should have received a copy of the GNU General Public License
// along with this software.  If not, see <http://www.gnu.org/licenses/>.
// 
///////////////////////////////////////////////////////////////////////////////

package bm

type bm struct {
    data []byte
    size int
}


func New(size int) (b * bm) {
    b = &bm{
        size : size,
        data : make([]byte, size),
    }
    return
}

func NewFromString(s string) (b * bm) {
    b = New(len(s) / 8)
    for i := 0; i < len(s); i++ {
        if s[i] == '1' {
            b.SetBit(uint(i + 1))
        }
    }
    return
}

func (b * bm) IsBitSet(idx uint) (bool) {
    idx--
    return ((b.data[idx / 8] >> ( 7 - (idx % 8))) & 1) != 0
}

func (b * bm) SetBit(idx uint) {
    idx--
    b.data[idx / 8] |= 1 << (7 - (idx % 8))
}

func (b * bm) ClearBit(idx uint) {
    idx--
    b.data[idx / 8] &= ^(1 << (7 - (idx % 8)))
}

func (b * bm) ToggleBit(idx uint) {
    idx--
    b.data[idx / 8] ^= 1 << (7 - (idx % 8))
}

func (b * bm) String() (s string) {
    for i := 1; i <= b.size * 8; i++ {
        if b.IsBitSet(uint(i)){
            s += "1"
        } else {
            s += "0"
        }
    }
    return
}

