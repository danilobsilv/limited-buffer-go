package processes

import (
	"fmt"
	"limited_buffer_golang/buffer"
	"math/rand"
	"sync"
	"time"
)

// Producer representa um produtor que gera itens e os adiciona ao buffer.
type Producer struct {
	ID            int           
	buffer        *buffer.Buffer 
	wg            *sync.WaitGroup 
	itemsToProduce int            
}

// NewProducer cria uma nova instância de Producer.
func NewProducer(id int, buf *buffer.Buffer, wg *sync.WaitGroup, items int) *Producer {
	return &Producer{
		ID:            id,
		buffer:        buf,
		wg:            wg,
		itemsToProduce: items,
	}
}

// Run inicia a execução do produtor, gerando e inserindo itens no buffer.
func (p *Producer) Run() {
	defer p.wg.Done() // Decrementa o contador do WaitGroup quando terminar

	for i := 0; i < p.itemsToProduce; i++ {
		item := rand.Intn(100) 

		// Simula um tempo aleatório de produção antes de adicionar ao buffer
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)

		p.buffer.Add(item) 
		fmt.Printf("Producer %d produced: %d\n", p.ID, item)
	}

	fmt.Printf("Producer %d finished\n", p.ID)
}
