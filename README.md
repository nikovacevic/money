# money
Go package supporting simply money types

# USD
At the time of this writing, `money` only supports USD. The following operations are supported:
```
func NewUSD() *USD  
func ToUSD(f float64) USD  
func (m USD) Float64() float64  
func (m USD) Multiply(f float64) USD  
func (m *USD) Scan(val interface{}) error  
func (m USD) String() string  
```
