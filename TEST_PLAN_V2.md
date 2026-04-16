# CLI 完整测试方案 (v2)

## 1. 测试范围

### 1.1 参数分类

| 类别 | 参数 | 选项/范围 |
|------|------|----------|
| **密码生成** | `-l, --length` | 8-128 |
| | `-c, --charset` | upper, lower, digit, symbol, all |
| | `-s, --strategy` | simple, pronounceable, passphrase |
| | `-n, --count` | 1-100 |
| | `--exclude-similar` | true/false |
| **输出控制** | `-q, --quiet` | true/false |
| | `-o, --output` | simple, json, csv, table |
| | `-v, --validate` | true/false |
| **语言** | `-L, --lang` | en, zh, zh-TW, ja, ko, es, fr |

---

## 2. 测试用例矩阵

### 2.1 边界值测试 (PV - Parameter Validation)

| 用例ID | 参数 | 输入值 | 预期结果 | **参数生效验证** |
|--------|------|--------|---------|-----------------|
| PV01 | -l | 7 | ❌ 错误 | - |
| PV02 | -l | 8 | ✅ 成功 | **len=8** |
| PV03 | -l | 16 | ✅ 成功 | **len=16 (默认)** |
| PV04 | -l | 64 | ✅ 成功 | **len=64** |
| PV05 | -l | 127 | ✅ 成功 | **len=127** |
| PV06 | -l | 128 | ✅ 成功 | **len=128** |
| PV07 | -l | 129 | ❌ 错误 | - |
| PV08 | -n | 0 | ❌ 错误 | - |
| PV09 | -n | 1 | ✅ 成功 | **输出1行** |
| PV10 | -n | 50 | ✅ 成功 | **输出50行** |
| PV11 | -n | 100 | ✅ 成功 | **输出100行** |
| PV12 | -n | 101 | ❌ 错误 | - |
| PV13 | -c | invalid | ❌ 错误 | - |
| PV14 | -s | invalid | ❌ 错误 | - |
| PV15 | -L | xx | ✅ 回退英语 | **显示英语** |

### 2.2 参数生效测试 (PE - Parameter Effect)

| 用例ID | 参数 | 验证逻辑 |
|--------|------|---------|
| **长度验证** |||
| PE01 | -l 8 | `len(password) == 8` |
| PE02 | -l 20 | `len(password) == 20` |
| PE03 | -l 32 | `len(password) == 32` |
| PE04 | -l 64 | `len(password) == 64` |
| PE05 | -l 100 | `len(password) == 100` |
| PE06 | -l 128 | `len(password) == 128` |
| **字符集验证** |||
| PE07 | -c upper | `regexp.MustCompile("^[A-Z]+$").MatchString(pwd)` |
| PE08 | -c lower | `regexp.MustCompile("^[a-z]+$").MatchString(pwd)` |
| PE09 | -c digit | `regexp.MustCompile("^[0-9]+$").MatchString(pwd)` |
| PE10 | -c symbol | 仅含 `!@#$%^&*()_+-=[]{}|;:,.<>?` |
| PE11 | -c all | 含大小写+数字+符号 |
| **排除相似字符** |||
| PE12 | --exclude-similar | 不含 `0`, `O`, `1`, `l`, `I`, `\|` |
| **策略验证** |||
| PE13 | -s simple | 随机字符，无规律 |
| PE14 | -s pronounceable | 含至少2个元音+辅音组合 |
| PE15 | -s passphrase | 包含 `-` 分隔符，word模式 |
| **数量验证** |||
| PE16 | -n 1 | `strings.Count(output, "\n") == 0` (无换行) |
| PE17 | -n 3 | `strings.Count(output, "\n") == 2` |
| PE18 | -n 10 | `strings.Count(output, "\n") == 9` |

### 2.3 输出格式测试 (OF - Output Format)

| 用例ID | 格式 | 参数 | 验证点 | **参数生效验证** |
|--------|------|------|--------|-----------------|
| OF01 | simple | 默认 | 无标签 | 不含 "Password:" |
| OF02 | simple | -v | 有标签+信息 | 含 "Password:", "Entropy:", "Strength:" |
| OF03 | simple | -q | 仅密码 | 不含 "Password:", 仅1行 |
| OF04 | json | 默认 | 有效JSON | `json.Valid(output)` |
| OF05 | json | -v | 含完整字段 | 含 password, entropy, strength |
| OF06 | csv | 默认 | 仅1列 | 首行仅 "password" |
| OF07 | csv | -v | 含4列 | 首行含 "password,entropy,strength,crack_time" |
| OF08 | table | 默认 | 无表头 | 仅密码行 |
| OF09 | table | -v | 有表头 | 含 "PASSWORD", "ENTROPY" |
| OF10 | json | -q | 忽略format | 仅密码，无JSON |

