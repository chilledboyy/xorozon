# xorozon Library

Гибридная библиотека для:
1. Генерации криптографических ключей на основе хаотических систем
2. XOR-шифрования данных

## Установка
```bash
go get github.com/chilledboyy/xorozon
```

## Использование

### 1. Генерация ключа
```go
key := xorozon.GenerateChaosKey(32, 3.9, 0.123, 0.8, 0.2)
```

### 2. Шифрование данных
```go
data := []byte("secret message")
encrypted, err := xorozon.Encrypt(data, key)
```

### 3. Дешифрование
```go
decrypted, err := xorozon.Decrypt(encrypted, key)
```

### 4. Автогенерация параметров
```go
r, x0, p, q, err := xorozon.GenerateSecureParams()
```

## Пример
```go
package main

import (
	"fmt"
	"github.com/chilledboyy/xorozon"
)

func main() {
	// Автогенерация параметров
	r, x0, p, q, _ := xorozon.GenerateSecureParams()
	
	// Шифрование строки
	msg := "Hello, Chaos!"
	encrypted, key, _ := xorozon.EncryptString(msg, r, x0, p, q)
	
	fmt.Printf("Key: %x\n", key)
	fmt.Printf("Encrypted: %x\n", encrypted)
}
```
