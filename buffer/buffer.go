package buffer

import "sync"

// Buffer representa um buffer limitado baseado em canal para comunicação entre produtores e consumidores.
type Buffer struct {
	queue    chan int  // Canal que armazena os itens do buffer
	mu       sync.Mutex // Mutex para garantir acesso seguro ao buffer (não usado no momento)
	capacity int       // Capacidade máxima do buffer
}

// NewBuffer cria e retorna uma nova instância de Buffer com a capacidade definida.
func NewBuffer(capacity int) *Buffer {
	return &Buffer{
		queue:    make(chan int, capacity), // Inicializa um canal com buffer de tamanho definido
		capacity: capacity,
	}
}

// Add insere um item no buffer. 
// Se o buffer estiver cheio, a goroutine será bloqueada até que haja espaço disponível.
func (b *Buffer) Add(item int) {
	b.queue <- item
}

// Remove retira um item do buffer. 
// Retorna false se o canal estiver fechado e vazio.
func (b *Buffer) Remove() (int, bool) {
	item, ok := <-b.queue
	return item, ok
}

// Close fecha o buffer quando todos os produtores terminam, 
// impedindo que mais itens sejam adicionados e sinalizando o fim do processamento.
func (b *Buffer) Close() {
	close(b.queue)
}
