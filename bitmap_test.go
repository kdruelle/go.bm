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

import(
    "testing"
)

func TestInitFromString(t *testing.T) {
    b := NewFromString("1001100001000110")
    
    for i := 1; i != 16; i++ {
        if i == 1 || i == 4 || i == 5 || i == 10 || i == 14 || i == 15 {
            if !b.IsBitSet(uint(i)){
                t.Errorf("Bit %d should be set, it isn't.", i)
            }
        }else  {
            if b.IsBitSet(uint(i)){
                t.Errorf("Bit %d should be unset, it isn't.", i)
            }
        }
    }
}


func TestSetBit(t * testing.T) {
    b := New(2)
    b.SetBit(1)
    b.SetBit(3)
    b.SetBit(11)
    for i := 1; i != 16; i++ {
        if (i == 1 || i == 3 || i == 11) && !b.IsBitSet(uint(i)) {
            t.Errorf("Bit %d should be set, it isn't.", i)
        }else if i != 1 && i != 3 && i != 11 && b.IsBitSet(uint(i)) {
            t.Errorf("Bit %d should be unset, it isn't.", i)
        }
    }
}

func TestClearBit(t * testing.T) {
    b := NewFromString("1111111111111111")
    b.ClearBit(1)
    b.ClearBit(3)
    b.ClearBit(11)
    for i := 1; i != 16; i++ {
        if (i == 1 || i == 3 || i == 11) && b.IsBitSet(uint(i)) {
            t.Errorf("Bit %d should be set, it isn't.", i)
        }else if i != 1 && i != 3 && i != 11 && !b.IsBitSet(uint(i)) {
            t.Errorf("Bit %d should be unset, it isn't.", i)
        }
    }
}

func TestToggleBit(t *testing.T) {
    b := NewFromString("1001100001000110")
    
    for i := 1; i != 16; i++ {
        b.ToggleBit(uint(i))
        if i == 1 || i == 4 || i == 5 || i == 10 || i == 14 || i == 15 {
            if b.IsBitSet(uint(i)){
                t.Errorf("Bit %d should be unset, it isn't.", i)
            }
        }else  {
            if !b.IsBitSet(uint(i)){
                t.Errorf("Bit %d should be set, it isn't.", i)
            }
        }
    }
}

func TestString(t *testing.T) {
    b := NewFromString("0000111111001100")
    if b.String() != "0000111111001100" {
        t.Errorf("Exepted %s, give %s", "0000111111001100", b)
    }
    b = New(2)
    b.SetBit(1)
    b.SetBit(3)
    b.SetBit(11)
    if b.String() != "1010000000100000" {
        t.Errorf("Exepted %s, give %s", "1010000000100000", b)
    }
}


