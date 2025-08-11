package d18

type Deque struct {
	data []*Cube
}

func NewQ() *Deque {
	d := Deque{make([]*Cube, 0)}
	return &d
}

func (d *Deque) Size() int {
	return len(d.data)
}

func (d *Deque) Popleft() *Cube {
	if d.Size() == 0 {
		return nil
	} else if d.Size() == 1 {
		c := d.data[0]
		d.data = make([]*Cube, 0)
		return c
	} else {
		c := d.data[0]
		d.data = d.data[1:]
		return c
	}
}

func (d *Deque) Popright() *Cube {
	if d.Size() == 0 {
		return nil
	} else if d.Size() == 1 {
		c := d.data[0]
		d.data = make([]*Cube, 0)
		return c
	} else {
		c := d.data[0]
		d.data = d.data[0 : len(d.data)-1]
		return c
	}
}

func (d *Deque) Pushleft(c *Cube) {
	x := make([]*Cube, 0)
	x = append(x, c)
	x = append(x, d.data...)
	d.data = x
}

func (d *Deque) Pushright(c *Cube) {
	d.data = append(d.data, c)
}
