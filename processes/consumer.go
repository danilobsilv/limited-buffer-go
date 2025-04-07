package processes

import (
	"fmt"
	"limited_buffer_golang/buffer"
	"math/rand"
	"sync"
	"time"
)

// Consumer representa um consumidor que retira itens do buffer para processá-los.
type Consumer struct {
	ID     int           // Identificador único do consumidor
	buffer *buffer.Buffer // Referência ao buffer compartilhado
	wg     *sync.WaitGroup // WaitGroup para sincronização das goroutines
}

// NewConsumer cria uma nova instância de Consumer.
func NewConsumer(id int, buf *buffer.Buffer, wg *sync.WaitGroup) *Consumer {
	return &Consumer{
		ID:     id,
		buffer: buf,
		wg:     wg,
	}
}

// Run executa o consumidor, processando itens do buffer até que ele seja fechado.
func (c *Consumer) Run() {
	defer c.wg.Done() // Decrementa o contador do WaitGroup quando terminar

	for {
		// Tenta remover um item do buffer
		item, ok := c.buffer.Remove()
		if !ok {
			break // Sai do loop se o buffer estiver fechado e vazio
		}

		// Simula o tempo de processamento do item
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

		// Exibe mensagem indicando o processamento do item
		fmt.Printf("Consumer %d processed: %d\n", c.ID, item)
	}

	// Mensagem indicando que o consumidor terminou a execução
	fmt.Printf("Consumer %d finished\n", c.ID)
}
