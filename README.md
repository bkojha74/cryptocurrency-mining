# Cryptocurrency Mining

Cryptocurrencies often use **proof-of-work (PoW)** algorithms to generate new blocks on the blockchain. These algorithms typically require generating a hash with specific properties.

This project demonstrates a simplified mining process in Go by finding a **nonce** that produces a SHA256 hash starting with two `'0'` digits.

---

## 📌 Problem Statement

Implement a function `FindNonce` in Go that computes the smallest `uint64` value (nonce) such that:

- When appended to a given block (in **little-endian order**, least significant byte first),
- The SHA256 hash of the resulting byte slice,
- Has a **hex string representation starting with two `'0'` digits**.

The function should return the minimum such nonce.

---

## 🔧 Function Description

```go
func FindNonce(input []byte) uint64

Parameters
- input []byte: A byte slice containing the new block data.
Return
- The minimum uint64 nonce value that results in a SHA256 hash starting with two zero digits.
```

## 📜 Constraints
- Input slice length < 1024 bytes
- Nonce is a uint64 value
- Guaranteed that the valid nonce ≤ 10^6

## 📥 Input Format (Custom Testing)
- Input contains a single line: a space-separated list of hexadecimal bytes.

## 🧩 Example
### Example Nonce Encoding
If the nonce is 0x01234567890ABCDE, then the appended slice should contain:
```bash
DE BC 0A 89 67 45 23 01
```



## 🧪 Sample Cases
### Sample Case 0
#### Input:
```hex
12 DE FF
```
#### Output:
```hex
17F
```

Explanation: The hash value with nonce 0x017F is:
```hex
005ab76d9afd76cbc5f795a83e305f430acaa31145727994e9fcf1dad4354243
```

It starts with two zeros.

### Sample Case 1
#### Input:
```hex
AB AB
```

#### Output:
```hex
33
```

Explanation: The hash value with nonce 0x33 is:
```hex
000bf7f2c9df3cbcb4b9a029f0c98d92104192458823148000f4bc191eb94125
```

It starts with two zeros.

## ⚙️ Compilation & Execution
1. Pull the code from 
```git
https://github.com/bkojha74/cryptocurrency-mining.git
```

2. Compile the program
Run:
```bash
go build main.go
```

This will produce an executable (main on Linux/macOS, main.exe on Windows).

3. Declare environment variable OUTPUT_PATH

on macos/Linux
```
export OUTPUT_PATH="output.txt"
```
on windows powershell
```
$env:OUTPUT_PATH="output.txt"
```

4. Execute the program
Provide input via standard input:
```bash
echo "12 DE FF" | ./main


Or run interactively:
./main


Then type:
12 DE FF


and press Enter.
```

5. Expected Output
```bash
17F
```

## 🚀 Conceptual Notes
- This simulates proof-of-work mining in a simplified way.
- Real cryptocurrencies (like Bitcoin) require hashes with many leading zeros, making mining computationally intensive.
- Here, the difficulty is reduced to just two leading zeros for demonstration.