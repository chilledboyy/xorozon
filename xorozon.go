package xorozon

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
)

// BLogisticMap генерирует последовательность значений с модифицированным логистическим отображением
// n - количество элементов
// r - параметр хаоса (3.57 < r < 4.0)
// x0 - начальное значение (0 < x0 < 1)
// p - верхняя граница множителя
// q - нижняя граница множителя
func BLogisticMap(n int, r, x0, p, q float64) []float64 {
	sequence := make([]float64, n)
	sequence[0] = x0

	for i := 1; i < n; i++ {
		d := q + (float64(i)+1)/float64(n)*(p-q)
		if (i+1)%2 == 1 {
			d = p - d + q
		}

		sequence[i] = r * sequence[i-1] * (1 - sequence[i-1]) * d
		if sequence[i] <= 0 || sequence[i] >= 1 {
			sequence[i] = 0.5
		}
	}

	return sequence
}

// GenerateChaosKey генерирует байтовый ключ заданного размера
func GenerateChaosKey(size int, r, x0, p, q float64) []byte {
	sequence := BLogisticMap(size*2, r, x0, p, q)
	key := make([]byte, size)
	for i := 0; i < size; i++ {
		val := sequence[size+i]
		key[i] = byte(val * 256)
	}
	return key
}

// GenerateKeyHex генерирует ключ в hex-формате
func GenerateKeyHex(size int, r, x0, p, q float64) string {
	key := GenerateChaosKey(size, r, x0, p, q)
	return hex.EncodeToString(key)
}

// Encrypt выполняет XOR-шифрование данных
func Encrypt(data, key []byte) ([]byte, error) {
	if len(key) == 0 {
		return nil, errors.New("key cannot be empty")
	}

	encrypted := make([]byte, len(data))
	for i := 0; i < len(data); i++ {
		encrypted[i] = data[i] ^ key[i%len(key)]
	}
	return encrypted, nil
}

// Decrypt выполняет XOR-дешифрование (аналогично Encrypt)
func Decrypt(data, key []byte) ([]byte, error) {
	return Encrypt(data, key)
}

// GenerateSecureParams генерирует случайные параметры для хаотической системы
func GenerateSecureParams() (r, x0, p, q float64, err error) {
	buf := make([]byte, 32)
	if _, err = rand.Read(buf); err != nil {
		return
	}

	r = 3.57 + 0.43*float64(buf[0])/255.0
	x0 = float64(buf[1])/256.0 + 0.001
	p = 0.5 + 0.4*float64(buf[2])/255.0
	q = 0.1 + 0.4*float64(buf[3])/255.0

	return r, x0, p, q, nil
}

// EncryptString шифрует строку с автоматической генерацией ключа
func EncryptString(text string, r, x0, p, q float64) ([]byte, []byte, error) {
	data := []byte(text)
	key := GenerateChaosKey(len(data), r, x0, p, q)
	encrypted, err := Encrypt(data, key)
	return encrypted, key, err
}
