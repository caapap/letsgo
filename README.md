# 🚀 云原生技术栈学习指南

> 面向多个技术岗位的系统化学习平台，高效学习，快速进阶！

## 🎯 支持岗位

| 岗位类型 | 核心技术栈 | 学习路径 |
|---------|-----------|---------|
| **Go后端开发** | Go + 微服务 + 数据库 | [Go开发路径](#go后端开发路径) |
| **Agent开发** | Go + 系统编程 + 网络 | [Agent开发路径](#agent开发路径) |
| **K8s工程师** | K8s + 容器 + 网络 | [K8s工程师路径](#k8s工程师路径) |
| **运维开发工程师** | K8s + Go + 监控 | [运维开发路径](#运维开发路径) |
| **SRE工程师** | 监控 + 自动化 + 故障处理 | [SRE路径](#sre工程师路径) |
| **区块链开发** | Go + 区块链 + 密码学 | [区块链路径](#区块链开发路径) |

## 📚 学习模式

### 🎯 快速复习模式（1-2小时）
适合面试前突击复习，涵盖高频考点

### 🎯 专项提升模式（半天）
针对特定岗位的核心技能深度学习

### 📖 系统学习模式（1-2天）
完整的技术栈体系化学习

---

## 🎯 岗位学习路径

### Go后端开发路径
**目标**: 成为优秀的Go后端开发工程师
**时长**: 4-6周
**核心技能**: Go语言 + Web框架 + 数据库 + 微服务

#### 快速复习（2小时）
- [Go语言核心概念](./docs/foundations/quick-reference.md#go语言要点)
- [高频算法题](./docs/foundations/algorithms.md)
- [后端面试题集](./docs/positions/go-backend-developer.md#面试重点)

#### 系统学习
- [完整学习路径](./docs/positions/go-backend-developer.md)

### K8s工程师路径
**目标**: 掌握容器编排和云原生技术
**时长**: 6-8周
**核心技能**: K8s + 容器 + 网络 + 存储 + 监控

#### 快速复习（2小时）
- [K8s核心概念](./docs/foundations/quick-reference.md#kubernetes核心问题)
- [K8s面试题](./docs/specializations/k8s-interview.md)

#### 系统学习
- [完整学习路径](./docs/positions/k8s-engineer.md)

### 运维开发路径
**目标**: 具备开发能力的运维工程师
**时长**: 6-8周
**核心技能**: K8s + Go + 监控 + 自动化

#### 快速复习（2小时）
- [运维核心知识](./docs/foundations/quick-reference.md)
- [运维面试题集](./docs/positions/devops-engineer.md)

#### 系统学习
- [完整学习路径](./docs/positions/devops-engineer.md)

### SRE工程师路径
**目标**: 站点可靠性工程师
**时长**: 8-10周
**核心技能**: 监控 + 自动化 + 故障处理 + 性能优化

#### 快速复习（2小时）
- [SRE核心理念](./docs/positions/sre-engineer.md#核心理念)
- [故障处理流程](./docs/practices/troubleshooting.md)
- [SRE面试题库](./interviews/questions/monitoring.md)
- [技术基础题库](./interviews/questions/technical-basics.md)

#### 系统学习
- [完整学习路径](./docs/positions/sre-engineer.md)
- [SRE面试准备中心](./interviews/)

### Agent开发路径
**目标**: 系统级Agent开发工程师
**时长**: 4-6周
**核心技能**: Go + 系统编程 + 网络 + 安全

#### 快速复习（2小时）
- [系统编程基础](./docs/positions/agent-developer.md#核心技能)
- [网络编程要点](./docs/specializations/network-programming.md)

#### 系统学习
- [完整学习路径](./docs/positions/agent-developer.md)

### 区块链开发路径
**目标**: 区块链应用开发工程师
**时长**: 8-12周
**核心技能**: Go + 区块链 + 密码学 + 智能合约

#### 快速复习（2小时）
- [区块链基础概念](./docs/positions/blockchain-developer.md#基础概念)
- [Go语言在区块链中的应用](./docs/specializations/blockchain-tech.md)

#### 系统学习
- [完整学习路径](./docs/positions/blockchain-developer.md)

## 🗂️ 目录结构

```
cloud-native-learning/
├── README.md                    # 项目总览
├── docs/                        # 文档目录
│   ├── foundations/             # 基础知识
│   │   ├── quick-reference.md   # 核心知识点速查
│   │   ├── algorithms.md        # 算法基础
│   │   ├── go-fundamentals.md   # Go语言基础
│   │   └── container-basics.md  # 容器技术基础
│   ├── positions/               # 岗位学习路径
│   │   ├── go-backend-developer.md    # Go后端开发
│   │   ├── agent-developer.md         # Agent开发
│   │   ├── k8s-engineer.md            # K8s工程师
│   │   ├── devops-engineer.md         # 运维开发工程师
│   │   ├── sre-engineer.md            # SRE工程师
│   │   └── blockchain-developer.md    # 区块链开发
│   ├── specializations/         # 专业技能深入
│   │   ├── k8s-interview.md     # K8s面试题
│   │   ├── k8s-deep-dive.md     # K8s核心原理
│   │   ├── go-concurrency.md    # Go并发编程
│   │   ├── microservices.md     # 微服务架构
│   │   └── blockchain-tech.md   # 区块链技术
│   └── practices/               # 实践经验
│       ├── project-templates.md # 项目经验模板
│       ├── performance-tuning.md # 性能调优
│       └── troubleshooting.md   # 故障排查
├── code/                        # 代码实践
│   ├── algorithms/              # 算法实现
│   ├── go-examples/             # Go语言示例
│   ├── k8s-configs/            # K8s配置文件
│   ├── blockchain/              # 区块链代码
│   └── agent-examples/          # Agent开发示例
└── resources/                   # 学习资源
    ├── books.md                 # 推荐书籍
    ├── courses.md               # 在线课程
    └── tools.md                 # 开发工具
```

## ⚡ 使用指南

### 🎯 选择学习路径
1. **确定目标岗位**: 根据职业规划选择对应的岗位路径
2. **评估当前水平**: 通过快速复习模式测试基础知识掌握程度
3. **制定学习计划**: 根据时间安排选择学习模式

### 📚 学习方式建议
```bash
# 1. 快速评估（30分钟）
cat docs/foundations/quick-reference.md

# 2. 岗位路径学习
# 选择对应岗位的学习路径文档

# 3. 专项技能深入
# 根据岗位要求深入学习特定技术

# 4. 实践项目
cd code/ && # 选择对应的代码示例进行实践
```

### 🔄 学习循环
1. **理论学习** → 2. **代码实践** → 3. **项目应用** → 4. **面试准备**

## 🎪 考试策略

### 技术问题回答框架
1. **理解确认** - 重复问题确保理解正确
2. **思路分析** - 说出解题思路和方案
3. **代码实现** - 边写边解释关键点
4. **测试验证** - 考虑边界情况
5. **优化改进** - 分析时空复杂度

### 项目经验描述(STAR)
- **Situation**: 项目背景和业务场景
- **Task**: 你负责的具体任务
- **Action**: 采取的技术方案
- **Result**: 最终效果和收益

## 🔥 技术栈图谱

### 基础技能（所有岗位必备）
- ✅ **编程语言**: Go语言基础、并发编程
- ✅ **算法基础**: 数据结构、常用算法
- ✅ **系统基础**: Linux、网络、存储
- ✅ **容器技术**: Docker基础使用

### 专业技能（按岗位分类）

#### 后端开发
- 🔹 Web框架、数据库、缓存
- 🔹 微服务架构、API设计
- 🔹 性能优化、监控告警

#### K8s/运维开发
- 🔹 K8s集群管理、网络存储
- 🔹 CI/CD、自动化运维
- 🔹 监控体系、故障排查

#### SRE工程师
- 🔹 可观测性、SLI/SLO
- 🔹 故障响应、容量规划
- 🔹 自动化、混沌工程

#### Agent/系统开发
- 🔹 系统编程、网络编程
- 🔹 安全防护、性能监控
- 🔹 跨平台、协议设计

#### 区块链开发
- 🔹 密码学、共识算法
- 🔹 智能合约、DeFi协议
- 🔹 链上数据分析

### 加分技能
- 🌟 **开源贡献**: GitHub活跃度、PR贡献
- 🌟 **技术分享**: 博客、演讲、技术文章
- 🌟 **大规模实践**: 高并发、大数据、多集群
- 🌟 **新技术探索**: WASM、eBPF、Serverless

## 🎯 面试准备

### 📝 系统化面试准备
- **[面试准备中心](./interviews/)** - 完整的面试准备平台
- **技术题库** - 59道精选面试题，覆盖6大技术领域
- **难度分级** - ⭐⭐⭐ 到 ⭐⭐⭐⭐⭐，循序渐进
- **实战经验** - 字节跳动、猿辅导等公司真实面试经验

### 🔥 热门题库
| 分类 | 题目数量 | 适合岗位 | 推荐用时 |
|------|----------|----------|----------|
| [技术基础](./interviews/questions/technical-basics.md) | 13题 | 所有岗位 | 2-3小时 |
| [Kubernetes](./interviews/questions/kubernetes.md) | 8题 | K8s/SRE | 3-4小时 |
| [监控告警](./interviews/questions/monitoring.md) | 10题 | SRE/运维 | 3-4小时 |
| [编程技术](./interviews/questions/programming.md) | 10题 | 开发/SRE | 4-5小时 |
| [系统设计](./interviews/questions/system-design.md) | 10题 | 高级岗位 | 5-6小时 |
| [案例分析](./interviews/questions/case-studies.md) | 8题 | 高级岗位 | 4-5小时 |

## 🚀 项目特色

### 📋 系统化学习
- **多岗位覆盖**: 6个主流技术岗位完整学习路径
- **分层设计**: 基础→专业→实践的递进式学习
- **模块化**: 可根据需要灵活组合学习内容

### 🎯 实用导向
- **面试友好**: 针对技术面试优化的知识结构
- **项目驱动**: 每个岗位都有对应的实战项目
- **经验总结**: 来自一线工程师的实践经验

### 🔄 持续更新
- **技术跟进**: 跟随云原生技术发展更新内容
- **社区驱动**: 欢迎贡献和反馈改进建议
- **版本管理**: 规范的文档版本和更新记录

## 📞 参与贡献

### 🤝 如何贡献
- 🐛 **发现问题**: 提交Issue报告错误或建议
- 📝 **完善内容**: 提交PR补充或优化文档
- 💡 **分享经验**: 贡献实战案例和最佳实践
- 🌟 **推广项目**: Star项目并分享给需要的朋友

### 📋 贡献指南
1. Fork本项目
2. 创建特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 开启Pull Request

---
**愿景**: 打造云原生时代最实用的技术学习平台，助力每个技术人的职业发展！🚀 

 