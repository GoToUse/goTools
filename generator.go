// generator.go imitates python's generator, get the value with Next method.

package goTools

type GenerateI interface {
	Generate()
	Next() (int, bool)
	Reset()
	Close()
}

// Generator 定义生成器类型
type Generator struct {
	start   int
	end     int
	step    int
	current int
	ch      chan int
	control chan struct{}
}

// NewGenerator 创建一个新的生成器实例
func NewGenerator(start, end, step int) *Generator {
	return &Generator{
		start:   start,
		end:     end,
		step:    step,
		current: start,
		ch:      make(chan int),
		control: make(chan struct{}),
	}
}

// Generate 启动生成过程
func (g *Generator) Generate() {
	go func() {
		for g.current < g.end {
			select {
			case g.ch <- g.current:
				g.current += g.step
			case <-g.control:
				return
			}
		}
		close(g.ch)
	}()
}

// Next 获取下一个值
func (g *Generator) Next() (int, bool) {
	val, open := <-g.ch
	return val, open
}

// Reset 重置生成器状态
func (g *Generator) Reset() {
	g.current = g.start
	g.ch = make(chan int)
}

// Close 停止生成过程
func (g *Generator) Close() {
	g.control <- struct{}{}
}
