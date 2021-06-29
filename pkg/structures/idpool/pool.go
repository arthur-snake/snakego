package idpool

type Pool struct {
	set map[string]struct{}
}

func NewPool() *Pool {
	return &Pool{
		set: make(map[string]struct{}),
	}
}

func (p *Pool) Remove(id string) {
	delete(p.set, id)
}

func (p *Pool) Put(id string) {
	p.set[id] = struct{}{}
}

func (p *Pool) Generate() string {
	id := ""
	for {
		id += string(generateChar())
		if _, has := p.set[id]; !has {
			break
		}
	}

	p.Put(id)
	return id
}
