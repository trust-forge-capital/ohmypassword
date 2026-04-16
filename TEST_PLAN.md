# 测试计划：功能完整性验证

## 执行日期: 2026-04-16

## 目标
验证所有功能完全符合文档设计 (api-spec.md, algorithm.md, i18n.md, security.md)

---

## 阶段 1：修复 Bug (P0)

### Bug: Entropy 未显示在输出中
- **文件**: `cmd/cli/generate.go`
- **问题**: PasswordResult 创建时未传递 Entropy 字段
- **修复**: 添加 `result.Entropy = strength.Entropy`

---

## 阶段 2：单元测试 (P1)

### 2.1 internal/generator/entropy_test.go (新建)
- 测试熵公式: `entropy = length × log2(charset_size)`
- 测试用例:
  - 16 chars, all (94) → ~103.8 bits
  - 8 chars, lower (26) → ~37.6 bits
  - 8 chars, digit (10) → ~26.7 bits

### 2.2 internal/validator/cracktime_test.go (新建)
- 测试 5 种攻击场景的 crack time 计算
- 测试用例:
  - Online throttled: 100/sec
  - Online unthrottled: 1,000/sec
  - Offline slow hash: 1,000,000/sec
  - Offline fast hash: 10B/sec
  - Offline GPU: 1T/sec

### 2.3 internal/validator/strength_boundary_test.go (新建)
- 测试强度边界值
- 测试用例:
  - 27 bits → very_weak
  - 28 bits → weak
  - 35 bits → weak
  - 36 bits → reasonable
  - 59 bits → reasonable
  - 60 bits → strong
  - 79 bits → strong
  - 80 bits → very_strong

### 2.4 internal/strategy/pattern_test.go (新建)
- Pronounceable: 验证辅音-元音交替模式
- Passphrase: 验证连字符分隔 + digit/symbol 后缀

---

## 阶段 3：集成测试 (P1)

### cmd/integration_test.go (新建)
- CLI 标志测试
- 输出格式验证 (JSON/CSV/table 结构)
- 错误处理和 exit code

---

## 阶段 4：手动验证清单

### CLI 命令 (30 条)
```bash
# 基础
./bin/ohmypassword generate
./bin/ohmypassword gen

# 长度
./bin/ohmypassword generate -l 8
./bin/ohmypassword generate -l 128
./bin/ohmypassword generate -l 5  # 应报错

# 字符集
./bin/ohmypassword generate -c upper
./bin/ohmypassword generate -c lower
./bin/ohmypassword generate -c digit
./bin/ohmypassword generate -c symbol
./bin/ohmypassword generate -c all
./bin/ohmypassword generate -c upper,lower,digit
./bin/ohmypassword generate -c invalid  # 应报错

# 策略
./bin/ohmypassword generate -s simple
./bin/ohmypassword generate -s pronounceable
./bin/ohmypassword generate -s passphrase

# 数量
./bin/ohmypassword generate -n 1
./bin/ohmypassword generate -n 100
./bin/ohmypassword generate -n 0  # 应报错

# 输出格式
./bin/ohmypassword generate -o simple
./bin/ohmypassword generate -o json
./bin/ohmypassword generate -o csv
./bin/ohmypassword generate -o table

# 验证
./bin/ohmypassword generate -v
./bin/ohmypassword generate --validate

# 安静模式
./bin/ohmypassword generate -q

# 排除相似
./bin/ohmypassword generate --exclude-similar

# 语言
./bin/ohmypassword generate -L en
./bin/ohmypassword generate -L zh
./bin/ohmypassword generate -L ja

# 命令
./bin/ohmypassword version
./bin/ohmypassword --help
./bin/ohmypassword generate --help

# 错误退出码
./bin/ohmypassword generate -l 5; echo $?
```

---

## 验证检查项

| 功能 | 验证方法 | 预期结果 |
|------|----------|----------|
| Entropy 显示 | `-v -o simple` | 显示 "Entropy: xx.xx bits" |
| JSON Entropy | `-v -o json` | 有 "entropy" 字段 |
| CSV Entropy | `-v -o csv` | 有 entropy 列 |
| Table Entropy | `-v -o table` | 有 ENTROPY 列 |
| 强度 5 级 | 边界值测试 | 正确分级 |
| Crack Time | 5 场景 | 正确计算 |
| crypto/rand | 代码审查 | 仅使用 crypto/rand |
| 7 种语言 | 切换 -L | 正确显示 |

---

## 实施顺序

1. 修复 Entropy Bug (generate.go)
2. 手动验证 Entropy 显示
3. 添加单元测试文件
4. 添加集成测试
5. 运行 `make test`
6. 执行手动验证清单
7. Git commit

---

## 预估工作量

| 任务 | 行数 |
|------|------|
| Bug 修复 | +2 |
| 单元测试 | ~160 |
| 集成测试 | ~100 |
| 总计 | ~260 |