### 2.4 多语言测试 (ML - Multi-Language)

| 用例ID | 语言 | 验证点 | **参数生效验证** |
|--------|------|--------|-----------------|
| ML01 | en | 英文帮助 | `strings.Contains(help, "Password length")` |
| ML02 | zh | 中文帮助 | `strings.Contains(help, "密码长度")` |
| ML03 | zh-TW | 繁体中文 | `strings.Contains(help, "密碼長度")` |
| ML04 | ja | 日语帮助 | `strings.Contains(help, "パスワード")` |
| ML05 | ko | 韩语帮助 | `strings.Contains(help, "비밀번호")` |
| ML06 | es | 西班牙语 | `strings.Contains(help, "Contraseña")` |
| ML07 | fr | 法语帮助 | `strings.Contains(help, "Mot de passe")` |
| ML08 | -L 无效 | 回退英语 | 显示英文而非其他语言 |

### 2.5 组合测试 (CB - Combination)

| 用例ID | 参数组合 | **验证点** |
|--------|---------|-----------|
| CB01 | `-l 20 -c lower -n 5` | 5个20字符小写密码 |
| CB02 | `-s passphrase -l 6 -v` | 6个单词+验证信息 |
| CB03 | `-l 32 -o json -v` | JSON格式+验证信息 |
| CB04 | `-L ja -s pronounceable` | 日语+可发音策略 |
| CB05 | `-q -o json` | quiet优先于format |
| CB06 | `-l 64 -c all -s simple -n 50 --exclude-similar` | 全部参数组合 |

---

## 3. 测试执行计划

### 3.1 执行顺序

```
Phase 1: 边界值测试 → 参数生效验证
Phase 2: 输出格式测试
Phase 3: 多语言测试
Phase 4: 组合测试
```

---

## 4. 验证方法

### 4.1 边界值测试脚本

```bash
# PV01-PV07: 长度边界
bin/ohmypassword generate -l 7 2>&1 | grep -q "invalid" && echo "PV01 PASS" || echo "PV01 FAIL"
bin/ohmypassword generate -l 8 | wc -c  # 应为 9 (8字符+换行)
bin/ohmypassword generate -l 16 | wc -c  # 应为 17
bin/ohmypassword generate -l 64 | wc -c  # 应为 65
bin/ohmypassword generate -l 127 | wc -c # 应为 128
bin/ohmypassword generate -l 128 | wc -c # 应为 129
bin/ohmypassword generate -l 129 2>&1 | grep -q "invalid" && echo "PV07 PASS" || echo "PV07 FAIL"

# PV08-PV12: 数量边界
bin/ohmypassword generate -n 0 2>&1 | grep -q "invalid"
bin/ohmypassword generate -n 1 | wc -l  # 应为 1
bin/ohmypassword generate -n 50 | wc -l # 应为 50
bin/ohmypassword generate -n 100 | wc -l # 应为 100
bin/ohmypassword generate -n 101 2>&1 | grep -q "invalid"
```

### 4.2 参数生效验证脚本

```bash
# PE01-PE06: 长度验证
for len in 8 20 32 64 100 128; do
  pwd=$(bin/ohmypassword generate -l $len -q)
  [ ${#pwd} -eq $len ] && echo "PE0x PASS: -l $len" || echo "PE0x FAIL: -l $len"
done

# PE07-PE11: 字符集验证
bin/ohmypassword generate -c upper -q | grep -qE '^[A-Z]+$' && echo "PE07 PASS"
bin/ohmypassword generate -c lower -q | grep -qE '^[a-z]+$' && echo "PE08 PASS"
bin/ohmypassword generate -c digit -q | grep -qE '^[0-9]+$' && echo "PE09 PASS"

# PE12: 排除相似字符
pwd=$(bin/ohmypassword generate --exclude-similar -q)
echo $pwd | grep -qvE '[0OlI|]' && echo "PE12 PASS" || echo "PE12 FAIL"

# PE13-PE15: 策略验证
bin/ohmypassword generate -s passphrase -q | grep -q '-' && echo "PE15 PASS"

# PE16-PE18: 数量验证
for count in 1 3 10; do
  lines=$(bin/ohmypassword generate -n $count -q | wc -l)
  [ $lines -eq $count ] && echo "PE PASS: -n $count"
done
```