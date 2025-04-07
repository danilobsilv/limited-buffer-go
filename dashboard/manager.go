package dashboard

import (
	"fmt"
	"limited_buffer_golang/processes"
	"limited_buffer_golang/buffer"
	"sync"
)

// Config define as configurações para o sistema de produtores e consumidores.
type Config struct {
	BufferSize   int
	NumProducers int
	NumConsumers int
	TotalItems   int
}

// Manager é responsável por gerenciar os produtores e consumidores.
type Manager struct {
	config Config
	buffer *buffer.Buffer
}

// NewManager cria uma nova instância do Manager.
func NewManager(config Config, buf *buffer.Buffer) *Manager {
	return &Manager{
		config: config,
		buffer: buf,
	}
}

// Start inicia o sistema, criando os produtores e consumidores.
func (m *Manager) Start(wg *sync.WaitGroup) {
	fmt.Println("Starting manager...")

	// WaitGroup separado para os produtores
	var prodWg sync.WaitGroup
	prodWg.Add(m.config.NumProducers)

	// Iniciar produtores
	for i := 1; i <= m.config.NumProducers; i++ {
		itemsToProduce := m.config.TotalItems / m.config.NumProducers
		if i == m.config.NumProducers {
			itemsToProduce += m.config.TotalItems % m.config.NumProducers
		}

		go func(id int, items int) {
			producer := processes.NewProducer(id, m.buffer, &prodWg, items)
			producer.Run()
		}(i, itemsToProduce)
	}

	// Iniciar goroutine que fechará o buffer após todos os produtores finalizarem
	go func() {
		prodWg.Wait()
		m.buffer.Close()
	}()

	// Iniciar consumidores
	for i := 1; i <= m.config.NumConsumers; i++ {
		go processes.NewConsumer(i, m.buffer, wg).Run()
	}
}
