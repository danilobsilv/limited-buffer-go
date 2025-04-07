package main

import (
	"fmt"
	"limited_buffer_golang/buffer"
	"limited_buffer_golang/dashboard"
	"sync"
	"time"
)

func runTest(testName string, config dashboard.Config) {
	fmt.Printf("\n========================\n")
	fmt.Printf(" Rodando %s\n", testName)
	fmt.Printf("========================\n")

	// Inicializar buffer e manager
	buf := buffer.NewBuffer(config.BufferSize)
	manager := dashboard.NewManager(config, buf)

	// Criar e configurar WaitGroup para os consumidores apenas
	var wg sync.WaitGroup
	wg.Add(config.NumConsumers)

	// Iniciar sistema
	manager.Start(&wg)

	// Esperar conclusão dos consumidores
	wg.Wait()

	fmt.Printf("%s finalizado\n", testName)
	// Espera para separar visualmente os testes no terminal
	time.Sleep(2 * time.Second)
}

func main() {
	fmt.Println(" Iniciando testes de buffer limitado...")

	tests := []struct {
		name   string
		config dashboard.Config
	}{
		{
			name: "Caso 1 - 1 Produtor, 1 Consumidor, Buffer pequeno",
			config: dashboard.Config{
				BufferSize:   2,
				NumProducers: 1,
				NumConsumers: 1,
				TotalItems:   10,
			},
		},
		{
			name: "Caso 2 - 3 Produtores, 2 Consumidores, Buffer médio",
			config: dashboard.Config{
				BufferSize:   5,
				NumProducers: 3,
				NumConsumers: 2,
				TotalItems:   20,
			},
		},
		{
			name: "Caso 3 - 5 Produtores, 1 Consumidor, Buffer pequeno",
			config: dashboard.Config{
				BufferSize:   3,
				NumProducers: 5,
				NumConsumers: 1,
				TotalItems:   25,
			},
		},
		{
			name: "Caso 4 - 1 Produtor, 5 Consumidores, Buffer grande",
			config: dashboard.Config{
				BufferSize:   10,
				NumProducers: 1,
				NumConsumers: 5,
				TotalItems:   20,
			},
		},
		{
			name: "Caso 5 - 4 Produtores, 4 Consumidores, Buffer muito pequeno",
			config: dashboard.Config{
				BufferSize:   1,
				NumProducers: 4,
				NumConsumers: 4,
				TotalItems:   16,
			},
		},
	}

	// Rodar cada teste
	for _, test := range tests {
		runTest(test.name, test.config)
	}

	fmt.Println("\n Todos os testes foram concluídos.")
}


/*
O que é uma Go Panic?
- é como um erro fatal em tempo de execução. Ele para a execução normal do programa imediatamente e começa a desenrolar a pilha (stack unwinding), ou seja, ele sai das funções ativas até encontrar um recover() (se houver) ou até o programa encerrar com erro.
É equivalente a um "crash controlado" que você pode ou não recuperar, dependendo do contexto.

=============

Por que às vezes dá panic e às vezes não?
Porque o comportamento pode depender de:

--- Concorrência (goroutines) – Em sistemas concorrentes, como o seu, a ordem de execução não é garantida. Então:
	Se um produtor tenta adicionar ao buffer depois que ele foi fechado, o Go dá panic: send on closed channel.
	Mas se todos os produtores terminarem direitinho antes do buffer ser fechado, não há problema.

--- Condições de corrida – Se duas goroutines tentam acessar/modificar a mesma coisa ao mesmo tempo sem sincronização adequada, pode dar erro às vezes e às vezes não (depende da ordem da execução).

--- Uso indevido de canais – Se você:
	fecha o canal (close(chan)) e depois tenta usar chan <- value, dá panic.
	consome de um canal fechado com value, ok := <-chan, não dá panic, ok só será false.

================

---Como fazer tratativas a fim de evitar erros?
	Fechar o buffer apenas depois de todos os produtores terminarem assim, o sistema não será fechado enquanto tiver produtor ativo.
	O sistema se encerra de forma segura, limpa e sem panics.
*/